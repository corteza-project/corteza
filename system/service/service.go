package service

import (
	"context"
	"errors"
	"time"

	automationService "github.com/cortezaproject/corteza-server/automation/service"
	"github.com/cortezaproject/corteza-server/pkg/actionlog"
	"github.com/cortezaproject/corteza-server/pkg/eventbus"
	"github.com/cortezaproject/corteza-server/pkg/healthcheck"
	"github.com/cortezaproject/corteza-server/pkg/id"
	"github.com/cortezaproject/corteza-server/pkg/logger"
	"github.com/cortezaproject/corteza-server/pkg/objstore"
	"github.com/cortezaproject/corteza-server/pkg/objstore/minio"
	"github.com/cortezaproject/corteza-server/pkg/objstore/plain"
	"github.com/cortezaproject/corteza-server/pkg/options"
	"github.com/cortezaproject/corteza-server/pkg/rbac"
	"github.com/cortezaproject/corteza-server/store"
	"github.com/cortezaproject/corteza-server/system/automation"
	"github.com/cortezaproject/corteza-server/system/types"
	"go.uber.org/zap"
)

type (
	websocketSender interface {
		Send(kind string, payload interface{}, userIDs ...uint64) error
	}

	Config struct {
		ActionLog options.ActionLogOpt
		Storage   options.ObjectStoreOpt
		Template  options.TemplateOpt
		Auth      options.AuthOpt
		RBAC      options.RBACOpt
	}

	eventDispatcher interface {
		WaitFor(ctx context.Context, ev eventbus.Event) (err error)
		Dispatch(ctx context.Context, ev eventbus.Event)
	}
)

var (
	DefaultObjectStore objstore.Store

	// DefaultStore is an interface to storage backend(s)
	// ng (next-gen) is a temporary prefix
	// so that we can differentiate between it and the file-only store
	DefaultStore store.Storer

	DefaultLogger *zap.Logger

	// DefaultSettings controls system's settings
	DefaultSettings *settings

	// DefaultAccessControl Access control checking
	DefaultAccessControl *accessControl

	DefaultAuthNotification AuthNotificationService

	// CurrentSettings represents current system settings
	CurrentSettings = &types.AppSettings{}

	DefaultActionlog actionlog.Recorder

	DefaultSink *sink

	DefaultAuth        *auth
	DefaultAuthClient  *authClient
	DefaultUser        *user
	DefaultRole        *role
	DefaultApplication *application
	DefaultReminder    ReminderService
	DefaultAttachment  AttachmentService
	DefaultRenderer    TemplateService
	DefaultQueue       *queue
	DefaultRoute       *route

	DefaultStatistics *statistics

	// wrapper around time.Now() that will aid service testing
	now = func() *time.Time {
		c := time.Now().Round(time.Second)
		return &c
	}

	// wrapper around nextID that will aid service testing
	nextID = func() uint64 {
		return id.Next()
	}
)

func Initialize(ctx context.Context, log *zap.Logger, s store.Storer, ws websocketSender, c Config) (err error) {
	var (
		hcd = healthcheck.Defaults()
	)

	// we're doing conversion to avoid having
	// store interface exposed or generated inside app package
	DefaultStore = s

	DefaultLogger = log.Named("service")

	{
		tee := zap.NewNop()
		policy := actionlog.MakeProductionPolicy()

		if !c.ActionLog.Enabled {
			policy = actionlog.MakeDisabledPolicy()
		} else if c.ActionLog.Debug {
			policy = actionlog.MakeDebugPolicy()
			tee = logger.MakeDebugLogger()
		}

		DefaultActionlog = actionlog.NewService(DefaultStore, log, tee, policy)
	}

	DefaultAccessControl = AccessControl()

	DefaultSettings = Settings(ctx, DefaultStore, DefaultLogger, DefaultAccessControl, CurrentSettings)

	if DefaultObjectStore == nil {
		const svcPath = "system"
		if c.Storage.MinioEndpoint != "" {
			var bucket = svcPath
			if c.Storage.MinioBucket != "" {
				bucket = c.Storage.MinioBucket + c.Storage.MinioBucketSep + svcPath
			}

			DefaultObjectStore, err = minio.New(bucket, minio.Options{
				Endpoint:        c.Storage.MinioEndpoint,
				Secure:          c.Storage.MinioSecure,
				Strict:          c.Storage.MinioStrict,
				AccessKeyID:     c.Storage.MinioAccessKey,
				SecretAccessKey: c.Storage.MinioSecretKey,

				ServerSideEncryptKey: []byte(c.Storage.MinioSSECKey),
			})

			log.Info("initializing minio",
				zap.String("bucket", bucket),
				zap.String("endpoint", c.Storage.MinioEndpoint),
				zap.Error(err))
		} else {
			path := c.Storage.Path + "/" + svcPath
			DefaultObjectStore, err = plain.New(path)
			log.Info("initializing store",
				zap.String("path", path),
				zap.Error(err))
		}

		if err != nil {
			return err
		}
	}

	hcd.Add(objstore.Healthcheck(DefaultObjectStore), "ObjectStore/System")

	DefaultRenderer = Renderer(c.Template)
	DefaultAuthNotification = AuthNotification(CurrentSettings, DefaultRenderer, c.Auth)
	DefaultAuth = Auth()
	DefaultAuthClient = AuthClient(DefaultStore, DefaultAccessControl, DefaultActionlog, eventbus.Service())
	DefaultUser = User()
	DefaultRole = Role()
	DefaultApplication = Application(DefaultStore, DefaultAccessControl, DefaultActionlog, eventbus.Service())
	DefaultReminder = Reminder(ctx, DefaultLogger.Named("reminder"), ws)
	DefaultSink = Sink()
	DefaultStatistics = Statistics()
	DefaultAttachment = Attachment(DefaultObjectStore)
	DefaultQueue = Queue()
	DefaultRoute = Route()

	if err = initRoles(ctx, log.Named("rbac.roles"), c.RBAC, eventbus.Service(), rbac.Global()); err != nil {
		return err
	}

	automationService.DefaultUser = DefaultUser

	automationService.Registry().AddTypes(
		automation.User{},
		automation.Role{},
		automation.Template{},
		automation.RenderOptions{},
		automation.RenderedDocument{},
	)

	automation.UsersHandler(
		automationService.Registry(),
		DefaultUser,
	)

	automation.TemplatesHandler(
		automationService.Registry(),
		DefaultRenderer,
	)

	automation.RolesHandler(
		automationService.Registry(),
		DefaultRole,
	)

	return
}

func Activate(ctx context.Context) (err error) {
	// Run initial update of current settings with super-user credentials
	err = DefaultSettings.UpdateCurrent(ctx)
	if err != nil {
		return
	}

	return
}

func Watchers(ctx context.Context) {
	DefaultReminder.Watch(ctx)
	return
}

// isGeneric returns true if given error is generic
func isGeneric(err error) bool {
	g, ok := err.(interface{ IsGeneric() bool })
	return ok && g != nil && g.IsGeneric()
}

// unwrapGeneric unwraps error if error is generic (and wrapped)
func unwrapGeneric(err error) error {
	for {
		if isGeneric(err) {
			err = errors.Unwrap(err)
			continue
		}

		return err
	}
}

// Data is stale when new date does not match updatedAt or createdAt (before first update)
//
// @todo This is the same as in compose.service; do we want to make an util thing?
func isStale(new *time.Time, updatedAt *time.Time, createdAt time.Time) bool {
	if new == nil {
		// Change to true for stale-data-check
		return false
	}

	if updatedAt != nil {
		return !new.Equal(*updatedAt)
	}

	return new.Equal(createdAt)
}
