package pkg

import (
	"strings"
	"unicode"
)

/**
 * Make a string's first character lowercase
 * @param string str <p>
 * The input string.
 * </p>
 * @return string the resulting string.
 */
func LCFirst(str string) string {
	for _, v := range str {
		u := string(unicode.ToLower(v))
		return u + str[len(u):]
	}
	return ""
}

/**
 * Uppercase the first character of each word in a string
 * @param string str <p>
 * The input string.
 * </p>
 * @param string delimiters [optional] <p>
 * @return string the modified string.
 */
func UCWords(str string) string {
	return strings.Title(str)
}

/**
 * Check for uppercase character(s)
 * @param string str <p>
 * The input string.
 * </p>
 * @param string<p>
 * @return bool.
 */
func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

/**
 * Check for lowercase character(s)
 * @param string str <p>
 * The input string.
 * </p>
 * @param string<p>
 * @return bool.
 */
func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

/**
 * Convert the given string to lower-case.
 *
 * @param  string  $value
 * @return string
 */
func Lower(value string) string {
	return strings.ToLower(value)
}
