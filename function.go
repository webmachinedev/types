package types

import "strings"

type Function struct {
	Name string
}

func (f *Function) GoPackageName() string {
	validchars := "abcdefghijklmnopqrstuvwxyz"
	name := ""
	for _, char := range f.Name {
		if strings.ContainsRune(validchars, char) {
			name += string(char)
		}
	}
	return name
}
