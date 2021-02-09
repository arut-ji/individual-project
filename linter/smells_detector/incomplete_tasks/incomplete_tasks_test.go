package incomplete_tasks

import (
	"github.com/arut-ji/individual-project/util"
	"testing"
)

func loadFixture(name string) string {
	return util.LoadFixture("./fixtures")(name)
}

func TestForMixedFiveInstances(t *testing.T) {
	script := loadFixture("contain_incomplete_tasks.yaml")
	if result, err := countInCompleteTasks(script); result != 5 || err != nil {
		t.Errorf("Detection result was incorrect, got: %v, want: %v.", result, 5)
	}
}

func TestForOneTODOInstance(t *testing.T) {
	script := loadFixture("contain_TODO.yaml")
	if result, err := countInCompleteTasks(script); result != 1 || err != nil {
		t.Errorf("Detection result was incorrect, got: %v, want: %v.", result, 1)
	}
}

func TestForOneFIXMEInstance(t *testing.T) {
	script := loadFixture("contain_FIXME.yaml")
	if result, err := countInCompleteTasks(script); result != 1 || err != nil {
		t.Errorf("Detection result was incorrect, got: %v, want: %v.", result, 1)
	}
}

func TestInCompleteTaskForNoSmells(t *testing.T) {
	script := loadFixture("no_smell.yaml")
	if result, err := countInCompleteTasks(script); result != 0 || err != nil {
		t.Errorf("Detection result was incorrect, got: %v, want: %v.", result, 0)
	}
}
