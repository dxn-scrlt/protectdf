package commands

import (
	"errors"
	"os"

	"github.com/dxn-scrlt/protectdf/internal/pdf"
	"github.com/dxn-scrlt/protectdf/internal/prompt"
)

func ChangePassword(input string) error {
	userErr := errors.New("change password unsuccessful")

	oldPassword, _ := prompt.Password(false, false)

	tempFile, err := os.CreateTemp("", "protectdf-*.pdf")

	if err != nil {
		err = userErr
	} else {
		temp := tempFile.Name()
		defer os.Remove(temp)

		err = tempFile.Close()
		if err != nil {
			err = userErr
		} else {
			err = pdf.DecryptFile(input, temp, oldPassword)
			if err != nil {
				err = userErr
			} else {
				var newPassword string

				newPassword, err = prompt.Password(true, true)

				if err == nil {
					output := input

					err = pdf.EncryptFile(temp, output, newPassword)
					if err != nil {
						err = userErr
					}
				}
			}
		}
	}

	return err
}
