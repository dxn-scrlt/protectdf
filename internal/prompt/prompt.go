package prompt

import (
	"errors"
	"fmt"
)

func Password(shouldRetype, isNew bool) (string, error) {
	var password string
	var retype string

	if !shouldRetype {
		fmt.Print("Enter current password: ")
		fmt.Scanln(&password)

		return password, nil
	}

	var promptName string

	if isNew {
		promptName = "new password"
	} else {
		promptName = "password"
	}

	fmt.Printf("Enter %s: ", promptName)
	fmt.Scanln(&password)

	fmt.Printf("Re-enter %s: ", promptName)
	fmt.Scanln(&retype)

	if password != retype {
		return "", errors.New("passwords do not match")
	}

	return password, nil
}
