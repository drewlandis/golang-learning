package main

import (
	"fmt"
	"strings"

	syslog "github.com/RackSec/srslog"
)

// view with netcat using a terminal
// nc -u -l 5514

func main() {
	priority := syslog.LOG_INFO | syslog.LOG_LOCAL7
	tag := "D3SCORE"
	content := "threat=foo.com score=8"

	servers := []string{"udp|127.0.0.1:5514", "udp|127.0.0.1:5513"}
	fmt.Println(servers)

	var writers []*syslog.Writer
	for _, server := range servers {
		addr := strings.Split(server, "|")
		if len(addr) != 2 {
			panic("Length of split not equal to two!")
		}

		w, err := syslog.Dial(addr[0], addr[1], priority, tag)
		if err != nil {
			fmt.Println(err)
		}

		w.SetFormatter(syslog.RFC5424Formatter)

		writers = append(writers, w)
	}

	for _, w := range writers {
		w.Write([]byte(content))
		w.Write([]byte(content))
	}
}
