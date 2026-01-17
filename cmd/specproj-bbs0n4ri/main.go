package main

import (
    "fmt"
    "io"
    "os"

    "'$name'/internal/analyze"
)

func main() {
    data, _ := io.ReadAll(os.Stdin)
    tokens := analyze.Tokens(string(data))
    counts := analyze.Count(tokens)
    for k, v := range counts {
        fmt.Printf("%s=%d\n", k, v)
    }
}
