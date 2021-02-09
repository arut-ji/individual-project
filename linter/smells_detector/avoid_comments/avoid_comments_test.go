package avoid_comments

import (
	"github.com/arut-ji/individual-project/util"
	"regexp"
	"testing"
)

func loadFixture(name string) string {
	return util.LoadFixture("./fixtures")(name)
}

func TestRegexPattern(t *testing.T) {
	line := "# A comment here"
	match, err := regexp.MatchString("#.*", line)
	if err != nil {
		t.Error("Pattern matching contains errors.", err)
	}
	if match != true {
		t.Errorf("Matching result was incorrect, got: %v, want: %v.", match, true)
	}
}

func TestForZeroInstances(t *testing.T) {
	script := loadFixture("no_smell.yaml")
	if result, err := countComments(script); result != 0 || err != nil {
		if err != nil {
			t.Error("Detection contains error: ", err)
		}
		t.Errorf("Detection result was incorrect, got: %v, want: %v.", result, 0)
	}
}

func TestForThreeInstances(t *testing.T) {
	script := loadFixture("has_comments.yaml")
	if result, err := countComments(script); result != 3 || err != nil {
		t.Errorf("Detection result was incorrect, got: %v, want: %v.", result, 3)
	}
}
