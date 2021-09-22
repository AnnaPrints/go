package main

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	fmt.Println(os.Hostname())
	if ans != -2 {
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 2, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func TestGoland(t *testing.T) {
	t.Run("first level", func(t *testing.T) {
		t.Run("second level", func(t *testing.T) {
			time.Sleep(100 * time.Millisecond)
			//t.Fail()
		})
	})
}


func ExampleEscape() {
	fmt.Println(Escape("15px;"))
	fmt.Println(Escape("<value>"))
	// Pre-wrapped with <![CDATA[ should be simply output "as is", even with spaces.
	fmt.Println(Escape(`  <![CDATA[@font-face { font-family: "helloworld"; }]]>`))
	// Wrap with Cdata for convenience
	fmt.Println(Escape(Cdata(`@font-face { font-family: "helloworld"; }  `)))
	fmt.Println(Escape(`</s>`))
	fmt.Println(Escape(`<s>`))
	fmt.Println(Escape(`</s><s>`))
	// Output: 15px;
	// &lt;value&gt;
	//   <![CDATA[@font-face { font-family: "helloworld"; }]]>
	// <![CDATA[@font-face { font-family: "helloworld"; }  ]]>
	// &lt;/s&gt;
	// &lt;s&gt;
	// &lt;/s&gt;&lt;s&gt;
}
// a.go
// Escape returns an escaped xml value for names and attributes.
// It detects if the value is already wrapped with <![CDATA[ in order to avoid
// additional escaping.
func Escape(value string) string {
	return "dummy"
}
// Cdata wraps a value in <![CDATA[ .. ]]>
// using standard library's approach
func Cdata(value string) string {
	return "dummy"
}