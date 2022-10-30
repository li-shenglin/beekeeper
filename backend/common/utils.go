package common

import (
	"encoding/hex"

	"github.com/google/uuid"
)

func SampleUUID() string {
	guid := uuid.New()
	var buf [32]byte
	hex.Encode(buf[:], guid[:4])
	hex.Encode(buf[8:12], guid[4:6])
	hex.Encode(buf[12:16], guid[6:8])
	hex.Encode(buf[16:20], guid[8:10])
	hex.Encode(buf[20:], guid[10:])
	return string(buf[:])
}

func UUID() []byte {
	guid := uuid.New()
	var buf [32]byte
	hex.Encode(buf[:], guid[:4])
	hex.Encode(buf[8:12], guid[4:6])
	hex.Encode(buf[12:16], guid[6:8])
	hex.Encode(buf[16:20], guid[8:10])
	hex.Encode(buf[20:], guid[10:])
	return buf[:]
}
