package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	//before test
	fmt.Println("THIS LINE EXECUTED BEFORE UNIT TEST")
	fmt.Println("Connecting to database...")
	m.Run() //run all the unit test in this file
	//after test
	fmt.Println("THIS LINE EXECUTED AFTER UNIT TEST")
	fmt.Println("Disconnecting from database...")
}

func TestSayHello(t *testing.T) {
	result := SayHello("Budiman")
	if result != "Hello Budiman" {
		t.Fatal("Result must be Hello Budiman") // used to stop the test
		//there are also t.Error(), t.Fail() and t.FailNow() to mark the test as failed
		//Fail() will mark the test as failed but continue the execution
		//FailNow() will mark the test as failed and stop the execution
		//Error("args") will mark the test as failed then call Fail() and log the error args
		//Fatal("args") will mark the test as failed then call FailNow() and log the error args
	}
}

func TestSayHelloAssert(t *testing.T) {
	result := SayHello("Bro")
	// syntax : assert.Equal(t, "Expected result", result value, "Message if failed")
	assert.Equal(t, "Hello Bro", result, "Result must be Hello Bro") //assert.Equal() call t.Fail() if the result is not equal
	fmt.Println("TestSayHelloAssert Done")                           //this line will be executed because assert.Equal() call t.Fail() not t.FailNow()
}

func TestSayHelloRequire(t *testing.T) {
	result := SayHello("Bro")
	// syntax : require.Equal(t, "Expected result", result value, "Message if failed")
	require.Equal(t, "Hello Bro", result, "Result must be Hello Bro") //require.Equal() call t.FailNow() if the result is not equal
	fmt.Println("TestSayHelloRequire Done")                           //this line will not be executed because require.Equal() call t.FailNow()
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unit test cannot run on windows") //used to skip the test, usually used for OS specific test
	}
	result := SayHello("Bro")
	assert.Equal(t, "Hello Bro", result, "Result must be Hello Bro")
	fmt.Println("TestSkip Done")
}

func TestSubtest(t *testing.T) {
	t.Run("!" /* this is the subtest name */, func(t *testing.T) {
		result := SayHello("!")
		assert.Equal(t, "Hello !", result, "Result must be Hello !")
	})
	t.Run("Bro", func(t *testing.T) {
		result := SayHello("Bro")
		assert.Equal(t, "Hello Bro", result, "Result must be Hello Bro")
	})
	//to run a specific subtest, use command : go test -v -run TestSubtest/Bro
}

func TestTableSayHello(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "SayHello(Bro)",
			request:  "Bro",
			expected: "Hello Bro",
		},
		{
			name:     "SayHello(!)",
			request:  "!",
			expected: "Hello !",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := SayHello(test.request)
			assert.Equal(t, test.expected, result)
			fmt.Println("Test Table", test.name, "Done")
		})
	}
}

func BenchmarkSayHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SayHello("Bro")
	}
}

func BenchmarkSayHello2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SayHello("Bro!")
	}
}
func BenchmarkSub(b *testing.B) {
	b.Run("Bro", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SayHello("Bro")
		}
	})
	b.Run("Bro!", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SayHello("Bro!")
		}
	})
}

func BenchmarkTableSayHello(b *testing.B) {
	benchmarks := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "SayHello(Bro)",
			request:  "Bro",
			expected: "Hello Bro",
		},
		{
			name:     "SayHello(!)",
			request:  "!",
			expected: "Hello !",
		},
	}
	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SayHello(benchmark.request)
			}
			fmt.Println("Benchmark Table /", benchmark.name, "Done")
		})
	}
}

// Run the test with command : go test -v
