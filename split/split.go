package asciiart

// SplitBySep returns slice of a slice of byte. Each element contains 8 lines of the ASCII graphical representation from the textfile
func SplitBySep(bannerAll []byte, sep string) [][]byte {
	var bannerLines [][]byte
	for i := 0; i < len(bannerAll); i++ {
		if string(bannerAll[i]) == sep {
			bannerLines = append(bannerLines, bannerAll[:i])
			if string(bannerAll[i+1]) == sep {
				bannerAll = bannerAll[i+2:]
			} else {
				bannerAll = bannerAll[i+1:]
			}
			i = 0
		}
	}
	bannerLines = append(bannerLines, bannerAll)
	return bannerLines
}
