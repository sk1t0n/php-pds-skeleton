package creator

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"text/template"
)

type Creator struct {
	ProjectName string
	innerDirs   []string
}

func New() *Creator {
	c := &Creator{}
	c.setProjectName()
	c.setInnerDirs()
	return c
}

func (c *Creator) setProjectName() {
	err := c.checkArgs()
	if err != nil {
		log.Fatalln(err)
	}

	c.ProjectName = os.Args[1]
}

func (*Creator) checkArgs() error {
	if len(os.Args) < 2 {
		return errors.New("checkArgs: the project name is not set")
	}
	return nil
}

func (c *Creator) setInnerDirs() {
	c.innerDirs = []string{
		"bin",
		"config",
		"public",
		"resources",
		"src",
		"tests",
	}
}

func (c *Creator) CreateProjectStructure() {
	c.createDirs()

	err := c.createPublicIndexPhp()
	if err != nil {
		log.Fatalln(err)
	}
	err = c.createConfigAppPhp()
	if err != nil {
		log.Fatalln(err)
	}
	err = c.createReadmeMd()
	if err != nil {
		log.Fatalln(err)
	}
	err = c.createGitIgnore()
	if err != nil {
		log.Fatalln(err)
	}
	err = c.createGitAttributes()
	if err != nil {
		log.Fatalln(err)
	}
	err = c.createEditorConfig()
	if err != nil {
		log.Fatalln(err)
	}
	err = c.createComposerJson()
	if err != nil {
		log.Fatalln(err)
	}
}

func (c *Creator) createDirs() {
	err := c.createDir(c.ProjectName)
	if err != nil && errors.Is(err, fs.ErrNotExist) {
		log.Fatalln(err)
	}

	for _, dir := range c.innerDirs {
		path := fmt.Sprintf("%s/%s", c.ProjectName, dir)
		err = c.createDir(path)
		if err != nil && errors.Is(err, fs.ErrNotExist) {
			log.Fatalln(err)
		}
	}
}

func (*Creator) createDir(dirName string) error {
	err := os.Mkdir(dirName, 0750)
	return err
}

func (c *Creator) createPublicIndexPhp() error {
	file := fmt.Sprintf("%s/public/index.php", c.ProjectName)
	content := "<?php\n"
	err := os.WriteFile(file, []byte(content), 0666)
	return err
}

func (c *Creator) createConfigAppPhp() error {
	file := fmt.Sprintf("%s/config/app.php", c.ProjectName)
	content := "<?php\n"
	err := os.WriteFile(file, []byte(content), 0666)
	return err
}

func (c *Creator) createReadmeMd() error {
	content := `# {{ .ProjectName }}

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

	path := fmt.Sprintf("%s/README.md", c.ProjectName)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	t := template.Must(template.New("README.md").Parse(content))
	err = t.Execute(file, c)
	return err
}

func (c *Creator) createGitIgnore() error {
	content := `# IDE
.vscode/
.idea/

# dependencies
vendor/
node_modules/
`
	file := fmt.Sprintf("%s/.gitignore", c.ProjectName)
	err := os.WriteFile(file, []byte(content), 0666)
	return err
}

func (c *Creator) createGitAttributes() error {
	content := `* text=auto

*.css diff=css
*.html diff=html
*.md diff=markdown
*.php diff=php

/.github export-ignore
`
	file := fmt.Sprintf("%s/.gitattributes", c.ProjectName)
	err := os.WriteFile(file, []byte(content), 0666)
	return err
}

func (c *Creator) createEditorConfig() error {
	content := `root = true

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
	file := fmt.Sprintf("%s/.editorconfig", c.ProjectName)
	err := os.WriteFile(file, []byte(content), 0666)
	return err
}

func (c *Creator) createComposerJson() error {
	content := `{
  "name": "vendor_name/{{ .ProjectName }}",
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

	path := fmt.Sprintf("%s/composer.json", c.ProjectName)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	t := template.Must(template.New("composer.json").Parse(content))
	err = t.Execute(file, c)
	return err
}
