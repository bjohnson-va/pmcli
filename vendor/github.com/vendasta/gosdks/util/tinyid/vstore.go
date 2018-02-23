package tinyid

import (
	"github.com/vendasta/gosdks/vstore"
	"regexp"
	"fmt"
	"golang.org/x/net/context"
	"github.com/vendasta/gosdks/logging"
	"math/rand"
	"errors"
	"time"
)

const tinyIDKind = "TinyID"

var (
	tinyIdVStoreNamespace string
	tinyIDCollisionErr = errors.New("TinyID Collision")
)

// Length controls the length of the tiny id
func Length(length int32) Opt {
	return func(o *options) {
		o.tinyIDLength = length
	}
}

// MaxTrials controls how many times we will retry on tiny id collisions.
func MaxTrials(maxTrails int32) Opt {
	return func(o *options) {
		o.maxTrials = maxTrails
	}
}


// Chars controls the set of characters that the tiny generator will pull from.
func Chars(chars []byte) Opt {
	return func(o *options) {
		o.chars = chars
	}
}

type options struct {
	prefix       string
	tinyIDLength int32
	maxTrials    int32
	chars        []byte
}

type Opt func(o *options)

// Must start with a letter, only contains words after that.
var validPrefixRegex = regexp.MustCompile(`^[a-zA-Z][\w]*$`)

var defaultChars = []byte{'B', 'C', 'D', 'F', 'G', 'H', 'J', 'K', 'L', 'M', 'N', 'P', 'Q', 'R', 'S', 'T', 'V', 'W', 'X', 'Z', '2', '3', '4', '5', '6', '7', '8', '9'}

// NewVStoreGenerator returns a tiny ID generator that is backed by vStore.
func NewVStoreGenerator(ctx context.Context, vstoreAPI vstore.Interface, vstoreNamespace string, vstoreServiceAccounts []string, tinyIDPrefix string, opts ...Opt) (Interface, error) {
	o := &options{
		tinyIDLength: 8,
		maxTrials:    100,
		chars:        defaultChars,
		prefix:       tinyIDPrefix,
	}
	for _, opt := range opts {
		opt(o)
	}
	if !validPrefixRegex.MatchString(o.prefix) {
		return nil, fmt.Errorf("Invalid prefix provided %s, must match regex %s", o.prefix, validPrefixRegex.String())
	}

	if o.tinyIDLength <= 1 {
		return nil, fmt.Errorf("Invalid Tiny ID Length. Must be > 1, provided %d", o.tinyIDLength)
	}
	if o.maxTrials <= 1 {
		return nil, fmt.Errorf("Invalid Tiny ID Max Trials. Must be > 1, provided %d", o.maxTrials)
	}
	if len(o.chars) <= 1 {
		return nil, fmt.Errorf("Invalid amount of Tiny ID Chars. Must be > 1, provided %d chars.", len(o.chars))
	}
	tinyIdVStoreNamespace = vstoreNamespace
	_, err := vstoreAPI.RegisterKind(ctx, vstoreNamespace, tinyIDKind, vstoreServiceAccounts, (*TinyID)(nil))
	if err != nil {
		logging.Errorf(ctx, "Error registering Tiny ID %s", err.Error())
		return nil, err
	}

	return &vstoreTinyID{Interface: vstoreAPI, opts: o, stringGenerator: newStringGenerator()}, nil
}

type vstoreTinyID struct {
	vstore.Interface
	opts *options
	stringGenerator
}

func (v *vstoreTinyID) GenerateTinyID(ctx context.Context) (string, error) {
	for i := int32(0); i < v.opts.maxTrials; i++ {

		tinyID := v.generateString(v.opts.prefix + "-", v.opts.tinyIDLength, v.opts.chars)
		err := v.claimTinyID(ctx, tinyID)

		if err == nil {
			return tinyID, nil
		}
		logging.Debugf(ctx, "Error claiming tiny id, retrying. On attempt %d. %s", i, err.Error())
	}
	return "", tinyIDCollisionErr
}

func (v *vstoreTinyID) claimTinyID(ctx context.Context, tinyID string) error {
	tinyIDEntity := &TinyID{
		Prefix: v.opts.prefix,
		TinyID: tinyID,
	}
	return v.Interface.Transaction(ctx, newKeySet(v.opts.prefix, tinyIDEntity.TinyID), func(t vstore.Transaction, m vstore.Model) error {
		if m != nil {
			return tinyIDCollisionErr
		}
		return t.Save(tinyIDEntity)
	})
}

type stringGenerator interface {
	generateString(prefix string, length int32, chars []byte) string
}

type stringGeneratorImpl struct {

}

func (s *stringGeneratorImpl) generateString(prefix string, length int32, chars []byte) string {
	prefixLength := len(prefix)

	tinyID := make([]byte, length + int32(prefixLength))
	copy(tinyID[0:prefixLength], []byte(prefix))

	for i := int32(0); i < length; i++ {
		tinyID[i + int32(prefixLength)] = chars[rand.Int31n(int32(len(chars) - 1))]
	}
	return string(tinyID[:])
}

func newStringGenerator() stringGenerator {
	rand.Seed(time.Now().UnixNano())
	return &stringGeneratorImpl{}
}


func newKeySet(prefix string, tinyID string) *vstore.KeySet {
	return vstore.NewKeySet(tinyIdVStoreNamespace, tinyIDKind, []string{prefix, tinyID})
}

// TinyID is a vStore model for persisting tiny ids.
type TinyID struct {
	Prefix string `vstore:"prefix"`
	TinyID string `vstore:"tiny_id"`
}

// Schema provides the vStore schema for this model.
func (t *TinyID) Schema() *vstore.Schema {
	p := vstore.NewPropertyBuilder().StringProperty(
		"prefix", vstore.Required(),
	).StringProperty(
		"tiny_id", vstore.Required(),
	).Build()
	backupConfig := vstore.NewBackupConfigBuilder().PeriodicBackup(vstore.MonthlyBackup).Build()
	return vstore.NewSchema(tinyIdVStoreNamespace, tinyIDKind, []string{"prefix", "tiny_id"}, p, nil, backupConfig)
}
