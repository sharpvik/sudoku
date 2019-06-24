# Sudoku Solver

This is a Sudoku Solver program that aims to solve **any** given sudoku puzzle.
It doesn't use any fancy neural networks or anything of that sort... just the
good old recursive algorithm with backtracking. 

Wrote it in `Go` as `C++` is no longer my favourite -- it's too inconsistent and
overcomplicated. Sorry, not sorry.

## How To Use

```terminal
make
sudoku path/to/input_file.txt optional/path/to/output_file.txt
```

Input must be strictly formatted for it to be processed properly. No spaces
between the numbers, no tabs or newlines to separate squares. Each file must
only have 10 lines -- 9 lines with numbers + one final line displayed due to
the newline `\n` character at the end of line 9. Empty cells are represented 
by `0`s.

This is a valid input example:

```
003020600
900305001
001806400
008102900
700000008
006708200
002609500
800203009
005010300

```

You can find more examples in the *inputs* folder. If you decide to output
the solution to a file, you are lucky -- I made outputs look ok, not just like
a mess of numbers... It looks decent, doesn't it?

```
4 8 3   9 2 1   6 5 7   
9 6 7   3 4 5   8 2 1   
2 5 1   8 7 6   4 9 3   

5 4 8   1 3 2   9 7 6   
7 2 9   5 6 4   1 3 8   
1 3 6   7 9 8   2 4 5   

3 7 2   6 8 9   5 1 4   
8 1 4   2 5 3   7 6 9   
6 9 5   4 1 7   3 8 2   

```

Here is an example of an [input file] and [output file] if 
you ever want to check it out.

[input file]: inputs/input01.txt
[output file]: outputs/output01.txt

## Materials

I have even provided some sudoku puzzles for you to solve if you ever dare to 
try. As you may have guessed already, they are all in the *inputs* folder.
There is also a Python script called *solve_all.py* that solves all puzzles from
the *inputs* folder and writes solutions to files in the *outputs* folder.