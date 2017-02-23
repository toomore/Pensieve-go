package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

var (
	shareFilms = [7]string{
		"20150920 大阪 京都 東京 [films]",
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
)

func getRand(n int) []int {
	rand.Seed(int64(now.YearDay() / n))
	return rand.Perm(n)
}

func output() {
	fmt.Printf("今天 %d %s, 是 %d 年的第 %s 天\n",
		now.Day(), now.Month(), now.Year(), color.YellowString("%d", now.YearDay()))
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
}
