package myutil

import (
	"fmt"
	"net"
)

func GetLocalIP() string {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("net.Interfaces failed, err:", err.Error())
	}
	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()

			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && ipnet.IP.IsGlobalUnicast() {
					if ipnet.IP.To4() != nil {
						return ipnet.IP.String()
					}
				}
			}
		}
	}
	return ""
}

//utf8下字符串长度感性计算，非实际长度
func GetStringLen(str string) int {
	cnt := 0
	for _, b := range []byte(str) {
		if b <= 0x7F { //首字节   UTF-8 占用1个字节
			cnt += 1
		} else if b >= 0xC0 && b <= 0xDF { //首字节   UTF-8 占用2个字节
			cnt += 1
		} else if b >= 0xE0 && b <= 0xEF { //首字节   UTF-8 占用3个字节
			cnt += 2
		} else if b >= 0xF0 && b <= 0xF7 { //首字节   UTF-8 占用4个字节
			cnt += 2
		} else {
			//非首字节
		}
	}
	return cnt
}
