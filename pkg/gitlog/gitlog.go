package gitlog

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strconv"
)

var (
	lineRegex = regexp.MustCompile(`(?im)^ *(closes?|related|relates|merges?|back\s?ports?|) +.*$`)
	idRegex   = regexp.MustCompile(`#(\d+)`)
)

func GetHistory(path, start, end string) string {
	logRange := fmt.Sprintf("%s..%s", start, end)

	// use git command til git lib implements range feature, see https://github.com/src-d/go-git/issues/1166
	command := exec.Command("git", "-C", path, "log", logRange, "--")
	log.Println(command)
	out, err := command.CombinedOutput()
	result := string(out)

	if err != nil {
		log.Fatal(result, err)
	}

	return result
}

func ExtractIssueIds(message string) []int {
	seen := map[int]bool{}
	var issueIds []int

	for _, line := range lineRegex.FindAllString(message, -1) {
		for _, match := range idRegex.FindAllStringSubmatch(line, -1) {
			issueId, err := strconv.Atoi(match[1])
			if err != nil {
				log.Fatalln("Cannot convert issue id", match[1], err)
			}

			if !seen[issueId] {
				seen[issueId] = true
				issueIds = append(issueIds, issueId)
			}
		}
	}

	return issueIds
}
