package serialization_test

import (
	"fmt"
	
	"utilgo/pkg/errs"
	"utilgo/pkg/stringutil"
)

type subObject struct {
	Flag bool `yaml:"Flag" json:"Flag"`
}

type TestObject struct {
	Name      string     `yaml:"Name" json:"Name"`
	Value     int        `yaml:"Value" json:"Value"`
	SubObject *subObject `yaml:"SubObject,omitempty" json:"SubObject,omitempty"`
}

func (o TestObject) Validate() error {
	var errs errs.Errors
	if o.Name == "Illegal" {
		errs = append(errs, fmt.Errorf("'Name' must not be illegal, was: %s", o.Name))
	}
	if o.Value < 0 {
		errs = append(errs, fmt.Errorf("'Value' must not be negative, was: %d", o.Value))
	}
	return errs.OrNil()
}

type IllegalObject struct {
	Self *IllegalObject `yaml:"Self" json:"Self"`
}

type TestDatum[T any] struct {
	Object     T
	Yaml, Json string
}

type TestDataDefs struct {
	Valid   TestDatum[TestObject]
	Invalid TestDatum[TestObject]
	Error   TestDatum[TestObject]
	Illegal TestDatum[IllegalObject]
}

var TestData TestDataDefs

func init() {
	illegal := IllegalObject{}
	illegal.Self = &illegal
	TestData = TestDataDefs{
		Valid: TestDatum[TestObject]{
			Object: TestObject{
				Name:      "Legal",
				Value:     1,
				SubObject: &subObject{false},
			},
			Yaml: stringutil.Indent(`
				Name: Legal
				Value: 1
				SubObject:
					Flag: true
			`),
			Json: `
			
			`,
		},
		Invalid: TestDatum[TestObject]{
			Object: TestObject{
				Name:      "Illegal",
				Value:     -1,
				SubObject: &subObject{true},
			},
			Yaml: stringutil.Indent(`
				Name: Illegal
				Value: -1
				SubObject:
					Flag: false
			`),
			Json: `
			
			`,
		},
		Error: TestDatum[TestObject]{
			Object: TestObject{
				Name:      "",
				Value:     0,
				SubObject: nil,
			},
			Yaml: stringutil.Indent(`
				Name:nope
				Value: nope"Nope
					SubNope:
			`),
			Json: `
			
			`,
		},
		Illegal: TestDatum[IllegalObject]{
			Object: illegal,
			Yaml: stringutil.Indent(`
				# Nothing
			`),
			Json: `
			
			`,
		},
	}
}
