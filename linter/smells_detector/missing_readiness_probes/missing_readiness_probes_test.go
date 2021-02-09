package missing_readiness_probes

import (
	"github.com/arut-ji/individual-project/util"
	"testing"
)

func loadFixture(name string) string {
	return util.LoadFixture("./fixtures")(name)
}

func TestForNoSmell(t *testing.T) {
	script := loadFixture("no_smell.yaml")
	if result, err := hasMissingReadinessProbes(script); result != false {
		if err != nil {
			t.Error("Detection returned errors: ", err)
		}
		t.Errorf("Detection result was incorrect, got: %v, want: %v.", result, false)
	}
}

func TestForOneSmell(t *testing.T) {
	script := loadFixture("one_smell.yaml")
	if result, err := hasMissingReadinessProbes(script); result != true {
		if err != nil {
			t.Error("Detection returned errors: ", err)
		}
		t.Errorf("Detection result was incorrect, got: %v, want: %v.", result, true)
	}
}
