package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	discovery "github.com/cortezaproject/corteza-server/discovery/types"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"strconv"
)

type (
	ModuleFieldOptions map[string]interface{}

	ModuleFieldOptionIndex struct {
		// index values in this field for public, protected or private  searches
		Access discovery.Access `json:"access"`

		// Applicable to ref field kinds where we use nested/object type
		// and we want to control how deep do we want to follow the references
		NestingDepth map[string]bool `json:"nestingDepth,omitempty"`

		// When embedding this field into a parent doc, do we skip it?
		SkipNested bool `json:"skipNested,omitempty"`

		// @todo we'll probably define aggregation/buckets params here
		//       so that we can handle faceting properly?

	}
)

const (
	moduleFieldOptionExpression         = "expression"
	moduleFieldOptionIsUnique           = "isUnique"
	moduleFieldOptionIsUniqueMultiValue = "isUniqueMultiValue"
	moduleFieldOptionIndex              = "index"

	moduleFieldNumberOptionPrecision         = "precision"
	moduleFieldNumberOptionPrecisionMin uint = 0
	moduleFieldNumberOptionPrecisionMax uint = 6
)

func (opt *ModuleFieldOptions) Scan(value interface{}) error {
	//lint:ignore S1034 This typecast is intentional, we need to get []byte out of a []uint8
	switch value.(type) {
	case nil:
		*opt = ModuleFieldOptions{}
	case []uint8:
		b := value.([]byte)
		if err := json.Unmarshal(b, opt); err != nil {
			return errors.Wrapf(err, "cannot scan '%v' into ModuleFieldOptions", string(b))
		}
	}

	return nil
}

func (opt ModuleFieldOptions) Value() (driver.Value, error) {
	return json.Marshal(opt)
}

// Bool returns option value for key as boolean true or false
//
// Invalid, non-existing are returned as false
func (opt ModuleFieldOptions) Bool(key string) bool {
	if _, has := opt[key]; has {
		if v, ok := opt[key].(bool); ok {
			return v
		}
	}

	return false
}

func (opt ModuleFieldOptions) Int64(key string) int64 {
	return opt.Int64Def(key, 0)
}

func (opt ModuleFieldOptions) Int64Def(key string, def int64) int64 {
	if val, has := opt[key]; has {
		switch conv := val.(type) {
		case int:
			return int64(conv)
		case int64:
			return conv
		default:
			// to avoid covering every possible type, just convert value into string
			strVal := fmt.Sprintf("%v", val)

			if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
				return intVal
			}

			if floatVal, err := strconv.ParseFloat(strVal, 64); err == nil {
				return int64(floatVal)
			}

		}

	}
	return def
}

// Uint64 returns option value for key casted to uint64
func (opt ModuleFieldOptions) Uint64(key string) uint64 {
	if val, has := opt[key]; has {
		return cast.ToUint64(val)
	}

	return 0
}

// Strings returns option value for key as slice of strings
//
// Invalid, non-existing are returned as nil
func (opt ModuleFieldOptions) Strings(key string) []string {
	if _, has := opt[key]; has {
		if v, ok := opt[key].([]string); ok {
			return v
		}
	}

	return nil
}

// String returns option value for key as single string
//
// Invalid, non-existing are returned as empty string ("")
func (opt ModuleFieldOptions) String(key string) string {
	if _, has := opt[key]; has {
		if v, ok := opt[key].(string); ok {
			return v
		}
	}

	return ""
}

// IsUnique - should value in this field be unique across records?
func (opt ModuleFieldOptions) IsUnique() bool {
	return opt.Bool(moduleFieldOptionIsUnique)
}

// SetIsUnique - should value in this field be unique across records?
func (opt ModuleFieldOptions) SetIsUnique(value bool) {
	opt[moduleFieldOptionIsUnique] = value
}

// IsUniqueMultiValue - should value in this field be unique in the multi-value set?
func (opt ModuleFieldOptions) IsUniqueMultiValue() bool {
	return opt.Bool(moduleFieldOptionIsUniqueMultiValue)
}

func (opt ModuleFieldOptions) SetIsUniqueMultiValue(value bool) {
	// SetIsUniqueMultiValue - should value in this field be unique in the multi-value set?
	opt[moduleFieldOptionIsUniqueMultiValue] = value
}

func (opt ModuleFieldOptions) Precision() (p uint) {
	p = uint(opt.Int64(moduleFieldNumberOptionPrecision))

	if p < moduleFieldNumberOptionPrecisionMin {
		p = moduleFieldNumberOptionPrecisionMin
	} else if p > moduleFieldNumberOptionPrecisionMax {
		p = moduleFieldNumberOptionPrecisionMax
	}

	return
}

func (opt ModuleFieldOptions) SetPrecision(p uint) {
	opt[moduleFieldNumberOptionPrecision] = p
}

// IsUnique - should value in this field be unique across records?
func (opt ModuleFieldOptions) Index() *ModuleFieldOptionIndex {
	if val, has := opt[moduleFieldOptionIndex]; has {
		if i, is := val.(*ModuleFieldOptionIndex); is {
			return i
		}
	}

	return nil
}
