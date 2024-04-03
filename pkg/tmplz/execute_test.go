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

func DoOrDie(assert *assert.Assertions, expected, actual, name string) {
	if assert.Equal(expected, actual) {
		os.Remove(name)
	} else {
		os.WriteFile(name, []byte(actual), 0644)
	}
}

func Test_Execute_Node(t *testing.T) {
	var testAssignment Assignments = make(Assignments)
	testAssignment.Assign("Fmt", "TORPOR")
	testAssignment.Assign("Op", "ReCode")
	testAssignment.Assign("Mod", "NonLocal")
	testAssignment.Assign("TORPORReCode", "For_Great_Justice!")
	testAssignment.Assign("TORPORReCodeNonLocal", "ZigZigZig")
	testAssignment.Assign("TORPORReCode_NonLocal", "ZIG!")

	tests := []struct {
		name             string
		template         string
		assignments      Assignments
		expectedNoChange bool
		expectedString   string
		expectedDebug    string
		expectedDot      string
		expectedResult   string
	}{
		{
			name:             "NoChange",
			template:         "@@Fmt_@Op__",
			assignments:      Assignments{},
			expectedNoChange: true,
			expectedString:   "{@@Fmt_@Op__{@Fmt_}{@Op_}}",
			expectedDebug:    "{'@@Fmt_@Op__'[0,11]{'@Fmt_'[1,6]}{'@Op_'[6,10]}}",
			expectedDot: stringutil.Indent(`
				digraph g {
					"@@Fmt_@Op__"
					"@@Fmt_@Op__" -> "@Fmt_" [label = "[1, 6]"]
					"@@Fmt_@Op__" -> "@Op_" [label = "[6, 10]"]
				}
			`),
			expectedResult: "@@Fmt_@Op__",
		},
		{
			name:           "Change1",
			template:       "@Fmt@Op__",
			assignments:    testAssignment,
			expectedString: "{@FmtReCode_}",
			expectedDebug:  "{'@FmtReCode_'[0,9]}",
			expectedDot: stringutil.Indent(`
				digraph g {
					"@FmtReCode_"
				}
			`),
			expectedResult: "@FmtReCode_",
		},
		{
			name:           "Change1.1",
			template:       "@@Op@Fmt___",
			assignments:    testAssignment,
			expectedString: "{@@OpTORPOR__{@OpTORPOR_}}",
			expectedDebug:  "{'@@OpTORPOR__'[0,11]{'@OpTORPOR_'[1,11]}}",
			expectedDot: stringutil.Indent(`
				digraph g {
					"@@OpTORPOR__"
					"@@OpTORPOR__" -> "@OpTORPOR_" [label = "[1, 11]"]
				}
			`),
			expectedResult: "@@OpTORPOR__",
		},
		{
			name:           "Change2",
			template:       "@@Fmt_@Op__",
			assignments:    testAssignment,
			expectedString: "{For_Great_Justice!}",
			expectedDebug:  "{'For_Great_Justice!'[0,11]}",
			expectedDot: stringutil.Indent(`
					digraph g {
						"For_Great_Justice!"
					}
				`),
			expectedResult: "For_Great_Justice!",
		},
		{
			name:           "Change3",
			template:       "@@Fmt_@Op_@_@Mod__",
			assignments:    testAssignment,
			expectedString: "{ZIG!}",
			expectedDebug:  "{'ZIG!'[0,18]}",
			expectedDot: stringutil.Indent(`
					digraph g {
						"ZIG!"
					}
				`),
			expectedResult: "ZIG!",
		},
		{
			name:           "Change3.1",
			template:       "@@Fmt_@Op_@Mod__",
			assignments:    testAssignment,
			expectedString: "{ZigZigZig}",
			expectedDebug:  "{'ZigZigZig'[0,16]}",
			expectedDot: stringutil.Indent(`
					digraph g {
						"ZigZigZig"
					}
				`),
			expectedResult: "ZigZigZig",
		},
	}

	t.Run("InPlace", func(t *testing.T) {
		for ti, tt := range tests {
			t.Run(fmt.Sprint(ti), func(t *testing.T) {
				must := require.New(t)
				node, _, _ := ParseNode(String(tt.template))
				must.NotPanics(func() {
					changed := node.ExecuteInPlace(tt.assignments)
					must.Equal(tt.expectedNoChange, !changed)

					are := assert.New(t)
					DoOrDie(are, tt.expectedString, node.String(), fmt.Sprintf("NODE_%d_%s_%d", ti, tt.name, 0))
					DoOrDie(are, tt.expectedDebug, node.Debug(), fmt.Sprintf("NODE_%d_%s_%d", ti, tt.name, 1))
					DoOrDie(are, tt.expectedDot, node.Dot(), fmt.Sprintf("NODE_%d_%s_%d", ti, tt.name, 2))
					DoOrDie(are, tt.expectedResult, node.Text.String(), fmt.Sprintf("NODE_%d_%s_%d", ti, tt.name, 3))
				})
			})
		}
	})

	t.Run("Execute", func(t *testing.T) {
		for ti, tt := range tests {
			t.Run(fmt.Sprint(ti), func(t *testing.T) {
				must := require.New(t)
				node, _, _ := ParseNode(String(tt.template))
				must.NotPanics(func() {
					result, newNode := node.Execute(tt.assignments)
					must.Equal(tt.template, node.Text.String())
					must.Equal(tt.expectedResult, result)

					are := assert.New(t)
					DoOrDie(are, tt.expectedString, newNode.String(), fmt.Sprintf("NODE_%d_%s_%d", ti, tt.name, 4))
					DoOrDie(are, tt.expectedResult, newNode.Text.String(), fmt.Sprintf("NODE_%d_%s_%d", ti, tt.name, 5))
				})
			})
		}
	})
}

func Test_Execute_Template(t *testing.T) {
	var testAssignment Assignments = make(Assignments)
	testAssignment.Assign("Fmt", "TORPOR")
	testAssignment.Assign("Op", "ReCode")
	testAssignment.Assign("Mod", "NonLocal")
	testAssignment.Assign("TORPORReCode", "For_Great_Justice!")
	testAssignment.Assign("TORPORReCodeNonLocal", "ZigZigZig")
	testAssignment.Assign("TORPORReCode_NonLocal", "ZIG!")
	testAssignment.Assign("Replace", "Maybe @Op?")

	tests := []struct {
		name             string
		template         string
		assignments      Assignments
		expectedNoChange bool
		expectedString   string
		expectedDebug    string
		expectedDot      string
		expectedResult   string
	}{
		{
			name:             "NoChange",
			template:         "Hello @: @@Fmt_@Op bob",
			assignments:      Assignments{},
			expectedNoChange: true,
			expectedString: stringutil.Indent(`
				Template:
					Hello @: @@Fmt_@Op bob
				Variables:
					{@@Fmt_@Op__{@Fmt_}{@Op_}}
			`),
			expectedDebug: stringutil.Indent(`
				Template:
					Hello @: @@Fmt_@Op bob
				Variables:
					{'@@Fmt_@Op__'[9,18]{'@Fmt_'[1,6]}{'@Op_'[6,10]}}
			`),
			expectedDot: stringutil.Indent(`
				digraph g {
					"tmpl root" [label="Hello @: @@Fmt_@Op bob", shape=diamond]
				
					"tmpl root" -> "@@Fmt_@Op__" [label = "[9, 18]"]
					"@@Fmt_@Op__" -> "@Fmt_" [label = "[1, 6]"]
					"@@Fmt_@Op__" -> "@Op_" [label = "[6, 10]"]
				}				
			`),
			expectedResult: `Hello @: @@Fmt_@Op bob`,
		},
		{
			name:        "Escape",
			template:    "o@_o",
			assignments: testAssignment,
			expectedString: stringutil.Indent(`
				Template:
					o_o
				Variables:
			`),
			expectedDebug: stringutil.Indent(`
				Template:
					o_o
				Variables:
			`),
			expectedDot: stringutil.Indent(`
				digraph g {
					"tmpl root" [label="o_o", shape=diamond]
				}
			`),
			expectedResult: "o_o",
		},
		{
			name:        "Escape2",
			template:    "->@Replace<-",
			assignments: testAssignment,
			expectedString: stringutil.Indent(`
				Template:
					->Maybe ReCode?<-
				Variables:
			`),
			expectedDebug: stringutil.Indent(`
			Template:
				->Maybe ReCode?<-
			Variables:
			`),
			expectedDot: stringutil.Indent(`
				digraph g {
					"tmpl root" [label="->Maybe ReCode?<-", shape=diamond]
				}
			`),
			expectedResult: "->Maybe ReCode?<-",
		},
		{
			name:        "Change1",
			template:    "func @Fmt@Op()",
			assignments: testAssignment,
			expectedString: stringutil.Indent(`
					Template:
						func @FmtReCode_()
					Variables:
						{@FmtReCode_}
				`),
			expectedDebug: stringutil.Indent(`
					Template:
						func @FmtReCode_()
					Variables:
						{'@FmtReCode_'[5,16]}
				`),
			expectedDot: stringutil.Indent(`
					digraph g {
						"tmpl root" [label="func @FmtReCode_()", shape=diamond]
					
						"tmpl root" -> "@FmtReCode_" [label = "[5, 16]"]
					}
				`),
			expectedResult: "func @FmtReCode_()",
		},
		{
			name:        "Change1.1",
			template:    "func @@Op@Fmt()",
			assignments: testAssignment,
			expectedString: stringutil.Indent(`
				Template:
					func @@OpTORPOR__()
				Variables:
					{@@OpTORPOR__{@OpTORPOR_}}
			`),
			expectedDebug: stringutil.Indent(`
				Template:
					func @@OpTORPOR__()
				Variables:
					{'@@OpTORPOR__'[5,17]{'@OpTORPOR_'[1,11]}}
			`),
			expectedDot: stringutil.Indent(`
				digraph g {
					"tmpl root" [label="func @@OpTORPOR__()", shape=diamond]
				
					"tmpl root" -> "@@OpTORPOR__" [label = "[5, 17]"]
					"@@OpTORPOR__" -> "@OpTORPOR_" [label = "[1, 11]"]
				}
			`),
			expectedResult: "func @@OpTORPOR__()",
		},
		{
			name:        "Change2",
			template:    "func @@Fmt_@Op()",
			assignments: testAssignment,
			expectedString: stringutil.Indent(`
				Template:
					func For_Great_Justice!()
				Variables:
			`),
			expectedDebug: stringutil.Indent(`
				Template:
					func For_Great_Justice!()
				Variables:
			`),
			expectedDot: stringutil.Indent(`
				digraph g {
					"tmpl root" [label="func For_Great_Justice!()", shape=diamond]
				}
			`),
			expectedResult: "func For_Great_Justice!()",
		},
		{
			name:        "ChangeChain",
			template:    "(@Fmt_@Op_@_@Mod)",
			assignments: testAssignment,
			expectedString: stringutil.Indent(`
				Template:
					(TORPORReCode_NonLocal)
				Variables:
			`),
			expectedDebug: stringutil.Indent(`
				Template:
					(TORPORReCode_NonLocal)
				Variables:
			`),
			expectedDot: stringutil.Indent(`
				digraph g {
					"tmpl root" [label="(TORPORReCode_NonLocal)", shape=diamond]
				}
			`),
			expectedResult: "(TORPORReCode_NonLocal)",
		},
	}

	t.Run("InPlace", func(t *testing.T) {
		for ti, tt := range tests {
			t.Run(fmt.Sprint(ti), func(t *testing.T) {
				tmpl := ParseTemplate(tt.template)
				must := require.New(t)
				must.NotPanics(func() {
					changed := tmpl.ExecuteInPlace(tt.assignments)
					must.Equal(tt.expectedNoChange, !changed)

					are := assert.New(t)
					DoOrDie(are, tt.expectedString, tmpl.String(), fmt.Sprintf("TMPL_%d_%s_%d", ti, tt.name, 0))
					DoOrDie(are, tt.expectedDebug, tmpl.Debug(), fmt.Sprintf("TMPL_%d_%s_%d", ti, tt.name, 1))
					DoOrDie(are, tt.expectedDot, tmpl.Dot(), fmt.Sprintf("TMPL_%d_%s_%d", ti, tt.name, 2))
					DoOrDie(are, tt.expectedResult, tmpl.Text.String(), fmt.Sprintf("TMPL_%d_%s_%d", ti, tt.name, 3))
				})
			})
		}
	})

	t.Run("Execute", func(t *testing.T) {
		for ti, tt := range tests {
			t.Run(fmt.Sprint(ti), func(t *testing.T) {
				tmpl := ParseTemplate(tt.template)
				must := require.New(t)
				must.NotPanics(func() {
					result, newTmpl := tmpl.Execute(tt.assignments)
					must.Equal(tt.template, tmpl.Text.String())
					must.Equal(tt.expectedResult, result)

					are := assert.New(t)
					DoOrDie(are, tt.expectedString, newTmpl.String(), fmt.Sprintf("TMPL_%d_%s_%d", ti, tt.name, 4))
					DoOrDie(are, tt.expectedResult, newTmpl.Text.String(), fmt.Sprintf("TMPL_%d_%s_%d", ti, tt.name, 5))
				})
			})
		}
	})
}

func Test_Execute_Templates(t *testing.T) {
	testTmpls_simple := map[string]string{
		"FuncName": "@Op_@Fmt_@Mod",
		"FuncArgs": "(@InArg_@OptArg_@OutArg)",
		"FuncRets": "(@OptRet_@OutRet_@ErrRet)",
		".":        "func @FuncName_@FuncArgs @FuncRets {}",
	}

	testAssignments_simple_A := Assignments{
		"Op":     "ReCode",
		"Fmt":    "TORPOR",
		"Mod":    "NonLocal",
		"InArg":  "n int, ",
		"OptArg": "",
		"OutArg": "Polo@_o",
		"OptRet": "X, ",
		"OutRet": "@Op_rRet, ",
		"ErrRet": "Error",
	}

	testAssignments_simple_B := Assignments{
		"Op":     "Operate|",
		"Fmt":    "zz",
		"Mod":    "@B@Fmt__",
		"Bzz":    "Abort|",
		"InArg":  "n int, ",
		"OptArg": "@Mod_@OutRet, ",
		"OutArg": "Error",
		"OptRet": "Out, ",
		"OutRet": "out@Op__",
		"ErrRet": "",
	}

	tests := []struct {
		templates   map[string]string
		assignments Assignments

		expectedString string
		expectedDebug  string
		expectedDot    string

		expectedDefaultResult string
		expectedRootResult    string
		expectedTmplResult    string
	}{
		{
			templates:   testTmpls_simple,
			assignments: testAssignments_simple_A,
			expectedString: stringutil.Indent(`
				.:
					Template:
						func ReCodeTORPORNonLocal(n int, Polo_o) (X, ReCoderRet, Error) {}
					Variables:

				FuncArgs:
					Template:
						(n int, Polo_o)
					Variables:
				
				FuncName:
					Template:
						ReCodeTORPORNonLocal
					Variables:
				
				FuncRets:
					Template:
						(X, ReCoderRet, Error)
					Variables:
			`),
			expectedDebug: stringutil.Indent(`
				.:
					Template:
						func ReCodeTORPORNonLocal(n int, Polo_o) (X, ReCoderRet, Error) {}
					Variables:

				FuncArgs:
					Template:
						(n int, Polo_o)
					Variables:
				
				FuncName:
					Template:
						ReCodeTORPORNonLocal
					Variables:
				
				FuncRets:
					Template:
						(X, ReCoderRet, Error)
					Variables:
			`),
			expectedDot: stringutil.Indent(`
				digraph g {

					"." [shape=rect]
					"." -> "tmpl root 0" [color=red]
					"tmpl root 0" [label="func ReCodeTORPORNonLocal(n int, Polo_o) (X, ReCoderRet, Error) {}", shape=diamond]
				
					"FuncArgs" [shape=rect]
					"FuncArgs" -> "tmpl root 1" [color=red]
					"tmpl root 1" [label="(n int, Polo_o)", shape=diamond]
				
					"FuncName" [shape=rect]
					"FuncName" -> "tmpl root 2" [color=red]
					"tmpl root 2" [label="ReCodeTORPORNonLocal", shape=diamond]
				
					"FuncRets" [shape=rect]
					"FuncRets" -> "tmpl root 3" [color=red]
					"tmpl root 3" [label="(X, ReCoderRet, Error)", shape=diamond]
				}
			`),
			expectedDefaultResult: `func ReCodeTORPORNonLocal(n int, Polo_o) (X, ReCoderRet, Error) {}`,
			expectedRootResult: stringutil.Indent(`
				?
			`),
			expectedTmplResult: stringutil.Indent(`
				?
			`),
		},
		{
			templates:   testTmpls_simple,
			assignments: testAssignments_simple_B,
			expectedString: stringutil.Indent(`
				.:
					Template:
						func Operate|zzAbort|(n int, Abort|outOperate|_, Error) (Out, outOperate|_) {}
					Variables:

				FuncArgs:
					Template:
						(n int, Abort|outOperate|_, Error)
					Variables:
				
				FuncName:
					Template:
						Operate|zzAbort|
					Variables:
				
				FuncRets:
					Template:
						(Out, outOperate|_)
					Variables:
			`),
			expectedDebug: stringutil.Indent(`
				.:
					Template:
						func Operate|zzAbort|(n int, Abort|outOperate|_, Error) (Out, outOperate|_) {}
					Variables:

				FuncArgs:
					Template:
						(n int, Abort|outOperate|_, Error)
					Variables:

				FuncName:
					Template:
						Operate|zzAbort|
					Variables:

				FuncRets:
					Template:
						(Out, outOperate|_)
					Variables:
			`),
			expectedDot: stringutil.Indent(`
				digraph g {

					"." [shape=rect]
					"." -> "tmpl root 0" [color=red]
					"tmpl root 0" [label="func Operate|zzAbort|(n int, Abort|outOperate|_, Error) (Out, outOperate|_) {}", shape=diamond]
				
					"FuncArgs" [shape=rect]
					"FuncArgs" -> "tmpl root 1" [color=red]
					"tmpl root 1" [label="(n int, Abort|outOperate|_, Error)", shape=diamond]
				
					"FuncName" [shape=rect]
					"FuncName" -> "tmpl root 2" [color=red]
					"tmpl root 2" [label="Operate|zzAbort|", shape=diamond]
				
					"FuncRets" [shape=rect]
					"FuncRets" -> "tmpl root 3" [color=red]
					"tmpl root 3" [label="(Out, outOperate|_)", shape=diamond]
				}
			`),
			expectedDefaultResult: `func Operate|zzAbort|(n int, Abort|outOperate|_, Error) (Out, outOperate|_) {}`,
			expectedRootResult: stringutil.Indent(`
				?
			`),
			expectedTmplResult: stringutil.Indent(`
				?
			`),
		},
	}
	t.Run("InPlace", func(t *testing.T) {
		for ti, tt := range tests {
			t.Run(fmt.Sprint(ti), func(t *testing.T) {
				tmpls := ParseTemplates(tt.templates)
				must := require.New(t)
				must.NotPanics(func() {
					changed := tmpls.ExecuteInPlace(tt.assignments)
					must.True(changed)

					are := assert.New(t)
					DoOrDie(are, tt.expectedString, tmpls.String(), fmt.Sprintf("IP_TMPLS_%d_%d", ti, 0))
					DoOrDie(are, tt.expectedDebug, tmpls.Debug(), fmt.Sprintf("IP_TMPLS_%d_%d", ti, 1))
					DoOrDie(are, tt.expectedDot, tmpls.Dot(), fmt.Sprintf("IP_TMPLS_%d_%d", ti, 2))

				})
			})
		}
	})

	t.Run("Execute", func(t *testing.T) {
		for ti, tt := range tests {
			t.Run(fmt.Sprint(ti), func(t *testing.T) {
				tmpls := ParseTemplates(tt.templates)
				must := require.New(t)
				must.NotPanics(func() {
					result, newTmpls := tmpls.Execute(tt.assignments)

					are := assert.New(t)
					DoOrDie(are, tt.expectedString, newTmpls.String(), fmt.Sprintf("E_TMPLS_%d_%d", ti, 0))
					DoOrDie(are, tt.expectedDebug, newTmpls.Debug(), fmt.Sprintf("E_TMPLS_%d_%d", ti, 1))
					DoOrDie(are, tt.expectedDot, newTmpls.Dot(), fmt.Sprintf("E_TMPLS_%d_%d", ti, 2))
					DoOrDie(are, tt.expectedDefaultResult, result, fmt.Sprintf("E_TMPLS_%d_%d", ti, 3))

				})
			})
		}
	})
}
