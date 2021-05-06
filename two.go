package two

import (
	one "github.com/bflad/go-module-one"
)

func Two() string {
	return "two is greater than " + one.One()
}
