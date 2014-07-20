package main

import "fmt"
import "net/url"

type params struct {
    values map[string]string
    result url.Values
}

func (p *params) String() string {
    return fmt.Sprintf("%s", p.values)
}

func (p *params) Update(data map[string]string) {
    if len(p.values) == 0 {
        p.values = make(map[string]string)
    }

    for key, value := range data {
        p.values[key] = value
    }
}

func (p *params) Encode() string {
    p.result = url.Values{}
    for key, value := range p.values {
        p.result.Add(key, value)
    }
    return p.result.Encode()
}


func main() {
    var p params
    //p.values = make(map[string]string)
    //p.values["name"] = "toomore"
    //p.values["age"] = "30"

    data := make(map[string]string)
    data["name"] = "Toomore"
    data["age"] = "30"

    fmt.Println(data)
    p.Update(data)
    fmt.Println(p)

    data["local"] = "Taiwan"
    p.Update(data)
    fmt.Println(p)

    fmt.Println(p.Encode())
    //fmt.Println(len(p.values))
    //fmt.Println(p.Encode())
}
