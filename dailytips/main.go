package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func main() {
	var (
		shareFilms = [4]string{
			"20150920 大阪 京都 東京 [films]",
			"20160131 北海道 小樽 千葉 求名 銚子 東京",
			"20160521 橫濱 東京 二日往返",
			"20160924 廣島 福岡",
		}
		shareTime = [2]string{
			"白天",
			"晚上",
		}
		now = time.Now()
	)
	fmt.Printf("今天 %d %s, 是 %d 年的第 %s 天\n",
		now.Day(), now.Month(), now.Year(), color.YellowString("%d", now.YearDay()))
	fmt.Printf("建議分享的是 \"%s\" 的主題，", color.YellowString("%s", shareFilms[now.YearDay()%len(shareFilms)]))
	fmt.Printf("在 %s 分享比較好\n", color.YellowString("%s", shareTime[int(now.Weekday())%len(shareTime)]))
	if int(now.Weekday())%3 == 0 {
		color.Cyan("也確認一下 Balance")
	}
}
