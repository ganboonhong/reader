package main

import (
    "fmt"
    "io/ioutil"
)

func main(){
    b, err := ioutil.ReadFile("debug.json")
    if err != nil {
        fmt.Println(err)
        return;
    }
    fmt.Println(string(b))
}