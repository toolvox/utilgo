package flags_test

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"utilgo/pkg/flags"
)

func ExampleOutputFileValue() {
	// Simulating setting the flag from command line arguments.
	// In a real scenario, this would come from the user input.
	f, err := os.CreateTemp(os.TempDir(), "example_output.txt")
	if err != nil {
		fmt.Printf("could not create temp file: %v\n", err)
		return
	}
	defer os.Remove(f.Name())
	os.Args = []string{"cmd", "-out", f.Name()}

	// Create a FlagSet and OutputFileValue flag to simulate command line flag parsing.
	fs := flag.NewFlagSet("", flag.ExitOnError)
	outputFile := flags.NewOutputFileValue("", os.O_CREATE|os.O_WRONLY|os.O_TRUNC)
	fs.Var(&outputFile, "out", "Output file path. Defaults to stdout if empty.")

	// Parse the flags to simulate command line input.
	fs.Parse(os.Args[1:])

	// Get the writer from OutputFileValue, which will be either a file writer or os.Stdout.
	writer, err := outputFile.Writer()
	if err != nil {
		fmt.Printf("Failed to open output writer: %v\n", err)
		return
	}
	defer writer.(*os.File).Close() // Ensure we close the file if one was opened.

	// Example writing operation to the writer.
	_, err = writer.Write([]byte("Hello, OutputFileValue!\n"))
	if err != nil {
		fmt.Printf("Failed to write to output: %v\n", err)
		return
	}

	// This is for the sake of example to show how you might check the file content.
	// In a real application, this verification step would be unnecessary.
	if outputFile.Filename != "" {
		// Re-open the file for reading to verify the content.
		file, err := os.Open(outputFile.Filename)
		if err != nil {
			fmt.Printf("Failed to re-open the output file: %v\n", err)
			return
		}
		defer file.Close()

		reader := bufio.NewReader(file)
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Printf("Failed to read from output file: %v\n", err)
			return
		}

		// Print the line to stdout for the example output verification.
		fmt.Printf("%s", line)
	}

	// Output: Hello, OutputFileValue!
}
