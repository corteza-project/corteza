package types

import (
	"encoding/json"
)

func ParseApigwfFunctionParams(ss []string) (p FuncParams, err error) {
	p = FuncParams{}
	return p, parseStringsInput(ss, p)
}

func ParseTemplateMeta(ss []string) (p TemplateMeta, err error) {
	p = TemplateMeta{}
	return p, parseStringsInput(ss, p)
}

func parseStringsInput(ss []string, p interface{}) (err error) {
	if len(ss) == 0 {
		return
	}

	return json.Unmarshal([]byte(ss[0]), &p)
}

func ParseAuthClientMeta(ss []string) (p *AuthClientMeta, err error) {
	p = &AuthClientMeta{}
	return p, parseStringsInput(ss, &p)
}

func ParseAuthClientSecurity(ss []string) (p *AuthClientSecurity, err error) {
	p = &AuthClientSecurity{}
	return p, parseStringsInput(ss, &p)
}

func ParseRoleMeta(ss []string) (p *RoleMeta, err error) {
	p = &RoleMeta{}
	return p, parseStringsInput(ss, &p)
}
