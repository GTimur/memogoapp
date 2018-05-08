package main

import (
	"fmt"
	"time"
)

// REMIND SCENARIO
// REPEAT_UNTIL_EVENT_START = Remind in Int1 before event DateStart and repeat remind every Int2 until event DateStart
// REPEAT_UNTIL_EVENT_END = Remind from DateStart to DateEnd with repeat inteval=Int2 (minutes)

// Remind options
type Remind struct {
	DateStart time.Time //Date start of event
	DateEnd   time.Time //Date end of event
	Int1      int       //Frequency of reminder messages (minutes)
	Int2      int       //Frequency of reminder messages (minutes)
}

// Memo definiton
type Memo struct {
	ID       int       //ID
	Date     time.Time //Date of creation
	Reminder Remind    //Remind scenario options
	Subj     string    //visible memo subject
	Memo     []string  //memo body
	Mails    []string  //emails list
}

func main() {
	memo := Memo{}
	rem := Remind{}

	loc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		panic(err)
	}

	rem.DateStart = time.Date(2018, time.May, 11, 10, 0, 0, 0, loc)
	rem.DateEnd = time.Date(2018, time.August, 16, 23, 59, 59, 0, loc)
	rem.Int1 = 5 //in * minutes
	rem.Int2 = 1 //every * minutes

	memo.ID = 100
	memo.Date = time.Now()
	memo.Subj = "Сертификат thawte ibank.ymkbank.ru истекает 17.08.2018"
	memo.Memo = []string{"Истекает срок действия сертификата для", "web-сервера ibank.ymkbank.ru"}
	memo.Mails = []string{"gtg@ymkbank.ru"}
	memo.Reminder = rem

	fmt.Println(memo)
}
