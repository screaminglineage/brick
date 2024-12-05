package main

import (
    "fmt"
    "os"
)

const TAPE_LEN = 30000

// TODO: report an error rather than returning -1 when no matching bracket is found
func find_matching_close_bracket(input string, index int) int {
    level := 0
    for i, ch := range input[index:] {
        if ch == '[' {
            level += 1
        } else if ch == ']' {
            level -= 1
            if level == 0 {
                return index + i
            }
        }
    }
    return -1
}

// TODO: report an error rather than returning -1 when no matching bracket is found
func find_matching_open_bracket(input string, index int) int {
    level := 0
    for i := index; i >= 0; i-- {
        ch := input[i]
        if ch == ']' {
            level += 1
        } else if ch == '[' {
            level -= 1
            if level == 0 {
                return i
            }
        }
    }
    return -1
}

func execute(input string, tape []byte) {
    ip := 0
    tp := 0
    for {
        if ip >= len(input) {
            return
        }
        if tp < 0 || tp >= TAPE_LEN {
            fmt.Println("The brainfuck got you!")
            return
        }
        op := input[ip]
        switch op {
            case '>': tp += 1
            case '<': tp -= 1
            case '+': tape[tp] += 1
            case '-': tape[tp] -= 1
            case '[':
                if tape[tp] == 0 {
                    ip = find_matching_close_bracket(input, ip)
                }
            case ']':
                if tape[tp] != 0 {
                    ip = find_matching_open_bracket(input, ip)
                }
            case '.': fmt.Print(string(tape[tp]))
            case ',':
                panic("Taking input is not yet implemented")
        }
        ip += 1
    }
}


func main() {
    if len(os.Args) < 2 {
        fmt.Println("Not enough arguments")
        return
    }
    input, err := os.ReadFile(os.Args[1])
    if err != nil {
        fmt.Println(err)
        return
    }
    tape := make([]byte, TAPE_LEN)
    execute(string(input), tape)
}





