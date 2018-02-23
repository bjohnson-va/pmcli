package subjectcontext

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NamespacedContext(t *testing.T) {
	iamContext := New("smb", "ABC").ToPB()
	namespacedContext := iamContext.GetNamespaced()
	assert.NotNil(t, namespacedContext)
	assert.Equal(t, "ABC", namespacedContext.Namespace)
	assert.Equal(t, "smb", namespacedContext.Type)
}

func Test_TypedContext(t *testing.T) {
	iamContext := New("partner", "").ToPB()
	typedContext := iamContext.GetTyped()
	assert.NotNil(t, typedContext)
	assert.Equal(t, "partner", typedContext.Type)
}
