package services

import "regexp"

// Validates a user id. A user id can only contain letters, numbers and underscores, and must be between 5 and 12 characters long.
func validateUserId(id string) bool {
	regex := regexp.MustCompile(`^[[:alnum:]_]{5,12}$`)
	return regex.MatchString(id)
}

// Validates a user name. A user name can only contain letters, numbers and underscores, and must be between 5 and 12 characters long.
func validateUserName(name string) bool {
	regex := regexp.MustCompile(`^[[:alnum:]_]{5,12}$`)
	return regex.MatchString(name)
}

// Validates a video id. A video id can only contain letters, numbers, dashes and underscores, and must be between 11 characters long.
func validateVidId(id string) bool {
	regex := regexp.MustCompile(`^[0-9A-Za-z_-]{5,10}[048AEIMQUYcgkosw]$`) // Remove "5,"
	return regex.MatchString(id)
}

// Validates a video title is not empty.
func validateVidTitle(name string) bool {
	return name != ""
}
