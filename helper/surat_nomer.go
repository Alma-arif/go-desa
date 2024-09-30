package helper

import (
	"fmt"
	"strconv"
	"strings"
)

func NomerSurat(nomer int) string {
	nomorStr := strconv.Itoa(nomer)
	nomorSurat := fmt.Sprintf("%s", strings.Repeat("0", 4-len(nomorStr))+nomorStr)
	return nomorSurat
}
