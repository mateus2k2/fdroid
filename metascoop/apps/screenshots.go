package apps

import (
	"io/fs"
	"path/filepath"
	"sort"
	"strings"
)

type RepoMetadata struct {
	Screenshots []string
}

var imageSuffixes = map[string]bool{
	"png":  true,
	"jpg":  true,
	"jpeg": true,
}

func hasImageSuffix(path string) bool {
	return imageSuffixes[strings.TrimPrefix(filepath.Ext(path), ".")]
}

// maxScreenshots caps how many screenshots are pulled from an upstream repo.
// Some projects (e.g. forks of Infinity for Reddit) vendor Fastlane
// screenshots for ~100 locales; without a cap metascoop copies thousands of
// images into the index and commits them to fdroid/repo/.
const maxScreenshots = 8

// FindMetadata locates phone screenshots in a cloned upstream repo. It only
// matches Fastlane-style "phoneScreenshots" directories (not every image with
// "screenshot" somewhere in its path), prefers the English locale when the
// repo ships localized sets, and returns at most maxScreenshots of them.
func FindMetadata(clonedRepoPath string) (r RepoMetadata, err error) {
	abs, err := filepath.Abs(clonedRepoPath)
	if err != nil {
		return
	}

	var all []string
	err = filepath.WalkDir(abs, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}

		lp := strings.ToLower(path)
		if strings.Contains(lp, "phonescreenshots") && hasImageSuffix(lp) {
			all = append(all, path)
		}

		return nil
	})
	if err != nil {
		return
	}

	var english []string
	for _, p := range all {
		lp := strings.ToLower(p)
		if strings.Contains(lp, "/en-us/") || strings.Contains(lp, "/en-rus/") ||
			strings.Contains(lp, "/en/") {
			english = append(english, p)
		}
	}
	if len(english) > 0 {
		all = english
	}

	sort.Strings(all)
	if len(all) > maxScreenshots {
		all = all[:maxScreenshots]
	}

	r.Screenshots = all
	return
}
