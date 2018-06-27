package main

import (
	"fmt"
	"log"
	"memogo"
	"time"
)

func main() {
	memogo.GlobalConfig = memogo.Config{
		Root: "root\\",
		SMTPSrv: memogo.SrvSMTP{
			Addr:     "10.20.20.6",
			Port:     25,
			Account:  "noti",
			Password: "Bank999",
			From:     "noti@ymkbank.ru",
			FromName: "Memo GO",
			UseTLS:   false,
		},
		MgrSrv: memogo.ManagerSrv{
			Addr: "127.0.0.1",
			Port: 8000,
		},
	}

	err := memogo.TasksReload()
	if err != nil {
		log.Fatal(err)
	}

	err = memogo.BuildTimeMap()
	if err != nil {
		log.Fatal(err)
	}

	m := memogo.GlobalTimeMap[100]

	for i := 0; i <= 10; i++ {
		fmt.Println("TimeMap:\n", m[int64(i)], i)
	}

	/*for i := len(m) - 30; i <= len(m); i++ {
		fmt.Println("TimeMap:\n", m[i], i)
	}*/

	err = memogo.Reader()
	if err != nil {
		log.Fatal(err)
	}

	res := memogo.Row(time.Now(), 10)
	fmt.Println(res)

	//n := memogo.NextTime(memogo.TasksGlobal[0].Memo, 60*60)

	//fmt.Println("NEXT:", n)

	//fmt.Println("TimeMapCounter:\n", memogo.GlobalTimeCount[100])

}
