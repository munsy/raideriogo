# raideriogo 
[![Build Status](https://travis-ci.org/Munsy/raideriogo.svg?branch=master)](https://travis-ci.org/Munsy/raideriogo)  [![Documentation](https://godoc.org/github.com/munsy/raideriogo?status.svg)](https://godoc.org/github.com/munsy/raideriogo)  [![Go Report Card](https://goreportcard.com/badge/github.com/munsy/raideriogo)](https://goreportcard.com/report/github.com/munsy/raideriogo)  [![license](https://img.shields.io/github/license/mashape/apistatus.svg)](LICENSE)  

Raider.io bindings written in golang.

## Install
```
$ go get github.com/munsy/raideriogo
```

## A Very Simple Example
```
package main

import(
        "fmt"

        "github.com/munsy/raideriogo"
)

func main() {
        client := raideriogo.New()

        character, err := client.GetCharacterProfile("us", "thrall", "munsy", "")

        if nil != err {
                panic(err)
        }

        fmt.Printf("Name: %v\nClass: %v\nRace: %v\n", character.Name, character.Class, character.Race)
}
```
```
$ go build
$ ./raideriotest
Name: Munsy
Class: Druid
Race: Troll
```

## Disclaimer
I am in no way associated with the fine folks over at https://www.raider.io other than being a fan of their work and wanting to use it with my own stuff.
