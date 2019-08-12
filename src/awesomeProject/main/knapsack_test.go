package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxItemsOnly(t *testing.T){
	t.Run("it packs max items with high worth", func(t *testing.T) {
		store :=[]Item{
			{Name: "kitten", Worth:40, Weight:240},
			{Name: "food", Worth:38, Weight:220},
			{Name: "water", Worth:39, Weight:120},
			{Name: "blanket", Worth:33, Weight:80},
			{Name: "toys", Worth:20, Weight:100},
			{Name: "book", Worth:24, Weight:90},
			{Name: "hat", Worth:28, Weight:30},
			{Name: "camera", Worth:26, Weight:110},
			{Name: "cat-food", Worth:37, Weight:160},
		}
		
		result, err := MaxItemsOnly(80, 540, store)
		fmt.Println(result)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.ElementsMatch(t,result, store)
		
	})
}


//func TestItem_String(t *testing.T) {
//	t.Run("string prints items", func(t *testing.T) {
//		testItem := &Item{Name:"burger", Weight:30, Worth:40}
//		result := testItem.String()
//		assert.NotNil(t,result)
//		assert.Equal(t,"burger: 30, 40", result )
//	})
//}

