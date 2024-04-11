package execwrap

import "strings"

func (ce *CmdExecutor) List(name string, args ...string) ([]string, error) {
	cmd := ce.Command(name, args...)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	outputStr := string(output)
	list := strings.Split(outputStr, "\n")
	return list, nil
}
