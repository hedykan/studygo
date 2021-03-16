package main

import "fmt"

func showe() string {
     return "ok"
}

func pro(f func() string) string {
     str := "hello"
     return str+f()
}
func main(){
     str := pro("showe")
     fmt.Println(str)
}