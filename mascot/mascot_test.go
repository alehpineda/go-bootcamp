package mascot_test

import (
	"testing"

	"github.com/alehpineda/go-bootcamp/mascot"
)

func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "Gopher" {
		t.Fatal("That's not the best mascot :(")
	}
}