package main

import (
	"github.com/codeclimate/cc-engine-go/engine"
	"strings"
	"os"
	"os/exec"
	"strconv"
	"fmt"
)

func main() {
	rootPath := "/code/"

	config, err := engine.LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
		os.Exit(1)
	}

	analysisFiles, err := engine.GoFileWalk(rootPath, engine.IncludePaths(rootPath, config))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing: %v\n", err)
		os.Exit(1)
	}

	for _, path := range analysisFiles {
		cmd := exec.Command("go", "tool", "vet", path)

		out, err := cmd.CombinedOutput()

		if err != nil && err.Error() != "exit status 1" {
			fmt.Fprintf(os.Stderr, "Error analyzing path: %v\n", path)
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)

			if out != nil {
				s := string(out[:])
				fmt.Fprintf(os.Stderr, "Govet output: %v\n", s)
			}

			return
		}

		lines := strings.Split(string(out), "\n")
		if len(lines) > 1 {
			for _, line := range lines[:len(lines)-1] {

				pieces := strings.Split(line, ":")

				if len(pieces) < 3 {
					fmt.Fprintf(os.Stderr, "Unexpected format for the following output: %v\n", line)
					return
				} else {
					lineNo, err := strconv.Atoi(pieces[1])
					if err != nil {
						fmt.Fprintf(os.Stderr, "Unexpected format for the following output: %v\n", line)
						fmt.Fprintf(os.Stderr, "\nError: %v\n", err)
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
