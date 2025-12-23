package main

// CyranoMessage represents a parsed Cyrano Protocol message
type CyranoMessage struct {
	ProtocolVersion ProtocolVersion
	Command         CommandType

	LeftPCard    string
	LeftReserve  string
	LeftMedical  string
	LeftWLight   string
	LeftLight    string
	LeftRCard    string
	LeftYCard    string
	LeftStatus   string
	LeftScore    string
	LeftNat      string
	LeftName     string
	LeftId       string // Left fencer
	RightPCard   string
	RightReserve string
	RightMedical string
	RightWLight  string
	RightLight   string
	RightRCard   string
	RightYCard   string
	RightStatus  string
	RightScore   string
	RightNat     string
	RightName    string
	RightId      string // Right fencer
	RefNat       string
	RefName      string
	RefId        string
	State        StateType
	Priority     string
	Weapon       string
	Type         CompetitionType
	Stopwatch    string
	Time         string
	Round        string
	Match        string
	PoulTab      string
	Phase        string
	Competition  string
	Piste        string
}
