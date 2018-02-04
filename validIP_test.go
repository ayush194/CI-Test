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
		{".9.8.9.0", false},
		{"fdghghfhr", false},
		{"99.9.9.", false},
		{".....", false},
		{".gd.56.hf.65", false},
		{"08.8.43.44", false},
		{".45.43.54.3.", false},
		{"43.4..54.54", false},
		{"32.3.66.00", false},
	}
	for _, pair := range table {
		if pair.ans != isValidIP(pair.ip) {
			t.Errorf("Expected the ip %v to be %v but instead got %v!", pair.ip, pair.ans, isValidIP(pair.ip))
		}
	}
}
