package mailcheck

import (
	"fmt"
	"github.com/markbates/pkger"
	"github.com/thoas/go-funk"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func walkUsingFilePath() []string {
	var files []string
	root := "resources"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(files); i++ {
		files[i] = filepath.ToSlash(files[i])
	}
	return files
}

func walkUsingPkger() []string {
	var files []string
	root, err := pkger.Open("/resources")
	if err != nil {
		panic(err)
	}
	err = pkger.Walk(root.Path().String(), func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}

func TestPkgedData(t *testing.T) {
	files := walkUsingFilePath()
	fmt.Println("----")
	pkgFiles := walkUsingPkger()
	for i := 0; i < len(pkgFiles); i++ {
		pkgFiles[i] = strings.TrimPrefix(pkgFiles[i], "github.com/deepakputhraya/mailcheck:/")
	}
	l, r := funk.DifferenceString(files, pkgFiles)
	if len(l) > 0 || len(r) > 0 {
		t.Fail()
	}
}
