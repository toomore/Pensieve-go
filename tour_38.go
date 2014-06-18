//http://tour.golang.org/#38
//https://gist.github.com/pyk/8c54534f7c6163e9178d
//data:image/png;base64,IMAGEBASE64
package main

import "code.google.com/p/go-tour/pic"
//import "fmt"

func Pic(dx, dy int) [][]uint8 {
    d := make([][]uint8, dx)
    for x := range d {
        d[x] = make([]uint8, dy)
        for y := range d[x] {
            d[x][y] = uint8(x+x*y-y)
        }
    }
    return d
}

func main() {
    pic.Show(Pic)
    //fmt.Println(Pic)
}
