package main

import "github.com/codeclimate/cc-engine-go/engine"
import "strings"
import "os"
import "os/exec"
import "strconv"

func main() {
	rootPath := "/code/"

	config, err := engine.LoadConfig()
	if err != nil {
		os.Exit(1)
	}

	analysisFiles, err := engine.GoFileWalk(rootPath, engine.IncludePaths(rootPath, config))
	if err != nil {
		os.Exit(1)
	}

	for _, path := range analysisFiles {
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
