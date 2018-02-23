package subjectcontext

import (
	"github.com/vendasta/gosdks/pb/iam/v1"
)

// Context scopes a request to a specific persona.  A persona can either be a typed or a typed/namespace unique pair.
type Context struct {
	// Type is the persona type (salesperson, partner, smb, etc)
	Type string

	// Namespace indicated that the persona is in a specific context (the context of a PID for example). This can be empty.
	Namespace string
}

//New creates and returns a new context.
func New(subjectType, namespace string) *Context {
	return &Context{
		Type:      subjectType,
		Namespace: namespace,
	}
}

// FromProto returns an iam context from the context proto
func FromProto(context *iam_v1.Context) *Context {
	namespaced := context.GetNamespaced()
	if namespaced == nil {
		typedContext := context.GetTyped()
		return New(typedContext.Type, "")
	}
	return New(namespaced.Type, namespaced.Namespace)
}

// ToPB returns the context in proto format.
func (c *Context) ToPB() *iam_v1.Context {
	if c == nil {
		return nil
	}

	if c.Namespace == "" {
		return &iam_v1.Context{
			Context: &iam_v1.Context_Typed{
				Typed: &iam_v1.TypedContext{
					Type: c.Type,
				},
			},
		}
	}
	return &iam_v1.Context{
		Context: &iam_v1.Context_Namespaced{
			Namespaced: &iam_v1.NamespacedContext{
				Type:      c.Type,
				Namespace: c.Namespace,
			},
		},
	}
}
