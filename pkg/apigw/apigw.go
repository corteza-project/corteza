package apigw

import (
	"context"

	"github.com/cortezaproject/corteza-server/system/types"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

type (
	storer interface {
		SearchApigwRoutes(ctx context.Context, f types.RouteFilter) (types.RouteSet, types.RouteFilter, error)
		SearchApigwFunctions(ctx context.Context, f types.FunctionFilter) (types.FunctionSet, types.FunctionFilter, error)
	}

	apigw struct {
		log        *zap.Logger
		reg        *registry
		routes     []*route
		dispatcher dispatcher
		storer     storer
		reload     chan bool
	}
)

var (
	// global service
	apiGw *apigw
)

func Service() *apigw {
	return apiGw
}

func Set(a *apigw) {
	apiGw = a
}

// Setup handles the singleton service
func Setup(opts interface{}, log *zap.Logger, dispatcher dispatcher, storer storer) {
	if apiGw != nil {
		return
	}

	apiGw = New(opts, log, dispatcher, storer)
}

func New(opts interface{}, logger *zap.Logger, dispatcher dispatcher, storer storer) *apigw {
	reg := NewRegistry()
	reg.Preload()

	return &apigw{
		log:        logger,
		dispatcher: dispatcher,
		storer:     storer,
		reload:     make(chan bool),
		reg:        reg,
	}
}

func (s *apigw) Reload(ctx context.Context) {
	go func() {
		s.reload <- true
	}()
}

func (s *apigw) loadRoutes(ctx context.Context) (rr []*route, err error) {
	routes, _, err := s.storer.SearchApigwRoutes(ctx, types.RouteFilter{Enabled: true})

	if err != nil {
		return
	}

	for _, r := range routes {
		route := &route{
			ID:       r.ID,
			endpoint: r.Endpoint,
			method:   r.Method,
		}

		rr = append(rr, route)
	}

	return
}

func (s *apigw) loadFunctions(ctx context.Context, route uint64) (ff []*types.Function, err error) {
	ff, _, err = s.storer.SearchApigwFunctions(ctx, types.FunctionFilter{})
	return
}

func (s *apigw) Router(ctx context.Context) func(r chi.Router) {
	return func(r chi.Router) {

		routes, err := s.loadRoutes(ctx)

		if err != nil {
			s.log.Error("could not load routes", zap.Error(err))
			return
		}

		s.Init(ctx, routes...)

		for _, route := range s.routes {
			r.Handle(route.endpoint, route)
		}

		go func() {
			for {
				select {
				case <-s.reload:
					s.log.Debug("got reload signal")

					routes, err := s.loadRoutes(ctx)

					if err != nil {
						s.log.Error("could not reload routes", zap.Error(err))
						return
					}

					s.Init(ctx, routes...)

					for _, route := range s.routes {
						r.Handle(route.endpoint, route)
					}

				case <-ctx.Done():
					s.log.Debug("done! getting out")
					return
				}
			}
		}()
	}
}

// init all the routes
func (s *apigw) Init(ctx context.Context, route ...*route) {
	s.routes = route

	s.log.Debug("initializing routes\n", zap.Int("num", len(s.routes)))

	for _, r := range s.routes {
		r.pipe = &pl{}
		regFuncs, err := s.loadFunctions(ctx, r.ID)

		if err != nil {
			s.log.Error("could not load functions for route", zap.String("route", r.endpoint), zap.Error(err))
			continue
		}

		r.pipe.ErrorHandler(errorHandler{
			name:   "error handler expediter",
			args:   []string{},
			weight: 5,
			step:   0,
		})

		for _, f := range regFuncs {
			fc := functionHandler{}

			h, err := s.reg.Get(f.Ref)

			if err != nil {
				s.log.Error("could not register function for route", zap.String("route", r.endpoint), zap.Error(err))
				continue
			}

			fc.Merge(ctx, h.Meta(f))
			fc.SetHandler(h.Handler())

			r.pipe.Add(fc, f.Params)
		}
	}
}

func (s *apigw) Funcs(kind string) (list functionMetaList) {
	list = s.reg.All()

	if kind != "" {
		list, _ = list.Filter(func(fm *functionMeta) (bool, error) {
			// return fm.
			return fm.Kind == kind, nil
		})
	}

	return
}
