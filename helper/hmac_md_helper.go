package helper

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"time"
)

func GenerateHMAC(data string) string {
	h := hmac.New(md5.New, []byte(os.Getenv("APP_SESSION_PASSWORD_KEY")))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func GenerateHMACTest(data string) string {
	h := hmac.New(md5.New, []byte("SIDESAdatadesaKepukBangsriJeparaDANjALANjALANkeMANA"))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func md5Hash(data []byte) []byte {
	hash := md5.New()
	hash.Write(data)
	return hash.Sum(nil)
}

// generateHMACMD5 generates HMAC using MD5 hashing algorithm manually
func GenerateHMACMD5(messages string) string {
	start := time.Now()

	fmt.Println("Password : ", messages)

	// ubah string menjadi byte
	message := []byte(messages)

	key := []byte("kunciaplikasisidesa")

	// ukuran blok of MD5
	blockSize := 64

	// Jika kunci lebih panjang dari ukuran blok atau ukurannya sama dengan blok lakukan hash. jika tidak, isi dengan angka nol
	if len(key) > blockSize {
		key = md5Hash(key)

	} else if len(key) < blockSize {
		padding := make([]byte, blockSize-len(key))
		key = append(key, padding...)
	}

	// Inner and outer pads
	innerPad := make([]byte, blockSize)

	outerPad := make([]byte, blockSize)

	for i := 0; i < blockSize; i++ {
		innerPad[i] = key[i] ^ 0x36
		outerPad[i] = key[i] ^ 0x5c
	}

	// Calculate inner hash
	innerHash := md5Hash(append(innerPad, message...))

	// Calculate outer hash
	outerHash := md5Hash(append(outerPad, innerHash...))

	duration := time.Since(start)

	fmt.Println("Waktu Hasing HMAC-MD5 (detik) : ", duration.Seconds())

	return hex.EncodeToString(outerHash)
}
