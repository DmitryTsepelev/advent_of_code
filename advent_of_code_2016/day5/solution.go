package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func findPassword(str string) string {
	password := ""

	idx := 0
	for len(password) < 8 {
		hash := GetMD5Hash(str + strconv.Itoa(idx))
		if strings.HasPrefix(hash, "00000") {
			password += string(hash[5])
		}
		idx++
	}

	return password
}

func findPassword2(str string) string {
	password := make([]rune, 8)
	filled := make([]bool, 8)

	idx := -1
	for {
		idx++
		hash := GetMD5Hash(str + strconv.Itoa(idx))

		if !strings.HasPrefix(hash, "00000") {
			continue
		}

		strPos := hash[5]
		pos, _ := strconv.Atoi(string(strPos))

		if strPos >= '0' && strPos <= '7' && !filled[pos] {
			password[pos] = rune(hash[6])
			filled[pos] = true

			done := true
			for _, isFilled := range filled {
				if !isFilled {
					done = false
				}
			}

			if done {
				break
			}
		}
	}

	return string(password)
}

func main() {
	str := "uqwqemis"

	fmt.Println("Solution 1 is", findPassword(str))
	fmt.Println("Solution 2 is", findPassword2(str))
}
