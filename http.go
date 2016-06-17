package main

import (
    "io"
    "fmt"
    "log"
    "net/http"
    "strconv"
)

func hello(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello world!")
}

func to_roman(n int)  string {
  m := make(map[int]string)

    m[1] = "I"
    m[2] = "II"
    m[3] = "III"
    m[5] = "V"
    m[10] = "X"
    m[50] = "L"
    m[100] = "C"

    fmt.Println("map:", m)

  //  for k, _ := range m {

  //  }




    if len(m[n]) != 0 {
        return m[n]
    }
    return "I"
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
