package security

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/twinj/uuid"
)

/*
  - When a user requests to change his password, a token is sent to that user's email.
    A function is written to hash the token, it will be used when we wire up the ResetPassword controller file.
*/
func TokenHash(text string) string {

	hasher := md5.New()
	hasher.Write([]byte(text))
	theHash := hex.EncodeToString(hasher.Sum(nil))

	//also use uuid
	u := uuid.NewV4()
	theToken := theHash + u.String()

	return theToken
}
