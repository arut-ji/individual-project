package smells_detector

import "testing"

const (
	AvoidCommentsFixtureDir = "./fixtures/avoid_comments"
)

func TestHasCommentsForNoSmell(t *testing.T) {
	script := loadFixture(AvoidCommentsFixtureDir)("no_smell.yaml")
	if result, err := hasComments(script); result != false || err != nil {
		t.Errorf("Detection result was incorrect, got: %v, want: %v.", result, false)
	}
}

func TestHasCommentsForContainingSmell(t *testing.T) {
	script := loadFixture(AvoidCommentsFixtureDir)("has_comments.yaml")
	if result, err := hasComments(script); result != true || err != nil {
		t.Errorf("Detection result was incorrect, got: %v, want: %v.", result, true)
	}
}
