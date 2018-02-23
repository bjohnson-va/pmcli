package vstore

//Model is capable of defining the schema for a specific Kind
type Model interface {
	Schema() *Schema
}
