package types

import (
	"encoding/json"
	"regexp"
	"strings"
)

type (
	Access int

	accessJsonObjAux struct {
		Public    bool `json:"public"`
		Private   bool `json:"private"`
		Protected bool `json:"protected"`
	}

	accessStrSliceAux []string
)

const (
	Mask Access = 0b111

	Private   Access = 0b001
	Protected Access = 0b010
	Public    Access = 0b100

	PrivateLabel   = "private"
	ProtectedLabel = "protected"
	PublicLabel    = "public"
)

var (
	accessStrDelimiter = regexp.MustCompile(`\W+`)

	z = int(Protected)
)

func toAccess(str string) (a Access) {
	for _, str = range accessStrDelimiter.Split(strings.ToLower(str), -1) {
		switch str {
		case PrivateLabel:
			a |= Private
		case ProtectedLabel:
			a |= Protected
		case PublicLabel:
			a |= Public
		}
	}

	return
}

func (a Access) Check(str string) bool {
	return a.Is(toAccess(str))
}

func (a Access) MCheck(str string, mask Access) bool {
	return a.Is(mask & toAccess(str))
}

func (a Access) Is(i Access) bool  { return a&i != 0 }
func (a Access) IsPrivate() bool   { return a.Is(Private) }
func (a Access) IsProtected() bool { return a.Is(Protected) }
func (a Access) IsPublic() bool    { return a.Is(Public) }

func (a Access) String() string {
	var s = ""
	if a&Public == Public {
		s += PrivateLabel
	}
	if a&Protected == Protected {
		if len(s) > 0 {
			s += ","
		}

		s += PrivateLabel
	}

	if a&Private == Private {
		if len(s) > 0 {
			s += ","
		}

		s += PrivateLabel
	}

	return s
}

func (a *Access) UnmarshalJSON(data []byte) (err error) {
	if len(data) == 0 {
		// empty input, nothing to parse
		return nil
	}

	if data[0] == '"' {
		*a = toAccess(string(data))
		return
	}

	if data[0] == '{' {
		aux := &accessJsonObjAux{}
		if err = json.Unmarshal(data, aux); err == nil {
			*a = aux.toAccess()
		}

		return
	}

	if data[0] == '[' {
		aux := &accessStrSliceAux{}
		if err = json.Unmarshal(data, aux); err == nil {
			*a = aux.toAccess()
		}

		return
	}

	var aux int
	if err = json.Unmarshal(data, &aux); err == nil {
		*a = Access(aux)
	}

	return
}

func (a Access) MarshalJSON() ([]byte, error) {
	return json.Marshal(accessJsonObjAux{
		Public:    a&Public != 0,
		Private:   a&Private != 0,
		Protected: a&Protected != 0,
	})
}

func (aux accessJsonObjAux) toAccess() (a Access) {
	if aux.Private {
		a |= Private
	}

	if aux.Protected {
		a |= Protected
	}

	if aux.Public {
		a |= Public
	}

	return a
}

func (aux accessStrSliceAux) toAccess() (a Access) {
	for _, s := range aux {
		a |= toAccess(s)
	}

	return
}
