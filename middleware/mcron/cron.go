package mcron

import "github.com/robfig/cron"

var mcron *cron.Cron

func InitCron() {
	mcron = cron.New()
}
