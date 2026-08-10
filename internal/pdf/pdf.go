package pdf

import (
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

const aesKeyLength = 256

func EncryptFile(input, output, password string) error {
	conf := model.NewAESConfiguration(password, password, aesKeyLength)

	err := api.EncryptFile(input, output, conf)

	return err
}

func DecryptFile(input, output, password string) error {
	conf := model.NewAESConfiguration(password, password, aesKeyLength)

	err := api.DecryptFile(input, output, conf)

	return err
}
