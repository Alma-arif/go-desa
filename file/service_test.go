package file

import (
	"app-desa-kepuk/config"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

var ValidateS = validator.New()
var DB = config.InitDBTest()
var RepositoryUji = NewRepository(DB)
var B = NewService(RepositoryUji, ValidateS)

func Test_Enkripsi(t *testing.T) {

	// input ID Dokumen
	var id uint = 3

	// fungsi metode
	_, err := B.EnkripsiRC(id)

	// fungsi testing
	assert.Nil(t, err, "Enkripsi")
}

func Test_Enkripsi_ID(t *testing.T) {

	// input ID Dokumen
	var id uint = 23

	// fungsi metode
	_, err := B.EnkripsiRC(id)

	// fungsi testing
	assert.NotNil(t, err, "Enkripsi")
}

func Test_Enkripsi_No_ID(t *testing.T) {

	// input ID
	var id uint

	// fungsi metode
	_, err := B.EnkripsiRC(id)

	// fungsi testing
	assert.NotNil(t, err, "Enkripsi")
}

func Test_Dekripsi(t *testing.T) {

	// input ID
	var id uint = 3

	// fungsi metode
	_, err := B.DekripsiRC(id)

	// fungsi testing
	assert.Nil(t, err, "Dekripsi")
}

func Test_Dekripsi_ID(t *testing.T) {

	// input ID
	var id uint = 45

	// fungsi metode
	_, err := B.DekripsiRC(id)

	// fungsi testing
	assert.NotNil(t, err, "Dekripsi")
}

func Test_Dekripsi_No_ID(t *testing.T) {
	// input ID
	var id uint

	// fungsi metode
	_, err := B.DekripsiRC(id)

	// fungsi testing

	assert.NotNil(t, err, "Dekripsi")
}
