package analyze

import "strings"

func Tokens(input string) []string {
    parts := strings.Fields(input)
    return parts
}

func Count(tokens []string) map[string]int {
    out := map[string]int{}
    for _, t := range tokens {
        out[t]++
    }
    return out
}
