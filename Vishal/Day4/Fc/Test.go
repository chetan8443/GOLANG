pacakage Fc
 
var AB int = 3  // global variable declared

func IsRotation(str1, str2 string) bool {
	// If strings have different lengths, return false
	if len(str1) != len(str2) {
		return false
	}

	
	concat := str1 + str1           // Concatenate str1 with itself to check for rotations

	// Check if str2 is a substring of the concatenated string
	if (len(concat) >= len(str2)) && (concat[:len(str2)] == str2) {
		return true
	}

	return false
}
