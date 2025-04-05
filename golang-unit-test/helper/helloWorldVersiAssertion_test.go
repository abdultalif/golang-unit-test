package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloTalifAssert(t *testing.T) {
	result := HelloWorld("Talif")
	assert.Equal(t, "Hello Talif", result, "Result must be 'Hello Talif'")
	fmt.Println("Test Hello Talif done")
	
}

func TestHelloAbdulRequire(t *testing.T) {
	result := HelloWorld("Abdul")
	if result != "Hello Abdul" {
		require.Equal(t, "Hello Abdul", result, "Result must be 'Hello Abdul'")
	}
	fmt.Println("Test Hello Abdul done")
}