package encrypt

func Nimbus(str string) string {
	encryptedStr := ""
	for _, c := range str {
		ascii_code := int(c)
		character := string(ascii_code + 3)
		encryptedStr += character
	}
	return encryptedStr
}
