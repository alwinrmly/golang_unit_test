package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Alwin",
			request: "Alwin",
		},
		{
			name:    "Ramli",
			request: "Ramli",
		},
		{
			name:    "Sasmita",
			request: "Sasmita",
		},
		{
			name:    "Alwin Ramli Sasmita",
			request: "Alwin Ramli Sasmita",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Alwin", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Alwin")
		}
	})
	b.Run("Sasmita", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Sasmita")
		}
	})
}
func BenchmarkHelloWorldAlwin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Alwin")
	}
}

func BenchmarkHelloWorldSasmita(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Sasmita")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Alwin",
			request:  "Alwin",
			expected: "Hello Alwin",
		},
		{
			name:     "Ramli",
			request:  "Ramli",
			expected: "Hello Ramli",
		},
		{
			name:     "Sasmita",
			request:  "Sasmita",
			expected: "Hello Sasmita",
		},
		{
			name:     "Asep",
			request:  "Asep",
			expected: "Hello Asep",
		},
		{
			name:     "Tejo",
			request:  "Tejo",
			expected: "Hello Tejo",
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
	t.Run("Alwin", func(t *testing.T) {
		result := HelloWorld("Alwin")
		require.Equal(t, "Hello Alwin", result, "This is must be 'Hello Alwin'")
	})
	t.Run("Sasmita", func(t *testing.T) {
		result := HelloWorld("Sasmita")
		require.Equal(t, "Hello Sasmita", result, "This is must be 'Hello 'Sasmita'")
	})
}

func TestMain(m *testing.M) {
	fmt.Println("BEFORE UNIT TEST")
	// Before
	m.Run()
	// After
	fmt.Println("AFTER UNIT TEST")

}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Cannot run on Mac OS")
	}

	result := HelloWorld("Alwin")
	require.Equal(t, "Hello Alwin", result, "This is must be 'Hello Alwin'")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Alwin")
	require.Equal(t, "Hello Alwin", result, "This is must be 'Hello Alwin'")
	fmt.Println("TestHelloWorld with Require is Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Alwin")
	assert.Equal(t, "Hello Alwin", result, "This is must be 'Hello Alwin'")
	fmt.Println("TestHelloWorld with Assert is Done")
}

func TestHelloWorldAlwin(t *testing.T) {
	result := HelloWorld("Alwin")

	if result != "Hello Alwin" {
		// unit test failed
		t.Error("Result Must Be 'Hello Alwin'")
	}
	fmt.Println("TestHelloWorldAlwin Done")
}

func TestHelloWorldSasmita(t *testing.T) {
	result := HelloWorld("Sasmita")

	if result != "Hello Sasmita" {
		// unit test failed
		t.Fatal("Result Must Be 'Hello Sasmita'")
	}
	fmt.Println("TestHelloWorldSasmita Done")
}
