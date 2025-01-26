package shell

import (
	"bytes"
	"os/exec"

	"github.com/charmbracelet/log"
)

const ShellToUse = "bash"

// func Exec(command string) (string, string, error) {
// 	var stdout bytes.Buffer
// 	var stderr bytes.Buffer
// 	cmd := exec.Command(ShellToUse, "-c", command)
// 	cmd.Stdout = &stdout
// 	cmd.Stderr = &stderr
// 	err := cmd.Run()
// 	return stdout.String(), stderr.String(), err
// }

func Exec(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	stdoutstr := stdout.String()
	stderrstr := stderr.String()
	if err != nil {
		log.Errorf("\n%s\n%s", stdoutstr, stderrstr)
	}
	return stdoutstr, stderrstr, err
}

func ExecPrint(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	stdoutstr := stdout.String()
	stderrstr := stderr.String()
	if stdoutstr != "" {
		log.Print(stdoutstr)
	}
	if stderrstr != "" {
		log.Errorf("\n%s", stderrstr)
	}
	return stdoutstr, stderrstr, err
}
