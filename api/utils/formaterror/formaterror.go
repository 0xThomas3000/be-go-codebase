package formaterror

import (
	"strings"
)

/*
  - Error formatting:
    We will like to handle errors nicely when they occur.
    The ORM(Object-Relational Mapping) that is used in the app is GORM.
    There are some error messages that are not displayed nicely, especially those that occurred when the database is hit.

    EX: when a user inputs someone else email that is already in our database,
    in an attempt to sign up, we need to prevent such action and politely
    tell the user that he can't use that email.
*/
var errorMessages = make(map[string]string)

var err error

func FormatError(errString string) map[string]string {

	if strings.Contains(errString, "username") {
		errorMessages["Taken_username"] = "Username Already Taken"
	}

	if strings.Contains(errString, "email") {
		errorMessages["Taken_email"] = "Email Already Taken"

	}
	if strings.Contains(errString, "title") {
		errorMessages["Taken_title"] = "Title Already Taken"

	}
	if strings.Contains(errString, "hashedPassword") {
		errorMessages["Incorrect_password"] = "Incorrect Password"
	}
	if strings.Contains(errString, "record not found") {
		errorMessages["No_record"] = "No Record Found"
	}

	if strings.Contains(errString, "double like") {
		errorMessages["Double_like"] = "You cannot like this post twice"
	}

	if len(errorMessages) > 0 {
		return errorMessages
	}

	if len(errorMessages) == 0 {
		errorMessages["Incorrect_details"] = "Incorrect Details"
		return errorMessages
	}

	return nil
}
