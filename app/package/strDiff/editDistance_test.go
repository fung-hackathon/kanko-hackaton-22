package strDiff

import "testing"

func TestEditDistance(t *testing.T) {
	if get := EditDistance("kitten", "sitting"); get != 3 {
		t.Errorf("failed: expects\"%d\" but output was \"%d\"", 3, get)
	}
}
