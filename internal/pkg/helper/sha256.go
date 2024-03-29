package helper

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"strconv"
	"strings"
	"time"
)

func encodeSHA256(str string) []byte {
	var secret string

	// TODO: will be supported later
	//secret := library.GetConfig.SecretKey

	strToEncode := str + secret
	h := sha256.New()
	h.Write([]byte(strToEncode))

	return h.Sum(nil)
}

// UnixMilli returns t as a Unix time, the number of milliseconds elapsed since January 1, 1970 UTC
func UnixMilliByTime(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

// GenSHA256 Helper to generate SHA256 with str + cert key combination with ttl expiration in minutes
func GenSHA256(str string, ttl time.Duration) string {
	t := time.Now()
	subtract := UnixMilliByTime(t.Add(time.Minute * ttl))

	enc := encodeSHA256(str)
	return base64.StdEncoding.EncodeToString([]byte(hex.EncodeToString(enc) + "." + strconv.FormatInt(subtract, 10)))
}

// ValidateSHA256 Validate Encoded Key is valid by time and compare them with the same string
func ValidateSHA256(str string, key string) bool {
	// Decode previous encryption string
	decode, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return false
	}

	encSplitStr := strings.Split(string(decode), ".") // split string [encode string, ttl]

	if len(encSplitStr) < 2 {
		return false
	}

	// Compare and make sure enc str still valid
	s, _ := strconv.ParseInt(encSplitStr[1], 10, 64)

	// Make sure encode key still valid
	isKeyValid := CompareTimeNow(s)
	if isKeyValid.isGt {
		return false // false if tim.Now() greater than ttl. enc key is invalid
	}

	//  Encode string to be compared
	encode := hex.EncodeToString(encodeSHA256(str))

	return encSplitStr[0] == encode
}
