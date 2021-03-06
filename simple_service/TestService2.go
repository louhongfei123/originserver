package simple_service

import (
	"fmt"
	"github.com/duanhf2012/origin/node"
	"github.com/duanhf2012/origin/service"
	"github.com/duanhf2012/origin/util/timer"
	"time"
)

func init(){
	node.Setup(&TestService2{})
}

type TestService2 struct {
	service.Service
}

func (slf *TestService2) OnInit() error {
	fmt.Printf("TestService2 OnInit.\n")

	//间隔时间定时器
	slf.AfterFunc(time.Second*1,slf.OnSecondTick)

	//crontab模式定时触发
	//NewCronExpr的参数分别代表:Seconds Minutes Hours DayOfMonth Month DayOfWeek
	//以下为每换分钟时触发
	cron,_:=timer.NewCronExpr("0 * * * * *")
	slf.CronFunc(cron,slf.OnCron)
	return nil
}

func (slf *TestService2) OnSecondTick(){
	fmt.Printf("tick.\n")
	slf.AfterFunc(time.Second*1,slf.OnSecondTick)
}

func (slf *TestService2) OnCron(){
	fmt.Printf("A minute passed!\n")
}