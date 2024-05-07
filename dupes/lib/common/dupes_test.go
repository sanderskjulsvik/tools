package common_test

import (
	"testing"

	common "github.com/sander-skjulsvik/tools/dupes/lib/common"
)

func getDupesA() *common.Dupes {
	a := common.NewDupes()
	a.D = map[string]*common.Dupe{
		"1": {
			Hash:  "1",
			Paths: []string{"1", "a/1", "a/a/1"},
		},
		"2": {
			Hash:  "2",
			Paths: []string{"2", "a/2", "b/a/2"},
		},
		"3": {
			Hash:  "3",
			Paths: []string{"3", "b/2", "b/a/2"},
		},
	}
	return &a
}

func getDupesB() *common.Dupes {
	b := common.NewDupes()
	b.D = map[string]*common.Dupe{
		"1": {
			Hash:  "1",
			Paths: []string{"1", "a/1", "a/a/1"},
		},
		"2": {
			Hash:  "2",
			Paths: []string{"2", "a/2", "b/a/2"},
		},
		"4": {
			Hash:  "4",
			Paths: []string{"4", "b/2", "b/a/2"},
		},
	}
	return &b
}

func TestOnlyInboth(t *testing.T) {
	a := getDupesA()
	b := getDupesB()
	both := a.OnlyInBoth(b)
	if len(both.D) != 2 {
		t.Errorf("Expected 2 dupes, got %d", len(both.D))
	}
	if _, ok := both.D["1"]; !ok {
		t.Errorf("Expected dupe 1, but it was not found")
	}
	if _, ok := both.D["2"]; !ok {
		t.Errorf("Expected dupe 2, but it was not found")
	}
	if _, ok := both.D["3"]; ok {
		t.Errorf("Did not expect dupe 3, but it was found")
	}
	if _, ok := both.D["4"]; ok {
		t.Errorf("Did not expect dupe 4, but it was found")
	}
}

func TestOnlyInSelf(t *testing.T) {
	a := getDupesA()
	b := getDupesB()
	onlyInA := a.OnlyInSelf(b)
	if len(onlyInA.D) != 1 {
		t.Errorf("Expected 1 dupe, got %d", len(onlyInA.D))
	}
	if _, ok := onlyInA.D["3"]; !ok {
		t.Errorf("Expected dupe 3, but it was not found")
	}
	if _, ok := onlyInA.D["4"]; ok {
		t.Errorf("Did not expect dupe 4, but it was found")
	}
}

func TestHasSameFiles(t *testing.T) {
	a := getDupesA()
	b := getDupesB()
	a2 := getDupesA()
	if a.HasSameFiles(b) {
		t.Errorf("Expected a and b to not have the same files")
	}
	if b.HasSameFiles(a) {
		t.Errorf("Expected b and a to not have the same files")
	}
	if !a.HasSameFiles(a2) {
		t.Errorf("Expected a and a2 to have the same files")
	}
	a.AppendDupes(b)
	if a.HasSameFiles(b) {
		t.Errorf("Expected a and b to not be have tee s ame files")
	}
	if !a.HasSameFiles(a) {
		t.Errorf("Expected a and a to have the same files")
	}
	if b.HasSameFiles(a) {
		t.Errorf("Expected b and a to not have the same files")
	}
}
