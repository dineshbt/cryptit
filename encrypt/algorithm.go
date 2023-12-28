package encrypt

func Nimbus(str string) string {
	encryptedStr := ""
	for _, char := range str {
		asciiCode := int(char)
		character := string(asciiCode + 3)
		encryptedStr += character
	}
	return encryptedStr
}
