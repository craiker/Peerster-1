package types

import (
	"fmt"
	"net"
)

// BufSize - Size of the UDP buffer
const BufSize = 2048

// TimeoutSec - Length of generic timeout
const TimeoutSec = 1

// Gossiper - Represents a gossiper
type Gossiper struct {
	ClientAddr    string                   // IP/Port on which the client talks (RO)
	GossipAddr    string                   // IP/Port on which to listen to other gossips (RO)
	Name          string                   // Name of that gossiper (RO)
	SimpleMode    bool                     // Indicate whether the gossiper operated in simple broadcast mode (RO)
	ClientChannel *net.UDPConn             // UDP channel to communicate with the client (Shared, thread-safe)
	GossipChannel *net.UDPConn             // UDP channel to communicate with the network (Shared, thread-safe)
	NameIndex     *NameIndex               // A dictionnary between peer names and received messages (Shared, thread-safe)
	PeerIndex     *PeerIndex               // A dictionnary between <ip:port> and peer addresses (Shared, thread-safe)
	Timeouts      *StatusResponseForwarder // Timeouts for RumorMessage's answer (Shared, thread-safe)
}

// Client - Represents a client
type Client struct {
	Addr *net.UDPAddr // Address on which to send
	Msg  string       // Message to send
}

// NewGossiper - Creates a new instance of Gossiper
func NewGossiper() *Gossiper {
	var gossip Gossiper

	gossip.NameIndex = NewNameIndex()
	gossip.PeerIndex = NewPeerIndex()
	// TODO: add timeout

	return &gossip
}

// GossiperToString -
func (ent *Gossiper) GossiperToString() string {
	return fmt.Sprintf("ClienAddr: %s\nGossipAddr: %s\nName: %s\nSimpleMode: %v\n",
		ent.ClientAddr, ent.GossipAddr, ent.Name, ent.SimpleMode)
}