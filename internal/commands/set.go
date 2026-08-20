package commands

import (
	"errors"
	"strings"

	"github.com/dxn-scrlt/protectdf/internal/pdf"
	"github.com/dxn-scrlt/protectdf/internal/prompt"
)

func SetPassword(input string) error {
	password, err := prompt.Password(true, false)

	if err == nil {
		output := strings.TrimSuffix(input, ".pdf") + "-protected.pdf"

		err = pdf.EncryptFile(input, output, password)
		if err != nil {
			err = errors.New("set password unsuccessful")
		}
	}

	return err
}
