package demo

import "fmt"

// TitleToken returns the task marker that must appear in the PR title.
func TitleToken(taskID string) string {
	return fmt.Sprintf("[task:%s]", taskID)
}

// PRTitle returns a PR title that satisfies the demo task contract.
func PRTitle(taskID, summary string) string {
	if summary == "" {
		return TitleToken(taskID)
	}
	return fmt.Sprintf("%s %s", TitleToken(taskID), summary)
}

// BranchName returns the expected feature branch for a task.
func BranchName(taskID string) string {
	return fmt.Sprintf("task/%s", taskID)
}
