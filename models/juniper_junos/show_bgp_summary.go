package juniper_junos

type ShowBgpSummary struct {
	Interface string `json:"INTERFACE"`
	Received  string `json:"RECEIVED"`
	State     string `json:"STATE"`
	Peer_ip   string `json:"PEER_IP"`
	Peer_as   string `json:"PEER_AS"`
	Up_down   string `json:"UP_DOWN"`
	Out_pkt   string `json:"OUT_PKT"`
	Damped    string `json:"DAMPED"`
	Last_flap string `json:"LAST_FLAP"`
	Active    string `json:"ACTIVE"`
	Accepted  string `json:"ACCEPTED"`
	In_pkt    string `json:"IN_PKT"`
	Out_q     string `json:"OUT_Q"`
	Flaps     string `json:"FLAPS"`
}

var ShowBgpSummary_Template string = `Value Required PEER_IP ([0-9a-f:\.]+)
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
