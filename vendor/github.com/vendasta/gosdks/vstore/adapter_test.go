package vstore

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"time"
	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/suite"
	"github.com/vendasta/gosdks/pb/vstorepb"
)

type I64Enum int64
type I32Enum int32
type IntEnum int
type F32Enum float32
type F64Enum float64
type StrEnum string
type BulEnum bool

const (
	I64 I64Enum = 0
	I32 I32Enum = 1
	Int IntEnum = 2
	F32 F32Enum = 3.0
	F64 F64Enum = 4.0
	Str StrEnum = "string"
	Bul BulEnum = false
)

func TestModelToStructPB_StringTypeHasExpectedOutput(t *testing.T) {
	s, err := ModelToStructPB(
		&struct{ Name string `vstore:"name"` }{Name: "Hello"},
	)
	assert.Nil(t, err)
	assert.Equal(t, &vstorepb.Struct{Values: map[string]*vstorepb.Value{"name": {&vstorepb.Value_StringValue{"Hello"}}}}, s)
}

func TestModelToStructPB_BytesTypeHasExpectedOutput(t *testing.T) {
	s, err := ModelToStructPB(
		&struct{ Name []byte `vstore:"name"` }{Name: []byte("Hello")},
	)
	assert.Nil(t, err)
	assert.Equal(t, &vstorepb.Struct{Values: map[string]*vstorepb.Value{"name": {&vstorepb.Value_BytesValue{[]byte("Hello")}}}}, s)
}

func TestModelToStructPB_RepeatedStringTypeHasExpectedOutput(t *testing.T) {
	s, err := ModelToStructPB(
		&struct{ Name []string `vstore:"name"` }{Name: []string{"Hello"}},
	)
	assert.Nil(t, err)
	assert.Equal(t, &vstorepb.Struct{Values: map[string]*vstorepb.Value{"name": {&vstorepb.Value_ListValue{
		ListValue: &vstorepb.ListValue{[]*vstorepb.Value{{&vstorepb.Value_StringValue{"Hello"}}}}},
	}}}, s)
}

func TestModelToStructPB_IntTypeHasExpectedOutput(t *testing.T) {
	s, err := ModelToStructPB(
		&struct{ Name int `vstore:"name"` }{Name: 123},
	)
	assert.Nil(t, err)
	assert.Equal(t, &vstorepb.Struct{Values: map[string]*vstorepb.Value{"name": {&vstorepb.Value_IntValue{int64(123)}}}}, s)
}

func TestModelToStructPB_RepeatedIntTypeHasExpectedOutput(t *testing.T) {
	s, err := ModelToStructPB(
		&struct{ Name []int `vstore:"name"` }{Name: []int{123}},
	)
	assert.Nil(t, err)
	assert.Equal(t, &vstorepb.Struct{Values: map[string]*vstorepb.Value{"name": {&vstorepb.Value_ListValue{
		ListValue: &vstorepb.ListValue{[]*vstorepb.Value{{&vstorepb.Value_IntValue{123}}}}},
	}}}, s)
}

func TestModelToStructPB_Int64TypeHasExpectedOutput(t *testing.T) {
	s, err := ModelToStructPB(
		&struct{ Name int64 `vstore:"name"` }{Name: 123},
	)
	assert.Nil(t, err)
	assert.Equal(t, &vstorepb.Struct{Values: map[string]*vstorepb.Value{"name": {&vstorepb.Value_IntValue{int64(123)}}}}, s)
}

func TestModelToStructPB_Int32TypeHasExpectedOutput(t *testing.T) {
	s, err := ModelToStructPB(
		&struct{ Name int32 `vstore:"name"` }{Name: 123},
	)
	assert.Nil(t, err)
	assert.Equal(t, &vstorepb.Struct{Values: map[string]*vstorepb.Value{"name": {&vstorepb.Value_IntValue{int64(123)}}}}, s)
}

func TestModelToStructPB_Float64TypeHasExpectedOutput(t *testing.T) {
	s, err := ModelToStructPB(
		&struct{ Name float64 `vstore:"name"` }{Name: 123.321},
	)
	assert.Nil(t, err)
	assert.Equal(t, &vstorepb.Struct{Values: map[string]*vstorepb.Value{"name": {&vstorepb.Value_DoubleValue{123.321}}}}, s)
}

func TestModelToStructPB_RepeatedFloat64TypeHasExpectedOutput(t *testing.T) {
	s, err := ModelToStructPB(
		&struct{ Name []float64 `vstore:"name"` }{Name: []float64{123.321}},
	)
	assert.Nil(t, err)
	assert.Equal(t, &vstorepb.Struct{Values: map[string]*vstorepb.Value{"name": {&vstorepb.Value_ListValue{
		ListValue: &vstorepb.ListValue{[]*vstorepb.Value{{&vstorepb.Value_DoubleValue{123.321}}}}},
	}}}, s)
}

func TestModelToStructPB_Float32TypeHasExpectedOutput(t *testing.T) {
	s, err := ModelToStructPB(
		&struct{ Name float32 `vstore:"name"` }{Name: 123},
	)
	assert.Nil(t, err)
	assert.Equal(t, &vstorepb.Struct{Values: map[string]*vstorepb.Value{"name": {&vstorepb.Value_DoubleValue{123}}}}, s)
}

func TestModelToStructPB_BoolHasExpectedOutput(t *testing.T) {
	s, err := ModelToStructPB(
		&struct{ Name bool `vstore:"name"` }{Name: true},
	)
	assert.Nil(t, err)
	assert.Equal(t, &vstorepb.Struct{Values: map[string]*vstorepb.Value{"name": {&vstorepb.Value_BoolValue{true}}}}, s)
}

func TestModelToStructPB_RepeatedBoolTypeHasExpectedOutput(t *testing.T) {
	s, err := ModelToStructPB(
		&struct{ Name []bool `vstore:"name"` }{Name: []bool{false}},
	)
	assert.Nil(t, err)
	assert.Equal(t, &vstorepb.Struct{Values: map[string]*vstorepb.Value{"name": {&vstorepb.Value_ListValue{
		ListValue: &vstorepb.ListValue{[]*vstorepb.Value{{&vstorepb.Value_BoolValue{false}}}}},
	}}}, s)
}

func TestModelToStructPB_GeoWithValueHasExpectedOutput(t *testing.T) {
	geo := &vstorepb.GeoPoint{Latitude: 100, Longitude: 120}
	s, err := ModelToStructPB(
		&struct{ Name *vstorepb.GeoPoint `vstore:"geo"` }{Name: geo},
	)
	assert.Nil(t, err)
	assert.Equal(t, &vstorepb.Struct{Values: map[string]*vstorepb.Value{"geo": {&vstorepb.Value_GeopointValue{geo}}}}, s)
}

func TestModelToStructPB_RepeatedGeoTypeHasExpectedOutput(t *testing.T) {
	geo := &vstorepb.GeoPoint{Latitude: 100, Longitude: 120}
	s, err := ModelToStructPB(
		&struct{ Name []*vstorepb.GeoPoint `vstore:"name"` }{Name: []*vstorepb.GeoPoint{geo}},
	)
	assert.Nil(t, err)
	assert.Equal(t, &vstorepb.Struct{Values: map[string]*vstorepb.Value{"name": {&vstorepb.Value_ListValue{
		ListValue: &vstorepb.ListValue{[]*vstorepb.Value{{&vstorepb.Value_GeopointValue{geo}}}}},
	}}}, s)
}

func TestModelToStructPB_AliasedInt64EnumHasExpectedOutput(t *testing.T) {
	s, err := ModelToStructPB(
		&struct{ Name I64Enum `vstore:"my_enum"` }{Name: I64},
	)
	assert.Nil(t, err)
	assert.Equal(t, &vstorepb.Struct{Values: map[string]*vstorepb.Value{"my_enum": {&vstorepb.Value_IntValue{int64(I64)}}}}, s)
}

func TestModelToStructPB_RepeatedAliasedInt64TypeHasExpectedOutput(t *testing.T) {
	s, err := ModelToStructPB(
		&struct{ Name []I64Enum `vstore:"name"` }{Name: []I64Enum{4, 5, 6}},
	)
	assert.Nil(t, err)
	assert.Equal(t, &vstorepb.Struct{Values: map[string]*vstorepb.Value{"name": {&vstorepb.Value_ListValue{
		ListValue: &vstorepb.ListValue{[]*vstorepb.Value{{&vstorepb.Value_IntValue{int64(4)}}, {&vstorepb.Value_IntValue{int64(5)}}, {&vstorepb.Value_IntValue{int64(6)}}}}},
	}}}, s)
}

func TestModelToStructPB_AliasedInt32EnumHasExpectedOutput(t *testing.T) {
	s, err := ModelToStructPB(
		&struct{ Name I32Enum `vstore:"my_enum"` }{Name: I32},
	)
	assert.Nil(t, err)
	assert.Equal(t, &vstorepb.Struct{Values: map[string]*vstorepb.Value{"my_enum": {&vstorepb.Value_IntValue{int64(I32)}}}}, s)
}

func TestModelToStructPB_AliasedIntEnumHasExpectedOutput(t *testing.T) {
	s, err := ModelToStructPB(
		&struct{ Name IntEnum `vstore:"my_enum"` }{Name: Int},
	)
	assert.Nil(t, err)
	assert.Equal(t, &vstorepb.Struct{Values: map[string]*vstorepb.Value{"my_enum": {&vstorepb.Value_IntValue{int64(Int)}}}}, s)
}

func TestModelToStructPB_AliasedFloat32EnumHasExpectedOutput(t *testing.T) {
	s, err := ModelToStructPB(
		&struct{ Name F32Enum `vstore:"my_enum"` }{Name: F32},
	)
	assert.Nil(t, err)
	assert.Equal(t, &vstorepb.Struct{Values: map[string]*vstorepb.Value{"my_enum": {&vstorepb.Value_DoubleValue{float64(F32)}}}}, s)
}

func TestModelToStructPB_AliasedFloat64EnumHasExpectedOutput(t *testing.T) {
	s, err := ModelToStructPB(
		&struct{ Name F64Enum `vstore:"my_enum"` }{Name: F64},
	)
	assert.Nil(t, err)
	assert.Equal(t, &vstorepb.Struct{Values: map[string]*vstorepb.Value{"my_enum": {&vstorepb.Value_DoubleValue{float64(F64)}}}}, s)
}

func TestModelToStructPB_AliasedStringEnumHasExpectedOutput(t *testing.T) {
	s, err := ModelToStructPB(
		&struct{ Name StrEnum `vstore:"my_enum"` }{Name: Str},
	)
	assert.Nil(t, err)
	assert.Equal(t, &vstorepb.Struct{Values: map[string]*vstorepb.Value{"my_enum": {&vstorepb.Value_StringValue{string(Str)}}}}, s)
}

func TestModelToStructPB_AliasedBoolEnumHasExpectedOutput(t *testing.T) {
	s, err := ModelToStructPB(
		&struct{ Name BulEnum `vstore:"my_enum"` }{Name: Bul},
	)
	assert.Nil(t, err)
	assert.Equal(t, &vstorepb.Struct{Values: map[string]*vstorepb.Value{"my_enum": {&vstorepb.Value_BoolValue{bool(Bul)}}}}, s)
}

func TestModelToStructPB_NilGeoHasExpectedOutput(t *testing.T) {
	s, err := ModelToStructPB(
		&struct{ Name *vstorepb.GeoPoint `vstore:"geo"` }{nil},
	)
	assert.Nil(t, err)
	assert.Equal(t, &vstorepb.Struct{Values: map[string]*vstorepb.Value{}}, s)
}

func TestModelToStructPB_TimeWithValueHasExpectedOutput(t *testing.T) {
	now := time.Now().UTC()
	nowPb, _ := ptypes.TimestampProto(now)
	s, err := ModelToStructPB(
		&struct{ Name time.Time `vstore:"now"` }{Name: now},
	)
	assert.Nil(t, err)
	assert.Equal(t, &vstorepb.Struct{Values: map[string]*vstorepb.Value{"now": {&vstorepb.Value_TimestampValue{nowPb}}}}, s)
}

func TestModelToStructPB_TimeWithEmptyValueHasExpectedOutput(t *testing.T) {
	s, err := ModelToStructPB(
		&struct{ Name time.Time `vstore:"now"` }{},
	)
	assert.Nil(t, err)
	assert.Equal(t, &vstorepb.Struct{Values: map[string]*vstorepb.Value{}}, s)
}

func TestModelToStructPB_PointerToTimeWithEmptyValueHasExpectedOutput(t *testing.T) {
	s, err := ModelToStructPB(
		&struct{ T *time.Time `vstore:"now"` }{T: &time.Time{}},
	)
	assert.Nil(t, err)
	assert.Equal(t, &vstorepb.Struct{Values: map[string]*vstorepb.Value{}}, s)
}

func TestModelToStructPB_PointerToTimeWithValueHasExpectedOutput(t *testing.T) {
	now := time.Now().UTC()
	nowPb, _ := ptypes.TimestampProto(now)
	s, err := ModelToStructPB(
		&struct{ Name *time.Time `vstore:"now"` }{Name: &now},
	)
	assert.Nil(t, err)
	assert.Equal(t, &vstorepb.Struct{Values: map[string]*vstorepb.Value{"now": {&vstorepb.Value_TimestampValue{nowPb}}}}, s)
}

func TestModelToStructPB_PointerToNilTimeWithValueHasExpectedOutput(t *testing.T) {
	s, err := ModelToStructPB(
		&struct{ Name *time.Time `vstore:"now"` }{},
	)
	assert.Nil(t, err)
	assert.Equal(t, &vstorepb.Struct{Values: map[string]*vstorepb.Value{}}, s)
}

type TestStruct struct {
	Name string `vstore:"name"`
}

func TestModelToStructPB_ReferencedStructuredPropertyWithValueHasExpectedOutput(t *testing.T) {
	s, err := ModelToStructPB(
		&struct{ Test *TestStruct `vstore:"test"` }{Test: &TestStruct{Name: "hi"}},
	)
	assert.Nil(t, err)
	v := map[string]*vstorepb.Value{
		"test": {
			&vstorepb.Value_StructValue{
				&vstorepb.Struct{
					map[string]*vstorepb.Value{
						"name": {&vstorepb.Value_StringValue{"hi"}},
					},
				},
			},
		},
	}
	assert.Equal(t, &vstorepb.Struct{Values: v}, s)
}

func TestModelToStructPB_StructuredPropertyWithValueHasExpectedOutput(t *testing.T) {
	s, err := ModelToStructPB(
		&struct{ Test TestStruct `vstore:"test"` }{Test: TestStruct{Name: "hi"}},
	)
	assert.Nil(t, err)
	v := map[string]*vstorepb.Value{
		"test": {
			&vstorepb.Value_StructValue{
				&vstorepb.Struct{
					map[string]*vstorepb.Value{
						"name": {&vstorepb.Value_StringValue{"hi"}},
					},
				},
			},
		},
	}
	assert.Equal(t, &vstorepb.Struct{Values: v}, s)
}
func TestModelToStructPB_StructuredPropertyWithNilValueHasExpectedOutput(t *testing.T) {
	ts := &YATS{}
	s, err := ModelToStructPB(
		ts,
	)
	assert.Nil(t, err)
	assert.Equal(t, &vstorepb.Struct{Values: map[string]*vstorepb.Value{}}, s)
}

type YATS struct {
	*TestStruct `vstore:"TestStruct"`
}

func TestModelToStructPB_ReferencedDoubleNestedStructuredPropertyWithValueHasExpectedOutput(t *testing.T) {
	s, err := ModelToStructPB(
		&struct{ Test *YATS `vstore:"test"` }{Test: &YATS{TestStruct: &TestStruct{"hi"}}},
	)
	assert.Nil(t, err)
	v := map[string]*vstorepb.Value{
		"test": {
			&vstorepb.Value_StructValue{
				&vstorepb.Struct{
					map[string]*vstorepb.Value{
						"TestStruct": {
							&vstorepb.Value_StructValue{
								&vstorepb.Struct{
									map[string]*vstorepb.Value{
										"name": {&vstorepb.Value_StringValue{"hi"}},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	assert.Equal(t, &vstorepb.Struct{Values: v}, s)
}

func TestModelToStructPB_RepeatedReferencedDoubleNestedStructuredPropertyWithValueHasExpectedOutput(t *testing.T) {
	s, err := ModelToStructPB(
		&struct{ Test []*YATS `vstore:"test"` }{Test: []*YATS{{TestStruct: &TestStruct{"hi"}}}},
	)
	assert.Nil(t, err)
	v := map[string]*vstorepb.Value{
		"test": &vstorepb.Value{
			&vstorepb.Value_ListValue{
				&vstorepb.ListValue{
					[]*vstorepb.Value{
						{&vstorepb.Value_StructValue{
							&vstorepb.Struct{
								map[string]*vstorepb.Value{
									"TestStruct": {
										&vstorepb.Value_StructValue{
											&vstorepb.Struct{
												map[string]*vstorepb.Value{
													"name": {&vstorepb.Value_StringValue{"hi"}},
												},
											},
										},
									},
								},
							},
						}},
					},
				},
			},
		},
	}
	assert.Equal(t, &vstorepb.Struct{Values: v}, s)
}

func TestProtoToStructPB_String(t *testing.T) {
	s, err := ModelToStructPB(
		&vstorepb.Value_StringValue{"Hello"},
	)
	assert.Nil(t, err)
	assert.Equal(t, &vstorepb.Struct{Values: map[string]*vstorepb.Value{"string_value": {&vstorepb.Value_StringValue{"Hello"}}}}, s)
}

func TestModelToStructPB_privateVariableReturnsCorrectError(t *testing.T) {
	_, err := ModelToStructPB(
		&struct{ name string `vstore:"name"` }{name: "Hello"},
	)
	assert.NotNil(t, err)
	assert.Equal(t, "Cant serialize private field: `name`.", err.Error())
}

func TestStructPBToModelTestSuite(t *testing.T) {
	suite.Run(t, new(StructPBToModelTestSuite))
}

type StructPBToModelTestSuite struct {
	suite.Suite
}

type Artist struct {
	Name string `vstore:"name"`
}

type Song struct {
	Name         string             `vstore:"name"`
	Duration     int64              `vstore:"duration"`
	Rating       float64            `vstore:"rating"`
	RecordedAt   *vstorepb.GeoPoint `vstore:"recorded_at"`
	Released     time.Time          `vstore:"released"`
	WentPlatinum *time.Time         `vstore:"went_platinum"`
	Genres       []string           `vstore:"genres"`
	Artist       *Artist            `vstore:"artist"`
	Featuring    []*Artist          `vstore:"featuring"`
	DatesPlayed  []time.Time        `vstore:"dates_played"`
	IntField     IntEnum            `vstore:"int_field"`
	I32Field     I32Enum            `vstore:"i32_field"`
	I64Field     I64Enum            `vstore:"i64_field"`
	I64ListField []I64Enum          `vstore:"i64_list_field"`
	F32Field     F32Enum            `vstore:"f32_field"`
	F64Field     F64Enum            `vstore:"f64_field"`
	StrField     StrEnum            `vstore:"str_field"`
	BulField     BulEnum            `vstore:"bul_field"`
	BytesField   []byte             `vstore:"bytes_field"`
}

func (t *Song) Schema() *Schema {
	return NewSchema("vstore", "Song", []string{"name"}, NewPropertyBuilder().StringProperty("name").Build(), nil, nil)
}

func (suite *StructPBToModelTestSuite) SetupTest() {
	RegisterModel("vstore", "Song", (*Song)(nil))
}

func (suite *StructPBToModelTestSuite) Test_AliasedIntFieldIsConverted() {
	s := &vstorepb.Struct{Values: map[string]*vstorepb.Value{"int_field": {&vstorepb.Value_IntValue{int64(Int)}}}}
	m, err := StructPBToModel("vstore", "Song", s)
	suite.Assert().Nil(err)
	song := m.(*Song)
	suite.Assert().Equal(Int, song.IntField)
}

func (suite *StructPBToModelTestSuite) Test_AliasedI32FieldIsConverted() {
	s := &vstorepb.Struct{Values: map[string]*vstorepb.Value{"i32_field": {&vstorepb.Value_IntValue{int64(I32)}}}}
	m, err := StructPBToModel("vstore", "Song", s)
	suite.Assert().Nil(err)
	song := m.(*Song)
	suite.Assert().Equal(I32, song.I32Field)
}

func (suite *StructPBToModelTestSuite) Test_AliasedI64FieldIsConverted() {
	s := &vstorepb.Struct{Values: map[string]*vstorepb.Value{"i64_field": {&vstorepb.Value_IntValue{int64(I64)}}}}
	m, err := StructPBToModel("vstore", "Song", s)
	suite.Assert().Nil(err)
	song := m.(*Song)
	suite.Assert().Equal(I64, song.I64Field)
}

func (suite *StructPBToModelTestSuite) Test_AliasedI64ListFieldIsConverted() {
	s := &vstorepb.Struct{Values: map[string]*vstorepb.Value{"i64_list_field": {&vstorepb.Value_ListValue{
		ListValue: &vstorepb.ListValue{[]*vstorepb.Value{{&vstorepb.Value_IntValue{int64(4)}}, {&vstorepb.Value_IntValue{int64(5)}}, {&vstorepb.Value_IntValue{int64(6)}}}}},
	}}}
	m, err := StructPBToModel("vstore", "Song", s)
	suite.Assert().Nil(err)
	song := m.(*Song)
	suite.Assert().Equal([]I64Enum{I64Enum(4), I64Enum(5), I64Enum(6)}, song.I64ListField)
}

func (suite *StructPBToModelTestSuite) Test_AliasedFloat32FieldIsConverted() {
	s := &vstorepb.Struct{Values: map[string]*vstorepb.Value{"f32_field": {&vstorepb.Value_DoubleValue{float64(F32)}}}}
	m, err := StructPBToModel("vstore", "Song", s)
	suite.Assert().Nil(err)
	song := m.(*Song)
	suite.Assert().Equal(F32, song.F32Field)
}

func (suite *StructPBToModelTestSuite) Test_AliasedFloat64FieldIsConverted() {
	s := &vstorepb.Struct{Values: map[string]*vstorepb.Value{"f64_field": {&vstorepb.Value_DoubleValue{float64(F64)}}}}
	m, err := StructPBToModel("vstore", "Song", s)
	suite.Assert().Nil(err)
	song := m.(*Song)
	suite.Assert().Equal(F64, song.F64Field)
}

func (suite *StructPBToModelTestSuite) Test_AliasedStringFieldIsConverted() {
	s := &vstorepb.Struct{Values: map[string]*vstorepb.Value{"str_field": {&vstorepb.Value_StringValue{string(Str)}}}}
	m, err := StructPBToModel("vstore", "Song", s)
	suite.Assert().Nil(err)
	song := m.(*Song)
	suite.Assert().Equal(Str, song.StrField)
}

func (suite *StructPBToModelTestSuite) Test_AliasedBoolFieldIsConverted() {
	s := &vstorepb.Struct{Values: map[string]*vstorepb.Value{"bul_field": {&vstorepb.Value_BoolValue{bool(Bul)}}}}
	m, err := StructPBToModel("vstore", "Song", s)
	suite.Assert().Nil(err)
	song := m.(*Song)
	suite.Assert().Equal(Bul, song.BulField)
}

func (suite *StructPBToModelTestSuite) Test_StringFieldIsConverted() {
	s := &vstorepb.Struct{Values: map[string]*vstorepb.Value{"name": {&vstorepb.Value_StringValue{"Hello"}}}}
	m, err := StructPBToModel("vstore", "Song", s)
	suite.Assert().Nil(err)
	song := m.(*Song)
	suite.Assert().Equal("Hello", song.Name)
}

func (suite *StructPBToModelTestSuite) Test_BytesFieldIsConverted() {
	s := &vstorepb.Struct{Values: map[string]*vstorepb.Value{"bytes_field": {&vstorepb.Value_BytesValue{[]byte("Hello")}}}}
	m, err := StructPBToModel("vstore", "Song", s)
	suite.Assert().Nil(err)
	song := m.(*Song)
	suite.Assert().Equal([]byte("Hello"), song.BytesField)
}

func (suite *StructPBToModelTestSuite) Test_RepeatedStringFieldIsConverted() {
	s := &vstorepb.Struct{Values: map[string]*vstorepb.Value{"genres": {&vstorepb.Value_ListValue{
		ListValue: &vstorepb.ListValue{[]*vstorepb.Value{{&vstorepb.Value_StringValue{"Hello"}}}}},
	}}}
	m, err := StructPBToModel("vstore", "Song", s)
	suite.Assert().Nil(err)
	song := m.(*Song)
	suite.Assert().Equal([]string{"Hello"}, song.Genres)
}

func (suite *StructPBToModelTestSuite) Test_IntegerFieldIsConverted() {
	s := &vstorepb.Struct{Values: map[string]*vstorepb.Value{"duration": {&vstorepb.Value_IntValue{120}}}}
	m, err := StructPBToModel("vstore", "Song", s)
	suite.Assert().Nil(err)
	song := m.(*Song)
	suite.Assert().Equal(int64(120), song.Duration)
}

func (suite *StructPBToModelTestSuite) Test_FloatFieldIsConverted() {
	s := &vstorepb.Struct{Values: map[string]*vstorepb.Value{"rating": {&vstorepb.Value_DoubleValue{4.5}}}}
	m, err := StructPBToModel("vstore", "Song", s)
	suite.Assert().Nil(err)
	song := m.(*Song)
	suite.Assert().Equal(4.5, song.Rating)
}

func (suite *StructPBToModelTestSuite) Test_GeoFieldIsConverted() {
	geo := vstorepb.GeoPoint{Latitude: 120, Longitude: 60}
	s := &vstorepb.Struct{Values: map[string]*vstorepb.Value{"recorded_at": {&vstorepb.Value_GeopointValue{&geo}}}}
	m, err := StructPBToModel("vstore", "Song", s)
	suite.Assert().Nil(err)
	song := m.(*Song)
	suite.Assert().Equal(&geo, song.RecordedAt)
}

func (suite *StructPBToModelTestSuite) Test_TimeIsConverted() {
	now := time.Now().UTC()
	nowPb, _ := ptypes.TimestampProto(now)
	s := &vstorepb.Struct{Values: map[string]*vstorepb.Value{"released": {&vstorepb.Value_TimestampValue{nowPb}}}}
	m, err := StructPBToModel("vstore", "Song", s)
	suite.Assert().Nil(err)
	song := m.(*Song)
	suite.Assert().Equal(now, song.Released)
}

func (suite *StructPBToModelTestSuite) Test_ReferenceToTimeIsConverted() {
	now := time.Now().UTC()
	nowPb, _ := ptypes.TimestampProto(now)
	s := &vstorepb.Struct{Values: map[string]*vstorepb.Value{"went_platinum": {&vstorepb.Value_TimestampValue{nowPb}}}}
	m, err := StructPBToModel("vstore", "Song", s)
	suite.Assert().Nil(err)
	song := m.(*Song)
	suite.Assert().Equal(&now, song.WentPlatinum)
}

func (suite *StructPBToModelTestSuite) Test_StructuredPropertyIsConverted() {
	s := &vstorepb.Struct{Values: map[string]*vstorepb.Value{
		"artist": {
			&vstorepb.Value_StructValue{
				&vstorepb.Struct{
					map[string]*vstorepb.Value{
						"name": {&vstorepb.Value_StringValue{"Ben"}},
					},
				},
			},
		},
	}}
	m, err := StructPBToModel("vstore", "Song", s)
	suite.Assert().Nil(err)
	song := m.(*Song)
	suite.Assert().Equal("Ben", song.Artist.Name)
}

func (suite *StructPBToModelTestSuite) Test_RepeatedReferencedStructuredPropertyIsConverted() {
	s := &vstorepb.Struct{Values: map[string]*vstorepb.Value{
		"featuring": {
			&vstorepb.Value_ListValue{
				&vstorepb.ListValue{
					[]*vstorepb.Value{
						{
							&vstorepb.Value_StructValue{
								&vstorepb.Struct{
									map[string]*vstorepb.Value{
										"name": {&vstorepb.Value_StringValue{"Ben"}},
									},
								},
							},
						},
					},
				},
			},
		},
	}}
	m, err := StructPBToModel("vstore", "Song", s)
	suite.Assert().Nil(err)
	song := m.(*Song)
	suite.Assert().Equal("Ben", song.Featuring[0].Name)
}

func (suite *StructPBToModelTestSuite) Test_RepeatedStructuredPropertyIsConverted() {
	now := time.Now().UTC()
	nowPb, _ := ptypes.TimestampProto(now)
	s := &vstorepb.Struct{Values: map[string]*vstorepb.Value{
		"dates_played": {
			&vstorepb.Value_ListValue{
				&vstorepb.ListValue{
					[]*vstorepb.Value{
						{
							&vstorepb.Value_TimestampValue{
								nowPb,
							},
						},
					},
				},
			},
		},
	}}
	m, err := StructPBToModel("vstore", "Song", s)
	suite.Assert().Nil(err)
	song := m.(*Song)
	suite.Assert().Equal(now, song.DatesPlayed[0])
}

func TestModelToByteArrayHasExpectedOutput(t *testing.T) {
	s := &struct{ Name string `vstore:"name"` }{Name: "Hello"}
	b, err := ModelToByteArray(s)
	assert.Nil(t, err)
	assert.Equal(t, []byte{0xa, 0xf, 0xa, 0x4, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x7, 0x22, 0x5, 0x48, 0x65, 0x6c, 0x6c, 0x6f}, b)
}

func TestByteArrayToModelHasExpectedOutput(t *testing.T) {
	RegisterModel("vstore", "Song", (*Song)(nil))
	now := time.Now().UTC()
	s := &Song{
		Name:     "Morbid Dimensions",
		Duration: 636,
		Rating:   4.0,
		RecordedAt: &vstorepb.GeoPoint{
			Latitude:  50.0,
			Longitude: 60.0,
		},
		Released: now,
		Genres:   []string{"death metal", "osdm", "norwegian"},
		Artist: &Artist{
			Name: "Execration",
		},
	}

	b, err := ModelToByteArray(s)
	assert.Nil(t, err)

	m, err := ByteArrayToModel("vstore", "Song", b)
	assert.Nil(t, err)

	song, ok := m.(*Song)
	assert.True(t, ok)

	assert.Equal(t, "Morbid Dimensions", song.Name)
	assert.Equal(t, int64(636), song.Duration)
	assert.Equal(t, 50.0, song.RecordedAt.Latitude)
	assert.Equal(t, 60.0, song.RecordedAt.Longitude)
	assert.Equal(t, now, song.Released)
	assert.Equal(t, []string{"death metal", "osdm", "norwegian"}, song.Genres)
	assert.Equal(t, "Execration", song.Artist.Name)
}
