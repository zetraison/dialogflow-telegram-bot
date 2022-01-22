package tgbot

import (
	"crypto/sha256"
	"fmt"
	"strconv"
)

func sha256HashFromInt(ID int) string {
	bs := []byte(strconv.Itoa(ID))
	hash := sha256.Sum256(bs)
	return fmt.Sprintf("%x", hash[:])
}
