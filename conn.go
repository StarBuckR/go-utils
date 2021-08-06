package sbutils

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

// Check if the ip address can be resolved or not
func CheckDNSAccess(dns string) (bool, error) {
	_, err := net.LookupIP(dns)
	return err == nil, err
}

// Check if the port is accessible
func CheckPortAccess(dns string, port int) (bool, error) {
	hostPort := net.JoinHostPort(dns, strconv.Itoa(port))
	conn, err := net.DialTimeout("tcp", hostPort, time.Second)
	if err != nil {
		return false, err
	}

	if conn != nil {
		defer conn.Close()
		return true, nil
	}
	return false, fmt.Errorf("cant access %s", hostPort)
}
