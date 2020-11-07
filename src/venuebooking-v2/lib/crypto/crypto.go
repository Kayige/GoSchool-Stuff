package crypto

import (
	"crypto/sha256"
	"math/rand"
	"strings"
)

// set crypto key settings
const (
	CRYPT_SETTINGS = "$P$BwQZDcQaNU9zAOF.6MOUdEhz9X68fL1"
)

// encode blocksize 64
func encode64(inp []byte, count int) string {
	// int to string
	const itoa64 = "./0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var outp string
	cur := 0
	for cur < count {
		value := uint(inp[cur])
		cur++
		outp += string(itoa64[value&0x3f])
		if cur < count {
			value |= (uint(inp[cur]) << 8)
		}
		outp += string(itoa64[(value>>6)&0x3f])

		if cur >= count {
			break
		}
		cur++
		if cur < count {
			value |= (uint(inp[cur]) << 16)
		}
		outp += string(itoa64[(value>>12)&0x3f])
		if cur >= count {
			break
		}
		cur++
		outp += string(itoa64[(value>>18)&0x3f])
	}
	return outp
}

// CryptPrivate creates a private key & hashes the password
func CryptPrivate(pw, setting string) string {
	const itoa64 = "./0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var outp = "*0"
	var countlog2 uint
	countlog2 = uint(strings.Index(itoa64, string(setting[3])))
	if countlog2 < 7 || countlog2 > 30 {
		return outp
	}
	count := 1 << countlog2
	salt := setting[4:12]
	if len(salt) != 8 {
		return outp
	}
	hasher := sha256.New()
	hasher.Write([]byte(salt + pw))
	hx := hasher.Sum(nil)
	for count != 0 {
		hasher := sha256.New()
		hasher.Write([]byte(string(hx) + pw))
		hx = hasher.Sum(nil)
		count--
	}
	return setting[:12] + encode64(hx, 16)
}

// PortableHashCheck takes in the password and verifies with the storedHash
func PortableHashCheck(pw, storedHash string) bool {
	hx := CryptPrivate(pw, storedHash)
	return hx == storedHash
}

// GenerateRandomString creates a random string of letters
func GenerateRandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyz")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// GenerateRandomNumber creates a random int of digits
func GenerateRandomNumber(n int) string {
	var letters = []rune("0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
