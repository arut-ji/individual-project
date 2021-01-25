package avoid_comments

import (
	"github.com/arut-ji/individual-project/util"
	"testing"
)

func loadFixture(name string) string {
	return util.LoadFixture("./fixtures")(name)
}

func TestHasCommentsForNoSmell(t *testing.T) {
	script := loadFixture("no_smell.yaml")
	if result, err := hasComments(script); result != false || err != nil {
		t.Errorf("Detection result was incorrect, got: %v, want: %v.", result, false)
	}
}

func TestHasCommentsForContainingSmell(t *testing.T) {
	script := loadFixture("has_comments.yaml")
	if result, err := hasComments(script); result != true || err != nil {
		t.Errorf("Detection result was incorrect, got: %v, want: %v.", result, true)
	}
}
