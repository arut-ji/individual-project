package duplicated_liveness_and_readiness

import (
	"github.com/arut-ji/individual-project/util"
	"testing"
)

func loadFixture(name string) string {
	return util.LoadFixture("./fixtures")(name)
}

func TestDuplicatedLivenessAndReadinessForNoSmell(t *testing.T) {
	script := loadFixture("no_smell.yaml")
	if result, err := hasDuplicatedLivenessAndReadiness(script); result != false {
		if err != nil {
			t.Error("Detection returned errors: ", err)
		}
		t.Errorf("Detection result was incorrect, got: %v, want: %v.", result, false)
	}
}

func TestDuplciatedLivenessAndReadinessForHasSmell(t *testing.T) {
	script := loadFixture("duplicated_liveness_and_readiness.yaml")
	if result, err := hasDuplicatedLivenessAndReadiness(script); result != true {
		if err != nil {
			t.Error("Detection returned errors: ", err)
		}
		t.Errorf("Detection result was incorrect, got: %v, want: %v.", result, true)
	}
}
