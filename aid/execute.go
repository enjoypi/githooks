package aid

import (
	"bytes"
	"os/exec"
	"strings"
)

func ExecuteToSlice(command string, arg ...string) ([]string, error) {
	cmd := exec.Command(command, arg...)
	cmd.Stdin = strings.NewReader(command)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	if out.Len() <= 0 {
		return nil, nil
	}

	ss := make([]string, 0)

	for _, s := range strings.Split(out.String(), "\n") {
		if len(s) > 0 {
			ss = append(ss, s)
		}
	}

	if len(ss) > 0 {
		return ss, nil
	}

	return nil, nil
}
