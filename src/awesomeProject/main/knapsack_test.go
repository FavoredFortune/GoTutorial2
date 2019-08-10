package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestItem_String(t *testing.T) {
	t.Run("string prints items", func(t *testing.T) {
		testItem := &Item{Name:"burger", Weight:30, Worth:40}
		result := testItem.String()
		assert.NotNil(t,result)
		assert.Equal(t,"burger: 30, 40", result )
	})
}
