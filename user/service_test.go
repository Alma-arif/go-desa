package user

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

func Test_Login_Succeed(t *testing.T) {
	// Input
	input := LoginInput{}
	input.Email = "kisah@gmail.com"
	input.Password = "123"

	_, err := B.Login(input)
	assert.Nil(t, err, "Login Succeed")
}

func Test_Login_Password(t *testing.T) {
	// Input
	input := LoginInput{}
	input.Email = ""
	input.Password = "123"

	_, err := B.Login(input)

	assert.NotNil(t, err, "Login Succeed")
}

func Test_Login_Email_Password(t *testing.T) {
	// Input
	input := LoginInput{}
	input.Email = ""
	input.Password = ""

	_, err := B.Login(input)

	assert.NotNil(t, err, "Login Succeed")
}
