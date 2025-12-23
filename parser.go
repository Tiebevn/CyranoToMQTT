package main

import (
	"fmt"
	"strings"
)

func ParseCyranoMessage(data []byte) (*CyranoMessage, error) {
	message := strings.TrimSpace(string(data))

	// Split by "%" to separate the three areas
	areas := strings.Split(message, "%")
	if len(areas) < 1 {
		return nil, fmt.Errorf("invalid message format: no areas found")
	}

	// Parse general area
	generalFields := strings.Split(areas[0], "|")
	msg := parseGeneralArea(generalFields)

	// Parse right fencer area if available
	if len(areas) > 1 && areas[1] != "" {
		rightFields := strings.Split(areas[1], "|")
		parseRightFencerArea(msg, rightFields)
	}

	// Parse left fencer area if available
	if len(areas) > 2 && areas[2] != "" {
		leftFields := strings.Split(areas[2], "|")
		parseLeftFencerArea(msg, leftFields)
	}

	return msg, nil
}

func parseGeneralArea(fields []string) *CyranoMessage {
	msg := &CyranoMessage{}

	// Fields are indexed from 1 in the protocol spec, but our array is 0-indexed
	// Skip empty leading field (index 0)
	if len(fields) > 1 {
		msg.ProtocolVersion = ParseProtocolVersion(strings.TrimSpace(fields[1]))
	}
	if len(fields) > 2 {
		msg.Command = ParseCommand(strings.TrimSpace(fields[2]))
	}
	if len(fields) > 3 {
		msg.Piste = strings.TrimSpace(fields[3])
	}
	if len(fields) > 4 {
		msg.Competition = strings.TrimSpace(fields[4])
	}
	if len(fields) > 5 {
		msg.Phase = strings.TrimSpace(fields[5])
	}
	if len(fields) > 6 {
		msg.PoulTab = strings.TrimSpace(fields[6])
	}
	if len(fields) > 7 {
		msg.Match = strings.TrimSpace(fields[7])
	}
	if len(fields) > 8 {
		msg.Round = strings.TrimSpace(fields[8])
	}
	if len(fields) > 9 {
		msg.Time = strings.TrimSpace(fields[9])
	}
	if len(fields) > 10 {
		msg.Stopwatch = strings.TrimSpace(fields[10])
	}
	if len(fields) > 11 {
		msg.Type = ParseCompetitionType(strings.TrimSpace(fields[11]))
	}
	if len(fields) > 12 {
		msg.Weapon = strings.TrimSpace(fields[12])
	}
	if len(fields) > 13 {
		msg.Priority = strings.TrimSpace(fields[13])
	}
	if len(fields) > 14 {
		msg.State = ParseStateType(strings.TrimSpace(fields[14]))
	}
	if len(fields) > 15 {
		msg.RefId = strings.TrimSpace(fields[15])
	}
	if len(fields) > 16 {
		msg.RefName = strings.TrimSpace(fields[16])
	}
	if len(fields) > 17 {
		msg.RefNat = strings.TrimSpace(fields[17])
	}

	return msg
}

func parseRightFencerArea(msg *CyranoMessage, fields []string) {
	// Right fencer area (R1-R12)
	// Skip empty leading field (index 0)
	if len(fields) > 1 {
		msg.RightId = strings.TrimSpace(fields[1])
	}
	if len(fields) > 2 {
		msg.RightName = strings.TrimSpace(fields[2])
	}
	if len(fields) > 3 {
		msg.RightNat = strings.TrimSpace(fields[3])
	}
	if len(fields) > 4 {
		msg.RightScore = strings.TrimSpace(fields[4])
	}
	if len(fields) > 5 {
		msg.RightStatus = strings.TrimSpace(fields[5])
	}
	if len(fields) > 6 {
		msg.RightYCard = strings.TrimSpace(fields[6])
	}
	if len(fields) > 7 {
		msg.RightRCard = strings.TrimSpace(fields[7])
	}
	if len(fields) > 8 {
		msg.RightLight = strings.TrimSpace(fields[8])
	}
	if len(fields) > 9 {
		msg.RightWLight = strings.TrimSpace(fields[9])
	}
	if len(fields) > 10 {
		msg.RightMedical = strings.TrimSpace(fields[10])
	}
	if len(fields) > 11 {
		msg.RightReserve = strings.TrimSpace(fields[11])
	}
	if len(fields) > 12 {
		msg.RightPCard = strings.TrimSpace(fields[12])
	}
}

func parseLeftFencerArea(msg *CyranoMessage, fields []string) {
	// Left fencer area (L1-L12)
	// Skip empty leading field (index 0)
	if len(fields) > 1 {
		msg.LeftId = strings.TrimSpace(fields[1])
	}
	if len(fields) > 2 {
		msg.LeftName = strings.TrimSpace(fields[2])
	}
	if len(fields) > 3 {
		msg.LeftNat = strings.TrimSpace(fields[3])
	}
	if len(fields) > 4 {
		msg.LeftScore = strings.TrimSpace(fields[4])
	}
	if len(fields) > 5 {
		msg.LeftStatus = strings.TrimSpace(fields[5])
	}
	if len(fields) > 6 {
		msg.LeftYCard = strings.TrimSpace(fields[6])
	}
	if len(fields) > 7 {
		msg.LeftRCard = strings.TrimSpace(fields[7])
	}
	if len(fields) > 8 {
		msg.LeftLight = strings.TrimSpace(fields[8])
	}
	if len(fields) > 9 {
		msg.LeftWLight = strings.TrimSpace(fields[9])
	}
	if len(fields) > 10 {
		msg.LeftMedical = strings.TrimSpace(fields[10])
	}
	if len(fields) > 11 {
		msg.LeftReserve = strings.TrimSpace(fields[11])
	}
	if len(fields) > 12 {
		msg.LeftPCard = strings.TrimSpace(fields[12])
	}
}
