package value_test

// import (
// 	"testing"

// 	"github.com/jamestunnell/go-setting/value"
// 	"github.com/stretchr/testify/assert"
// )

// func TestStringKnownTypes(t *testing.T) {
// 	strings := []string{
// 		value.TypeInt.TypeString(),
// 		value.TypeIntSlice.TypeString(),
// 		value.TypeUInt.TypeString(),
// 		value.TypeUIntSlice.TypeString(),
// 		value.TypeFloatSlice.TypeString(),
// 		value.TypeFloatSlice.TypeString(),
// 		value.TypeBoolSlice.TypeString(),
// 		value.TypeBoolSlice.TypeString(),
// 		value.TypeString.TypeString(),
// 		value.TypeStringSlice.TypeString(),
// 	}

// 	for i := 0; i < len(strings); i++ {
// 		assert.NotEmpty(t, strings[i])

// 		for j := 0; j < len(strings); j++ {
// 			if i != j {
// 				assert.NotEqual(t, strings[j], strings[i])
// 			}
// 		}
// 	}
// }

// func TestStringUnknownType(t *testing.T) {
// 	assert.Empty(t, value.Type(-1).TypeString())
// }
