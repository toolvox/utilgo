package reflectutil_test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"utilgo/pkg/reflectutil"
)

func Test_IsZero(t *testing.T) {
	t.Run("Sanity", func(t *testing.T) {
		t.Run("int", func(t *testing.T) {
			var v int
			assert.True(t, reflectutil.IsZero(v))
			v = 1
			assert.False(t, reflectutil.IsZero(v))
			v = 0
			assert.True(t, reflectutil.IsZero(v))
		})
		t.Run("complex64", func(t *testing.T) {
			var v complex64
			assert.True(t, reflectutil.IsZero(v))
			v = 1
			assert.False(t, reflectutil.IsZero(v))
			v = 1i
			assert.False(t, reflectutil.IsZero(v))
			v = 3 - 1i
			assert.False(t, reflectutil.IsZero(v))
			v = real(0 - 1i)
			assert.True(t, reflectutil.IsZero(v))
		})
		t.Run("string", func(t *testing.T) {
			var v string
			assert.True(t, reflectutil.IsZero(v))
			v = "asd"
			assert.False(t, reflectutil.IsZero(v))
			v = ""
			assert.True(t, reflectutil.IsZero(v))
		})
		t.Run("slice", func(t *testing.T) {
			var v []string
			assert.True(t, reflectutil.IsZero(v))
			v = []string{"asd", "bsd", "asx"}
			assert.False(t, reflectutil.IsZero(v))
			v = []string{}
			assert.True(t, reflectutil.IsZero(v))
			v = []string{"", "", "a"}
			assert.False(t, reflectutil.IsZero(v))
			v = []string{"", ""}
			assert.True(t, reflectutil.IsZero(v))
			v = []string{""}
			assert.True(t, reflectutil.IsZero(v))
		})
		t.Run("array", func(t *testing.T) {
			v0 := [0]uint16{}
			assert.True(t, reflectutil.IsZero(v0))
			v1 := [3]uint16{5, 7, 43}
			assert.False(t, reflectutil.IsZero(v1))
			v2 := [1]uint16{}
			assert.True(t, reflectutil.IsZero(v2))
			v3 := [4]uint16{0, 5, 0}
			assert.False(t, reflectutil.IsZero(v3))
			v4 := [3]uint16{0, 0}
			assert.True(t, reflectutil.IsZero(v4))
			v5 := [1]uint16{0}
			assert.True(t, reflectutil.IsZero(v5))
		})
		t.Run("func", func(t *testing.T) {
			var v func(int)
			assert.True(t, reflectutil.IsZero(v))
			v = func(i int) {}
			assert.False(t, reflectutil.IsZero(v))
			v = func(i int) { i += i }
			assert.False(t, reflectutil.IsZero(v))
			v = nil
			assert.True(t, reflectutil.IsZero(v))
		})
		t.Run("struct", func(t *testing.T) {
			type testStruct struct{ Field string }
			var v testStruct
			assert.True(t, reflectutil.IsZero(v))
			v = testStruct{Field: "value"}
			assert.False(t, reflectutil.IsZero(v))
			v = testStruct{}
			assert.True(t, reflectutil.IsZero(v))
		})
		t.Run("map", func(t *testing.T) {
			var v map[string]int
			assert.True(t, reflectutil.IsZero(v))
			v = make(map[string]int)
			v["key0"] = 0
			assert.True(t, reflectutil.IsZero(v))
			v["key"] = 1
			assert.False(t, reflectutil.IsZero(v))
			delete(v, "key")
			assert.True(t, reflectutil.IsZero(v))
		})
		t.Run("misc", func(t *testing.T) {
			var v fmt.Stringer
			assert.True(t, reflectutil.IsZero(v))
			v = reflect.ValueOf(v)
			assert.True(t, reflectutil.IsZero(v))
			v = reflect.ValueOf(v)
			assert.True(t, reflectutil.IsZero(v))
			v = nil
			assert.True(t, reflectutil.IsZero(v))
			v = &strings.Builder{}
			assert.True(t, reflectutil.IsZero(v))
			v.(*strings.Builder).WriteString("ASD")
			assert.False(t, reflectutil.IsZero(v))
		})
		t.Run("interface", func(t *testing.T) {
			var v Xer
			assert.True(t, reflectutil.IsZero(v))
			v = ZX{}
			assert.True(t, reflectutil.IsZero(v))
			v = &ZX2{4}
			assert.False(t, reflectutil.IsZero(v))
			v = &ZX2{}
			assert.True(t, reflectutil.IsZero(v))
		})
		t.Run("chan", func(t *testing.T) {
			var v chan string
			assert.True(t, reflectutil.IsZero(v))
			v = make(chan string, 2)
			assert.True(t, reflectutil.IsZero(v))
			v <- "asd"
			assert.False(t, reflectutil.IsZero(v))
			<-v
			assert.True(t, reflectutil.IsZero(v))
		})
	})
}

type Xer interface{ X() }

type ZX struct{}

func (x ZX) X() {}

type ZX2 struct{ Num int }

func (x *ZX2) X() {}
