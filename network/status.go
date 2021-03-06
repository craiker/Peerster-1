package network

import (
	"Peerster/entities"
	"Peerster/fail"
	"Peerster/messages"
	"Peerster/peers"
	"net"

	"github.com/dedis/protobuf"
)

// OnSendStatus - Sends a status
func OnSendStatus(vectorClock *messages.StatusPacket, channel *net.UDPConn, target *net.UDPAddr) error {

	// Create the packet
	pkt := messages.GossipPacket{Status: vectorClock}
	buf, err := protobuf.Encode(&pkt)
	if err != nil {
		return &fail.CustomError{Fun: "OnSendStatus", Desc: "failed to encode GossipPacket"}
	}

	// Send the packet
	if _, err = channel.WriteToUDP(buf, target); err != nil {
		return &fail.CustomError{Fun: "OnSendStatus", Desc: "failed to send StatusPacket"}
	}
	return nil
}

// OnReceiveStatus - Called when a status is received
func OnReceiveStatus(g *entities.Gossiper, status *messages.StatusPacket, sender *net.UDPAddr, threadID uint32) {

	// Print to the console
	fail.LeveledPrint(0, "", status.StatusPacketToString(peers.UDPAddressToString(sender)))
	fail.LeveledPrint(0, "", g.PeerIndex.PeersToString())

	// See if we must propagate a rumor
	rumorToPropagate := g.NameIndex.GetUnknownMessageTarget(status)

	if rumorToPropagate == nil { // We don't have anything to propagate
		if g.NameIndex.IsLocalStatusComplete(status) { // We are in sync with the other
			fail.LeveledPrint(0, "", "IN SYNC WITH %s", peers.UDPAddressToString(sender))
		} else { // We must send back our own Status
			vectorClock := g.NameIndex.GetVectorClock()
			OnSendStatus(vectorClock, g.GossipChannel, sender)
		}
	} else { // We must propagate a rumor to the sender
		OnSendRumor(g, rumorToPropagate, sender, threadID)
	}

}
