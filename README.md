# Brainfuck Interpreter in Go

## Overview
This is a simple Brainfuck interpreter written in Go.
It currently is able to:
- Run a Brainfuck program from a `.b` or `.bf` file
- Execute Brainfuck code directly from the command line

A couple example programs are included in /programs.

Note that some do not automatically terminate by design, they have to be cancelled with ctrl+C.

## Installation
Ensure you have Go installed, then build the binary:

```sh
git clone https://github.com/Nomi-y/brainfuck
cd ./brainfuck
go build -o bf
```

## Running brainfuck
Run a brainfuck program with one of these commands:
```sh
./bf run program.bf
```
```sh
./bf run program.b
```
```sh
./bf "+[-->-[>>+>-----<<]<--<---]>-.>>>+."
```

For docs on brainfuck refer to https://gist.github.com/roachhd/dce54bec8ba55fb17d3a
