package TCP_Chat

import (
	"fmt"
	"time"
)

func Format(name, msg string) string {
	Time := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(Time)
	return Time + "[" + name + "] : " + msg
}
