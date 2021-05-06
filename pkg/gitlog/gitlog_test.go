package gitlog

import "testing"
import "github.com/stretchr/testify/assert"

func TestGitHistory(t *testing.T) {
	tests := map[string]struct {
		path  string
		start string
		end   string
		size  int
	}{
		"First commit": {path: ".", start: "7b86247", end: "7ab8381", size: 177},
		"Between tags": {path: ".", start: "0.1.0", end: "0.2.0", size: 207},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			log := GetHistory(tc.path, tc.start, tc.end)
			assert.Equal(t, tc.size, len(log))
		})
	}
}

func TestExtractIssueIds(t *testing.T) {
	tests := map[string]struct {
		message  string
		issueIds []int
	}{
		"No issue id":                 {message: "hello world", issueIds: nil},
		"No keyword":                  {message: "#1234", issueIds: nil},
		"Close keyword":               {message: "close #1234", issueIds: []int{1234}},
		"Closes keyword":              {message: "closes #1234", issueIds: []int{1234}},
		"Related keyword":             {message: "related #1234", issueIds: []int{1234}},
		"Merge keyword":               {message: "merge #1234", issueIds: []int{1234}},
		"Merges keyword":              {message: "merges #1234", issueIds: []int{1234}},
		"Relates keyword":             {message: "relates to #1234", issueIds: []int{1234}},
		"Backport keyword":            {message: "backport #1234", issueIds: []int{1234}},
		"Backports keyword":           {message: "backports #1234", issueIds: []int{1234}},
		"Back ports keyword":          {message: "back ports #1234", issueIds: []int{1234}},
		"Keyword uppercase":           {message: "Closes #1234", issueIds: []int{1234}},
		"Spacing in front of keyword": {message: "  \t closes #1234", issueIds: []int{1234}},
		"Multiple issues":             {message: "closes #1234, #5678, #9 and #123", issueIds: []int{1234, 5678, 9, 123}},
		"Duplicate issue ids":         {message: "closes #123, #234, #123 and #23", issueIds: []int{123, 234, 23}},
		"Multiple lines":              {message: "foo bar\n\ncloses #1234\ntest", issueIds: []int{1234}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			issueIds := ExtractIssueIds(tc.message)
			assert.Equal(t, tc.issueIds, issueIds)
		})
	}
}
