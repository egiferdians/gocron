package cron

import "gocron/cron/cronmaker"

func Run() {
	cronmaker.InitCrons().StartCronJobs()
}
