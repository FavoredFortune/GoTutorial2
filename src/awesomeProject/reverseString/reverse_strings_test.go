package reverseString

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseStrings(t *testing.T){
	t.Run("it reverses a byte string as expected", func(t *testing.T) {
		testByteA := []byte{'h','e','l','l','o'}
		
		result := reverseString(testByteA)
		
		assert.NotNil(t, result)
		assert.Equal(t, []byte{'o','l','l','e','h'}, result)
	})
	
	t.Run("it can handle a single byte string as expected", func(t *testing.T) {
		testByte := []byte{'b'}
		
		result := reverseString(testByte)
		
		assert.NotNil(t, result)
		assert.Equal(t, testByte, result)
	})
}
