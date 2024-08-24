package main

import (
	"errors"
	"strconv"
	"strings"
)

// Caesar Cipher
func caesarCipher(text, keyStr, action string) (string, error) {
	shift, err := strconv.Atoi(keyStr)
	if err != nil {
		return "", err
	}

	if action == "decrypt" {
		shift = -shift
	}

	shift = shift % 26
	result := ""
	for _, char := range text {
		if char >= 'a' && char <= 'z' {
			// Convert rune to int for arithmetic
			shifted := (int(char-'a') + shift + 26) % 26
			result += string(rune(shifted + 'a'))
		} else if char >= 'A' && char <= 'Z' {
			// Convert rune to int for arithmetic
			shifted := (int(char-'A') + shift + 26) % 26
			result += string(rune(shifted + 'A'))
		} else {
			result += string(char)
		}
	}
	return result, nil
}

// Affine Cipher
func affineCipher(text, keyStr, action string) (string, error) {
	keys := strings.Split(keyStr, ",")
	if len(keys) != 2 {
		return "", errors.New("key must be two comma-separated integers")
	}

	a, err := strconv.Atoi(keys[0])
	if err != nil {
		return "", err
	}
	b, err := strconv.Atoi(keys[1])
	if err != nil {
		return "", err
	}

	if action == "decrypt" {
		a = modInverse(a, 26)
		b = -b
	}

	result := ""
	for _, char := range text {
		if char >= 'a' && char <= 'z' {
			// Convert rune to int for arithmetic
			shifted := (a*(int(char-'a')) + b) % 26
			if shifted < 0 {
				shifted += 26
			}
			result += string(rune(shifted + 'a'))
		} else if char >= 'A' && char <= 'Z' {
			// Convert rune to int for arithmetic
			shifted := (a*(int(char-'A')) + b) % 26
			if shifted < 0 {
				shifted += 26
			}
			result += string(rune(shifted + 'A'))
		} else {
			result += string(char)
		}
	}
	return result, nil
}

func modInverse(a, m int) int {
	m0, x0, x1 := m, 0, 1
	if m == 1 {
		return 0
	}
	for a > 1 {
		q := a / m
		m, a = a%m, m
		x0, x1 = x1-q*x0, x0
	}
	if x1 < 0 {
		x1 += m0
	}
	return x1
}
