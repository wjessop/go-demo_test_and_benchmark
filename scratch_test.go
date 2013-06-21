package main

import (
	"github.com/remogatto/prettytest"
	"regexp"
	"strings"
	"testing"
)

// Start of setup
type testSuite struct {
	prettytest.Suite
}

func TestRunner(t *testing.T) {
	prettytest.RunWithFormatter(
		t,
		new(prettytest.TDDFormatter),
		new(testSuite),
	)
}

// End of setup

func Benchmark_AnchoredRegex(b *testing.B) { //benchmark function starts with "Benchmark" and takes a pointer to type testing.B
	str := "counters something else"
	for i := 0; i < b.N; i++ { //use b.N for looping
		regexp.MatchString("^counters|^gauges", str)
	}
}

func Benchmark_Prefix(b *testing.B) { //benchmark function starts with "Benchmark" and takes a pointer to type testing.B
	stringGenerator := countersOrGuages()
	for i := 0; i < b.N; i++ { //use b.N for looping
		str := stringGenerator()
		if strings.HasPrefix(str, "counters") || strings.HasPrefix(str, "gauges") {}
	}
}

func Benchmark_Contains(b *testing.B) {
	str := "This is the string we want to know contains the string 'value' or not"
	for i := 0; i < b.N; i++ { //use b.N for looping
		strings.Contains(str, "value")
	}
}

func Benchmark_NonAnchoredRegex(b *testing.B) { //benchmark function starts with "Benchmark" and takes a pointer to type testing.B
	str := "This is the string we want to know contains the string 'value' or not"
	for i := 0; i < b.N; i++ { //use b.N for looping
		regexp.MatchString("value", str)
	}
}

type generator func() string

func countersOrGuages() generator {
	count := 1
	strings := [...]string{"counters is the first word in the string", "guages is the first word in the string"}
	return func() string {
		count = 1 - count
		return strings[count]
	}
}

func (t *testSuite) TestCountersOrGuages() {
	stringGenerator := countersOrGuages()
	t.Equal(stringGenerator(), "counters is the first word in the string")
	t.Equal(stringGenerator(), "guages is the first word in the string")
	t.Equal(stringGenerator(), "counters is the first word in the string")
	t.Equal(stringGenerator(), "guages is the first word in the string")
}
