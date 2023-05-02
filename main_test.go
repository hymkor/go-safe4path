package safe4path

import (
	"testing"
)

const (
	source = `a<b>c"d/e\f|g?h*i.j:k`
	expect = `a%3Cb%3Ec%22d%2Fe%5Cf%7Cg%3Fh%2Ai%2Ej%3Ak`
)

func TestAll(t *testing.T) {
	result := ToSafe(source, '%')
	if result != expect {
		t.Fatalf("ToSafe: `%s` != `%s`", result, expect)
	}

	result2, err := FromSafe(result, '%')
	if err != nil {
		t.Fatalf("FromSafe: %s", err.Error())
	}
	if result2 != source {
		t.Fatalf("FromSafe: `%s` != `%s`", result2, source)
	}
}
