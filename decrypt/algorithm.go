package decrypt

func Nimbus(str string) string {
	decryptedString := ""
	for _, c := range str {
		decryptedString += string(int(c) - 3)
	}
	return decryptedString
}
