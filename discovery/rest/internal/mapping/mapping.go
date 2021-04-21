package mapping

type (
	Mapping struct {
		Index   string               `json:"index"`
		Mapping map[string]*property `json:"mapping"`
	}

	property struct {
		// https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-types.html
		Type string `json:"type,omitempty"`

		// Boost factor
		// https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-boost.html
		Boost float32 `json:"boost,omitempty"`

		Properties map[string]*property `json:"properties,omitempty"`
	}

	Context struct {
		AccessRestriction string
	}
)

func change() *property {
	return &property{
		Type: "nested",
		Properties: map[string]*property{
			"at": {Type: "date"},
			"by": {Type: "nested", Properties: map[string]*property{
				"id":     {Type: "unsigned_long"},
				"email":  {Type: "keyword"},
				"name":   {Type: "keyword"},
				"handle": {Type: "keyword"},
			}},
		},
	}
}

func security() *property {
	return &property{
		Type: "nested",
		Properties: map[string]*property{
			"allowedRoles": {Type: "unsigned_long"},
			"deniedRoles":  {Type: "unsigned_long"},
		},
	}
}
