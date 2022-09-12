package utils

import "fmt"

func FailOnError(err error, msg string) error {

	return fmt.Errorf("%s: %s", msg, err)
}
