package BY

import (
	"fmt"
	geoip2db "github.com/cc14514/go-geoip2-db"
	"net"
)

var db, _ = geoip2db.NewGeoipDbByStatik()

// GetAddr 获取IP属地
func (_Other) GetAddr(ip string) string {
	record, _ := db.City(net.ParseIP(ip))
	defer db.Close()
	var province string
	if len(record.Subdivisions) > 0 {
		province = record.Subdivisions[0].Names["zh-CN"]
	}
	city := record.City.Names["zh-CN"]

	if province == "" && city == "" {
		return "内网IP"
	}
	return fmt.Sprintf("%s %s", province, city)
}

// GetNetworkCard 获取网卡所有地址
func (_Other) GetNetworkCard() (ipList []string) {
	interfaces, _ := net.Interfaces()
	for _, i := range interfaces {
		addr, _ := i.Addrs()
		// 过滤IPv6
		for _, ip := range addr {
			ipNet, _ := ip.(*net.IPNet)
			ip4 := ipNet.IP.To4()
			if ip4 != nil {
				ipList = append(ipList, ip4.String())
				//Log(ip4)
			}
		}
	}
	return
}
