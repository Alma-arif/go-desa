package helper

import (
	"crypto/rc4"
)

func Rc4Encrypt(data, key []byte) ([]byte, error) {
	var Filebyte []byte

	cipher, err := rc4.NewCipher(key)
	if err != nil {
		return Filebyte, err
	}
	encrypted := make([]byte, len(data))
	cipher.XORKeyStream(encrypted, data)

	return encrypted, nil
}

func Rc4Decrypt(data, key []byte) ([]byte, error) {
	var Filebyte []byte

	cipher, err := rc4.NewCipher(key)
	if err != nil {
		return Filebyte, err
	}
	dekrisi := make([]byte, len(data))
	cipher.XORKeyStream(dekrisi, data)

	return dekrisi, nil
}

func ksa(key []byte) []byte {
	s := make([]byte, 256)
	for i := range s {
		s[i] = byte(i)
	}

	j := byte(0)
	for i := range s {
		j += s[i] + key[i%len(key)]
		s[i], s[j] = s[j], s[i]
	}

	return s
}

func prgaEnkripsi(s []byte, plaintext []byte) []byte {
	i, j := byte(0), byte(0)
	ciphertext := make([]byte, len(plaintext))

	for k := range plaintext {
		i += 1
		j += s[i]
		s[i], s[j] = s[j], s[i]
		ciphertext[k] = plaintext[k] ^ s[(s[i]+s[j])&0xff]
	}

	return ciphertext
}

func prgaDekripsi(s []byte, ciphertext []byte) []byte {
	i, j := byte(0), byte(0)
	plaintext := make([]byte, len(ciphertext))

	for k := range ciphertext {
		i += 1
		j += s[i]
		s[i], s[j] = s[j], s[i]
		ciphertext[k] = ciphertext[k] ^ s[(s[i]+s[j])&0xff]
	}

	return plaintext
}

func Rc4Data(key []byte, data []byte) []byte {
	s := ksa(key)
	return prgaEnkripsi(s, data)
}

func Rc4DataDekripsi(key []byte, data []byte) []byte {
	s := ksa(key)
	return prgaDekripsi(s, data)
}
