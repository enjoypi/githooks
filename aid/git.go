package aid

import (
	"strings"

	git "github.com/go-git/go-git/v5"
)

func GitStatus() (staged []string, untracked []string, err error) {
	lines, err := ExecuteToSlice("git", "status", "--short")
	if err != nil {
		return nil, nil, err
	}

	if len(lines) <= 0 {
		return nil, nil, nil
	}

	staged = make([]string, 0)
	untracked = make([]string, 0)
	for _, line := range lines {
		b := []byte(line)
		statusCode := git.StatusCode(b[0])
		switch statusCode {
		case git.Untracked:
			untracked = append(untracked, string(b[3:]))
		case git.Modified:
			fallthrough
		case git.Added:
			fallthrough
		case git.Deleted:
			staged = append(staged, string(b[3:]))
		case git.Renamed:
			fallthrough
		case git.Copied:
			content := strings.Split(string(b[3:]), " ")
			orig := content[0]
			staged = append(staged, orig)
		}
	}
	return
}
