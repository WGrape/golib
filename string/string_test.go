package string

import (
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestReverse(t *testing.T) {
	s := "hello world"
	if Reverse(s) != "dlrow olleh" {
		t.Fail()
		return
	}
}

func TestIsValidIdentifierName(t *testing.T) {
	if IsValidIdentifierName("") != false {
		t.Fail()
		return
	}
	if IsValidIdentifierName(" s ") != false {
		t.Fail()
		return
	}
	if IsValidIdentifierName("s(") != false {
		t.Fail()
		return
	}
	if IsValidIdentifierName("s2") != true {
		t.Fail()
		return
	}
	if IsValidIdentifierName("g") != true {
		t.Fail()
		return
	}
}

func TestIsCamelCase(t *testing.T) {
	var name = "isHello"
	if IsCamelCase(name) != true {
		t.Fail()
		return
	}

	name = "IsHello"
	if IsCamelCase(name) != false {
		t.Fail()
		return
	}

	name = "isNBA"
	if IsCamelCase(name) != false {
		t.Fail()
		return
	}

	name = "isNba"
	if IsCamelCase(name) != true {
		t.Fail()
		return
	}

	name = "hello world"
	if IsCamelCase(name) != false {
		t.Fail()
		return
	}

	name = "()({}{]s["
	if IsCamelCase(name) != false {
		t.Fail()
		return
	}
}

func TestIsLower(t *testing.T) {
	if IsLower(rune('p')) != true {
		t.Fail()
		return
	}

	if IsLower(rune('P')) != false {
		t.Fail()
		return
	}

	if IsLower(rune('(')) != false {
		t.Fail()
		return
	}

	if IsLower(rune('乐')) != false {
		t.Fail()
		return
	}
}

func TestIsUpper(t *testing.T) {
	if IsUpper(rune('p')) != false {
		t.Fail()
		return
	}

	if IsUpper(rune('P')) != true {
		t.Fail()
		return
	}

	if IsUpper(rune('(')) != false {
		t.Fail()
		return
	}

	if IsUpper(rune('乐')) != false {
		t.Fail()
		return
	}
}

func TestIsLetter(t *testing.T) {
	if IsLetter('2') != false {
		t.Fail()
		return
	}
	if IsLetter('s') != true {
		t.Fail()
		return
	}
}
func TestIsDigit(t *testing.T) {
	if IsDigit('2') != true {
		t.Fail()
		return
	}
	if IsDigit('s') != false {
		t.Fail()
		return
	}
}
