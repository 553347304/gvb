package Time

import (
	"github.com/robfig/cron/v3"
	"time"
)

func Cron() *cron.Cron {
	timezone, _ := time.LoadLocation("Asia/Shanghai")
	_cron := cron.New(cron.WithSeconds(), cron.WithLocation(timezone))
	return _cron
}

func Now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func NowDay(day int) string {
	return time.Now().AddDate(0, 0, day).Format("2006-01-02")
}

func Sleep(s time.Duration) {
	time.Sleep(s * time.Second)
}
