package del2

import (
	"rsc.io/quote"
)

func Quote() string {
	return quote.Hello() + "\n" + quote.Go()
}