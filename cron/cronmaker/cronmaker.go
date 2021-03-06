package cronmaker

import (
	"gocron/cron/service/deadline"
	"gocron/cron/service/flush"

	"github.com/jasonlvhit/gocron"
)

type cronJob struct{}

type ICronJob interface {
	StartCronJobs()
}

func InitCrons() ICronJob {
	return &cronJob{}
}

func (cj cronJob) StartCronJobs() {
	cj.initDeadlineCron()
	cj.initFlushCron()
	<-gocron.Start()
}

func (cj *cronJob) initDeadlineCron() {
	action := func() { deadline.DoAction() }
	_ = gocron.Every(2).Seconds().Do(action)
}

func (cj *cronJob) initFlushCron() {
	action := func() { flush.DoAction() }
	_ = gocron.Every(1).Seconds().Do(action)
}
