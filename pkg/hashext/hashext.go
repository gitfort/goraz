package hashext

import (
	"crypto/sha256"
	"fmt"
)

func Encode(text string) string {
	hash := sha256.Sum256([]byte(text))
	return fmt.Sprintf("%x", hash)
}
