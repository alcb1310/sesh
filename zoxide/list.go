package zoxide

import (
	"fmt"
	"os"
	"strings"

	"github.com/joshmedeski/sesh/convert"
	"github.com/joshmedeski/sesh/models"
)

func List() ([]*models.ZoxideResult, error) {
	output, err := zoxideCmd([]string{"query", "-ls"})
	if err != nil {
		return []*models.ZoxideResult{}, nil
	}
	cleanOutput := strings.TrimSpace(string(output))
	list := strings.Split(cleanOutput, "\n")
	listLen := len(list)
	if listLen == 1 && list[0] == "" {
		return []*models.ZoxideResult{}, nil
	}

	results := make([]*models.ZoxideResult, 0, listLen)
	for _, line := range list {
		trimmed := strings.Trim(line, "[]")
		trimmed = strings.Trim(trimmed, " ")
		fields := strings.SplitN(trimmed, " ", 2)
		if len(fields) != 2 {
			fmt.Println("Zoxide entry has invalid number of fields (expected 2)")
			os.Exit(1)
		}
		path := fields[1]
		results = append(results, &models.ZoxideResult{
			Score: convert.StringToFloat(fields[0]),
			Name:  convert.PathToPretty(path),
			Path:  path,
		})
	}

	return results, nil
}
