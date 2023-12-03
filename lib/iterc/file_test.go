package iterc

import (
	"os"
	"testing"
)

func TestMustReadLines(t *testing.T) {
	f, err := os.CreateTemp("", "iterc")
	if err != nil {
		t.Error()
	}
	_, _ = f.WriteString("line1\n")
	_, _ = f.WriteString("line2\n")
	f.Close()

	l := MustReadLines(f.Name()).List()
	if len(l) != 2 {
		t.Error(l)
	}
	if l[0] != "line1" || l[1] != "line2" {
		t.Error(l)
	}

	err = os.Remove(f.Name())
	if err != nil {
		t.Error(err)
	}
}

func TestMustReadParagraphs(t *testing.T) {
	f, err := os.CreateTemp("", "iterc")
	if err != nil {
		t.Error()
	}
	_, _ = f.WriteString("line1\n")
	_, _ = f.WriteString("line2\n")
	_, _ = f.WriteString("\n")
	_, _ = f.WriteString("line3\n")
	_, _ = f.WriteString("\n")
	_, _ = f.WriteString("\n")
	_, _ = f.WriteString("line4\n")
	f.Close()

	p := MustReadParagraphs(f.Name()).List()
	if len(p) != 3 {
		t.Error(p)
	}
	if len(p[0]) != 2 || len(p[1]) != 1 || len(p[2]) != 1 {
		t.Error(p)
	}
	if p[0][0] != "line1" || p[0][1] != "line2" || p[1][0] != "line3" || p[2][0] != "line4" {
		t.Error(p)
	}

	err = os.Remove(f.Name())
	if err != nil {
		t.Error(err)
	}
}
