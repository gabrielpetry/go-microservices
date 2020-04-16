package data

import (
	"testing"
)

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "nics",
		Price: 1.2,
		SKU:   "aaa-bbb",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
