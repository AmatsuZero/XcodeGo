package Utils

import (
	"crypto/md5"
	"fmt"
	"github.com/satori/go.uuid"
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
