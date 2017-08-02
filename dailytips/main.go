package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

var (
	shareFilms = [7]string{
		"Share day",
		"20160131 北海道 小樽 千葉 求名 銚子 東京",
		"20160521 橫濱 東京 二日往返",
		"20160924 廣島 福岡",
		"20161105 Guam, US",
		"20161119 東京蚤の市",
		"20170122 稚內 札幌 東京",
	}
	shareTime = [2]string{
		"白天",
		"晚上",
	}
	now        = time.Now()
	randList   = getRand(len(shareFilms))
	yearDayMod = now.YearDay() % len(shareFilms)
	lastday    = time.Date(now.Year(), 1, 1-1, 0, 0, 0, 0, time.FixedZone("Asia/Taipei", 8*3600))
)

func getRand(n int) []int {
	rand.Seed(int64(now.YearDay() / n))
	return rand.Perm(n)
}

func showProcessBar(p float64, width int) ([]byte, []byte) {
	full := int(float64(width) * p)
	fullbar := make([]byte, full)
	dotbar := make([]byte, width-full)
	for i := 0; i < full; i++ {
		fullbar[i] = 'l'
	}
	for i := 0; i < width-full; i++ {
		dotbar[i] = '-'
	}
	return fullbar, dotbar
}

func output() {
	yeardayPercent := float64(now.YearDay()) / float64(lastday.YearDay()) * 100
	fullbar, dotbar := showProcessBar(yeardayPercent/100, 70)
	fmt.Printf("今天 %d %s, 是 %d 年的第 %s 天(%0.4f%%)\n[%s%s]\n",
		now.Day(), now.Month(), now.Year(), color.YellowString("%d", now.YearDay()),
		yeardayPercent, color.GreenString("%s", fullbar), color.BlueString("%s", dotbar))
	fmt.Printf("建議分享的是 \"%s\" 的主題，",
		color.YellowString("%s", shareFilms[randList[yearDayMod]]))
	fmt.Printf("在 %s 分享比較好\n",
		color.YellowString("%s", shareTime[int(now.Weekday())%len(shareTime)]))
	if now.Weekday()%3 == 0 {
		color.Cyan("也確認一下 Balance")
	}
	if 3 <= now.Weekday() || now.Weekday() <= 4 {
		color.Cyan("\n- Write Nore !")
	}
}

func diff(a, b time.Time) (year, month, day, hour, min, sec int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)
	hour = int(h2 - h1)
	min = int(m2 - m1)
	sec = int(s2 - s1)

	// Normalize negative values
	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return
}

func countDown(cyear int, cmonth time.Month, cday int) {
	year, month, day, hour, min, sec := diff(time.Date(cyear, cmonth, cday, 0, 0, 0, 0, time.FixedZone("Asia/Taipei", 8*3600)), time.Now())
	fmt.Printf("To %d-%02d-%02d\n", cyear, cmonth, cday)
	fmt.Printf("%d Year, %d Month, %d Day, %d Hour, %d Min, %d Sec\n", year, month, day, hour, min, sec)
}

func main() {
	output()
	fmt.Println("")
	fmt.Println("This time list:")
	for i, v := range randList {
		if i == yearDayMod {
			fmt.Println(color.YellowString(" = %d %s", v, shareFilms[v]))
		} else {
			fmt.Println(" =", v, shareFilms[v])
		}
	}
	countDown(2017, 8, 22)
	countDown(2017, 9, 29)
	countDown(2017, 10, 17)
}
