package main

import "fmt"
import "io/ioutil"
import "net/http"
import "net/url"


type params map[string]string

func (p params) String() string {
    result := url.Values{}
    for key, value := range p {
        result.Add(key, value)
    }
    return result.Encode()
}

func httpGet(url string, params params) []byte {
    resp, _ := http.Get(fmt.Sprintf("%s?%s", url, params))
    defer resp.Body.Close()
    str, _ := ioutil.ReadAll(resp.Body)
    return str
}

func main() {
    //fmt.Printf("%s\n", httpGet("http://httpbin.org/get"))
    //v := url.Values{}
    //v.Add("name", "toomore")
    //v.Add("age", "30")
    //fmt.Println(v.Encode())
    p := params{"name": "toomore", "age": "30"}
    fmt.Println(p)
    fmt.Printf("%s\n", httpGet("http://httpbin.org/get", p))
}
