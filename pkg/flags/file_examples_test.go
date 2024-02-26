package flags_test

import (
	"flag"
	"fmt"
	"io"
	"os"

	"utilgo/pkg/flags"
)

// Typical use-case for [utilgo/pkg/flags.FileValue]
func ExampleFileValue() {
	// Setup: initialization for example, can be ignored
	tmpPath := os.TempDir() + "/example.txt"
	os.WriteFile(tmpPath, []byte("Hello, world!"), 0644)
	defer os.Remove(tmpPath)
	os.Args = []string{"path/to/cmd", "-file", tmpPath}

	// Example: this is what you would typically find in your code
	var fv flags.FileValue
	flag.Var(&fv, "file", "file to read")
	flag.Parse() // FileValue.Set will be called by flag.Parse

	// Now you can get all the content by calling Get()
	bytes := fv.Get().([]byte)
	fmt.Println(string(bytes))
	// Or you can get an io.Reader by calling Reader()
	reader := fv.Reader()
	bytes, _ = io.ReadAll(reader)
	fmt.Println(string(bytes))
	// Output:
	// Hello, world!
	// Hello, world!
}

// Using the Reader method to read content from [utilgo/pkg/flags.FileValue].
func ExampleFileValue_Reader() {
	// Setup: Create a mock FileValue
	fv := flags.FileValue{"example.txt", []byte("Hello, world!")}

	// Example: using Reader
	buf, _ := io.ReadAll(fv.Reader())
	fmt.Println(string(buf))
	// Output: Hello, world!
}
