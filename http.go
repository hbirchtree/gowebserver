package main

import (
    "io"
    "fmt"
    "log"
    "net/http"
    "strconv"

)

var m = map[int]string{1000:"M",900:"CM",500:"D",400:"CD",100:"C",90:"XC",50:"L",40:"XL",10:"X",9:"IX",5:"V",4:"IV",1:"I"}

func hello(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello world!")
}

func to_roman(i int)  string {



    liste  := []int{1000,900,500,400,100,90,50,40,10,9,5,4,1}

    s:=""
    for k  := range(liste){
    antall := i/liste[k]
    i-=antall*liste[k]

    for it := 1; it <= antall; it++ {
        s+=m[liste[k]]
        }
    }

    //fmt.Println(s)
    //fmt.Println(liste)
    return s
}

type romanGenerator int
func (n romanGenerator) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    ascii_num := r.URL.Path[7:]
    i, err := strconv.Atoi(ascii_num)
    if err != nil {
        log.Print(err)
    }
    fmt.Fprintf(w, "Here's your number: %s\n", to_roman(i))
}



func main() {
    h := http.NewServeMux()

    h.Handle("/roman/", romanGenerator(1))
    h.HandleFunc("/", hello)

    err := http.ListenAndServe(":8000", h)
    log.Fatal(err)
}
