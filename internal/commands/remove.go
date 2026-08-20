package commands

import (
	"errors"
	"strings"

	"github.com/dxn-scrlt/protectdf/internal/pdf"
	"github.com/dxn-scrlt/protectdf/internal/prompt"
)

func RemovePassword(input string) error {
	password, err := prompt.Password(true, false)

	if err == nil {
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
