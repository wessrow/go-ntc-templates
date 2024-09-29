package models

type JuniperJunosShowBgpSummary struct {
	Peer_ip	string	`json:"PEER_IP"`
	Peer_as	string	`json:"PEER_AS"`
	In_pkt	string	`json:"IN_PKT"`
	Out_pkt	string	`json:"OUT_PKT"`
	Out_q	string	`json:"OUT_Q"`
	Flaps	string	`json:"FLAPS"`
	Last_flap	string	`json:"LAST_FLAP"`
	Up_down	string	`json:"UP_DOWN"`
	State	string	`json:"STATE"`
	Interface	string	`json:"INTERFACE"`
	Active	string	`json:"ACTIVE"`
	Received	string	`json:"RECEIVED"`
	Accepted	string	`json:"ACCEPTED"`
	Damped	string	`json:"DAMPED"`
}

var JuniperJunosShowBgpSummary_Template = `Value Required PEER_IP ([0-9a-f:\.]+)
Value Required PEER_AS (\d+)
Value IN_PKT (\d+)
Value OUT_PKT (\d+)
Value OUT_Q (\d+)
Value FLAPS (\d+)
Value LAST_FLAP (\S+)
Value UP_DOWN (\S+)
Value STATE (\S+)
Value INTERFACE (\S+)
Value ACTIVE (\d+)
Value RECEIVED (\d+)
Value ACCEPTED (\d+)
Value DAMPED (\d+)

Start
  ^Groups: -> BGP
  ^\s+. -> Error

BGP
  ^Peer\s+AS\s+InPkt\s+OutPkt\s+OutQ\s+Flaps\s+Last\s+Up\/Down\s+State\|#Active\/Received\/Accepted\/Damped...$$
  ^${PEER_IP}\s+${PEER_AS}\s+${IN_PKT}\s+${OUT_PKT}\s+${OUT_Q}\s+${FLAPS}\s+${LAST_FLAP}\s+${UP_DOWN}\s+${STATE} 
  ^\s+${INTERFACE}:\s+${ACTIVE}\/${RECEIVED}\/${ACCEPTED}\/${DAMPED} -> Record
  ^\s+. -> Error

`