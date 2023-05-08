package main

import (
	"os"
	"log"
	"regexp"
	"path/filepath"

	"github.com/bitfield/script"
)

const (
	Product = "mmshare"
	InitPyFile = "__init__.py"
)

var (
	SCHRODINGER string = os.Getenv("SCHRODINGER")
	SCHRODINGER_SRC string = os.Getenv("SCHRODINGER_SRC")
)


func main() {
	testDirGlob := filepath.Join(SCHRODINGER, Product + "*", "python", "test")
	paths, err := filepath.Glob(testDirGlob)
	if len(paths) != 1 || err != nil {
		log.Fatal("did not find exactly 1 path to " + testDirGlob)
	}
	testDir := paths[0]

	os.Chdir(testDir)
	pyFile := regexp.MustCompile(`.*\.py$`)

	script.FindFiles(".").
		   MatchRegexp(pyFile).
		   Reject(InitPyFile).
		   Exec("fzf").
		   Stdout()
}

