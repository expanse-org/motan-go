package serialize

import (
	motan "github.com/weibocom/motan-go/core"
)

const (
	Simple = "simple"
	ProtoBuf = "protobuf"
	Gob = "gob"
)

func RegistDefaultSerializations(extFactory motan.ExtentionFactory) {
	extFactory.RegistryExtSerialization(Simple, 6, func() motan.Serialization {
		return &SimpleSerialization{}
	})
	extFactory.RegistryExtSerialization(ProtoBuf, 7, func() motan.Serialization {
		return &ProtoBufSerialization{}
	})
	extFactory.RegistryExtSerialization(Gob, 8, func() motan.Serialization {
		return &GobSerialization{}
	})
}
