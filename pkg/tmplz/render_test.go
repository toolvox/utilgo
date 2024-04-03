package tmplz_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/toolvox/utilgo/pkg/stringutil"
	. "github.com/toolvox/utilgo/pkg/tmplz"
)

func Test_Renders(t *testing.T) {
	t.Run("Node", func(t *testing.T) {
		tests := []struct {
			varTemplate    string
			expectedString string
			expectedDebug  string
			expectedDot    string
		}{
			{"", "", "", ""},

			{"@?@1@@@$nope", "", "", ""},

			{"@Hello", "{@Hello_}", "{'@Hello_'[0,6]}", "digraph g {\n\t\"@Hello_\"\n}\n"},

			{"@In@Too@Deep", "{@In@Too@Deep___{@Too@Deep__{@Deep_}}}",
				"{'@In@Too@Deep___'[0,12]{'@Too@Deep__'[3,14]{'@Deep_'[4,10]}}}",
				stringutil.Indent(`
					digraph g {
						"@In@Too@Deep___"
						"@In@Too@Deep___" -> "@Too@Deep__" [label = "[3, 14]"]
						"@Too@Deep__" -> "@Deep_" [label = "[4, 10]"]
					}
				`),
			},

			{"@Where@Here_@There@Not@Here",
				"{@Where@Here_@There@Not@Here____{@Here_}{@There@Not@Here___{@Not@Here__{@Here_}}}}",
				"{'@Where@Here_@There@Not@Here____'[0,27]{'@Here_'[6,12]}{'@There@Not@Here___'[12,30]{'@Not@Here__'[6,17]{'@Here_'[4,10]}}}}",
				stringutil.Indent(`
					digraph g {
						"@Where@Here_@There@Not@Here____"
						"@Where@Here_@There@Not@Here____" -> "@Here_" [label = "[6, 12]"]
						"@Where@Here_@There@Not@Here____" -> "@There@Not@Here___" [label = "[12, 30]"]
						"@There@Not@Here___" -> "@Not@Here__" [label = "[6, 17]"]
						"@Not@Here__" -> "@Here_" [label = "[4, 10]"]
					}
				`),
			},
		}

		for ti, tt := range tests {
			t.Run(fmt.Sprint(ti), func(t *testing.T) {
				must := require.New(t)
				must.NotPanics(func() {
					node, rest, err := ParseNode(String(tt.varTemplate))
					if tt.expectedString == "" {
						must.Error(err)
						must.Equal(rest, tt.varTemplate)
						must.Nil(node)
						return
					}
					must.NoError(err)
					must.Empty(rest)

					are := assert.New(t)
					are.Equal(tt.expectedString, node.String(), "String")
					are.Equal(tt.expectedDebug, node.Debug(), "Debug")
					are.Equal(tt.expectedDot, node.Dot(), "Dot")
				})
			})
		}
	})

	t.Run("Template", func(t *testing.T) {
		t.Run("Empty", func(t *testing.T) {
			var (
				expectedString = stringutil.Indentf(`
					Template:
						%s
					Variables:
				`, "")

				expectedDot = stringutil.Indent(`
					digraph g {
						"tmpl root" [label="", shape=diamond]
					}
				`)
			)

			must := require.New(t)
			must.NotPanics(func() {
				tmpl := ParseTemplate("")
				must.NotNil(tmpl)
				must.Equal("", tmpl.Text.String())

				are := assert.New(t)
				are.Equal(expectedString, tmpl.String())
				are.Equal(expectedString, tmpl.Debug())
				are.Equal(expectedDot, tmpl.Dot())
				t.Run("EmptyNoLog", func(t *testing.T) {
					must := require.New(t)
					must.NotPanics(func() {
						tmpl := ParseTemplate("")
						must.NotNil(tmpl)
						must.Equal("", tmpl.Text.String())

						are := assert.New(t)
						are.Equal(expectedString, tmpl.String())
						are.Equal(expectedString, tmpl.Debug())
						are.Equal(expectedDot, tmpl.Dot())
					})
				})
			})
		})

		t.Run("Sanity", func(t *testing.T) {
			sanityTestTemplate := "Hello, @Who! @Where@Here_@Or@Here @Or@Where to go?"
			sanityExpectedString := stringutil.Indent(`
				Template:
					Hello, @Who! @Where@Here_@Or@Here @Or@Where to go?
				Variables:
					{@Who_}
					{@Where@Here_@Or@Here___{@Here_}{@Or@Here__{@Here_}}}
					{@Or@Where__{@Where_}}
			`)

			sanityExpectedDebug := stringutil.Indent(`
				Template:
					Hello, @Who! @Where@Here_@Or@Here @Or@Where to go?
				Variables:
					{'@Who_'[7,11]}
					{'@Where@Here_@Or@Here___'[13,33]{'@Here_'[6,12]}{'@Or@Here__'[12,22]{'@Here_'[3,9]}}}
					{'@Or@Where__'[34,43]{'@Where_'[3,10]}}
			`)

			sanityExpectedDot := stringutil.Indent(`
				digraph g {
					"tmpl root" [label="Hello, @Who! @Where@Here_@Or@Here @Or@Where to go?", shape=diamond]
				
					"tmpl root" -> "@Who_" [label = "[7, 11]"]
				
					"tmpl root" -> "@Where@Here_@Or@Here___" [label = "[13, 33]"]
					"@Where@Here_@Or@Here___" -> "@Here_" [label = "[6, 12]"]
					"@Where@Here_@Or@Here___" -> "@Or@Here__" [label = "[12, 22]"]
					"@Or@Here__" -> "@Here_" [label = "[3, 9]"]
				
					"tmpl root" -> "@Or@Where__" [label = "[34, 43]"]
					"@Or@Where__" -> "@Where_" [label = "[3, 10]"]
				}
			`)

			must := require.New(t)
			must.NotPanics(func() {
				tmpl := ParseTemplate(sanityTestTemplate)
				must.NotNil(tmpl)
				must.Equal(sanityTestTemplate, tmpl.Text.String())

				are := assert.New(t)
				are.Equal(sanityExpectedString, tmpl.String())
				are.Equal(sanityExpectedDebug, tmpl.Debug())
				are.Equal(sanityExpectedDot, tmpl.Dot())
			})

			t.Run("NoLog", func(t *testing.T) {
				must := require.New(t)
				must.NotPanics(func() {
					tmpl := ParseTemplate(sanityTestTemplate)
					must.NotNil(tmpl)
					must.Equal(sanityTestTemplate, tmpl.Text.String())

					are := assert.New(t)
					are.Equal(sanityExpectedString, tmpl.String())
					are.Equal(sanityExpectedDebug, tmpl.Debug())
					are.Equal(sanityExpectedDot, tmpl.Dot())
				})
			})

			t.Run("Templates", func(t *testing.T) {
				insanityTestTemplate := "example@gmail.com is @What@_@Maybe @is"
				insanityExpectedString := stringutil.Indent(`
					Template:
						example@gmail.com is @What@_@Maybe @is
					Variables:
						{@gmail_}
						{@What@_@Maybe__{@_}{@Maybe_}}
						{@is_}
				`)

				insanityExpectedDebug := stringutil.Indent(`
					Template:
						example@gmail.com is @What@_@Maybe @is
					Variables:
						{'@gmail_'[7,13]}
						{'@What@_@Maybe__'[21,34]{'@_'[5,7]}{'@Maybe_'[7,14]}}
						{'@is_'[35,38]}
				`)

				insanityExpectedDot := stringutil.Indent(`
					digraph g {
						"tmpl root" [label="example@gmail.com is @What@_@Maybe @is", shape=diamond]
					
						"tmpl root" -> "@gmail_" [label = "[7, 13]"]
					
						"tmpl root" -> "@What@_@Maybe__" [label = "[21, 34]"]
						"@What@_@Maybe__" -> "@_" [label = "[5, 7]"]
						"@What@_@Maybe__" -> "@Maybe_" [label = "[7, 14]"]
					
						"tmpl root" -> "@is_" [label = "[35, 38]"]
					}
				`)

				t.Run("Insanity_Alone", func(t *testing.T) {
					must := require.New(t)
					must.NotPanics(func() {
						tmpl := ParseTemplate(insanityTestTemplate)
						must.NotNil(tmpl)
						must.Equal(insanityTestTemplate, tmpl.Text.String())

						are := assert.New(t)
						are.Equal(insanityExpectedString, tmpl.String())
						are.Equal(insanityExpectedDebug, tmpl.Debug())
						are.Equal(insanityExpectedDot, tmpl.Dot())
					})
				})

				t.Run("Empty", func(t *testing.T) {
					must := require.New(t)
					must.NotPanics(func() {
						tmpls := ParseTemplates(nil)
						must.NotNil(tmpls)
						must.Len(tmpls.Templates, 0)

						are := assert.New(t)
						are.Equal("", tmpls.String())
						are.Equal("", tmpls.Debug())
						are.Equal("digraph g {\n}\n", tmpls.Dot())
					})
				})

				finalExpectedString := stringutil.Indent(`
					Insanity:
						Template:
							example@gmail.com is @What@_@Maybe @is
						Variables:
							{@gmail_}
							{@What@_@Maybe__{@_}{@Maybe_}}
							{@is_}

					Sanity:
						Template:
							Hello, @Who! @Where@Here_@Or@Here @Or@Where to go?
						Variables:
							{@Who_}
							{@Where@Here_@Or@Here___{@Here_}{@Or@Here__{@Here_}}}
							{@Or@Where__{@Where_}}
				`)

				finalExpectedDebug := stringutil.Indent(`
					Insanity:
						Template:
							example@gmail.com is @What@_@Maybe @is
						Variables:
							{'@gmail_'[7,13]}
							{'@What@_@Maybe__'[21,34]{'@_'[5,7]}{'@Maybe_'[7,14]}}
							{'@is_'[35,38]}
					
					Sanity:
						Template:
							Hello, @Who! @Where@Here_@Or@Here @Or@Where to go?
						Variables:
							{'@Who_'[7,11]}
							{'@Where@Here_@Or@Here___'[13,33]{'@Here_'[6,12]}{'@Or@Here__'[12,22]{'@Here_'[3,9]}}}
							{'@Or@Where__'[34,43]{'@Where_'[3,10]}}
				`)

				finalExpectedDot := stringutil.Indent(`
					digraph g {

						"Insanity" [shape=rect]
						"Insanity" -> "tmpl root 0" [color=red]
						"tmpl root 0" [label="example@gmail.com is @What@_@Maybe @is", shape=diamond]
					
						"tmpl root 0" -> "@gmail_" [label = "[7, 13]"]
					
						"tmpl root 0" -> "@What@_@Maybe__" [label = "[21, 34]"]
						"@What@_@Maybe__" -> "@_" [label = "[5, 7]"]
						"@What@_@Maybe__" -> "@Maybe_" [label = "[7, 14]"]
					
						"tmpl root 0" -> "@is_" [label = "[35, 38]"]
					
						"Sanity" [shape=rect]
						"Sanity" -> "tmpl root 1" [color=red]
						"tmpl root 1" [label="Hello, @Who! @Where@Here_@Or@Here @Or@Where to go?", shape=diamond]
					
						"tmpl root 1" -> "@Who_" [label = "[7, 11]"]
					
						"tmpl root 1" -> "@Where@Here_@Or@Here___" [label = "[13, 33]"]
						"@Where@Here_@Or@Here___" -> "@Here_" [label = "[6, 12]"]
						"@Where@Here_@Or@Here___" -> "@Or@Here__" [label = "[12, 22]"]
						"@Or@Here__" -> "@Here_" [label = "[3, 9]"]
					
						"tmpl root 1" -> "@Or@Where__" [label = "[34, 43]"]
						"@Or@Where__" -> "@Where_" [label = "[3, 10]"]
					}
				`)

				must := require.New(t)
				must.NotPanics(func() {
					tmpls := ParseTemplates(map[string]string{
						"Sanity":   sanityTestTemplate,
						"Insanity": insanityTestTemplate,
					})

					must.NotNil(tmpls)
					must.Equal(sanityTestTemplate, tmpls.Templates["Sanity"].Text.String())
					must.Equal(insanityTestTemplate, tmpls.Templates["Insanity"].Text.String())

					are := assert.New(t)
					if !are.Equal(finalExpectedString, tmpls.String()) {
						os.WriteFile("0", []byte(tmpls.String()), 0644)
					}

					if !are.Equal(finalExpectedDebug, tmpls.Debug()) {
						os.WriteFile("1", []byte(tmpls.Debug()), 0644)
					}

					if !are.Equal(finalExpectedDot, tmpls.Dot()) {
						os.WriteFile("2", []byte(tmpls.Dot()), 0644)
					}
				})
			})
		})
	})
}
