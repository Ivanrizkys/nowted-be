package helper

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ErrMsgFormat(err error) string {
	t := cases.Title(language.English)
	errStringSlice := strings.Split(err.Error(), "\n")
	if len(errStringSlice) < 2 {
		return errStringSlice[0]
	}
	errStringSlice[0] = t.String(errStringSlice[0])
	return strings.Join(errStringSlice[:], " : ")
}
