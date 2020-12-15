package option

import (
	"fmt"

	"github.com/jamestunnell/go-settings/value"
)

// var matchOption = regexp.MustCompile(`^[A-Za-z]+\(.*\)$`)

// Parse attempts to parse the given string into option type and a param value string.
func Parse(valType value.Type, optType Type, paramStr string) (*Option, error) {
	if !valType.Valid() {
		return nil, fmt.Errorf("value type %d is not valid", valType)
	}

	if !optType.Valid() {
		return nil, fmt.Errorf("option type %d is not valid", optType)
	}

	// if !matchOption.MatchString(s) {
	// 	return nil, errors.New("invalid format")
	// }

	// openParenIdx := strings.IndexRune(s, '(')
	// optTypeStr := s[:openParenIdx]
	// paramStr := s[openParenIdx+1 : len(s)-1]

	// optType, ok := optionTypeFromStr(optTypeStr)
	// if !ok {
	// 	return nil, fmt.Errorf("unknown option type %s", optTypeStr)
	// }

	// this will be the case for most options
	valTypeToParse := valType

	switch optType {
	case MinLen, MaxLen:
		valTypeToParse = value.UInt64
	case OneOf:
		switch valType {
		case value.Int64:
			valTypeToParse = value.Int64s
		case value.UInt64:
			valTypeToParse = value.UInt64s
		case value.Float64:
			valTypeToParse = value.Float64s
		case value.String:
			valTypeToParse = value.Strings
		}
	case Less, LessEqual, Greater, GreaterEqual:
		switch valType {
		case value.Int64s:
			valTypeToParse = value.Int64
		case value.UInt64s:
			valTypeToParse = value.UInt64
		case value.Float64s:
			valTypeToParse = value.Float64
		}
	}

	paramVal, err := value.Parse(paramStr, valTypeToParse)
	if err != nil {
		err := fmt.Errorf("failed to parse %s as %s: %v", paramStr, valType.String(), err)
		return nil, err
	}

	return &Option{Type: optType, Param: paramVal}, nil
}
