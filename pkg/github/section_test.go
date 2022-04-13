package github

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSection_GetBrokerIssues(t *testing.T) {
	tests := map[string]struct {
		section *Section
		size    int
	}{
		"No Issues":       {section: NewSection(), size: 0},
		"Different Issue": {section: NewSection().AddIssue(createIssueWithLabel("unknown")), size: 0},
		"One Issue":       {section: NewSection().AddIssue(createIssueWithLabel(brokerLabel)), size: 1},
		"Multiple Issues": {
			section: NewSection().
				AddIssue(createIssueWithLabel(brokerLabel, featureLabel)).
				AddIssue(createIssueWithLabel(bugLabel, brokerLabel)).
				AddIssue(createIssueWithLabel(javaClientLabel)),
			size: 2,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.size, len(tc.section.GetBrokerIssues()))
		})
	}
}

func TestSection_GetGatewayIssues(t *testing.T) {
	tests := map[string]struct {
		section *Section
		size    int
	}{
		"No Issues":       {section: NewSection(), size: 0},
		"Different Issue": {section: NewSection().AddIssue(createIssueWithLabel("unknown")), size: 0},
		"One Issue":       {section: NewSection().AddIssue(createIssueWithLabel(gatewayLabel)), size: 1},
		"Multiple Issues": {
			section: NewSection().
				AddIssue(createIssueWithLabel(gatewayLabel, featureLabel)).
				AddIssue(createIssueWithLabel(bugLabel, gatewayLabel)).
				AddIssue(createIssueWithLabel(javaClientLabel)),
			size: 2,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.size, len(tc.section.GetGatewayIssues()))
		})
	}
}

func TestSection_GetJavaClientIssues(t *testing.T) {
	tests := map[string]struct {
		section *Section
		size    int
	}{
		"No Issues":       {section: NewSection(), size: 0},
		"Different Issue": {section: NewSection().AddIssue(createIssueWithLabel("unknown")), size: 0},
		"One Issue":       {section: NewSection().AddIssue(createIssueWithLabel(javaClientLabel)), size: 1},
		"Multiple Issues": {
			section: NewSection().
				AddIssue(createIssueWithLabel(javaClientLabel, featureLabel)).
				AddIssue(createIssueWithLabel(bugLabel, javaClientLabel)).
				AddIssue(createIssueWithLabel(brokerLabel)),
			size: 2,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.size, len(tc.section.GetJavaClientIssues()))
		})
	}
}

func TestSection_GetGoClientIssues(t *testing.T) {
	tests := map[string]struct {
		section *Section
		size    int
	}{
		"No Issues":       {section: NewSection(), size: 0},
		"Different Issue": {section: NewSection().AddIssue(createIssueWithLabel("unknown")), size: 0},
		"One Issue":       {section: NewSection().AddIssue(createIssueWithLabel(goClientLabel)), size: 1},
		"Multiple Issues": {
			section: NewSection().
				AddIssue(createIssueWithLabel(goClientLabel, featureLabel)).
				AddIssue(createIssueWithLabel(bugLabel, goClientLabel)).
				AddIssue(createIssueWithLabel(brokerLabel)),
			size: 2,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.size, len(tc.section.GetGoClientIssues()))
		})
	}
}

func TestSection_GetMiscIssues(t *testing.T) {
	tests := map[string]struct {
		section *Section
		size    int
	}{
		"No Issues":       {section: NewSection(), size: 0},
		"Different Issue": {section: NewSection().AddIssue(createIssueWithLabel(brokerLabel)), size: 0},
		"One Issue":       {section: NewSection().AddIssue(createIssueWithLabel()), size: 1},
		"Multiple Issues": {
			section: NewSection().
				AddIssue(createIssueWithLabel(featureLabel)).
				AddIssue(createIssueWithLabel(bugLabel)).
				AddIssue(createIssueWithLabel(brokerLabel)),
			size: 2,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.size, len(tc.section.GetMiscIssues()))
		})
	}
}

func TestSection_String(t *testing.T) {
	tests := map[string]struct {
		section  *Section
		expected string
	}{
		"Empty section": {section: NewSection(), expected: ""},
		"Broker section": {
			section:  NewSection().AddIssue(createIssueWithLabel(brokerLabel)),
			expected: createSectionString(brokerSection, 1)},
		"Java Client Section": {
			section:  NewSection().AddIssue(createIssueWithLabel(javaClientLabel)),
			expected: createSectionString(javaClientSection, 1)},
		"Go Client Section": {
			section:  NewSection().AddIssue(createIssueWithLabel(goClientLabel)),
			expected: createSectionString(goClientSection, 1)},
		"Misc Section": {
			section:  NewSection().AddIssue(createIssueWithLabel(bugLabel)),
			expected: createSectionString(miscSection, 1)},
		"All Sections": {
			section: NewSection().
				AddIssue(createIssueWithLabel(goClientLabel)).
				AddIssue(createIssueWithLabel(bugLabel)).
				AddIssue(createIssueWithLabel(javaClientLabel)).
				AddIssue(createIssueWithLabel(brokerLabel)),
			expected: createSectionString(brokerSection, 1) +
				createSectionString(javaClientSection, 1) +
				createSectionString(goClientSection, 1) +
				createSectionString(miscSection, 1)},
		"Multiple Issues": {
			section: NewSection().
				AddIssue(createIssueWithLabel(goClientLabel, brokerLabel)).
				AddIssue(createIssueWithLabel(bugLabel)).
				AddIssue(createIssueWithLabel(javaClientLabel, goClientLabel, brokerLabel)).
				AddIssue(createIssueWithLabel(docsLabel)).
				AddIssue(createIssueWithLabel(brokerLabel)),
			expected: createSectionString(brokerSection, 3) +
				createSectionString(javaClientSection, 1) +
				createSectionString(goClientSection, 2) +
				createSectionString(miscSection, 2)},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.section.String())
		})
	}
}

func createIssueWithLabel(labels ...string) *Issue {
	return createIssue("test", 123, "test", false, labels...)
}

func createSectionString(title string, size int) string {
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("### %s\n", title))
	for i := 0; i < size; i++ {
		b.WriteString("* test ([#123](test))\n")
	}
	return b.String()
}
