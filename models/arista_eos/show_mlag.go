package arista_eos 

type ShowMlag struct {
	Domain	string	`json:"DOMAIN"`
	Interface	string	`json:"INTERFACE"`
	Peer_addr	string	`json:"PEER_ADDR"`
	Peer_link	string	`json:"PEER_LINK"`
	State	string	`json:"STATE"`
}

var ShowMlag_Template = `Value DOMAIN (.*)
Value INTERFACE (\S+)
Value PEER_ADDR (\S+)
Value PEER_LINK (\S+)
Value STATE (\S+)

Start
  ^domain-id\s+:\s+${DOMAIN} 
  ^local-interface\s+:\s+${INTERFACE}
  ^peer-address\s+:\s+${PEER_ADDR}
  ^peer-link\s+:\s+${PEER_LINK} 
  ^state\s+:\s+${STATE} -> Record
`