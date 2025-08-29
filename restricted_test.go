package restricted

import (
	"flag"
	"testing"
)

func TestInt(t *testing.T) {
	n := NewNumber[int](0, 0, 4)
	v := n.Get()
	if v != 0 {
		t.Fatalf("value must be 0 but %v", v)
	}

	err := n.Set("2")
	if err != nil {
		t.Fatal(err)
	}
	v = n.Get()
	if v != 2 {
		t.Fatalf("value must be 2 but %v", v)
	}
}

func TestIntFlag(t *testing.T) {
	n := NewNumber(0, 1, 4)

	fset := flag.NewFlagSet("test", flag.ContinueOnError)
	fset.Var(n, "number", "number between 1 to 4")
	err := fset.Parse([]string{"-number", "5"})
	if err == nil {
		t.Fatal("must be error")
	}

	err = fset.Parse([]string{"-number", "1"})
	if err != nil {
		t.Fatal(err)
	}

	v := n.Get()
	if v != 1 {
		t.Fatalf("value must be 1 but %v", v)
	}
}

func TestFloatFlag(t *testing.T) {
	n := NewNumber(0.0, 1.2, 3.4)

	fset := flag.NewFlagSet("test", flag.ContinueOnError)
	fset.Var(n, "number", "number between 1.2 to 3.4")
	err := fset.Parse([]string{"-number", "4.0"})
	if err == nil {
		t.Fatal("must be error")
	}

	err = fset.Parse([]string{"-number", "1.3"})
	if err != nil {
		t.Fatal(err)
	}

	v := n.Get()
	if v != 1.3 {
		t.Fatalf("value must be 1.3 but %v", v)
	}

}

func TestString(t *testing.T) {
	n := NewString("xx", 2, 4)
	v := n.Get()
	if v != "xx" {
		t.Fatalf("length must be 2 but %v", len(v))
	}

	err := n.Set("xxxxx")
	if err == nil {
		t.Fatal("must be error")
	}
}
