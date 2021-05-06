package github

import (
	"bytes"
	"fmt"
)

const (
	brokerSection     = "Broker"
	gatewaySection    = "Gateway"
	javaClientSection = "Java Client"
	goClientSection   = "Go Client"
	miscSection       = "Misc"
)

type Section struct {
	sections map[string][]*Issue
}

func NewSection() *Section {
	return &Section{
		sections: make(map[string][]*Issue),
	}
}

func (s *Section) AddIssue(issue *Issue) *Section {
	if issue.HasBrokerLabel() {
		s.addIssueToSection(brokerSection, issue)
	}

	if issue.HasGatewayLabel() {
		s.addIssueToSection(gatewaySection, issue)
	}

	if issue.HasJavaClientLabel() {
		s.addIssueToSection(javaClientSection, issue)
	}

	if issue.HasGoClientLabel() {
		s.addIssueToSection(goClientSection, issue)
	}

	isMisc := !(issue.HasBrokerLabel() || issue.HasJavaClientLabel() || issue.HasGoClientLabel())
	if isMisc {
		s.addIssueToSection(miscSection, issue)
	}
	return s
}

func (s *Section) addIssueToSection(section string, issue *Issue) {
	s.sections[section] = append(s.sections[section], issue)
}

func (s *Section) GetBrokerIssues() []*Issue {
	return s.getIssues(brokerSection)
}

func (s *Section) GetGatewayIssues() []*Issue {
	return s.getIssues(gatewaySection)
}

func (s *Section) GetJavaClientIssues() []*Issue {
	return s.getIssues(javaClientSection)
}

func (s *Section) GetGoClientIssues() []*Issue {
	return s.getIssues(goClientSection)
}

func (s *Section) GetMiscIssues() []*Issue {
	return s.getIssues(miscSection)
}

func (s *Section) getIssues(section string) []*Issue {
	return s.sections[section]
}

func (s *Section) IsEmpty() bool {
	return len(s.sections) == 0
}

func (s *Section) String() string {
	var b bytes.Buffer

	b.WriteString(sectionToString(brokerSection, s.GetBrokerIssues()))
	b.WriteString(sectionToString(gatewaySection, s.GetGatewayIssues()))
	b.WriteString(sectionToString(javaClientSection, s.GetJavaClientIssues()))
	b.WriteString(sectionToString(goClientSection, s.GetGoClientIssues()))
	b.WriteString(sectionToString(miscSection, s.GetMiscIssues()))

	return b.String()
}

func sectionToString(title string, issues []*Issue) string {
	var b bytes.Buffer

	if issues != nil {
		b.WriteString(fmt.Sprintf("### %s\n", title))
		for _, issue := range issues {
			b.WriteString(fmt.Sprintf("* %s\n", issue.String()))
		}
	}

	return b.String()
}
