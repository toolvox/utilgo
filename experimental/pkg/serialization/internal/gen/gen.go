package main

import (
	"fmt"
	"strings"

	"utilgo/pkg/sets"
	"utilgo/pkg/tmplz"
)

type Format = string

const (
	Json Format = "Json"
	Yaml Format = "Yaml"
)

var AllFormats = []Format{Json, Yaml}

type Operation = string

const (
	Unmarshal Operation = "Unmarshal"
	Marshal   Operation = "Marshal"
	Decode    Operation = "Decode"
	Encode    Operation = "Encode"
)

var AllOperations = []Operation{Unmarshal, Marshal, Decode, Encode}

type Target = string

const (
	Mem    Target = ""
	File   Target = "File"
	FS     Target = "FS"
	Reader Target = "Reader"
	Writer Target = "Writer"
)

var InitialTargets = []Target{Mem, File}

type Option = string

const (
	Must   Option = "Must"
	Indent Option = "Indent"
	Valid  Option = "Valid"
	Into   Option = "Into"
)

func IsPrefix(o Option) bool { return o == Must }
func IsInfix(o Option) bool  { return o == Valid }
func IsSuffix(o Option) bool { return o == Indent || o == Into }

type OptionSet []Option

const (
	Prefix = iota
	Infix
	Suffix
)

func (s OptionSet) Affixes() [3]OptionSet {
	res := [3]OptionSet{}
	if len(s) == 0 {
		return res
	}
	for _, o := range s {
		switch {
		case IsPrefix(o):
			res[Prefix] = append(res[Prefix], o)
		case IsInfix(o):
			res[Infix] = append(res[Infix], o)
		case IsSuffix(o):
			res[Suffix] = append(res[Suffix], o)
		}
	}
	return res
}

func (s OptionSet) String() string {
	var sb strings.Builder
	for _, v := range s {
		sb.Write([]byte(v))
	}
	return sb.String()
}

const (
	Go = iota
	External
	Mine
	This
)

var ImportGroupNames = [4]string{"Go", "External", "Mine", "This"}

type Imports struct {
	Sets [4]sets.TinySet[string]
}

func (i Imports) Clone() Imports {
	var clone Imports
	for index, tinySet := range i.Sets {
		clone.Sets[index] = sets.NewTinySet[string](tinySet...)
	}
	return clone
}

func (i Imports) Assign(assignments tmplz.Assignments) {
	for si, set := range i.Sets {
		var sb strings.Builder
		for _, imp := range set {
			sb.WriteString(fmt.Sprintf("\"%s\"\n", imp))
		}
		assignments.Assign(fmt.Sprintf("Imports%s", ImportGroupNames[si]), sb.String())
	}
}
