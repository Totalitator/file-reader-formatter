package pkg

import (
	"strings"
)

const (
	methodPrint     = "print"
	methodUpperCase = "uppercase"
	methodLowerCase = "lowercase"
)

type Formatter interface {
	Format() string
}

type Print struct {
	Text string
}

type UpperCase struct {
	Text string
}

type LowerCase struct {
	Text string
}

func (pr *Print) Format() string {
	return pr.Text
}

func (up *UpperCase) Format() string {
	return strings.ToUpper(up.Text)
}

func (lw *LowerCase) Format() string {
	return strings.ToLower(lw.Text)
}
