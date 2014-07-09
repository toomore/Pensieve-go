package main

import "fmt"
import "math"
import "net/http"
import "sort"
import "strconv"
import "strings"

type statistics struct {
    numbers []float64
    mean float64
    median float64
}

func Sum(numbers []float64) (sum float64) {
    for _, value := range numbers {
        sum += value
    }
    return
}

func Average(numbers []float64) float64 {
    return Sum(numbers) / float64(len(numbers))
}

func Median(numbers []float64) float64 {
    sort.Float64s(numbers)
    return numbers[int(math.Ceil(float64(len(numbers)) / 2))-1]
}

func getStatis(numbers []float64) (stats statistics) {
    stats.numbers = numbers
    sort.Float64s(stats.numbers)
    stats.mean = Average(numbers)
    stats.median = Median(numbers)
    return
}

func homePage(writer http.ResponseWriter, request *http.Request) {
    err := request.ParseForm()
    fmt.Println(request.Header)
    tpl := `
        <!DOCTYPE HTML>
        <form action="/" method="POST">
            <input type="text" name="numbers"><br>
            <input type="submit" value="Calcular">
        </form>`
    fmt.Fprint(writer, tpl)
    fmt.Fprint(writer, err)
    if slice := request.Form["numbers"]; len(slice) > 0 {
        fmt.Fprint(writer, request.Form)
        fmt.Fprint(writer, processData(slice))
    }
}

func processData(data []string) string {
    text := strings.Replace(data[0], ",", " ", -1)
    var numbers []float64
    for _, value := range strings.Fields(text) {
        float64_value, _ := strconv.ParseFloat(value, 64)
        numbers = append(numbers, float64_value)
    }
    tpl := `
        <table>
            <tr>
                <td colspan="2">Resule</td>
            </tr>
            <tr>
                <td>Numbers</td>
                <td>%v</td>
            </tr>
            <tr>
                <td>Meam</td>
                <td>%f</td>
            </tr>
            <tr>
                <td>Median</td>
                <td>%f</td>
            </tr>
        </table>`
    result := getStatis(numbers)
    return fmt.Sprintf(tpl, result.numbers, result.mean, result.median)
}

func main() {
    numbers := []float64{5, 2, 3, 4, 1}
    fmt.Println(getStatis(numbers))
    http.HandleFunc("/", homePage)
    http.ListenAndServe(":9001", nil)
}
