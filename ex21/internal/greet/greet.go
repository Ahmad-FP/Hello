package greet

import(
	"strings"
)

func Hello(name string) string{
	return "Halo, " + name
}

func shout(s string) string{
	return strings.ToUpper(s)
}