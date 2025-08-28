package restricted

import (
	"flag"
	"testing"
)

func TestSimple(t *testing.T) {
	n := NewNumber[int](0, 4)
	v := n.Get()
	if v != 0 {
		t.Fatalf("must be 0 but %v", v)
	}

	n = NewNumber[int](1, 4)
	v = n.Get()
	if v != 1 {
		t.Fatalf("must be 1 but %v", v)
	}

	err := n.Set("2")
	if err != nil {
		t.Fatal(err)
	}
	v = n.Get()
	if v != 2 {
		t.Fatalf("must be 2 but %v", v)
	}
}

func TestInt(t *testing.T) {
	n := NewNumber(1, 4)

	flag.CommandLine = flag.NewFlagSet("test", flag.ContinueOnError)
	flag.Var(n, "number", "number between 1 to 4")

	args := []string{
		"test",
		"-number",
		"0",
	}
	err := flag.CommandLine.Parse(args)
	if err != nil {
		t.Fatal(err)
	}

	v := n.Get()
	if v != 1 {
		t.Fatalf("must be 1 but %v", v)
	}

}

func TestFloat(t *testing.T) {
	n := NewNumber(1.2, 3.4)

	flag.CommandLine = flag.NewFlagSet("test", flag.ContinueOnError)
	flag.Var(n, "number", "number between 1.2 to 3.4")

	args := []string{
		"test",
		"-number",
		"0",
	}
	err := flag.CommandLine.Parse(args)
	if err != nil {
		t.Fatal(err)
	}

	v := n.Get()
	if v != 1.2 {
		t.Fatalf("must be 1 but %v", v)
	}

}
