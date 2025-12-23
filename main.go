package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Create UDP address
	addr, err := net.ResolveUDPAddr("udp", ":50103")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error resolving UDP address: %v\n", err)
		os.Exit(1)
	}

	// Listen on UDP port
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error listening on UDP port: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Initialize MQTT publisher
	publisher, err := NewMQTTPublisher()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing MQTT: %v\n", err)
		os.Exit(1)
	}
	defer publisher.Close()

	fmt.Println("UDP listener started on port 50103; MQTT connected")

	// Buffer for receiving messages
	buffer := make([]byte, 4096)

	for {
		// Read from UDP connection
		n, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading from UDP: %v\n", err)
			continue
		}

		// Parse the Cyrano Protocol message
		cyranoMsg, err := ParseCyranoMessage(buffer[:n])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing message from %s: %v\n", remoteAddr.String(), err)
			continue
		}

		// Log parsed message details
		logMessage(cyranoMsg, remoteAddr)

		// Publish to MQTT
		if err := publisher.PublishMessage(cyranoMsg); err != nil {
			fmt.Fprintf(os.Stderr, "MQTT publish error: %v\n", err)
		}
	}
}

func logMessage(msg *CyranoMessage, addr *net.UDPAddr) {
	fmt.Printf("=== Message from %s ===\n", addr.String())
	fmt.Printf("Protocol: %s\n", msg.ProtocolVersion.String())
	fmt.Printf("Command: %s\n", msg.Command.String())
	fmt.Printf("Piste: %s, Competition: %s\n", msg.Piste, msg.Competition)
	fmt.Printf("Phase: %s, PoulTab: %s, Match: %s, Round: %s\n",
		msg.Phase, msg.PoulTab, msg.Match, msg.Round)
	fmt.Printf("Time: %s, Stopwatch: %s\n", msg.Time, msg.Stopwatch)
	fmt.Printf("Type: %s, Weapon: %s, State: %s\n", msg.Type, msg.Weapon, msg.State)

	if msg.RightId != "" || msg.LeftId != "" {
		fmt.Printf("\nRight Fencer: %s (%s/%s) - Score: %s, Status: %s\n",
			msg.RightName, msg.RightId, msg.RightNat, msg.RightScore, msg.RightStatus)
		fmt.Printf("Left Fencer: %s (%s/%s) - Score: %s, Status: %s\n",
			msg.LeftName, msg.LeftId, msg.LeftNat, msg.LeftScore, msg.LeftStatus)
	}
	fmt.Println()
}
