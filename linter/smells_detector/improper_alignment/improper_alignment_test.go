package improper_alignment

import (
	"github.com/arut-ji/individual-project/util"
	"testing"
)

func loadFixture(name string) string {
	return util.LoadFixture("./fixtures")(name)
}

func TestImproperAlignmentWithNoSmell(t *testing.T) {
	script := loadFixture("no_smell.yaml")
	if result, err := countSmellInstances(script); result != 0 || err != nil {
		t.Errorf("Detection result was incorrect, got: %v, want: %v.", result, 0)
	}
}

func TestImproperAlignmentWithInvalidIndentation(t *testing.T) {
	script := loadFixture("invalid_indentation.yaml")
	if result, err := countSmellInstances(script); result != 1 {
		if err != nil {
			t.Error("Detection returned error", err)
		}
		t.Errorf("Detection result was incorrect, got: %v, want: %v.", result, 1)
	}
}
