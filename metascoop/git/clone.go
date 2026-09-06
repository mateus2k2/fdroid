package git

import (
	"os"
	"os/exec"
)

func CloneRepo(gitUrl string) (dirPath string, err error) {
	dirPath, err = os.MkdirTemp("", "git-*")
	if err != nil {
		return
	}

	// Shallow, single-branch: the repo is cloned only to scrape screenshots
	// and an optional icon, so full history and all branches are wasted work
	// (some upstreams are hundreds of MB).
	cloneCmd := exec.Command("git", "clone", "--depth", "1", "--single-branch", "--no-tags", gitUrl, dirPath)
	err = cloneCmd.Run()
	if err != nil {
		_ = os.RemoveAll(dirPath)
		return
	}

	return
}
