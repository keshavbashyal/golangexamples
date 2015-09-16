package main

import (
	"fmt"
	"reflect"
	"sort"
	"os"
)

type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i,j int){
	s[i], s[j]=s[j],s[i]
}

func (s ByLength) Less(i,j int) bool{
	return len(s[i]) < len(s[j])
}


type person struct{

	name string
	age int
}

func main() {
	fmt.Println(person{"Bob",20})
	fmt.Println(person{name:"Alice",age:30})
	fmt.Println(person{name:"keshav"})
	fmt.Println(&person{name: "Ann", age: 40})

	s := person{name: "Sean", age: 50}
    fmt.Println(s.name)

    d := &person{name: "Sean", age: 50}
    fmt.Println(d)

    messages := make(chan string)

    go func(){messages <- "ping"}()

    msg:= <- messages

    fmt.Println(msg)
    fmt.Println(reflect.TypeOf(msg))

    str:=[]string{"c","a","b"}
    sort.Strings(str)
    fmt.Println("Strings: ",str)

    fruits:=[]string{"peach", "banana","kiwi"}
    sort.Sort(ByLength(fruits))
    fmt.Println(fruits)

	panic("a problem")
	
	//A common use of panic is to abort if a function returns an error value that we don’t know how to (or want to) handle. Here’s an example of panicking if we get an unexpected error when creating a new file.
    _, err := os.Create("/tmp/file")
    if err != nil {
        panic(err)
    }

}