package main

import "fmt"
import "net/http"
import "io/ioutil"

func main() {
    reap, err := http.Get("http://httpbin.org/get")
    if err != nil {
        fmt.Println("Error:", err)
    }
    defer reap.Body.Close()
    fmt.Println(
        reap, "\n",
        reap.Status, "\n",
        reap.Cookies(), "\n",
    )
    content, err := ioutil.ReadAll(reap.Body)
    fmt.Printf("%s", content)
}
