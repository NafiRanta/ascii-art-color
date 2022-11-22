package asciiart

//colorReset := "\033[0m"

func ColorPallete(s string) string {

	var colorCode string

	switch s {
	case "red":
		colorCode = string("\033[31m")
	case "green":
		colorCode = string("\033[32m")
	case "yellow":
		colorCode = string("\033[33m")
	case "blue":
		colorCode = string("\033[34m")
	case "purple":
		colorCode = string("\033[35m")
	case "cyan":
		colorCode = string("\033[36m")
	case "white":
		colorCode = string("\033[37m")
	case "grey":
		colorCode = string("\033[30m")
	}
	return colorCode
}
