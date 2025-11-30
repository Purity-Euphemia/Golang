package main

func IsCapitalized(s string) bool {
    if s == "" {
        return false
    }

    words := []rune{}
    str := s + " " 
    start := 0

    var wordList []string
    for i, r := range str {
        if r == ' ' {
            if start != i { 
                wordList = append(wordList, s[start:i])
            }
            start = i + 1
        }
    }

        for _, word := range wordList {
        first := rune(word[0])

        if first >= 'a' && first <= 'z' {
            return false
        }

    }

    return true
}
