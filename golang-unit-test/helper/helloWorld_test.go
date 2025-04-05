package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)


func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Talif")
	}
}

func TestMain(m *testing.M) {
	fmt.Println("Before unit test")
	
	// ini untuk eksekusi semua unit test yg ada
	m.Run()
	
	fmt.Println("After unit test")
}


func TestHelloTalif(t *testing.T) {
	result := HelloWorld("Talif")
	if result != "Hello Talif" {
		t.Fail()
	}

	fmt.Println("Test Hello Talif done")
	
}

func TestHelloSiti(t *testing.T) {
	result := HelloWorld("Siti")
	if result != "Hello Siti" {
		t.FailNow()
	}
	fmt.Println("Test Hello Siti done")
}
func TestHelloIlham(t *testing.T) {
	result := HelloWorld("Ilham")
	if result != "Hello Ilham" {
		t.Error("Harus nya Hai Ilham")
	}

	fmt.Println("Test Hello Ilham done")
	
}

func TestHelloAbdul(t *testing.T) {
	result := HelloWorld("Abdul")
	if result != "Hello Abdul" {
		t.FailNow()
	}
	fmt.Println("Test Hello Abdul done")
}


func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Talif",
			request:  "Talif",
			expected: "Hello Talif",
		},
		{
			name:     "Abdul",
			request:  "Abdul",
			expected: "Hello Abdul",
		},
		{
			name:     "Parinduri",
			request:  "Parinduri",
			expected: "Hello Parinduri",
		},
		{
			name:     "Arif",
			request:  "Arif",
			expected: "Hello Arif",
		},
		{
			name:     "Solah",
			request:  "Solah",
			expected: "Hello Solah",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Dedo", func(t *testing.T) {
		result := HelloWorld("Dedo")
		require.Equal(t, "Hello Dedo", result, "Result must be 'Hello Dedo'")
	})
	t.Run("Ryan", func(t *testing.T) {
		result := HelloWorld("Ryan")
		require.Equal(t, "Hello Ryan", result, "Result must be 'Hello Ryan'")
	})
	t.Run("Ilham", func(t *testing.T) {
		result := HelloWorld("Ilham")
		require.Equal(t, "Hello Ilham", result, "Result must be 'Hello Ilham'")
	})
}
