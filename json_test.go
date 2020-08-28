package mangadex_test

import (
	"testing"

	"github.com/bake/mangadex"
)

func TestMaybeNumber(t *testing.T) {
	tt := []struct {
		in  []byte
		out string
	}{
		{[]byte(``), "0"},
		{[]byte(`""`), "0"},
		{[]byte(`0`), "0"},
		{[]byte(`"0"`), "0"},
		{[]byte(`1`), "1"},
		{[]byte(`"1"`), "1"},
		{[]byte(`3.14`), "3.14"},
		{[]byte(`"3.14"`), "3.14"},
	}
	for _, tc := range tt {
		var n mangadex.MaybeNumber
		if err := n.UnmarshalJSON(tc.in); err != nil {
			t.Fatal(err)
		}
		if n.String() != tc.out {
			t.Fatalf("expected %q to be %q, got %q", tc.in, tc.out, n)
		}
	}
}

func TestDynamicType(t *testing.T) {
	var v mangadex.DynamicType

	if err := v.UnmarshalJSON([]byte(`false`)); err != nil {
		t.Fatal(err)
	}

	if v.Value != false {
		t.Fatalf("expected boolean false, got %q", v.Value)
	}

	if err := v.UnmarshalJSON([]byte(`0`)); err != nil {
		t.Fatal(err)
	}

	if v.Value != 0 {
		t.Fatalf("expected int 0, got %q", v.Value)
	}
}
