package abbreviate

import (
	"strings"
)

var trigraphs = [...]string{ "chr", "sch" }
var triblends = [...]string{ "shr", "spl", "spr", "squ", "str", "thr" }
var digraphs = [...]string{ "ch", "gh", "gn", "kn", "ph", "qu", "sh", "th", "wh", "wr" }
var diblends = [...]string{ "bl", "br", "cl", "cr", "fl", "fr", "gl", "gr", "pl", "pr", "sc", "sl", "sm", "sn", "sp", "st" }

func Abbreviate(str string, length int) string {
	if length <= 0 { length = 3 }

	// Sanity check
	if length >= len(str) { return str }

	// Trim the string
	str = strings.Trim(str, " ")
	if length >= len(str) { return str }

	stringSlices := strings.Split(str, "")
	runeImportance := make([]int, len(stringSlices))
	scannedRunesCount := 0

	// Mark (di/tri)graphs and (di/tri)blends as high importance
	triDiImportance := 1
	for i := 0; i < len(stringSlices); i++ {
		if i > 0 {
			pre := stringSlices[i-1]
			if !strings.ContainsAny(pre, " _-") { continue }
		}

		// Definitely a new word (hopefully)

		// Trigraphs and Triblends
		if i + 2 < len(stringSlices) {
			if runeImportance[i] > 0 { continue }
			c := strings.ToLower(strings.Join(stringSlices[i:i+3], ""))

			tri := append(trigraphs[:0], triblends[:]...)
			for _, t := range tri {
				if strings.HasPrefix(c, t) {
					runeImportance[i+0] = triDiImportance+0
					runeImportance[i+1] = triDiImportance+1
					runeImportance[i+2] = triDiImportance+2
					triDiImportance += 3
					scannedRunesCount += 3
					break
				}
			}
		}

		// Digraphs and Diblends
		if i + 1 < len(stringSlices) {
			if runeImportance[i] > 0 { continue }
			c := strings.ToLower(strings.Join(stringSlices[i:i+2], ""))

			di := append(digraphs[:0], diblends[:]...)
			for _, d := range di {
				if strings.HasPrefix(c, d) {
					runeImportance[i+0] = triDiImportance+0
					runeImportance[i+1] = triDiImportance+1
					triDiImportance += 2
					scannedRunesCount += 2
					break
				}
			}
		}
	}

	currentImportance := len(runeImportance)
	containsChecks := []string{
		" _-",
		"aeiou",
		"bcdfghjklmnpqrstvwxyz",
		"AEIOU",
		"BCDFGHJKLMNPQRSTVWXYZ0123456789",
	}

	outCheck:
	for _, check := range containsChecks {
		for i := len(stringSlices) - 1; i >= 0; i-- {
			if runeImportance[i] > 0 { continue; }
			v := stringSlices[i]

			if strings.ContainsAny(v, check) {
				runeImportance[i] = currentImportance
				currentImportance -= 1
				scannedRunesCount += 1
			}
			
			if scannedRunesCount == len(stringSlices) { break outCheck }
		}
	}

	// Select runes that are below threshold (length)
	var finalString strings.Builder
	for i := range stringSlices {
		if runeImportance[i] > length { continue }

		finalString.WriteString(stringSlices[i])
	}

	return finalString.String()
}
