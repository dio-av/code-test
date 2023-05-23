package parser

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParse(t *testing.T) {
	t.Parallel()

	data, err := os.ReadFile("./testData.json")
	if err != nil {
		t.Fatal("err")
	}

	var p PlanData
	err = json.Unmarshal(data, &p)
	if err != nil {
		t.Fatal(err)
	}
	got := p.Plans[0].Name
	want := "Família 50GB"
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
