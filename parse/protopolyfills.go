package parse

import "github.com/emicklei/proto"

//
// EVERYTHING IN THIS FILE SHOULD EVENTUALLY BE PROVIDED BY THE PROTO LIB
// https://github.com/emicklei/proto/issues/58
//

func Fields(protoElements []proto.Visitee) []proto.NormalField {
	list := []proto.NormalField{}
	for _, each := range protoElements {
		if c, ok := each.(*proto.NormalField); ok {
			list = append(list, *c)
		}
	}
	return list
}

func Package(protoElements []proto.Visitee) *proto.Package {
	for _, each := range protoElements {
		if c, ok := each.(*proto.Package); ok {
			return c
		}
	}
	return nil
}

func Messages(protoElements []proto.Visitee) []proto.Message {
	list := []proto.Message{}
	for _, each := range protoElements {
		if c, ok := each.(*proto.Message); ok {
			list = append(list, *c)
		}
	}
	return list
}

func Imports(protoElements []proto.Visitee) []proto.Import {
	list := []proto.Import{}
	for _, each := range protoElements {
		if c, ok := each.(*proto.Import); ok {
			list = append(list, *c)
		}
	}
	return list
}

func Enums(protoElements []proto.Visitee) []proto.Enum {
	list := []proto.Enum{}
	for _, each := range protoElements {
		if c, ok := each.(*proto.Enum); ok {
			list = append(list, *c)
		}
	}
	return list
}

func EnumFields(protoElements []proto.Visitee) []proto.EnumField {
	list := []proto.EnumField{}
	for _, each := range protoElements {
		if c, ok := each.(*proto.EnumField); ok {
			list = append(list, *c)
		}
	}
	return list
}

func Services(protoElements []proto.Visitee) []proto.Service {
	list := []proto.Service{}
	for _, each := range protoElements {
		if c, ok := each.(*proto.Service); ok {
			list = append(list, *c)
		}
	}
	return list
}

func RPCs(elements []proto.Visitee) []proto.RPC {
	list := []proto.RPC{}
	for _, each := range elements {
		if c, ok := each.(*proto.RPC); ok {
			list = append(list, *c)
		}
	}
	return list
}
