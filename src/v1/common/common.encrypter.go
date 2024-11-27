package common

import (
	"crypto/sha256"
	"encoding/hex"
)

func UserEncrypt(text string) (encrypt string, err error) {
	hash := sha256.Sum256([]byte(text))

	encrypt = hex.EncodeToString(hash[:])

	return encrypt, nil
}
