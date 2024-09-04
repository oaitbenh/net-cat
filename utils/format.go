package TCP_Chat

import "time"

func Format(name, msg string) string {
	Time := time.Now().Format("[2020-01-20 16:03:43]")
	return Time + "[" + name + "] : " + msg
}
