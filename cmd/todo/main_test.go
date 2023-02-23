/* try to write integration tests instead unit as did on api
1) compile binary
2) execute binary with diff args and assert correct behaviour
recommended way for executing extra setup before tests is TestMain function.

*/

package main

import (
	"fmt"
	"os"
	"os/exec"       // exec external commands
	"path/filepath" // deal with directory path
	"runtime"       // identify running os
	"strings"       // compare strings
	"testing"
)

var (
	binName  = "todo"
	fileName = ".todo.json"
)

func TestMain(m *testing.M) {
	fmt.Println("building tool...")

	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	build := exec.Command("go", "build", "-o", binName)

	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "cannot build tool %s: %s", binName, err)
		os.Exit(1)
	}

	fmt.Println("running tests...")
	result := m.Run()

	fmt.Println("cleaning up...")
	os.Remove(binName)
	os.Remove(fileName)
	os.Exit(result)
}

// use the subtests feature to execute tests
// that depend on each other by using the t.Run method of the testing package.
func TestTodoCLItask(t *testing.T) {
	task := "Test task #1"

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	cmdPath := filepath.Join(dir, binName)

	// ensure tool can add new task via t.Run
	t.Run("Add new task", func(t *testing.T) {
		// execute binary with splitted task var
		cmd := exec.Command(cmdPath, strings.Split(task, " ")...)

		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	// ensure tool can list tasks
	t.Run("List tasks", func(t *testing.T) {
		cmd := exec.Command(cmdPath)

		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		expected := task + "\n"

		if expected != string(out) {
			t.Errorf("expected %q, got %q instead \n", expected, string(out))
		}

	})

}
