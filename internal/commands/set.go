package commands

import (
	"errors"
	"fmt"
	"strings"

	"github.com/dxn-scrlt/protectdf/internal/pdf"
)

func SetPassword(input string) error {
	var password string
	var retype string

	fmt.Print("Enter password: ")
	fmt.Scanln(&password)

	fmt.Print("Re-enter password: ")
	fmt.Scanln(&retype)

	var err error

	if password != retype {
		err = errors.New("passwords do not match")
	} else {
		output := strings.TrimSuffix(input, ".pdf") + "-protected.pdf"

		err = pdf.EncryptFile(input, output, password)
		if err != nil {
			err = errors.New("set password unsuccessful")
		}
	}

	return err
}
