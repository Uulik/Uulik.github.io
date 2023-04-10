package main

import (
    "fmt"
)

func main() {
    // Take input from user
    var n int
    fmt.Println("Enter the length of the list:")
    fmt.Scan(&n)
    fmt.Println("Enter the list of integers:")
    list := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&list[i])
    }

    // Sort the list using bubble sort algorithm
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            if list[i] > list[j] {
                list[i], list[j] = list[j], list[i]
            }
        }
    }

    // Print the sorted list
    fmt.Println("Sorted list of integers:")
    for _, val := range list {
        fmt.Printf("%d ", val)
    }
}
//Write a program that takes a list of integers as input and sorts the list in ascending order
