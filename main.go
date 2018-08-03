package main

import (
	"fmt"
	"log"
	"memogo"
	"net"
	"os"
	"os/signal"
	"strconv"
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

	// Инициализация web-сервера
	memogo.NeedExit = false // флаг для завершения работы
	var web memogo.WebCtl
	//memogo.GlobalConfig.SetManagerSrv("127.0.0.1", 4400)
	web.SetHost(net.ParseIP(memogo.GlobalConfig.ManagerSrvAddr()))
	web.SetPort(memogo.GlobalConfig.ManagerSrvPort())

	fmt.Println("Web control configured: " + "http://" + memogo.GlobalConfig.ManagerSrvAddr() + ":" + strconv.Itoa(int(memogo.GlobalConfig.ManagerSrvPort())))

	/* Запускаем сервер обслуживания WebCtl */
	err := web.StartServe()
	if err != nil {
		log.Println("HTTP сервер: Ошибка. ", err)
		os.Exit(1)
	}

	/* Перехват CTRL+C для завершения приложения */
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			fmt.Printf("\nReceived %v, shutdown procedure initiated.\n\n", sig)
			memogo.Quit <- 1
			memogo.NeedExit = true
		}
	}()

	// Цикл с таймером для ожидания команды завершения
	ticker := time.NewTicker(time.Second * 1) // Запускаем обработчик каждую секунду

	// Зациклимся с таймером посекундно пока не получим команду завершения работы.
	for range ticker.C {
		if !memogo.NeedExit {
			continue
		}
		break
	}

	ticker.Stop()

}
