package cron

import (
	"github.com/robfig/cron"
	"log"
)

func Setup() {
	go func() {
		log.Println("crontab starting...")
		c := cron.New()
		// 每天1点清理超过1天的导出记录
		if err := c.AddFunc("0 0 1 * * *", CESHI); err != nil {
			log.Printf("WriteIntoFile crontab failed：%v", err)
		}

	}()
}

func CESHI()  {
	log.Println("试试")
}




