package two

import (
	one "github.com/bflad/go-module-one"
	three "github.com/bflad/go-module-three"
)

func Two() string {
	return "two is greater than " + one.One() + " and less than " + three.Three()
}
