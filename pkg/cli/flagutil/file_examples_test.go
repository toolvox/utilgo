package flagutil_test

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/toolvox/utilgo/pkg/cli/flagutil"
)

// Typical use-case for [github.com/toolvox/utilgo/pkg/flagutil.FileValue]
func ExampleFileValue() {
	// Setup: initialization for example, can be ignored
	tempDir := os.TempDir()
	pathExample := filepath.Join(tempDir, "example.txt")
	pathDefault := filepath.Join(tempDir, "default.txt")
	defer func() {
		os.Remove(pathExample)
		os.Remove(pathDefault)
	}()
	os.WriteFile(pathExample, []byte("Hello, world!"), 0111)
	os.WriteFile(pathDefault, []byte("This\n  Is\n    DEFAULT!"), 0111)

	// Example: file1 set, file2 default
	os.Args = []string{"path/to/cmd", "-file1", pathExample}
	// this is what you would typically find in your code:
	var fv1, fv2 flagutil.FileValue
	flag.Var(&fv1, "file1", "file to read")
	flag.Var(flagutil.FileDefault(&fv2, pathDefault), "file2", "file to read")
	flag.Parse() // [FileValue.Set] will be called by flag.Parse

	// Now you can get an [io.Reader] by calling Reader()
	reader1 := fv1.Reader()
	bytes1, _ := io.ReadAll(reader1)
	fmt.Println(string(bytes1))

	reader2 := fv2.Reader()
	bytes2, _ := io.ReadAll(reader2)
	fmt.Println(string(bytes2))

	// Output:
	// Hello, world!
	// This
	//   Is
	//     DEFAULT!
}

// Using the Reader method to read content from [github.com/toolvox/utilgo/pkg/flagutil.FileValue].
func ExampleFileValue_Reader() {
	// Setup: Create a mock FileValue
	fv := flagutil.FileValue{"example.txt", []byte("Hello, world!")}

	// Example: using Reader
	buf, _ := io.ReadAll(fv.Reader())
	fmt.Println(string(buf))
	// Output: Hello, world!
}
