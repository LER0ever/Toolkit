package main

import (
	"fmt"
	// /"os"
	"os/exec"
	"strings"
	"time"
)

type seat struct {
	flyno string
	stype  string
	sprice int
}

func main() {
	for j := 0; ; j++ {
		var seats []seat
		//var oldseats []seat
		fmt.Println("\x1b[0;32mRequesting Webpages ... in about 30 seconds\n中国上海 -> 芝加哥 往返机票 2016-08-26\x1b[0m")
		cmd := exec.Command("phantomjs", "main.js", "http://flights.ctrip.com/international/round-shanghai-chicago-sha-chi?2016-08-26&2016-09-14&y", "webshot.png")
		buf, _ := cmd.Output()

		fmt.Println("===================\n\x1b[47;30m" + time.Now().Format("2006-01-02 15:04:05") + "\x1b[0m\n===================")

		strbuf := string(buf)
		bufs := strings.Split(strbuf, "\n")
		//fmt.Println(len(bufs))
		for i := 0; i < len(bufs)-2; i++ {
			var tmp seat
			fmt.Sscanf(bufs[i], "%s %s %d",&tmp.flyno, &tmp.stype, &tmp.sprice)
			seats = append(seats, tmp)
		}
		for i := 0; i < len(seats); i++ {
			fmt.Printf("\x1b[41;37m%d\x1b[0m RMB : 航班号 %s %s\n", seats[i].sprice, seats[i].flyno, seats[i].stype)
			if seats[i].sprice < 8000 && seats[i].flyno == "MU717" {
				fmt.Sprintln("\x1b[45;32m **** Price Under 8000 ! ****\x1b[0m")
			}
		}
		//strings.Replace(strbuf, "\r\n", " ; ", -1)
		strbuf = strings.Replace(strbuf, "\n", " ; ", -1)
		//fmt.Printf("DBG_FROM_PHANTOM:: %s", strbuf)
		//fmt.Printf("%s", strbuf)
		//fmt.Println("Sleeping for 2 Minute")
		//time.Sleep(2 * time.Minute)
	}
}
