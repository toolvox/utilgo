package flags_test

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"utilgo/pkg/flags"
)

func ExampleOutputFileValue() {
	// Simulating setting the flag from command line arguments.
	// In a real scenario, this would come from the user input.
	tempDir := os.TempDir()
	pathExample := filepath.Join(tempDir, "example.txt")
	pathDefault := filepath.Join(tempDir, "default.txt")
	defer func() {
		os.Remove(pathExample)
		os.Remove(pathDefault)
	}()

	// Example: file1 set, file2 default
	os.Args = []string{"path/to/cmd", "-out1", pathExample}
	// this is what you would typically find in your code:
	var ofv1, ofv2 flags.OutputFileValue
	flag.Var(&ofv1, "out1", "usage 1")
	flag.Var(flags.OutputFileDefault(&ofv2, pathDefault, 0), "out2", "usage 2")
	flag.Parse()

	// now we get the [io.Writer] and use it to write to the files.
	writer1 := ofv1.Writer()
	writer1.Write([]byte("Example"))
	writer1.Close()

	writer2 := ofv2.Writer()
	writer2.Write([]byte("Default"))
	writer2.Close()

	content1, _ := os.ReadFile(pathExample)
	fmt.Println(string(content1))
	content2, _ := os.ReadFile(pathDefault)
	fmt.Println(string(content2))

	// Output:
	// Example
	// Default
}
