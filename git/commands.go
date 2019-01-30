package git

import (
	"fmt"
	"os/exec"
)

// GetCommitMsg returns the short hash and commit message of the last commit.
func GetCommitMsg(gitPath string) (string, error) {

	gitCmd := exec.Command("git", "show", "-s", "--format=%h %s")
	gitCmd.Dir = gitPath
	output, err := gitCmd.Output()
	if err != nil {
		return "", fmt.Errorf("cannot get commit information for %s: %v", gitPath, err)
	}
	return string(output), nil
}
