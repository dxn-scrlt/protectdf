package commands

import (
	"errors"
	"fmt"
	"os"

	"github.com/dxn-scrlt/protectdf/internal/pdf"
)

func ChangePassword(input string) error {
	userErr := errors.New("change password unsuccessful")

	var oldPassword string

	fmt.Print("Enter current password: ")
	fmt.Scanln(&oldPassword)

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
				var retype string

				fmt.Print("Enter new password: ")
				fmt.Scanln(&newPassword)

				fmt.Print("Re-enter new password: ")
				fmt.Scanln(&retype)

				if newPassword != retype {
					err = errors.New("passwords do not match")
				} else {
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
