package net

import (
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

var (
	once     sync.Once
	clientIP = "127.0.0.1"
)

func IntranetIP() (ips []string, err error) {
	ips = make([]string, 0)

	ifaces, e := net.Interfaces()
	if e != nil {
		return ips, e
	}

	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}

		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}

		// ignore docker and warden bridge
		if strings.HasPrefix(iface.Name, "docker") || strings.HasPrefix(iface.Name, "w-") {
			continue
		}

		addrs, e := iface.Addrs()
		if e != nil {
			return ips, e
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip == nil || ip.IsLoopback() {
				continue
			}

			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}

			ipStr := ip.String()
			if IsIntranet(ipStr) {
				ips = append(ips, ipStr)
			}
		}
	}

	return ips, nil
}

func IsIntranet(ipStr string) bool {
	if strings.HasPrefix(ipStr, "10.") || strings.HasPrefix(ipStr, "192.168.") {
		return true
	}

	if strings.HasPrefix(ipStr, "172.") {
		// 172.16.0.0-172.31.255.255
		arr := strings.Split(ipStr, ".")
		if len(arr) != 4 {
			return false
		}

		second, err := strconv.ParseInt(arr[1], 10, 64)
		if err != nil {
			return false
		}

		if second >= 16 && second <= 31 {
			return true
		}
	}

	return false
}

// GetLocalIP 获取本地内网IP
func GetLocalIP() string {
	once.Do(func() {
		ips, _ := IntranetIP()
		if len(ips) > 0 {
			clientIP = ips[0]
		} else {
			clientIP = "127.0.0.1"
		}
	})
	return clientIP
}
func GetRealIP(request *http.Request) (ip string) {
	var header = request.Header
	var index int
	if ip = header.Get("X-Forwarded-For"); ip != "" {
		index = strings.IndexByte(ip, ',')
		if index < 0 {
			return ip
		}
		if ip = ip[:index]; ip != "" {
			return ip
		}
	}
	if ip = header.Get("X-Real-Ip"); ip != "" {
		index = strings.IndexByte(ip, ',')
		if index < 0 {
			return ip
		}
		if ip = ip[:index]; ip != "" {
			return ip
		}
	}
	if ip = header.Get("Proxy-Forwarded-For"); ip != "" {
		index = strings.IndexByte(ip, ',')
		if index < 0 {
			return ip
		}
		if ip = ip[:index]; ip != "" {
			return ip
		}
	}
	ip, _, _ = net.SplitHostPort(request.RemoteAddr)
	return ip
}
