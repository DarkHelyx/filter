package main

import (
    "fmt"
    "log"
    "os"
    "bufio"
    "regexp"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        line := scanner.Bytes()
        numArgs := len(os.Args)
        filtered := false

        for i := 0; i < numArgs; i++ {
            regex, err := regexp.Compile(os.Args[i])
            if err != nil {
                log.Fatalln(err)
            }

            // If it matches the filter,
            if regex.Match(line) {
                filtered = true
                break
            }
        }

        if !filtered {
            fmt.Printf("%s\n", line)
        }
    }
}
