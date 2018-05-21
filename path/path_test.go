package path

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os/exec"
	"path/filepath"
	"testing"
)

func runCommand(t *testing.T, command string, args []string) {
	cmd := exec.Command(command, args...)
	out, err := cmd.CombinedOutput()
	if err != nil && testing.Verbose() {
		fmt.Println("command:", command, " ", args, ", output:", string(out))
	}
	assert.NoError(t, err)
}

func copyFileWithSudo(t *testing.T, from, to string) {
	// Create directory
	todir := filepath.Dir(to)
	runCommand(t, "sudo", []string{"mkdir", "-p", todir})
	for todir != "/" {
		runCommand(t, "sudo", []string{"chmod", "755", todir})
		todir = filepath.Dir(todir)
	}
	todir = filepath.Dir(todir)

	runCommand(t, "sudo", []string{"mkdir", "-p", todir})
	runCommand(t, "sudo", []string{"cp", from, to})
}

func TestPath(t *testing.T) {
	copyFileWithSudo(t, "path.go", "/opt/path/path.go")
	fmt.Println("thisFileDir():", thisFileDir())
}
