package asciiart

import (
	checker "asciiart/checker"
	color "asciiart/color"
	"strconv"
	"strings"
)

// ========= PRINT THE FINAL RESULT =========
func PrintFinalResult(inputSlice []string, bannerLines [][]byte, args []string) string {
	var result string
	var temp string
	colorName := strings.TrimPrefix(args[0], "--color=")
	if checker.AllEmpty(inputSlice) {
		for l := 1; l < len(inputSlice); l++ {
			result += "\n"
		}
		return result
	}
	for i := 0; i < len(inputSlice); i++ {
		if len(inputSlice[i]) == 0 {
			result += "\n"
		} else {
			// ========== if no substring specified, print colored word ==========
			if len(args) == 2 {
				for j := 0; j < 8; j++ {
					for k := 0; k < len(inputSlice[i]); k++ {
						result += color.ColorPallete(colorName) + (string(bannerLines[(int(inputSlice[i][k]-32)*8)+j]))
					}
					result += "\n"
				}
				// ========== if substring is specified ==========
			} else if len(args) == 3 {
				// ========== if substring is only numeral ==========
				if checker.IsNumeral(args[2]) {
					for j := 0; j < 8; j++ {
						for k := 0; k < len(inputSlice[i]); k++ {
							if k == atoi(args[2]) {
								result += color.ColorPallete(colorName) + (string(bannerLines[(int(inputSlice[i][k]-32)*8)+j])) + string("\033[0m")
							} else {
								result += (string(bannerLines[(int(inputSlice[i][k]-32)*8)+j])) + string("\033[0m")

							}
						}
						result += "\n"
					}
					return result
				}
				// ========== if substring is numeral with colon ==========
				if checker.IsNumeralOrColon(args[2]) {
					firstnm := atoi(args[2][:strings.Index(args[2], ":")])
					secondnm := atoi(args[2][strings.Index(args[2], ":")+1:])
					second := false
					if secondnm > 0 {
						second = true
					}
					for j := 0; j < 8; j++ {
						for k := 0; k < len(inputSlice[i]); k++ {
							if k >= firstnm && (k < secondnm || !second) {
								result += color.ColorPallete(colorName) + (string(bannerLines[(int(inputSlice[i][k]-32)*8)+j])) + string("\033[0m")
							} else {
								result += (string(bannerLines[(int(inputSlice[i][k]-32)*8)+j])) + string("\033[0m")
							}
						}
						result += "\n"
					}
					return result
				}
				for j := 0; j < 8; j++ { // for each line of a character
					for k := 0; k < len(inputSlice[i]); k++ { // for each character in the word
						if inputSlice[i][k] == args[2][0] || len(temp) > 0 { // if the character is the same as the substrings first character
							if len(inputSlice[i])-k >= len(args[2]) { // if the length of the rest of the word is greater than the length of the substring
								for l := 0; l < len(args[2]); l++ { // for each character in the substring
									if k < len(inputSlice[i]) {
										if inputSlice[i][k+l] == args[2][l] { // if the character is the same as the substrings character
											temp += string(args[2][l])
										}
									}
									if len(temp) == len(args[2]) { // if the substring is found
										for m := 0; m < len(temp); m++ { // add the colored substring to the result
											result += color.ColorPallete(colorName) + (string(bannerLines[(int(temp[m]-32)*8)+j])) + string("\033[0m")
										}
										temp = ""
										k += len(args[2])
									}

								}
							}
						}
						if len(temp) > 0 {
							for m := 0; m < len(temp); m++ {
								result += (string(bannerLines[(int(temp[m]-32)*8)+j])) + string("\033[0m")
							}
							k += len(temp)
							temp = ""
						}
						if k < len(inputSlice[i]) {
							if !(inputSlice[i][k] == args[2][0]) {
								result += (string(bannerLines[(int(inputSlice[i][k]-32)*8)+j])) + string("\033[0m")
							} else {
								k--
							}

						}
					}
					result += "\n"
				}
			}
		}
	}
	return result
}

func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
