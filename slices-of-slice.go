package main

import (
    "fmt"
    "strings"
)

func main() {
    // create a tic-tack-toe board
    game := [][]string{
        []string{"_","_","_"},
        []string{"_","_","_"},
        []string{"_","_","_"},
    }
    
    // The players take turns.
    game[0][0] = "X"
    game[2][2] = "O"
    game[2][0] = "X"
    game[1][0] = "0"
    game[0][2] = "X"
    
    printBoard(game)   

    //Another example
    values := [][]int{}

    // These are the first two rows.
    row1 := []int{1, 2, 3}
    row2 := []int{4, 5, 6}

    // Append each row to the two-dimensional slice.
    values = append(values, row1)
    values = append(values, row2)

    // Display first row.
    fmt.Println("Row 1")
    fmt.Println(values[0])

    // Display second row.
    fmt.Println("Row 2")
    fmt.Println(values[1])

    // Access an element.
    fmt.Println("First element")
    fmt.Println(values[0][0])

    // Display entire slice.
    fmt.Println("Values")
    fmt.Println(values) 
}

func printBoard(s [][]string) {
    for i := 0; i < len(s); i++ {
        fmt.Printf("%s\n", strings.Join(s[i], " "))
    }
}