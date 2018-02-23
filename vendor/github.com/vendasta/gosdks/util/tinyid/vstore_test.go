package tinyid

import (
	"testing"
	"context"
	"github.com/vendasta/gosdks/vstore"
	"github.com/stretchr/testify/assert"
	"strings"
)

func Test_VStoreGeneratesTinyID(t *testing.T) {
	// arrange
	hps := vstore.NewHappyPathStub()
	gen, _ := NewVStoreGenerator(context.TODO(), hps, "test-prefix", nil, "AG", Length(6))

	// act
	tinyID, err := gen.GenerateTinyID(context.TODO())

	// assert
	assert.Nil(t, err)
	assert.True(t, strings.HasPrefix(tinyID, "AG-"))
	assert.Equal(t, 9, len(tinyID))
}

func Test_claimTinyID_ReturnsErrorOnCollision(t *testing.T) {
	// arrange
	hps := vstore.NewHappyPathStub()
	gen, _ := NewVStoreGenerator(context.TODO(), hps, "test-prefix", nil, "AG", Length(6))
	tinyIdGen := gen.(*vstoreTinyID)

	//act
	nilErr := tinyIdGen.claimTinyID(context.TODO(), "AG-PFRTXQ")
	err := tinyIdGen.claimTinyID(context.TODO(), "AG-PFRTXQ")

	// assert
	assert.Nil(t, nilErr)
	assert.NotNil(t, err)
	assert.Equal(t, "TinyID Collision", err.Error())
}

type stringGeneratorMock struct {
	// forces the mock to always return this string
	alwaysReturn string

	// will return each string until consumed
	returnList []string

}

func (s *stringGeneratorMock) generateString(prefix string, length int32, chars []byte) string {
	if s.alwaysReturn != "" {
		return prefix + s.alwaysReturn
	}
	if len(s.returnList) == 0 {
		panic("ran out of items to return")
	}
	toReturn := s.returnList[0]
	s.returnList = s.returnList[1:]
	return prefix + toReturn
}

func Test_ReachingMaxCollisionReturnsError(t *testing.T) {
	// arrange
	hps := vstore.NewHappyPathStub()
	gen, _ := NewVStoreGenerator(context.TODO(), hps, "test-prefix", nil, "AG", Length(4))
	tinyIdGen := gen.(*vstoreTinyID)
	tinyIdGen.stringGenerator = &stringGeneratorMock{alwaysReturn:"same"}

	//act
	tinyIdGen.claimTinyID(context.TODO(), "AG-same") // manually claim tiny id to force retries below
	_, err := tinyIdGen.GenerateTinyID(context.TODO())

	// assert
	assert.NotNil(t, err)
	assert.Equal(t, "TinyID Collision", err.Error())
}


func Test_OnRetryWillReturnTinyID(t *testing.T) {
	// arrange
	hps := vstore.NewHappyPathStub()
	gen, _ := NewVStoreGenerator(context.TODO(), hps, "test-prefix", nil, "AG", Length(6))
	tinyIdGen := gen.(*vstoreTinyID)
	tinyIdGen.stringGenerator = &stringGeneratorMock{returnList: []string{"same", "different"}}

	//act
	tinyIdGen.claimTinyID(context.TODO(), "AG-same") // manually claim tiny id to force retries below
	tinyID, err := tinyIdGen.GenerateTinyID(context.TODO())

	// assert
	assert.Nil(t, err)
	assert.Equal(t, "AG-different", tinyID)
}

func Test_NewVStoreGenerator_ReturnsErrorOnInvalidOptions(t *testing.T) {
	hps := vstore.NewHappyPathStub()
	tests := []struct {
		opts []Opt
		err string
	}{
		{
			opts: []Opt{
				Length(1),
			},
			err: "Invalid Tiny ID Length. Must be > 1, provided 1",
		},
		{
			opts: []Opt{
				Length(0),
			},
			err: "Invalid Tiny ID Length. Must be > 1, provided 0",
		},
		{
			opts: []Opt{
				MaxTrials(1),
			},
			err: "Invalid Tiny ID Max Trials. Must be > 1, provided 1",
		},
		{
			opts: []Opt{
				MaxTrials(0),
			},
			err: "Invalid Tiny ID Max Trials. Must be > 1, provided 0",
		},
		{
			opts: []Opt{
				Chars(nil),
			},
			err: "Invalid amount of Tiny ID Chars. Must be > 1, provided 0 chars.",
		},{
			opts: []Opt{
				Chars([]byte{'1'}),
			},
			err: "Invalid amount of Tiny ID Chars. Must be > 1, provided 1 chars.",
		},
	}
	for _, test := range tests {
		_, err := NewVStoreGenerator(context.TODO(), hps, "test-prefix", nil, "AG", test.opts...)
		assert.Equal(t, test.err, err.Error())
	}

}