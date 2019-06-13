package util

import (
	"net/http"
	"strings"
)

func ClientPublicIP(r *http.Request) string {
	var ips string
	ips = strings.TrimSpace(r.Header.Get("X-Forwarded-For"))
	if ips == "" || ips == "unknown" {
		ips = strings.TrimSpace(r.Header.Get("Proxy-Client-IP"))
	}
	if ips == "" || ips == "unknown" {
		ips = strings.TrimSpace(r.Header.Get("WL-Proxy-Client-IP"))
	}
	if ips == "" || ips == "unknown" {
		ips = r.RemoteAddr
	}
	for _, ip := range strings.Split(ips, ",") {
		ip = strings.TrimSpace(ip)
		if ip != "" && ip != "unknown" {
			return ip
		}
	}
	return ""
}
