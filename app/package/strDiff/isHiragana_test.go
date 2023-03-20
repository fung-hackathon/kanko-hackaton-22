package strDiff

import "testing"

func TestIsHiragana(t *testing.T) {
	if get := IsHiragana("函館"); get {
		t.Errorf("failed: expects\"%t\" but output was \"%t\"", false, get)
	}
	if get := IsHiragana("はこだて未来大"); get {
		t.Errorf("failed: expects\"%t\" but output was \"%t\"", false, get)
	}
	if get := IsHiragana("はこだて"); !get {
		t.Errorf("failed: expects\"%t\" but output was \"%t\"", true, get)
	}
}
