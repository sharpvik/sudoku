package main

import (
    "bufio"
    "fmt"
    "os"
    "time"
)


type board struct {
    grid    [9][9]int
    poss    [9][9][10]bool
}


func New(input [9][9]int) board {
    b := board{}
    b.grid = input
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            b.poss[i][j] = [10]bool{false, true, true, true, true, true, true, true, true, true}
        }
    }
    return b
}


func Read(filename string) (board, error) {
    var input [9][9]int
    file, err := os.Open(filename)
    if err != nil {
        return New(input), err
    }
    defer file.Close()
    rdr := bufio.NewReader(file)
    for i := 0; i < 9; i++ {
        for j := 0; j < 10; j++ {
            bte, _ := rdr.ReadByte()
            if '0' <= bte && bte <= '9' {
                n := bte - '0'
                input[i][j] = int(n)
            }
        }
    }
    return New(input), nil
}


func (b board) Write(filename string) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    wrtr := bufio.NewWriter(file)
    for i := 0; i < 9; i++ {
        if i % 3 == 0 && i != 0 {
            wrtr.WriteByte('\n')
        }
        for j := 0; j < 10; j++ {
            if j == 9 {
                wrtr.WriteByte('\n')
                wrtr.Flush()
            } else {
                wrtr.WriteByte( '0' + byte(b.grid[i][j]) )
                wrtr.WriteByte(' ')
                if (j + 1) % 3 == 0 {
                    wrtr.WriteByte(' ')
                    wrtr.WriteByte(' ')
                }
            }
        }
    }
    return nil
}


func (b board) Print() {
    fmt.Println()
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if (j + 1) % 3 != 0 {
                fmt.Printf("%d ", b.grid[i][j])
            } else {
                fmt.Printf("%d\t", b.grid[i][j])
            }
        }
        if (i +  1) % 3 != 0 {
            fmt.Print("\n")
        } else {
            fmt.Print("\n\n")
        }
    }
    fmt.Println()
}


func (b board) Solved() bool {
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if b.grid[i][j] == 0 {
                return false
            }
        }
    }
    return true
}


func (b board) EnhanceCell(y, x int) board {
    // Row
    for i := 0; i < 9; i++ {
        b.poss[y][x][ b.grid[y][i] ] = false
    }
    
    // Column
    for i := 0; i < 9; i++ {
        b.poss[y][x][ b.grid[i][x] ] = false
    }
    
    // Segment
    itsSegment := SEGMENTSGRID[y][x]
    segmentsCoords := SEGMENTSCOORDS[itsSegment]
    for _, pair := range segmentsCoords {
        _y := pair[0]
        _x := pair[1]
        b.poss[y][x][ b.grid[_y][_x] ] = false
    }
    
    return b
}


func (b board) Enhanced() (board, int, bool) {
    changes := 0
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            
            // for each empty cell
            if b.grid[i][j] == 0 {
                b = b.EnhanceCell(i, j)
                poss := b.poss[i][j]
                var arr []int
                for i, v := range poss {
                    if v {
                        arr = append(arr, i)
                    }
                }
                if len(arr) == 0 {
                    return board{}, 0, false
                } else if len(arr) == 1 {
                    b.grid[i][j] = arr[0]
                    changes++
                }
            }
            
        }
    }
    return b, changes, true
}


func (b board) FindEmptyCell() (y, x int) {
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if b.grid[i][j] == 0 {
                return i, j
            }
        }
    }
    return -1, -1
}


func (b board) Insert(v, y, x int) board {
    b.grid[y][x] = v
    return b
}


func (b board) GuessRecursively() (board, bool) {
    _y, _x := b.FindEmptyCell()
    var poss []int
    for i, v := range b.poss[_y][_x] {
        if v {
            poss = append(poss, i)
        }
    }
    
    if DEBUG {
        fmt.Println("Guessing recursively...")
        b.Print()
        fmt.Printf("Empty cell found at (%d, %d)\n", _y, _x)
        fmt.Print("Possible guesses: ")
        fmt.Println(poss)
        time.Sleep(8 * time.Second)
    }
    
    for _, p := range poss {
        n := b.Insert(p, _y, _x)
        
        if DEBUG {
            fmt.Printf("After inserting %d we've got...\n", p)
            n.Print()
            time.Sleep(8 * time.Second)
        }
            
        n, ok := n.JustSolveIt()
        if ok {
            return n, true
        }
    }
    return board{}, false
}


func (b board) JustSolveIt() (board, bool) {
    var (
        changes int
        ok      bool
    )
    for !b.Solved() {
        b, changes, ok = b.Enhanced()
        if !ok {
            return b, false
        }
        
        if changes == 0 {
            b, ok = b.GuessRecursively()
            if !ok {
                return b, false
            }
        }
        
    }
    return b, true
}
