gobuf

![Go](https://github.com/joesonw/gobuf/workflows/Go/badge.svg)
[![Documentation](https://godoc.org/github.com/joesonw/gobuf?status.svg)](http://godoc.org/github.com/joesonw/gobuf)

# Installation

`go get github.com/joesonw/gobuf`

# Usage

```go
buf := gobuf.New(nil, gobuf.WithAutoGrowMemory(gobuf.FixedGrow(64)))
buf.WriteString("hello world")
fmt.Println(buf.ReadString())
```

# Buffers

## gobuf.New

can read and write

## gobuf.Read

read only buffer backed by io.Reader

## gobuf.Write

write only buffer backed by io.Writer


# Memory 

## SliceMemory

backed by a single slice, can be expansive when write heavily with constant memory copying

## LinkedList

backed by a linked list, can be efficient when writing heavily, but performs poor when reading with sparse nodes.
