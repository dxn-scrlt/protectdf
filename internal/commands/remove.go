package commands

import (
	"errors"
	"fmt"
	"strings"

	"github.com/dxn-scrlt/protectdf/internal/pdf"
)

func RemovePassword(input string) error {
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
		var output string
		protectedSuffix := "-protected.pdf"

		if strings.HasSuffix(input, protectedSuffix) {
			output = strings.TrimSuffix(input, protectedSuffix) + ".pdf"
		} else {
			output = input
		}

		err = pdf.DecryptFile(input, output, password)
		if err != nil {
			err = errors.New("remove password unsuccessful")
		}
	}

	return err
}
