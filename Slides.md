# Slides for daily sessions

## Day 1

### Agenda

- Intro
- Topics to be covered
- Why Go?
- History
- Go playground
- Hello world example
- Hello web example

### Environment setup

- Install Go
- Install Visual Studio Code and Analysis Tools

## Day 2

### Git basics

- git commit is for your local repository
- git push is to origin (remote repository)
- git clone is to clone the repo to local

### To do

- Create a github account
- install git in your system
- Create SSH pub / private keys

```ssh-keygen -t ed25519 -C "your_email@example.com"```

- add the pub key to github.com at settings
- cheat sheets for reference

https://training.github.com/downloads/github-git-cheat-sheet
https://about.gitlab.com/images/press/git-cheat-sheet.pdf

### Understading the basic program

- The three main areas to fous
  - package main-> every file needs to define a package, package main is mandatory for any program to run
  - import -> to get and use other packages, only needed if you are using other packages
  - function main()-> you need a main() function to run the program

### Understanding packages

- main package -> Executable package, every program needs a main package if we have to run the program. Only the main packages compiles and generates binaries (ex: main.exe)
- reusable packages -> these the packages to share / reuse the code between modules
- Typically, the folder name would be the package name

### Go CLI

- Command line interface, you get the cli when you install go.
- You use this cli to work with Go
  - `go build` - This will compiles and generates the binary. You need a main package to generate binary
  - `go run` - This will build and run the binary
  - `go fmt` - Formats the code
  - `go install` - To install a go library / tool
  - `go get` - To get a package to our system
  - `go test` - To run the unit tests

## Day 3

### Buit in types

- bool, string
- int int8 int16 int32 int64
- uint uint8 uint16 uint32 uint64 uintptr
- byte // alias for uint8
- rune // alias for int32
- float32 float64
- complex64 complex128

### Variables

- `var` is the keyword
- Varibles can be decaled once
- make sure to use `:=` to declare and initalize
- Go is a strictly typed lang -> you can't change the type once declared

### Zero values

- Variables declared without an explicit initial value are given their zero value.
- `0` for numeric types,
- `false` for the boolean type
- `""` (the empty string) for strings

| Type          | Zero Value  |
| -----------   | ----------- |
| Integer       | 0           |
| Floating point| 0.0         |
| Boolean       | false       |
| String        | ""          |
| Pointer       | nil         |
| Interface     | nil         |
| Slice         | nil         |
| Map           | nil         |
| Channel       | nil         |
| Function      | nil         |

### Constants

- used with `const` keyword
- It's even if a const is declared and not used

## Day 4

### Type conversions

- For converting compatible types you just use `new_type(old_value)`

### Functions

- Go supports first class functions
- different ways of creating and using the functions in examples

### Use Case Playing cards

- Deck of cards functionality
  - New Deck
  - Print Deck
  - Shuffle
  - Deal
  - Save to file
  - Read from file

### Assignement 1

- Return a new deck of cards with suits and numbers

## Day 5

### Slices / Arrays

### Flow control statements
