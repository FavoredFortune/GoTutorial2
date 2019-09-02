package validAnagram

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsAnagram(t *testing.T){
	t.Run("it correctly finds an anagram with lowercase letters", func(t *testing.T) {
		stringA := "anagram"
		stringB := "nagaram"
		
		result := isAnagram(stringA,stringB)
		
		assert.NotNil(t, result)
		assert.True(t, result)
	})
	
	t.Run("it correctly finds when two strings aren't an anagram with lowercase letters", func(t *testing.T) {
		stringA := "anagram"
		stringB := "bagaram"
		
		result := isAnagram(stringA,stringB)
		
		assert.NotNil(t, result)
		assert.False(t, result)
	})
	t.Run("it correctly finds when two strings aren't an anagram with lowercase letters", func(t *testing.T) {
		stringA := "anagram"
		stringB := "agaram"
		
		result := isAnagram(stringA,stringB)
		
		assert.NotNil(t, result)
		assert.False(t, result)
	})
	t.Run("it correctly finds when two strings aren't an anagram with lowercase letters when one string is empty",
		func(t *testing.T) {
		stringA := "anagram"
		stringB := ""
		
		result := isAnagram(stringA,stringB)
		
		assert.NotNil(t, result)
		assert.False(t, result)
	})
	t.Run("it correctly finds when two strings are an anagram when both are empty", func(t *testing.T) {
		stringA := ""
		stringB := ""
		
		result := isAnagram(stringA,stringB)
		
		assert.NotNil(t, result)
		assert.True(t, result)
	})
	
	t.Run("it correctly finds when two strings are an anagram when upper and lower case are used", func(t *testing.T) {
		stringA := "Abba"
		stringB := "abbA"
		
		result := isAnagram(stringA,stringB)
		
		assert.NotNil(t, result)
		assert.True(t, result)
	})
	
	t.Run("it correctly finds when two strings are an anagram when upper and lower case are used more complex",
		func(t *testing.T) {
		stringA := "CompLex"
		stringB := "LexompC"
		
		result := isAnagram(stringA,stringB)
		
		assert.NotNil(t, result)
		assert.True(t, result)
	})
	
	t.Run("it correctly finds when two strings are an anagram when upper and lower case and characters",
		func(t *testing.T) {
			stringA := "CompLex!"
			stringB := "Lexomp!C"
			
			result := isAnagram(stringA,stringB)
			
			assert.NotNil(t, result)
			assert.True(t, result)
		})
}