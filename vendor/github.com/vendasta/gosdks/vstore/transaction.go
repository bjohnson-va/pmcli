package vstore

import (
	"errors"
)

// Transaction is any type capable of saving a model to VStore
type Transaction interface {
	Save(Model) error
}

// ProtoTransaction is capable of saving a proto representation of a VStore entity to vstore
type ProtoTransaction interface {
	SavePB(p *ProtoStruct) error
}

type transaction struct {
	toSave Model
	pbToSave *ProtoStruct
}

var (
	//ErrInvalidTx is raised when we detect an invalid transaction - when nothing is being saved or when transactions are being nested
	ErrInvalidTx = errors.New("Invalid transaction.")
)

//Save sets the provided model to be saved to VStore
func (t *transaction) Save(m Model) error {
	if t.toSave != nil || t.pbToSave != nil {
		return ErrInvalidTx
	}
	t.toSave = m
	return nil
}

//SavePB sets the provided vstore Struct to be saved to VStore
func (t *transaction) SavePB(m *ProtoStruct) error {
	if t.toSave != nil || t.pbToSave != nil {
		return ErrInvalidTx
	}
	t.pbToSave = m
	return nil
}

type txOpts struct {
	pbTransactionCallback func(ProtoTransaction, *ProtoStruct) error
}

// TransactionOption parametrizes the execution of a transaction
type TransactionOption func(o *txOpts)

// WithProtoTransaction executes the transaction given the provided callback operating over a ProtoStruct.
// This allows the transaction to be agnostic of the entity schema and work in terms of the raw proto format.
// You probably don't need to use this.
func WithProtoTransaction(f func(ProtoTransaction, *ProtoStruct) error) TransactionOption {
	return func(o *txOpts) {
		o.pbTransactionCallback = f
	}
}