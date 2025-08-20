package scrabble
import "strings"
func Score(word string) int {
    word = strings.ToUpper(word)
	runes := []rune(word)
    ans:=0
    for i:=0;i<len(runes);i++{
        switch runes[i]{
            case 'A','E','I','O','U','L','N','R','S','T':
                ans+=1
            case 'D','G':
                ans+=2
            case 'B','C','M','P':
                ans+=3
            case 'F','H','V','W','Y':
                ans+=4
            case 'K':
                ans+=5
            case 'J','X':
                ans+=8
            case 'Q','Z':
                ans+=10
        }
    }
    return ans
}
