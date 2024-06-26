package crontab
import (
	"testing"
	"github.com/alfarady/letgo/cron/context"
	"github.com/alfarady/letgo/lib"
	"fmt"
	"time"
)
func TestCrontab(t *testing.T) {
	AddCron("cron1","*/6 * * * * *",func(ctx *context.Context){
		fmt.Println("dddd", lib.Now())
	})
	AddCron("cron2","*/3 * * * * *",func(ctx *context.Context){
		fmt.Println("xxxx", lib.Now())
	})
	go func(){
		time.Sleep(10*time.Second)
		Stop()
	}()
	StartAndWait()
}