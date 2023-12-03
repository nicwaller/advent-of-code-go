package iterc

import (
	"fmt"
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

	for line := range MustReadLines(f.Name()).C {
		fmt.Println(line)
	}

	err = os.Remove(f.Name())
	if err != nil {
		t.Error(err)
	}
}
