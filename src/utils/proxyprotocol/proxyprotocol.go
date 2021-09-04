package proxyprotocol

import (
	"fmt"
	"net"

	proxyproto "github.com/pires/go-proxyproto"
)

func getAddrIp(addr net.Addr) (ip net.IP, err error) {
	ipString, _, err := net.SplitHostPort(addr.String())
	if err != nil {
		return
	}

	ip = net.ParseIP(ipString)
	if ip == nil {
		err = fmt.Errorf("Could not parse IP")
		return
	}

	return
}

/// SendProxyProtocolV1 sends a proxy protocol v1 header to initialize the connection
/// https://www.haproxy.org/download/1.8/doc/proxy-protocol.txt
func SendProxyProtocolV1(client net.Conn, backend net.Conn) error {

	h := proxyproto.Header{
		Version:         1,
		SourceAddr:      client.RemoteAddr(),
		DestinationAddr: client.LocalAddr(),
	}

	sourceIP, err := getAddrIp(client.RemoteAddr())
	if err != nil {
		return err
	}
	if sourceIP.To4() != nil {
		h.TransportProtocol = proxyproto.TCPv4
	} else {
		h.TransportProtocol = proxyproto.TCPv6
	}

	_, err = h.WriteTo(backend)
	if err != nil {
		return nil
	}
	return nil
}
