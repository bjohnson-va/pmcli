package iam

import (
	"github.com/vendasta/gosdks/pb/iam/v1"
	"github.com/vendasta/gosdks/iam/subject"
	"github.com/vendasta/gosdks/iam/subjectcontext"
)

// FromPersona converts a Persona proto into a Subject
func FromPersona(p *iam_v1.Persona) (subject.Subject, error) {
	subjectContext := subjectcontext.FromProto(p.Context)
	return subject.New(subjectContext, p.Subject)
}
