package cng

import (
	"fmt"
	"strings"
)

func errorWith(err error, msg ...string) error {
	if err != nil {
		m := strings.Join(msg, " ")
		return fmt.Errorf("%s:%w", m, err)
	}
	return nil
}
