package missing_readiness_probes

import (
	"github.com/arut-ji/individual-project/util"
	"testing"
)

func loadFixture(name string) string {
	return util.LoadFixture("./fixtures")(name)
}

func TestForNoInstance(t *testing.T) {
	script := loadFixture("no_smell.yaml")
	if result, err := countMissingReadinessProbes(script); result != 0 {
		if err != nil {
			t.Error("Detection returned errors: ", err)
		}
		t.Errorf("Detection result was incorrect, got: %v, want: %v.", result, 0)
	}
}

func TestForOneInstance(t *testing.T) {
	script := loadFixture("one_smell.yaml")
	if result, err := countMissingReadinessProbes(script); result != 1 {
		if err != nil {
			t.Error("Detection returned errors: ", err)
		}
		t.Errorf("Detection result was incorrect, got: %v, want: %v.", result, 1)
	}
}
