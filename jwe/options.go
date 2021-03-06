package jwe

import (
	"context"

	"github.com/lestrrat-go/option"
)

type Option = option.Interface
type identPrettyFormat struct{}
type identProtectedHeader struct{}
type SerializerOption interface {
	Option
	serializerOption()
}

type serializerOption struct {
	Option
}

func (*serializerOption) serializerOption() {}

type EncryptOption interface {
	Option
	encryptOption()
}

type encryptOption struct {
	Option
}

func (*encryptOption) encryptOption() {}

// WithPrettyFormat specifies if the `jwe.JSON` serialization tool
// should generate pretty-formatted output
func WithPrettyFormat(b bool) SerializerOption {
	return &serializerOption{option.New(identPrettyFormat{}, b)}
}

// Specify contents of the protected header. Some fields such as
// "enc" and "zip" will be overwritten when encryption is performed.
func WithProtectedHeaders(h Headers) EncryptOption {
	cloned, _ := h.Clone(context.Background())
	return &encryptOption{option.New(identProtectedHeader{}, cloned)}
}
