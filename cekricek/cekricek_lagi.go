package main

func cipher(text string, direction int) string {

	shift, offset := rune(3), rune(26)

	// string->rune conversion
	runes := []rune(text)

	for index, char := range runes {
		switch direction {
		case -1: // encoding
			if char >= 'a'+shift && char <= 'z' ||
				char >= 'A'+shift && char <= 'Z' {
				char = char - shift
			} else if char >= 'a' && char < 'a'+shift ||
				char >= 'A' && char < 'A'+shift {
				char = char - shift + offset
			}
		case +1: // decoding
			if char >= 'a' && char <= 'z'-shift ||
				char >= 'A' && char <= 'Z'-shift {
				char = char + shift
			} else if char > 'z'-shift && char <= 'z' ||
				char > 'Z'-shift && char <= 'Z' {
				char = char + shift - offset
			}
		}

		// Above `if`s handle both upper and lower case ASCII
		// characters; anything else is returned as is (includes
		// numbers, punctuation and space).
		runes[index] = char
	}

	return string(runes)
}

// encode and decode provide the API for encoding and decoding text using
// the Caesar Cipher algorithm.
func encode(text string) string { return cipher(text, -1) }
func decode(text string) string { return cipher(text, +1) }

// A simple test
func main() {
	//println("the text is `das fuchedes 666`")
	//encoded := encode("abcd")
	//println("  encoded: " + encoded)
	//decoded := decode(encoded)
	//println("  decoded: " + decoded)
}
