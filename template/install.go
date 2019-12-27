package template

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/state-alchemists/zaruba/command"
	"github.com/state-alchemists/zaruba/config"
)

// Install template
func Install(gitURL, dirName string) (err error) {
	templateDir := config.GetTemplateDir()
	log.Printf("[INFO] Install template from `%s` to `%s`", gitURL, templateDir)
	// run git init
	if err = command.Run(templateDir, "git", "clone", gitURL, dirName, "--depth=1"); err != nil {
		return
	}
	// install-template should be exists
	if !isScriptExists(templateDir, dirName, "install-template") {
		os.RemoveAll(filepath.Join(templateDir, dirName))
		err = errors.New("Cannot find `install-template` script")
		return
	}
	// create-component should be exists
	if !isScriptExists(templateDir, dirName, "create-component") {
		os.RemoveAll(filepath.Join(templateDir, dirName))
		err = errors.New("Cannot find `create-component` script")
		return
	}
	// make the file executable
	os.Chmod(filepath.Join(templateDir, dirName, "install-template.zaruba"), 0555)
	os.Chmod(filepath.Join(templateDir, dirName, "create-component.zaruba"), 0555)
	// run install
	log.Printf("[INFO] Execute `./install-template.zaruba`")
	err = command.Run(filepath.Join(templateDir, dirName), "./install-template.zaruba")
	return
}

func isScriptExists(templateDir, dirName, actionName string) (exist bool) {
	// imperative
	if _, err := os.Stat(filepath.Join(templateDir, dirName, actionName+".zaruba")); err == nil {
		return true
	}
	// declarative yml
	if _, err := os.Stat(filepath.Join(templateDir, dirName, actionName+".zaruba.yml")); err == nil {
		return true
	}
	// declarative yaml
	if _, err := os.Stat(filepath.Join(templateDir, dirName, actionName+".zaruba.yml")); err == nil {
		return true
	}
	return false
}
