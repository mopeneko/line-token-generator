package android

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"strings"
	"time"
)

func GenerateToken(authKey string) string {
	iat := generateIat()

	encodedIat := base64.StdEncoding.EncodeToString(([]byte)(iat))

	splitedAuthKey:= strings.Split(authKey, ":")

	mid := splitedAuthKey[0]
	key, _ := base64.StdEncoding.DecodeString(splitedAuthKey[1])

	message := encodedIat + "."

	mac := generateHmacIat(message, key)

	authToken := fmt.Sprintf("%s:%s.%s", mid, message, mac)
	return authToken
}

func generateIat() string {
	unixTime := time.Now().UnixNano() / int64(time.Millisecond)
	iat := fmt.Sprintf("iat: %d\n", unixTime)
	return iat
}

func generateHmacIat(iat string, key []byte) string {
	hasher := hmac.New(sha1.New, key)
	hasher.Write([]byte(iat))
	mac := base64.StdEncoding.EncodeToString(hasher.Sum(nil))
	return mac
}
