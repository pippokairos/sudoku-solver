# Sudoku solver

A sudoku solver written in Go.

## Prerequisites

- Install [Go](https://go.dev)

## Clone the project

```
git clone https://github.com/pippokairos/sudoku-solver.git
```

## Build the executable (optional)

```
go build
```

## Run the script

```
go run ./sudoku-solver -path ./files/sudokus
```

or, if you compiled

```
./sudoku-solver -path ./files/sudokus
```

## Output

```
2024/01/03 22:38:25 Found 54 grid(s) in ./files/sudokus
2024/01/03 22:38:25 Solving grid 1...
2024/01/03 22:38:25 Solved:
0 1 0 | 0 2 0 | 7 0 6         4 1 9 | 8 2 5 | 7 3 6
7 0 0 | 9 1 3 | 0 4 0         7 5 6 | 9 1 3 | 2 4 8
3 8 0 | 0 0 4 | 0 0 1         3 8 2 | 6 7 4 | 9 5 1
- - - + - - - + - - -         - - - + - - - + - - -
0 0 0 | 0 0 7 | 0 1 0         6 3 4 | 2 8 7 | 5 1 9
5 0 0 | 1 0 9 | 0 0 3   --->  5 2 7 | 1 4 9 | 8 6 3
0 9 0 | 5 0 0 | 0 0 0         8 9 1 | 5 3 6 | 4 2 7
- - - + - - - + - - -         - - - + - - - + - - -
2 0 0 | 3 0 0 | 0 9 4         2 7 8 | 3 5 1 | 6 9 4
0 4 0 | 7 6 2 | 0 0 5         9 4 3 | 7 6 2 | 1 8 5
1 0 5 | 0 9 0 | 0 7 0         1 6 5 | 4 9 8 | 3 7 2

2024/01/03 22:38:25 Solving grid 2...
2024/01/03 22:38:25 Solved:
0 0 0 | 0 0 0 | 8 0 0         1 9 2 | 4 3 6 | 8 7 5
7 8 5 | 0 9 0 | 0 0 6         7 8 5 | 2 9 1 | 4 3 6
0 4 0 | 7 0 0 | 0 0 2         6 4 3 | 7 5 8 | 1 9 2
- - - + - - - + - - -         - - - + - - - + - - -
0 6 8 | 3 0 0 | 0 5 1         2 6 8 | 3 7 4 | 9 5 1
4 0 0 | 0 0 0 | 0 0 7   --->  4 3 9 | 1 2 5 | 6 8 7
5 7 0 | 0 0 9 | 3 2 0         5 7 1 | 6 8 9 | 3 2 4
- - - + - - - + - - -         - - - + - - - + - - -
9 0 0 | 0 0 2 | 0 6 0         9 1 7 | 8 4 2 | 5 6 3
8 0 0 | 0 6 0 | 7 1 9         8 2 4 | 5 6 3 | 7 1 9
0 0 6 | 0 0 0 | 0 0 0         3 5 6 | 9 1 7 | 2 4 8

2024/01/03 22:38:25 Solving grid 3...
2024/01/03 22:38:25 Solved:
3 2 0 | 0 0 0 | 8 0 9         3 2 6 | 1 7 4 | 8 5 9
4 0 0 | 6 0 0 | 0 2 7         4 9 1 | 6 5 8 | 3 2 7
0 0 5 | 0 0 0 | 0 4 0         8 7 5 | 3 2 9 | 6 4 1
- - - + - - - + - - -         - - - + - - - + - - -
0 0 0 | 4 0 1 | 0 0 0         7 3 8 | 4 9 1 | 2 6 5
0 0 2 | 0 0 0 | 9 0 0   --->  5 6 2 | 8 3 7 | 9 1 4
0 0 0 | 5 0 2 | 0 0 0         9 1 4 | 5 6 2 | 7 8 3
- - - + - - - + - - -         - - - + - - - + - - -
0 8 0 | 0 0 0 | 1 0 0         2 8 3 | 7 4 5 | 1 9 6
6 4 0 | 0 0 3 | 0 0 8         6 4 9 | 2 1 3 | 5 7 8
1 0 7 | 0 0 0 | 0 3 2         1 5 7 | 9 8 6 | 4 3 2

...
```

## Input file structure

The file should contain one or more input sudoku grids with `0`s for missing values. The grids must be separated by one or more empty lines.
Refer to the [sudokus](files/sudokus) file for an example.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
