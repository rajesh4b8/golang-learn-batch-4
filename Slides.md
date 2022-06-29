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
