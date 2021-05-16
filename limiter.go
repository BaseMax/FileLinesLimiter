// Max Base
// https://github.com/BaseMax/FileLinesLimiter

package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "bufio"
)

func readFileLines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}

func myError() {
    fmt.Print("Press 'Enter' to continue...")
    bufio.NewReader(os.Stdin).ReadBytes('\n')
    os.Exit(1)
}

func main() {
    fmt.Println("Enter input path file: ")
    var FILE string
    _, err := fmt.Scanf("%s", &FILE)
    if err != nil {
        log.Fatalf("read user path file!: %s", err)
        myError()
    }

    // TODO: remove " from first and last of the file path!

    lines, err := readFileLines(FILE)
    if err != nil {
        log.Fatalf("readInput: %s", err)
        myError()
    }

    // for i, _ := range lines {
    //     fmt.Println(i, lines[i])
    // }

    fmt.Println("Enter start line: ")
    var START int
    _, err = fmt.Scanf("%d", &START)
    if err != nil {
        log.Fatalf("read start index!: %s", err)
        myError()
    }

    fmt.Println("Enter end line: ")
    var END int
    _, err = fmt.Scanf("%d", &END)
    if err != nil {
        log.Fatalf("read end index!: %s", err)
        myError()
    }

    fmt.Printf("Creating result output file with lines of [%d, %d]\n", START, END)

    file, err := os.Create("limiter.out.txt")
    if err != nil {
        log.Fatalf("writeOutput: %s", err)
        myError()
    }

    w := bufio.NewWriter(file)
    for i := START-1; i < END; i++ {
      fmt.Fprintln(w, lines[i])
    }
    
    w.Flush()
    file.Close()

    fmt.Println("Done.")
}
