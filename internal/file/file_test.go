package file_test

import (
	"testing"

	"github.com/gustavocd/gotask/internal/file"
)

func TestRead(t *testing.T) {
	_, err := file.ReadDeviceInputs("../../challenge.in")
	if err != nil {
		t.Fatal(err)
	}
}
