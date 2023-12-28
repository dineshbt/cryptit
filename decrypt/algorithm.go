package decrypt

func Nimbus(str string) string {
	decryptedStr := ""
	for _, char := range str {
		asciiCode := int(char)
		character := string(asciiCode - 3)
		decryptedStr += character
	}
	return decryptedStr
}
