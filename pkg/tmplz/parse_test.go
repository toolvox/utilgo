package tmplz_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"utilgo/pkg/errs"
	. "utilgo/pkg/tmplz"
)

func Test_Parse(t *testing.T) {
	t.Run("ParseNode", func(t *testing.T) {
		tests := []struct {
			input         string
			expectedNode  string
			expectedRest  string
			expectedError error
		}{
			{"!@#", "", "!@#", errs.New("node must start with '@'")},
			{"@#!", "", "@#!", errs.Wrap("'@#!' does not begin a valid node", errs.New("could not complete variable from: '@#!'"))},
			{"@@Fmt_@Op", "@@Fmt_@Op__", "", nil},
		}

		for ti, tt := range tests {
			t.Run(fmt.Sprint(ti), func(t *testing.T) {
				must := require.New(t)
				must.NotPanics(func() {
					node, rest, err := ParseNode(String(tt.input))
					if tt.expectedError != nil {
						must.ErrorIs(err, tt.expectedError, "bad error")
					} else {
						must.NoError(err)
					}

					if tt.expectedNode == "" {
						must.Nil(node)
					} else {
						must.NotNil(node)
						actualNode := node.Text.String()
						must.Equal(tt.expectedNode, actualNode, "node")
					}

					must.Equal(tt.expectedRest, rest, "rest")
				})
			})
		}
		t.Run("ParseNodeNoLog", func(t *testing.T) {
			for ti, tt := range tests {
				t.Run(fmt.Sprint(ti), func(t *testing.T) {
					must := require.New(t)
					must.NotPanics(func() {
						node, rest, err := ParseNode(String(tt.input))
						if tt.expectedError != nil {
							must.ErrorIs(err, tt.expectedError, "bad error")
						} else {
							must.NoError(err)
						}

						if tt.expectedNode == "" {
							must.Nil(node)
						} else {
							must.NotNil(node)
							actualNode := node.Text.String()
							must.Equal(tt.expectedNode, actualNode, "node")
						}

						must.Equal(tt.expectedRest, rest, "rest")
					})
				})
			}
		})
	})

	t.Run("ParseTemplate", func(t *testing.T) {
		tests := []struct {
			name          string
			input         string
			expectedNodes []string
		}{
			{"Bad", "Hello, @?Who!", []string{}},
			{"Escape", "o@_o", []string{"@_"}},
			{"Hello", "Hello, @Who!", []string{"@Who_"}},
			{"Greet", "@Greet, @What!", []string{"@Greet_", "@What_"}},
			{"Complex", "What the @@Format_@Operation@Mod?", []string{"@@Format_@Operation@Mod___"}},
		}

		for ti, tt := range tests {
			t.Run(fmt.Sprintln(ti, tt.name, tt.input), func(t *testing.T) {
				must := require.New(t)
				must.NotPanics(func() {
					tmpl := ParseTemplate(tt.input)
					must.NotNil(tmpl)
					must.Equal(tt.input, tmpl.Text.String())
					must.Len(tmpl.Children, len(tt.expectedNodes))
					for i, c := range tmpl.Children {
						must.Equal(tt.expectedNodes[i], c.Text.String())
					}
				})
			})
		}
		t.Run("ParseTemplateNoLog", func(t *testing.T) {
			for ti, tt := range tests {
				t.Run(fmt.Sprintln(ti, tt.input), func(t *testing.T) {
					must := require.New(t)
					must.NotPanics(func() {
						tmpl := ParseTemplate(tt.input)
						must.NotNil(tmpl)
						must.Equal(tt.input, tmpl.Text.String())
						must.Len(tmpl.Children, len(tt.expectedNodes))
						for i, c := range tmpl.Children {
							must.Equal(tt.expectedNodes[i], c.Text.String())
						}
					})
				})
			}
		})
		t.Run("ParseTemplates", func(t *testing.T) {
			testTemplates := map[string]string{}
			expectedTemplateNodes := map[string][]string{}
			for _, tt := range tests {
				testTemplates[tt.name] = tt.input
				expectedTemplateNodes[tt.name] = tt.expectedNodes
			}

			must := require.New(t)
			must.NotPanics(func() {
				tmpls := ParseTemplates(testTemplates)
				must.NotNil(tmpls)
				for name, tmpl := range tmpls.Templates {
					must.NotNil(tmpl)
					must.Equal(tmpl.Text.String(), testTemplates[name])
					expectedNodes := expectedTemplateNodes[name]
					must.Len(tmpl.Children, len(expectedNodes))
					for i, c := range tmpl.Children {
						must.Equal(expectedNodes[i], c.Text.String())
					}
				}
			})

			t.Run("ParseTemplatesNoLog", func(t *testing.T) {
				must := require.New(t)
				must.NotPanics(func() {
					tmpls := ParseTemplates(testTemplates)
					must.NotNil(tmpls)
					for name, tmpl := range tmpls.Templates {
						must.NotNil(tmpl)
						must.Equal(tmpl.Text.String(), testTemplates[name])
						expectedNodes := expectedTemplateNodes[name]
						must.Len(tmpl.Children, len(expectedNodes))
						for i, c := range tmpl.Children {
							must.Equal(expectedNodes[i], c.Text.String())
						}
					}
				})
			})
		})
	})
}
