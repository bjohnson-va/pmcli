package mirage

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type copyByPathsTestSuite struct {
	suite.Suite
}

func TestCopyByPathsTestSuite(t *testing.T) {
	suite.Run(t, &copyByPathsTestSuite{})
}

func (suite *copyByPathsTestSuite) TestCopiesNestedPaths() {
	paths := []string{"pointer.nested_int", "string_field", "pointer.deeply_nested.end"}
	src := &TestStruct{
		StringField: "stringtest",
		PointerField: &NestedStruct{
			NestedIntField: 42,
			DeepNestedField: &DeeplyNestedStruct{
				End: "end of all things",
			},
		},
	}
	dest := &TestStruct{PointerField: &NestedStruct{DeepNestedField: &DeeplyNestedStruct{}}}
	err := CopyByPaths("mirage", paths, src, dest)
	suite.Nil(err)
	suite.Equal(src, dest)
}

func (suite *copyByPathsTestSuite) TestCopiesNestedPathsWhereDestIsNil() {
	paths := []string{"pointer.nested_int", "string_field", "pointer.deeply_nested.end"}
	src := &TestStruct{
		StringField: "stringtest",
		PointerField: &NestedStruct{
			NestedIntField: 42,
			DeepNestedField: &DeeplyNestedStruct{
				End: "end of all things",
			},
		},
	}
	dest := &TestStruct{PointerField: nil}
	err := CopyByPaths("mirage", paths, src, dest)
	suite.Nil(err)
	suite.Equal(src, dest)
}

func (suite *copyByPathsTestSuite) TestCopiesEachFieldFromFieldMask() {
	src := &TestStruct{
		StringField: "hello",
		IntField:    42,
		BoolField:   true,
		PointerField: &NestedStruct{
			NestedStringField: "nestedPointer",
		},
		StringSlice: []string{"one", "two"},
		NestedField: NestedStruct{
			NestedStringField: "nestedStruct",
		},
	}
	fieldMask := []string{
		"string_field",
		"pointer",
		"string_slice_field",
	}

	expected := &TestStruct{
		StringField: "hello",
		PointerField: &NestedStruct{
			NestedStringField: "nestedPointer",
		},
		StringSlice: []string{"one", "two"},
	}
	dest := &TestStruct{}

	err := CopyByPaths("mirage", fieldMask, src, dest)

	suite.Nil(err)
	suite.Equal(expected, dest)
}
