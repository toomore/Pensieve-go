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

type params map[string]string

func (p params) String() string {
    result := url.Values{}
    for key, value := range p {
        result.Add(key, value)
    }
    return result.Encode()
}

func main() {
    //fmt.Printf("%s\n", httpGet("http://httpbin.org/get"))
    //v := url.Values{}
    //v.Add("name", "toomore")
    //v.Add("age", "30")
    //fmt.Println(v.Encode())
    p := params{"name": "toomore", "age": "30"}
    fmt.Println(p)
}
