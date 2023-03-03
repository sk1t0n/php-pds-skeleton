package creator

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"testing"
)

const projectName = "test_project"

var creator = &Creator{ProjectName: projectName}
var runTest = createRunTest(setUp, tearDown)

func setUp() {
	creator.setInnerDirs()
	creator.createDirs()
}

func tearDown() {
	cmd := exec.Command("rm", "-rf", projectName)
	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

func createRunTest(setUp func(), tearDown func()) func(func()) {
	return func(testFunc func()) {
		setUp()
		testFunc()
		tearDown()
	}
}

func TestNew(t *testing.T) {
	runTest(func() {
		New()
	})
}

func TestCheckArgsWhenNoError(t *testing.T) {
	runTest(func() {
		os.Args = []string{"creator_test", projectName}
		err := creator.checkArgs()
		if err != nil {
			t.Error(err)
		}
	})
}

func TestCheckArgsWhenError(t *testing.T) {
	runTest(func() {
		os.Args = []string{"creator_test"}
		err := creator.checkArgs()
		expectedErrorMessage := "checkArgs: the project name is not set"
		if err.Error() != expectedErrorMessage {
			t.Error("the error message does not match the expected error message")
		}
		if err == nil {
			t.Error(errors.New("the error should not be nil"))
		}
	})
}

func TestCreateProjectStructure(t *testing.T) {
	runTest(func() {
		creator.CreateProjectStructure()
	})
}

func TestCreatePublicIndexPhp(t *testing.T) {
	runTest(func() {
		creator.createPublicIndexPhp()
		file := fmt.Sprintf("%s/public/index.php", projectName)
		indexPhp, err := os.ReadFile(file)
		expectedContent := "<?php\n"
		if err != nil {
			t.Error(err)
		} else if string(indexPhp) != expectedContent {
			t.Error("the content of the file does not match the expected content")
		}
	})
}

func TestCreateConfigAppPhp(t *testing.T) {
	runTest(func() {
		creator.createConfigAppPhp()
		file := fmt.Sprintf("%s/config/app.php", projectName)
		appPhp, err := os.ReadFile(file)
		expectedContent := "<?php\n"
		if err != nil {
			t.Error(err)
		} else if string(appPhp) != expectedContent {
			t.Error("the content of the file does not match the expected content")
		}
	})
}

func TestCreateReadmeMd(t *testing.T) {
	runTest(func() {
		creator.createReadmeMd()
		file := fmt.Sprintf("%s/README.md", projectName)
		readmeMd, err := os.ReadFile(file)
		expectedContent := `# %s

[Source Link](https://github.com/php-pds/skeleton#root-level-directories)

## dir

If the package provides a root-level directory for command-line executable files, it MUST be named bin/.  
This publication does not otherwise define the structure and contents of the directory.

## config

If the package provides a root-level directory for configuration files, it MUST be named config/.  
This publication does not otherwise define the structure and contents of the directory.

## public

If the package provides a root-level directory for web server files, it MUST be named public/.  
This publication does not otherwise define the structure and contents of the directory.

## resources

If the package provides a root-level directory for other resource files, it MUST be named resources/.  
This publication does not otherwise define the structure and contents of the directory.

## src

If the package provides a root-level directory for PHP source code files, it MUST be named src/.  
This publication does not otherwise define the structure and contents of the directory.

## tests

If the package provides a root-level directory for test files, it MUST be named tests/.  
This publication does not otherwise define the structure and contents of the directory.
`
		expectedContent = fmt.Sprintf(expectedContent, projectName)
		if err != nil {
			t.Error(err)
		} else if string(readmeMd) != expectedContent {
			t.Error("the content of the file does not match the expected content")
		}
	})
}

func TestCreateReadmeMdWhenFailedToCreateFile(t *testing.T) {
	runTest(func() {
		creator.ProjectName = "dir_not_exists"
		err := creator.createReadmeMd()
		creator.ProjectName = "test_project"
		if err == nil {
			t.Error("the error should not be nil")
		}
	})
}

func TestCreateGitIgnore(t *testing.T) {
	runTest(func() {
		creator.createGitIgnore()
		file := fmt.Sprintf("%s/.gitignore", projectName)
		gitignore, err := os.ReadFile(file)
		expectedContent := `# IDE
.vscode/
.idea/

# dependencies
vendor/
node_modules/
`
		if err != nil {
			t.Error(err)
		} else if string(gitignore) != expectedContent {
			t.Error("the content of the file does not match the expected content")
		}
	})
}

func TestCreateGitAttributes(t *testing.T) {
	runTest(func() {
		creator.createGitAttributes()
		file := fmt.Sprintf("%s/.gitattributes", projectName)
		gitattributes, err := os.ReadFile(file)
		expectedContent := `* text=auto

*.css diff=css
*.html diff=html
*.md diff=markdown
*.php diff=php

/.github export-ignore
`
		if err != nil {
			t.Error(err)
		} else if string(gitattributes) != expectedContent {
			t.Error("the content of the file does not match the expected content")
		}
	})
}

func TestCreateEditorConfig(t *testing.T) {
	runTest(func() {
		creator.createEditorConfig()
		file := fmt.Sprintf("%s/.editorconfig", projectName)
		editorconfig, err := os.ReadFile(file)
		expectedContent := `root = true

[*]
charset = utf-8
end_of_line = lf
insert_final_newline = true
trim_trailing_whitespace = true
indent_style = space
indent_size = 4

[*.md]
trim_trailing_whitespace = false

[Makefile]
indent_style = tab

[*.{js,jsx,ts,tsx,vue,json,yml,yaml}]
indent_style = space
indent_size = 2
`
		if err != nil {
			t.Error(err)
		} else if string(editorconfig) != expectedContent {
			t.Error("the content of the file does not match the expected content")
		}
	})
}

func TestCreateComposerJson(t *testing.T) {
	runTest(func() {
		creator.createComposerJson()
		file := fmt.Sprintf("%s/composer.json", projectName)
		composerJson, err := os.ReadFile(file)
		expectedContent := `{
  "name": "vendor_name/%s",
  "type": "project",
  "autoload": {
    "psr-4": {
      "App\\": "src/"
    }
  },
  "autoload-dev": {
    "psr-4": {
      "Tests\\": "tests/"
    }
  },
  "prefer-stable": true
}
`
		expectedContent = fmt.Sprintf(expectedContent, projectName)
		if err != nil {
			t.Error(err)
		} else if string(composerJson) != expectedContent {
			t.Error("the content of the file does not match the expected content")
		}
	})
}

func TestCreateComposerJsonWhenFailedToCreateFile(t *testing.T) {
	runTest(func() {
		creator.ProjectName = "dir_not_exists"
		err := creator.createComposerJson()
		creator.ProjectName = "test_project"
		if err == nil {
			t.Error("the error should not be nil")
		}
	})
}
