package pkoderlite

import "fmt"

func ListenResource(protocol string, ip string, port int) string {
	if protocol == "srt" {
		return fmt.Sprintf("%s://%s:%d?mode=listener", protocol, ip, port)
	}
	return fmt.Sprintf("%s://%s:%d", protocol, ip, port)
}
