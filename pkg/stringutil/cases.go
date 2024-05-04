package stringutil

import (
	"strings"
	"unicode"
)

type StringCase int

const (
	// camelCaseText
	CamelCase StringCase = iota
	// PascalCaseText
	PascalCase
	// snake_case_text
	SnakeCase
	// Title Case Text
	TitleCase
)

func (c StringCase) Render(words []string) string {
	sb := strings.Builder{}
	switch c {
	case CamelCase:
		for i, w := range words {
			if i == 0 {
				sb.WriteString(strings.ToLower(w))
				continue
			}
			sb.WriteString(strings.ToUpper(w[:1]))
			sb.WriteString(strings.ToLower(w[1:]))
		}
	case PascalCase:
		for _, w := range words {
			sb.WriteString(strings.ToUpper(w[:1]))
			sb.WriteString(strings.ToLower(w[1:]))
		}
	case SnakeCase:
		sb.WriteString(strings.Join(Lowercase(words...), "_"))
	case TitleCase:
		sb.WriteString(strings.Join(Capitalize(words...), " "))
	default:
		return ""
	}
	return sb.String()
}

func Lowercase(words ...string) []string {
	result := make([]string, len(words))
	for i, w := range words {
		result[i] = strings.ToLower(w)
	}
	return result
}

func Capitalize(words ...string) []string {
	result := make([]string, len(words))
	for i, w := range words {
		result[i] = strings.ToUpper(w[:1]) + strings.ToLower(w[1:])
	}
	return result
}

func (c StringCase) Parse(phrase string) []string {
	switch c {
	case PascalCase, CamelCase:
		return parseCamelPascal(phrase)
	case SnakeCase:
		return strings.Split(phrase, "_")
	case TitleCase:
		return Lowercase(strings.Split(phrase, " ")...)
	default:
		return nil
	}
}

func parseCamelPascal(phrase string) []string {
	var result []string
	last := 0
	numbering := false
	for i, r := range phrase {
		if i == 0 && unicode.IsDigit(r) {
			numbering = true
		}
		if i == 0 || !unicode.IsUpper(r) && !unicode.IsDigit(r) {
			continue
		}
		if unicode.IsDigit(r) {
			if !numbering {
				result = append(result, strings.ToLower(phrase[last:i]))
				last = i
			}
			numbering = true
			continue
		}
		if numbering && !unicode.IsDigit(r) {
			result = append(result, strings.ToLower(phrase[last:i]))
			last = i
			numbering = false
			continue
		}
		result = append(result, strings.ToLower(phrase[last:i]))
		last = i
	}
	result = append(result, strings.ToLower(phrase[last:]))
	return result
}

func TransCase(input string, from, to StringCase) string {
	words := from.Parse(input)
	return to.Render(words)
}
