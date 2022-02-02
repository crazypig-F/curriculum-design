package decoding

import "fmt"

func CheckIn(keys []string, key string) bool {
	for _, k := range keys {
		if k == key {
			return true
		}
	}
	return false
}

func Decoding(ciphertext string) {
	Codebook := NewCodebook("data/passwordFile.txt")
	keys := make([]string, 0)
	for key := range Codebook {
		keys = append(keys, key)
	}
	for i := 0; i < len(ciphertext); i++ {
		for j := i + 1; j <= len(ciphertext); j++ {
			if CheckIn(keys, ciphertext[i:j]) {
				c := Codebook[ciphertext[i:j]]
				if c == "\\s" {
					fmt.Print(" ")
				} else if c == "\\n" {
					fmt.Println()
				} else {
					fmt.Print(c)
				}
				i = j-1
				break
			}
		}
	}
}
