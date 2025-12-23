package main

// Protocol version enumeration
type ProtocolVersion int

const (
	ProtocolEFP1 ProtocolVersion = iota
	ProtocolEFP1_1
)

// String returns the string representation of the protocol version
func (p ProtocolVersion) String() string {
	switch p {
	case ProtocolEFP1:
		return "Cyrano 1.0"
	case ProtocolEFP1_1:
		return "Cyrano 1.1"
	default:
		return "unknown"
	}
}

// ParseProtocolVersion converts a string to ProtocolVersion
func ParseProtocolVersion(s string) ProtocolVersion {
	switch s {
	case "EFP1":
		return ProtocolEFP1
	case "EFP1.1":
		return ProtocolEFP1_1
	default:
		return ProtocolEFP1 // default to EFP1
	}
}

// Command type enumeration
type CommandType int

const (
	CommandHELLO CommandType = iota
	CommandDISP
	CommandACK
	CommandNAK
	CommandINFO
	CommandNEXT
	CommandPREV
)

// String returns the string representation of the command type
func (c CommandType) String() string {
	switch c {
	case CommandHELLO:
		return "HELLO"
	case CommandDISP:
		return "DISP"
	case CommandACK:
		return "ACK"
	case CommandNAK:
		return "NAK"
	case CommandINFO:
		return "INFO"
	case CommandNEXT:
		return "NEXT"
	case CommandPREV:
		return "PREV"
	default:
		return "unknown"
	}
}

// ParseCommand converts a string to CommandType
func ParseCommand(s string) CommandType {
	switch s {
	case "HELLO":
		return CommandHELLO
	case "DISP":
		return CommandDISP
	case "ACK":
		return CommandACK
	case "NAK":
		return CommandNAK
	case "INFO":
		return CommandINFO
	case "NEXT":
		return CommandNEXT
	case "PREV":
		return CommandPREV
	default:
		return CommandHELLO // default to HELLO
	}
}

// CompetitionType enumeration
type CompetitionType int

const (
	CompetitionIndividual CompetitionType = iota
	CompetitionTeam
)

// String returns the string representation of the competition type
func (c CompetitionType) String() string {
	switch c {
	case CompetitionIndividual:
		return "Individual"
	case CompetitionTeam:
		return "Team"
	default:
		return "unknown"
	}
}

// ParseCompetitionType converts a string to CompetitionType
func ParseCompetitionType(s string) CompetitionType {
	switch s {
	case "I":
		return CompetitionIndividual
	case "T":
		return CompetitionTeam
	default:
		return CompetitionIndividual // default to Individual
	}
}

// StateType enumeration
type StateType int

const (
	StateFencing StateType = iota
	StateHalt
	StatePause
	StateWaiting
	StateEnding
)

// String returns the string representation of the state type
func (s StateType) String() string {
	switch s {
	case StateFencing:
		return "Fencing"
	case StateHalt:
		return "Halt"
	case StatePause:
		return "Pause"
	case StateWaiting:
		return "Waiting"
	case StateEnding:
		return "Ending"
	default:
		return "unknown"
	}
}

// ParseStateType converts a string to StateType
func ParseStateType(s string) StateType {
	switch s {
	case "F":
		return StateFencing
	case "H":
		return StateHalt
	case "P":
		return StatePause
	case "W":
		return StateWaiting
	case "E":
		return StateEnding
	default:
		return StateFencing // default to Fencing
	}
}
