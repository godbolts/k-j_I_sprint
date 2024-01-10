package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput() (toEncrypt bool, encoding string, message string) {
	fmt.Println("Welcome to the Cypher Tool!")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Select operation (1/2): \n1. Encrypt.\n2. Decrypt.\n")
		sisend, _ := reader.ReadString('\n')
		sisend = strings.ToLower(strings.TrimSpace(sisend))
		if sisend == "1" || sisend == "encrypt" {
			toEncrypt = true
			break
		} else if sisend == "2" || sisend == "decrypt" {
			toEncrypt = false
			break
		}
	}
	for {
		fmt.Print("Select cypher (1/3): \n1. ROT13.\n2. Reverse.\n3. Caesar\n")
		siffer, _ := reader.ReadString('\n')
		siffer = strings.ToLower(strings.TrimSpace(siffer))
		if siffer == "1" || siffer == "rot13" {
			encoding = "rot13"
			break
		} else if siffer == "2" || siffer == "reverse" {
			encoding = "reverse"
			break
		} else if siffer == "3" || siffer == "caesar" {
			encoding = "caesar"
			break
		}
	}

	fmt.Print("Enter the message:\n")
	message, _ = reader.ReadString('\n')
	message = strings.TrimSpace(message)

	return toEncrypt, encoding, message
}

func encrypt_rot13(s string) string {
	result := []byte(s)

	for i := 0; i < len(result); i++ {
		if result[i] >= 'A' && result[i] <= 'Z' {
			result[i] = (result[i]-'A'+13)%26 + 'A'
		} else if result[i] >= 'a' && result[i] <= 'z' {
			result[i] = (result[i]-'a'+13)%26 + 'a'
		}
	}

	return string(result)
}

func encrypt_reverse(s string) string {
	const a, z = 'a', 'z'
	const A, Z = 'A', 'Z'
	reverseCharacter := func(ch rune) rune {
		if ch >= a && ch <= z {
			reversed := z - (ch - a)
			return reversed
		} else if ch >= A && ch <= Z {
			reversed := Z - (ch - A)
			return reversed
		}
		return ch
	}
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = reverseCharacter(runes[j]), reverseCharacter(runes[i])
	}
	return string(runes)
}

func decrypt_rot13(s string) string {
	return encrypt_rot13(s)
}

func decrypt_reverse(s string) string {
	return encrypt_reverse(s)
}

func encrypt_caesar(message string) string {
	shift := 3
	encrypted := make([]rune, len(message))
	for i, char := range message {
		if char >= 'A' && char <= 'Z' {
			encrypted[i] = (char-'A'+rune(shift))%26 + 'A'
		} else if char >= 'a' && char <= 'z' {
			encrypted[i] = (char-'a'+rune(shift))%26 + 'a'
		} else {
			encrypted[i] = char
		}
	}
	return string(encrypted)
}

func decrypt_caesar(message string) string {
	shift := -3
	encrypted := make([]rune, len(message))
	for i, char := range message {
		if char >= 'A' && char <= 'Z' {
			encrypted[i] = (char-'A'+rune(shift))%26 + 'A'
		} else if char >= 'a' && char <= 'z' {
			encrypted[i] = (char-'a'+rune(shift))%26 + 'a'
		} else {
			encrypted[i] = char
		}
	}
	return string(encrypted)
}

func main() {
	toEncrypt, encoding, message := getInput()
	millega := "using " + encoding
	var tegevus string
	var vastus string

	if toEncrypt {
		tegevus = "Encrypted"
		switch encoding {
		case "rot13":
			vastus = encrypt_rot13(message)
		case "reverse":
			vastus = encrypt_reverse(message)
		case "caesar":
			vastus = encrypt_caesar(message)
		}
	} else {
		tegevus = "Decrypted"
		switch encoding {
		case "rot13":
			vastus = decrypt_rot13(message)
		case "reverse":
			vastus = decrypt_reverse(message)
		case "caesar":
			vastus = decrypt_caesar(message)
		}
	}
	fmt.Printf("%s the message %s:\n%s\n", tegevus, millega, vastus)
}
