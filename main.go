package main

import (
    "fmt"
    "os"
    "time"
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
    
    s := time.Now()
    b, ok := b.JustSolveIt()
    f := time.Now()
    elapsed := f.Sub(s)
    
    if !ok {
        fmt.Println("ERROR: CANNOT SOLVE PUZZLE.")
        os.Exit(1)
    }
    b.Print()
    
    fmt.Printf("%v taken to solve the puzzle.\n\n", elapsed)
    
    if len(os.Args) == 3 {
        writeto := os.Args[2]
        b.Write(writeto)
    }
}
