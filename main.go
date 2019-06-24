package main

import (
    "fmt"
    "os"
)


var DEBUG bool = false


func main() {
    if len(os.Args) < 2 || len(os.Args) > 3  {
        fmt.Println("ERROR: WRONG NUMBER OF COMMAND LINE ARGUMENTS.")
        os.Exit(1)
    }
    
    filename := os.Args[1]
    b, err := Read(filename)
    if err != nil {
        fmt.Println("ERROR: CANNOT READ FILE.")
        os.Exit(1)
    }
    
    b.Print()
    b, ok := b.JustSolveIt()
    if !ok {
        fmt.Println("ERROR: CANNOT SOLVE PUZZLE.")
        os.Exit(1)
    }
    b.Print()
    
    if len(os.Args) == 3 {
        writeto := os.Args[2]
        b.Write(writeto)
    }
}
