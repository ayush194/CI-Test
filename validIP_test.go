package validIP

import (
	"testing"
)

func TestIsValidIP(t *testing.T) {
	var table = []struct {
		ip  string
		ans bool
	}{
		{"123.89.11.2", true},
		{"008.008.008.008", false},
		{"0.6.54.0", true},
	}
	for _, pair := range table {
		if pair.ans != isValidIP(pair.ip) {
			t.Errorf("Expected the ip %v to be %v but instead got %v!", pair.ip, pair.ans, isValidIP(pair.ip))
		}
	}
}
