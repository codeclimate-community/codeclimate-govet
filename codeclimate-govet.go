package main

import "github.com/codeclimate/cc-engine-go/engine"
import "strings"
import "os"
import "os/exec"
import "strconv"
import "sort"

func main() {
	rootPath := "/code/"
	analysisFiles, err := engine.GoFileWalk(rootPath)
	if err != nil {
		os.Exit(1)
	}

	config, err := engine.LoadConfig()
	if err != nil {
		os.Exit(1)
	}

	excludedFiles := []string{}
	if config["exclude_paths"] != nil {
		for _, file := range config["exclude_paths"].([]interface{}) {
			excludedFiles = append(excludedFiles, file.(string))
		}
		sort.Strings(excludedFiles)
	}

	for _, path := range analysisFiles {
		relativePath := strings.SplitAfter(path, rootPath)[1]
		i := sort.SearchStrings(excludedFiles, relativePath)
		if i < len(excludedFiles) && excludedFiles[i] == relativePath {
			continue
		}

		cmd := exec.Command("go", "tool", "vet", path)

		out, err := cmd.CombinedOutput()
		if err != nil && err.Error() != "exit status 1" {
			// engine.PrintWarning()
			return
		}

		lines := strings.Split(string(out), "\n")
		if len(lines) > 1 {
			for _, line := range lines[:len(lines)-1] {

				pieces := strings.Split(line, ":")

				if len(pieces) < 3 {
					// engine.PrintWarning()
					return
				} else {
					lineNo, err := strconv.Atoi(pieces[1])
					if err != nil {
						// engine.PrintWarning()
						return
					}

					issue := &engine.Issue{
						Type:              "issue",
						Check:             "GoVet/BugRisk",
						Description:       strings.Join(pieces[2:], ":"),
						RemediationPoints: 50000,
						Categories:        []string{"Bug Risk"},
						Location: &engine.Location{
							Path: strings.SplitAfter(path, rootPath)[1],
							Lines: &engine.LinesOnlyPosition{
								Begin: lineNo,
								End:   lineNo,
							},
						},
					}
					engine.PrintIssue(issue)
				}
			}
		}
	}

}
