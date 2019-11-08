package github

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChangelog_String(t *testing.T) {
	tests := map[string]struct {
		changelog *Changelog
		expected  string
	}{
		"Empty section": {changelog: NewChangelog("Test"), expected: "# Test\n"},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.changelog.String())
		})
	}
}
