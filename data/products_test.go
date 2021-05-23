package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "test",
		Price: 12,
		SKU:   "abc-abc-abs",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
