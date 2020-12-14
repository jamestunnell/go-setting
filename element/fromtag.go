package element

import (
	"fmt"

	"github.com/jamestunnell/go-settings/option"
	"github.com/jamestunnell/go-settings/value"
)

// StructTag provides lookup for struct field tag key-value pairs.
type StructTag interface {
	// Get looks up the struct field tag value using the given key
	Get(key string) string
}

// FromTag makes an element from the given struct tag.
// Returns a non-nil error in case of failure.
// Failure can be caused by: invalid option format, using an option that is
// not allowed with the given value type, invalid param value format, or
// using options that are not compatible with eachother.
func FromTag(fieldName string, tag StructTag, valType value.Type) (*Element, error) {
	const notApplicableFmt = "option type %s is not applicable to value type %s"

	options := []*option.Option{}

	for _, optType := range option.AllTypes() {
		key := optType.String()
		val := tag.Get(key)

		// Empty string means the option is not present or is set to empty
		if val == "" {
			continue
		}

		opt, err := option.Parse(valType, optType, val)
		if err != nil {
			err := fmt.Errorf("failed to parse option '%s': %w", key, err)
			return nil, err
		}

		if !opt.Type.ApplicableTo(valType) {
			err := fmt.Errorf(notApplicableFmt, key, valType.String())
			return nil, err
		}

		if err = checkAgainstExistingOptions(options, opt); err != nil {
			return nil, err
		}

		options = append(options, opt)
	}

	elem := New(fieldName, valType, options...)

	return elem, nil
}

func checkAgainstExistingOptions(options []*option.Option, o *option.Option) error {
	for _, existingOption := range options {
		compatible, err := existingOption.CompatibleWith(o)
		if err != nil {
			return err
		}

		if !compatible {
			return fmt.Errorf(
				"%s is not compatible with %s", o.String(), existingOption.String())
		}
	}

	return nil
}
