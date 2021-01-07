package incomplete_tasks

import (
	"github.com/arut-ji/individual-project/linter/smells_detector"
	"testing"
)

func loadFixture(name string) string {
	return smells_detector.LoadFixture("./fixtures")(name)
}

func TestInCompleteTasks(t *testing.T) {
	script := loadFixture("contain_incomplete_tasks.yaml")
	if result, err := inCompleteTasks(script); result != true || err != nil {
		t.Errorf("Detection result was incorrect, got: %v, want: %v.", result, true)
	}
}

func TestInCompleteTaskForTODO(t *testing.T) {
	script := loadFixture("contain_TODO.yaml")
	if result, err := inCompleteTasks(script); result != true || err != nil {
		t.Errorf("Detection result was incorrect, got: %v, want: %v.", result, true)
	}
}

func TestInCompleteTaskForFIXME(t *testing.T) {
	script := loadFixture("contain_FIXME.yaml")
	if result, err := inCompleteTasks(script); result != true || err != nil {
		t.Errorf("Detection result was incorrect, got: %v, want: %v.", result, true)
	}
}

func TestInCompleteTaskForNoSmells(t *testing.T) {
	script := loadFixture("no_smell.yaml")
	if result, err := inCompleteTasks(script); result != false || err != nil {
		t.Errorf("Detection result was incorrect, got: %v, want: %v.", result, false)
	}
}
