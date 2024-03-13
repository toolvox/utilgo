package stringutil

import (
	"fmt"
	"slices"
	"strings"
	"unicode"
)

// IndentEmptyLineFlag defines flags for handling empty lines during indentation.
type IndentEmptyLineFlag uint

const (
	// EmptyLine_TrimFirst trims the first line if it's empty.
	EmptyLine_TrimFirst IndentEmptyLineFlag = 1 << iota
	// EmptyLine_TrimHead trims all leading empty lines.
	EmptyLine_TrimHead
	// EmptyLine_TrimBody trims empty lines within the text body.
	EmptyLine_TrimBody
	// EmptyLine_TrimTail trims all trailing empty lines.
	EmptyLine_TrimTail
	// EmptyLine_TrimLast trims the last line if it's empty.
	EmptyLine_TrimLast

	// EmptyLine_TrimAll is a convenience flag to trim all empty lines.
	EmptyLine_TrimAll = EmptyLine_TrimHead | EmptyLine_TrimBody | EmptyLine_TrimTail
)

// IndentOption specifies options for indenting a block of text.
// It includes settings for handling empty lines and defining indentation strings.
type IndentOption struct {
	EmptyLines   IndentEmptyLineFlag // Flags for which empty lines to trim.
	IndentPrefix string              // Prefix for all non-empty lines.
	IndentString string              // The string used for indentation, e.g., "  " or "\t".
}

// Indent applies the [IndentOption] to a given block of text, returning the indented text.
// It processes the text to indent non-empty lines according to the specified options and trims empty lines based on the EmptyLines flags.
func (io IndentOption) Indent(text string) string {
	lines := strings.Split(text, "\n")
	var i int
	for ; i < len(lines); i++ {
		noSpaceLine := strings.TrimSpace(lines[i])
		if noSpaceLine != "" {
			break
		}
	}
	firstNonHead, lastNonEmpty := i, i
	var bodyEmptyLines, batchEmptyLines []int
	var minNonemptyIndent int = 1 << 31
	for ; i < len(lines); i++ {
		noSpaceLine := strings.TrimSpace(lines[i])
		if noSpaceLine == "" {
			batchEmptyLines = append(batchEmptyLines, i)
			continue
		}
		bodyEmptyLines = append(bodyEmptyLines, batchEmptyLines...)
		batchEmptyLines = []int{}
		lastNonEmpty = i
		existingIndent := CountPrefix(lines[i], unicode.IsSpace)
		if existingIndent < minNonemptyIndent {
			minNonemptyIndent = existingIndent
		}
	}

	offset := 0
	if io.EmptyLines&EmptyLine_TrimHead != 0 {
		offset = firstNonHead
	} else if io.EmptyLines&EmptyLine_TrimFirst != 0 {
		offset = min(1, firstNonHead)
	}
	lines = lines[offset:]

	if io.EmptyLines&EmptyLine_TrimBody != 0 {
		for _, idx := range bodyEmptyLines {
			idx -= offset
			lines = slices.Delete(lines, idx, idx+1)
			offset++
		}
	}

	if len(lines) > 0 && io.EmptyLines&EmptyLine_TrimTail != 0 {
		lines = lines[:lastNonEmpty-offset+1]
	} else if len(lines) > 0 && io.EmptyLines&EmptyLine_TrimLast != 0 {
		lines = lines[:max(len(lines)-1, lastNonEmpty-offset+1)]
	}

	var sb strings.Builder
	for _, line := range lines {
		if len(line) == 0 {
			sb.WriteRune('\n')
			continue
		}
		currentIndent := CountPrefix(line[minNonemptyIndent:], unicode.IsSpace)

		sb.WriteString(io.IndentPrefix)
		sb.WriteString(strings.Repeat(io.IndentString, currentIndent))
		sb.WriteString(strings.TrimSpace(line))
		sb.WriteRune('\n')
	}

	return sb.String()

}

// Indent a block of text using a "default" [IndentOption]:
//
//	EmptyLines:   EmptyLine_TrimHead | EmptyLine_TrimTail
//	IndentPrefix: ""
//	IndentString: "\t"
func Indent(text string) string {
	return IndentOption{
		EmptyLines:   EmptyLine_TrimHead | EmptyLine_TrimTail,
		IndentPrefix: "",
		IndentString: "\t",
	}.Indent(text)
}

// Indentf a block of text using a "default" [IndentOption] and format it:
//
//	EmptyLines:   EmptyLine_TrimHead | EmptyLine_TrimTail
//	IndentPrefix: ""
//	IndentString: "\t"
func Indentf(text string, args ...any) string {
	return fmt.Sprintf(IndentOption{
		EmptyLines:   EmptyLine_TrimHead | EmptyLine_TrimTail,
		IndentPrefix: "",
		IndentString: "\t",
	}.Indent(text), args...)
}
