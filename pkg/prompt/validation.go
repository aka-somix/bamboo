package prompt

import "errors"


func ValidateName(name string) error {

	if len(name) > 50 {
		return errors.New("name is too long")
	}

	return nil;
}


func ValidateDescription(description string) error {

	if len(description) > 100 {
		return errors.New("description is too long")
	}

	return nil;
}


func ValidatePath(path string) error {

	return nil;
}
