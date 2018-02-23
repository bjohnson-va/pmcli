package mirage

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type copyByPathTestSuite struct {
	suite.Suite
}

func TestCopyByPathTestSuite(t *testing.T) {
	suite.Run(t, &copyByPathTestSuite{})
}

type TestStruct struct {
	StringField  string        `mirage:"string_field"`
	IntField     int           `mirage:"int_field"`
	BoolField    bool          `mirage:"bool_field"`
	StringSlice  []string      `mirage:"string_slice_field"`
	NestedField  NestedStruct  `mirage:"nested_struct"`
	PointerField *NestedStruct `mirage:"pointer"`
}

type NestedStruct struct {
	NestedStringField string              `mirage:"nested_string_field"`
	NestedIntField    int                 `mirage:"nested_int"`
	DeepNestedField   *DeeplyNestedStruct `mirage:"deeply_nested"`
}

type DeeplyNestedStruct struct {
	End string `mirage:"end"`
}

func (suite *copyByPathTestSuite) TestReturnsErrorIfPassedNonPointerForSrcOrDest() {
	err := CopyByPath("mirage", "string_field", TestStruct{}, &TestStruct{})
	suite.NotNil(err)

	err = CopyByPath("mirage", "string_field", &TestStruct{}, TestStruct{})
	suite.NotNil(err)
}

func (suite *copyByPathTestSuite) TestReturnsErrorIfNilAsSrcOrDest() {
	err := CopyByPath("mirage", "string_field", nil, &TestStruct{})
	suite.NotNil(err)

	err = CopyByPath("mirage", "string_field", &TestStruct{}, nil)
	suite.NotNil(err)
}

type copyFieldTestCase struct {
	src       *TestStruct
	tagToCopy string
}

func (suite *copyByPathTestSuite) TestCopiesBasicFields() {
	testCases := []copyFieldTestCase{
		{&TestStruct{StringField: "test"}, "string_field"},
		{&TestStruct{IntField: 42}, "int_field"},
		{&TestStruct{StringSlice: []string{"one", "two"}}, "string_slice_field"},
		{&TestStruct{BoolField: true}, "bool_field"},
		{&TestStruct{NestedField: NestedStruct{NestedStringField: "nested value"}}, "nested_struct"},
		{&TestStruct{PointerField: &NestedStruct{NestedStringField: "nested value"}}, "pointer"},
	}

	for _, testCase := range testCases {
		dest := &TestStruct{}
		err := CopyByPath("mirage", testCase.tagToCopy, testCase.src, dest)
		suite.Nil(err)
		suite.Equal(testCase.src, dest)
	}
}

func (suite *copyByPathTestSuite) TestOnlyCopiesFieldsMatchingKey() {
	src := &TestStruct{
		StringField: "hello",
		IntField:    42,
		BoolField:   true,
		PointerField: &NestedStruct{
			NestedStringField: "nestedPointer",
			NestedIntField:    11,
		},
		StringSlice: []string{"one", "two"},
		NestedField: NestedStruct{
			NestedStringField: "nestedStruct",
		},
	}
	dest := &TestStruct{}
	expected := &TestStruct{
		PointerField: &NestedStruct{
			NestedIntField: 11,
		},
	}

	err := CopyByPath("mirage", "pointer.nested_int", src, dest)

	suite.Nil(err)
	suite.Equal(expected, dest)
}

func (suite *copyByPathTestSuite) TestReturnsErrorIfFieldNotFound() {
	src := &TestStruct{}
	dest := &TestStruct{}

	err := CopyByPath("mirage", "not_found", src, dest)
	suite.Error(err)
	suite.Equal(err.Error(), "could not find field `mirage:\"not_found\" in type mirage.TestStruct")
}

type DiffStruct struct {
	FieldLabelledWithWrongType int               `mirage:"bool_field"`
	DiffNestedStruct           *DiffNestedStruct `mirage:"pointer"`
}

type DiffNestedStruct struct {
	DiffNestedStringField string `mirage:"nested_string_field"`
}

func (suite *copyByPathTestSuite) TestCopiesBetweenDifferingStructTypes() {
	src := &TestStruct{
		PointerField: &NestedStruct{
			NestedStringField: "word",
		},
	}
	dest := &DiffStruct{}

	err := CopyByPath("mirage", "pointer.nested_string_field", src, dest)
	suite.Nil(err)
	suite.Equal(src.PointerField.NestedStringField, dest.DiffNestedStruct.DiffNestedStringField)
}

func (suite *copyByPathTestSuite) TestReturnsErrorIfValuesAtEndOfPathsAreNotSameType() {
	src := &TestStruct{
		BoolField: true,
	}
	dest := &DiffStruct{}

	err := CopyByPath("mirage", "bool_field", src, dest)
	suite.NotNil(err)
	suite.Equal("bool_field > cannot copy differing types src: bool, dest: int", err.Error())
}

func (suite *copyByPathTestSuite) TestReturnsErrorIfPathTraversesNilInSrc() {
	src := &TestStruct{}
	dest := &TestStruct{}

	err := CopyByPath("mirage", "pointer.nested_string_field", src, dest)
	suite.NotNil(err)
	suite.Equal("encountered nil along path in src at key \"pointer\"", err.Error())
}

func (suite *copyByPathTestSuite) TestFillsNilsWithEmptyAlongPathInDest() {
	src := &TestStruct{
		PointerField: &NestedStruct{
			NestedStringField: "needle",
		},
	}
	dest := &TestStruct{PointerField: nil}

	err := CopyByPath("mirage", "pointer.nested_string_field", src, dest)
	suite.Nil(err)
	suite.Equal(src, dest)
}

func (suite *copyByPathTestSuite) TestCopiesPointers() {
	src := &TestStruct{
		PointerField: &NestedStruct{
			NestedStringField: "needle",
		},
	}
	dest := &TestStruct{PointerField: nil}

	err := CopyByPath("mirage", "pointer", src, dest)
	suite.Nil(err)
	suite.Equal(src, dest)
}

func (suite *copyByPathTestSuite) TestCopiesNilPointersIfAtEndOfPath() {
	src := &TestStruct{PointerField: nil}
	dest := &TestStruct{PointerField: &NestedStruct{}}

	err := CopyByPath("mirage", "pointer", src, dest)
	suite.Nil(err)
	suite.Equal(src, dest)
}
