package types

import (
	"net"
)

// Client - Represents a client
type Client struct {
	Addr     *net.UDPAddr // Address on which to send
	Msg      string       // Message to send
	Dst      string       // A private message's destination
	Filename string       // A file to index
	Request  []byte       // A hash
}
