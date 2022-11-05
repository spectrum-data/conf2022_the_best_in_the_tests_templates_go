package conf2022_the_best_in_the_tests_templates_go

import (
	"github.com/spectrum-data/conf2022_the_best_in_the_tests_templates_go/test"
	"os"
	"path/filepath"
	"testing"
)

func Test_local(t *testing.T) {
	currentDir, _ := os.Getwd()

	files := []test.TestDescFile{
		{Path: filepath.Join(currentDir, "base.csv"), Type: test.BASE},
		{Path: filepath.Join(currentDir, "local.csv"), Type: test.LOCAL},
	}

	testBase := test.TestBase{TestFiles: files}

	testBase.Run(t)
}
