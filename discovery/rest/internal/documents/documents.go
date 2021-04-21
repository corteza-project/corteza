package documents

import (
	"github.com/cortezaproject/corteza-server/pkg/filter"
	"time"
)

type (
	Response struct {
		Filter    Filter     `json:"filter"`
		Documents []Document `json:"documents"`
	}

	Document struct {
		ID     uint64      `json:"id,string"`
		URL    string      `json:"url,omitempty"`
		Source interface{} `json:"source"`
	}

	Filter struct {
		Limit    uint                 `json:"limit"`
		NextPage *filter.PagingCursor `json:"nextPage"`
	}

	docPartialUser struct {
		UserID uint64 `json:"userID,string"`
		Email  string `json:"email,omitempty"`
		Name   string `json:"name,omitempty"`
		Handle string `json:"handle,omitempty"`
	}

	docPartialChange struct {
		At *time.Time      `json:"at,omitempty"`
		By *docPartialUser `json:"by,omitempty"`
	}

	docUser struct {
		ResourceType string            `json:"resourceType"`
		UserID       uint64            `json:"userID,string"`
		Email        string            `json:"email"`
		Name         string            `json:"name,omitempty"`
		Handle       string            `json:"handle,omitempty"`
		Suspended    *time.Time        `json:"suspendedAt,omitempty"`
		Updated      *docPartialChange `json:"updated,omitempty"`
		Created      *docPartialChange `json:"created,omitempty"`
		Deleted      *docPartialChange `json:"deleted,omitempty"`
		Security     docSecurity       `json:"security"`
	}

	docComposeNamespace struct {
		ResourceType string                         `json:"resourceType"`
		NamespaceID  uint64                         `json:"namespaceID,string"`
		Name         string                         `json:"name,omitempty"`
		Handle       string                         `json:"handle,omitempty"`
		Enabled      bool                           `json:"enabled"`
		Meta         docPartialComposeNamespaceMeta `json:"meta"`

		Updated  *docPartialChange `json:"updated,omitempty"`
		Created  *docPartialChange `json:"created,omitempty"`
		Deleted  *docPartialChange `json:"deleted,omitempty"`
		Security docSecurity       `json:"security"`
	}

	docPartialComposeNamespaceMeta struct {
		Subtitle    string `json:"subtitle,omitempty"`
		Description string `json:"description,omitempty"`
	}

	docPartialComposeNamespace struct {
		NamespaceID uint64 `json:"namespaceID,string"`
		Name        string `json:"name,omitempty"`
		Handle      string `json:"handle,omitempty"`
	}

	docComposeModule struct {
		ResourceType string                          `json:"resourceType"`
		ModuleID     uint64                          `json:"moduleID,string"`
		Name         string                          `json:"name,omitempty"`
		Handle       string                          `json:"handle,omitempty"`
		Fields       []*docPartialComposeModuleField `json:"fields"`
		Namespace    docPartialComposeNamespace      `json:"namespace"`
		Updated      *docPartialChange               `json:"updated,omitempty"`
		Created      *docPartialChange               `json:"created,omitempty"`
		Deleted      *docPartialChange               `json:"deleted,omitempty"`
		Security     docSecurity                     `json:"security"`
	}

	docPartialComposeModuleField struct {
		Name  string `json:"name,omitempty"`
		Label string `json:"label,omitempty"`
	}

	docPartialComposeModule struct {
		ModuleID uint64 `json:"moduleID,string"`
		Name     string `json:"name,omitempty"`
		Handle   string `json:"handle,omitempty"`
	}

	docComposeRecord struct {
		ResourceType string                     `json:"resourceType"`
		RecordID     uint64                     `json:"recordID,string"`
		Namespace    docPartialComposeNamespace `json:"namespace"`
		Module       docPartialComposeModule    `json:"module"`
		Values       map[string][]interface{}   `json:"values"`
		Updated      *docPartialChange          `json:"updated,omitempty"`
		Created      *docPartialChange          `json:"created,omitempty"`
		Deleted      *docPartialChange          `json:"deleted,omitempty"`
		Security     docSecurity                `json:"security"`
	}

	docSecurity struct {
		// list of roles that are allowed to read the resource
		AllowedRoles []uint64 `json:"allowedRoles"`

		// list of roles that are disallowed to read the resource
		DeniedRoles []uint64 `json:"deniedRoles"`
	}
)

func makePartialChange(at *time.Time) *docPartialChange {
	if at == nil {
		return nil
	}

	// @todo handle unreadable (access-control) modules
	// @todo attach security info (allow/deny roles)
	return &docPartialChange{At: at}
}
