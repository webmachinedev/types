package types

type Function struct{
	Name string
	GoPackageName string
	Inputs []struct{
		Name string
		Type string
	}
	Outputs []struct{
		Name string
		Type string
	}
}

// func (f *Function) GoPackageName() string {
// 	validchars := "abcdefghijklmnopqrstuvwxyz"
// 	name := ""
// 	for _, char := range strings.ToLower(f.Name) {
// 		if strings.ContainsRune(validchars, char) {
// 			name += string(char)
// 		}
// 	}
// 	return name
// }
