package isogram

import "unicode"

func IsIsogram(word string) bool {
    mp := make(map[rune]bool)
    for _, ch := range word {
        if unicode.IsLetter(ch) {
            ch = unicode.ToLower(ch)
            if mp[ch] {
                return false
            }
            mp[ch] = true
        }
    }
    return true
}
