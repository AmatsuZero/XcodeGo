package Utils

import (
	"crypto/md5"
	"github.com/satori/go.uuid"
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
)

type KeyBuilder struct {
	value [16]byte
}

func Unique() KeyBuilder {
	uid := uuid.Must(uuid.NewV4())
	return create([]byte(uid.String()))
}

func create(bytes []byte) KeyBuilder {
	return KeyBuilder{value: md5.Sum(bytes)}
}

func (t KeyBuilder) Build() (stringValue string) {
	return fmt.Sprintf("%x", t.value)
}
