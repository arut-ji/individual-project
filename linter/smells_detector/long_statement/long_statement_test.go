package long_statement

import (
	"fmt"
	"math/rand"
	"testing"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func TestForNoInstance(t *testing.T) {
	script := randSeq(10)
	if result, err := countLongStatement(script); result != 0 {
		if err != nil {
			t.Error("Detection returned errors: ", err)
		}
		t.Errorf("Detection result was incorrect, got: %v, want: %v.", result, 0)
	}
}

func TestForOneInstance(t *testing.T) {
	script := fmt.Sprintf("%v\n%v\n", randSeq(101), randSeq(10))
	if result, err := countLongStatement(script); result != 1 {
		if err != nil {
			t.Error("Detection returned errors: ", err)
		}
		t.Errorf("Detection result was incorrect, got: %v, want: %v.", result, 1)
	}
}
