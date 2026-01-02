package funcs_test

import (
	"bytes"
	"html/template"
	"testing"

	"github.com/toolvox/utilgo/pkg/tmplutil/funcs"
)

func TestDict(t *testing.T) {
	tmpl, err := template.New("test").Funcs(template.FuncMap{
		"dict": funcs.DictFunc,
	}).Parse(`{{ with dict "name" "John" "age" 30 "city" "New York" }}Name: {{ .name }}, Age: {{ .age }}, City: {{ .city }}{{ else }}{{ . }}{{ end }}`)
	if err != nil {
		t.Fatalf("Failed to parse template: %v", err)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, nil)
	if err != nil {
		t.Fatalf("Failed to execute template: %v", err)
	}

	expected := "Name: John, Age: 30, City: New York"
	if buf.String() != expected {
		t.Fatalf("Unexpected output: got %v, want %v", buf.String(), expected)
	}
}

func TestDictOddArgs(t *testing.T) {
	tmpl, err := template.New("test").Funcs(template.FuncMap{
		"dict": funcs.DictFunc,
	}).Parse(`{{ dict "name" "John" "age" }}`)
	if err != nil {
		t.Fatalf("Failed to parse template: %v", err)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, nil)
	if err == nil {
		t.Fatalf("Expected error due to odd number of arguments, got nil")
	}
}

func TestDictNonStringKey(t *testing.T) {
	tmpl, err := template.New("test").Funcs(template.FuncMap{
		"dict": funcs.DictFunc,
	}).Parse(`{{ dict 123 "John" }}`)
	if err != nil {
		t.Fatalf("Failed to parse template: %v", err)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, nil)
	if err == nil {
		t.Fatalf("Expected error due to non-string key, got nil")
	}
}

func TestContainsKeyFunc(t *testing.T) {
	tests := []struct {
		name     string
		tmplStr  string
		data     interface{}
		expected string
		err      bool
	}{
		{
			name:     "Key Exists",
			tmplStr:  `{{ if containsKey .Map .Key }}Key exists{{ else }}Key does not exist{{ end }}`,
			data:     map[string]interface{}{"Map": map[string]string{"foo": "bar"}, "Key": "foo"},
			expected: "Key exists",
			err:      false,
		},
		{
			name:     "Key Does Not Exist",
			tmplStr:  `{{ if containsKey .Map .Key }}Key exists{{ else }}Key does not exist{{ end }}`,
			data:     map[string]interface{}{"Map": map[string]string{"foo": "bar"}, "Key": "baz"},
			expected: "Key does not exist",
			err:      false,
		},
		{
			name:     "Invalid Map Type",
			tmplStr:  `{{ if containsKey .Map .Key }}Key exists{{ else }}Key does not exist{{ end }}`,
			data:     map[string]interface{}{"Map": "not a map", "Key": "foo"},
			expected: "",
			err:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpl, err := template.New("test").Funcs(template.FuncMap{
				"containsKey": funcs.ContainsKeyFunc,
			}).Parse(tt.tmplStr)
			if err != nil {
				t.Fatalf("Failed to parse template: %v", err)
			}

			var buf bytes.Buffer
			err = tmpl.Execute(&buf, tt.data)
			if (err != nil) != tt.err {
				t.Fatalf("Expected error: %v, got: %v", tt.err, err)
			}

			if buf.String() != tt.expected {
				t.Fatalf("Unexpected output: got %v, want %v", buf.String(), tt.expected)
			}
		})
	}
}
