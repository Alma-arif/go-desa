package helper

import (
	"encoding/hex"
	"errors"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func GetSessionID(sessionID string) (uint, error) {
	var idUser uint

	// idDekripsiString, err := hex.DecodeString(sessionID)

	// if err != nil {
	// 	return idUser, err
	// }

	// idEnkripsi := Rc4Decrypt(idDekripsiString, []byte(os.Getenv("APP_SESSION_key")))

	// idStras := string(idEnkripsi)
	// idStrasa := strings.ReplaceAll(idStras, "\n", "")
	// idStrasa = strings.TrimSpace(idStrasa)

	idUint, err := strconv.Atoi(sessionID)
	if err != nil {
		return idUser, err
	}

	idUser = uint(idUint)

	return idUser, nil
}

func GetFileNameEnkrip(fileName string) (string, error) {
	var file string

	extension := filepath.Ext(fileName)
	nameWithoutExt := fileName[:len(fileName)-len(extension)]
	resultFileName := strings.ReplaceAll(nameWithoutExt, ".", "")

	fileNameEnkripsi, err := Rc4Encrypt([]byte(resultFileName), []byte("SIDESAdatadesaKepukBangsriJeparaFILEterEnkripsiRc4"))

	// fileNameEnkripsi, err := Rc4Encrypt([]byte(resultFileName), []byte(os.Getenv("APP_FILE_NAME_SECRET_KEY")))
	if err != nil {
		return file, err
	}

	fileNameEnkripsiString := hex.EncodeToString(fileNameEnkripsi)
	unixtime := time.Now().UnixNano()
	unixtimeStr := strconv.FormatInt(unixtime, 10)

	fileString := fmt.Sprintf("%s-%s%s", unixtimeStr, fileNameEnkripsiString, extension)

	return fileString, nil
}

func GetFileNameDekrip(fileName string) (string, error) {
	var file string

	fileSplit := strings.Split(fileName, "-")

	fmt.Println(fileSplit)
	if fileSplit[1] == "" {
		return fileName, errors.New("file tidak di temukan!")
	}

	file = fileSplit[1]

	extension := filepath.Ext(file)
	nameWithoutExt := file[:len(file)-len(extension)]
	resultFileName := strings.ReplaceAll(nameWithoutExt, ".", "")

	fileNameDecod, err := hex.DecodeString(resultFileName)
	if err != nil {
		return file, err
	}
	fileNameDekrip, err := Rc4Decrypt(fileNameDecod, []byte("SIDESAdatadesaKepukBangsriJeparaFILEterEnkripsiRc4"))

	// fileNameDekrip, err := Rc4Decrypt(fileNameDecod, []byte(os.Getenv("APP_FILE_NAME_SECRET_KEY")))
	if err != nil {
		return file, err
	}
	fileNameStringOne := string(fileNameDekrip)
	fileNameString := strings.ReplaceAll(fileNameStringOne, "\n", "")
	fileNameString = strings.TrimSpace(fileNameString)

	file = fmt.Sprintf("%s%s", fileNameString, extension)

	return file, nil
}
