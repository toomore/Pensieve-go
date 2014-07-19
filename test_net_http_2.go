package main

import "fmt"
import "io/ioutil"
import "net/http"
import "net/url"


func httpGet(url string) []byte {
    resp, _ := http.Get(url)
    defer resp.Body.Close()
    str, _ := ioutil.ReadAll(resp.Body)
    return str
}

//type params map[string]string

func main() {
    fmt.Printf("%s\n", httpGet("http://httpbin.org/get"))
    v := url.Values{}
    v.Add("name", "toomore")
    v.Add("age", "30")
    fmt.Println(v.Encode())
}
