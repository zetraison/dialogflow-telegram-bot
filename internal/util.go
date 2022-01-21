package internal

import (
	"crypto/sha256"
	"fmt"
	"strconv"
)

func sessionFromUserID(ID int) string {
	bs := []byte(strconv.Itoa(ID))
	hash := sha256.Sum256(bs)
	return fmt.Sprintf("%x", hash[:])
}
