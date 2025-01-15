package pkoderlite

import (
	"fmt"
	"net"
)

func PrivateIPv4() (net.IP, error) {
	ia, e := net.InterfaceAddrs()
	if e != nil {
		return nil, e
	}
	for _, a := range ia {
		if ip, ok := a.(*net.IPNet); ok {
			if ip.IP.IsPrivate() && ip.IP.To4() != nil {
				return ip.IP, nil
			}
		}
	}
	return nil, fmt.Errorf("unable to find a private ipv4 address")
}

func LoopbackIPv4() (net.IP, error) {
	ia, e := net.InterfaceAddrs()
	if e != nil {
		return nil, e
	}
	for _, a := range ia {
		if ip, ok := a.(*net.IPNet); ok {
			if ip.IP.IsLoopback() && ip.IP.To4() != nil {
				return ip.IP, nil
			}
		}
	}
	return nil, fmt.Errorf("unable to find a private ipv4 address")
}

// LoopbackListener represents a TCP listener that binds to a loopback address.
type LoopbackListener struct {
	// Server represents the underlying TCP server.
	Server *net.TCPListener
	// Path is the specific path or endpoint associated with this listener.
	Path string
}

// ListenLoopback initializes a new LoopbackListener on a given path. It binds the
// listener to a loopback IPv4 address and an available port.
// Returns the LoopbackListener instance or an error if one occurs during binding.
func ListenLoopback(path string) (*LoopbackListener, error) {
	ip, _ := LoopbackIPv4()
	if server, err := net.ListenTCP("tcp4", &net.TCPAddr{IP: ip}); err == nil {
		return &LoopbackListener{Server: server, Path: path}, nil
	} else {
		return nil, err
	}
}

// URI constructs a full URI string for the LoopbackListener, combining the listener's
// address and the specified path.
func (l *LoopbackListener) URI() string {
	return "tcp://" + l.Server.Addr().String() + l.Path
}

// AcceptTCP waits for and returns the next connection to the listener, specifically
// as a TCP connection.
func (l *LoopbackListener) AcceptTCP() (*net.TCPConn, error) {
	return l.Server.AcceptTCP()
}

// Accept waits for and returns the next connection to the listener.
func (l *LoopbackListener) Accept() (net.Conn, error) {
	return l.Server.Accept()
}

// Close shuts down the LoopbackListener, terminating all established connections and
// freeing any resources.
func (l *LoopbackListener) Close() error {
	return l.Server.Close()
}
