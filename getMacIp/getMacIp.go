package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/getMacAddr", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "成功", "status": http.StatusOK, "data": gin.H{"mac": getMacAddrs(), "ip": getIPs()}})
	})
	router.Run(":6240")
}

func getMacAddrs() (macAddrs []string) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("fail to get net interfaces: %v", err)
		return macAddrs
	}
	for _, netInterface := range netInterfaces {
		macAddr := netInterface.HardwareAddr.String()
		if len(macAddr) == 0 {
			continue
		}
		macAddrs = append(macAddrs, macAddr)
	}
	return macAddrs
}

func getIPs() (ips []string) {

	interfaceAddr, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Printf("fail to get net interface addrs: %v", err)
		return ips
	}

	for _, address := range interfaceAddr {
		ipNet, isValidIpNet := address.(*net.IPNet)
		if isValidIpNet && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String())
			}
		}
	}
	return ips
}
