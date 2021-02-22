## Scoring APP

Written in Golang. A default input.txt file is provided. With and sc executable build on Max OS.

To build your own executable you to do a ```go build```

To provide your own input file
```shell
./sc -file input.txt //or short hand -f input.txt
```
### The Problem

We want you to create a command-line application that will calculate the ranking table for a
league.
Input/output
The input and output will be text. Either using stdin/stdout or taking filenames on the command
line is fine.
The input contains results of games, one per line. See “Sample input” for details.
The output should be ordered from most to least points, following the format specified in
“Expected output”.
You can expect that the input will be well-formed. There is no need to add special handling for
malformed input files.

### The rules
In this league, a draw (tie) is worth 1 point and a win is worth 3 points. A loss is worth 0 points.
If two or more teams have the same number of points, they should have the same rank and be
printed in alphabetical order (as in the tie for 3rd place in the sample data).
