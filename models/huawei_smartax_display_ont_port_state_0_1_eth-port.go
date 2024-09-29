package models

type HuaweiSmartaxDisplayOntPortState01EthPort struct {
	Ont_id	string	`json:"ONT_ID"`
	Ont_port_id	string	`json:"ONT_PORT_ID"`
	Ont_port_type	string	`json:"ONT_PORT_TYPE"`
	Speed_mbps	string	`json:"SPEED_MBPS"`
	Duplex	string	`json:"DUPLEX"`
	Link_state	string	`json:"LINK_STATE"`
	Ring_status	string	`json:"RING_STATUS"`
}

var HuaweiSmartaxDisplayOntPortState01EthPort_Template = `Value Key ONT_ID (\d+)
Value ONT_PORT_ID (\d+)
Value ONT_PORT_TYPE (\S+)
Value SPEED_MBPS (\S+)
Value DUPLEX (\S+)
Value LINK_STATE (up|down)
Value RING_STATUS (\S+)

Start
  ^\s+ONT-ID\s*ONT\s*ONT\s*Speed\(Mbps\)\s*Duplex\s*LinkState\s*RingStatus\s*$$
  ^\s+port-ID\s*Port-type\s*$$
  ^\s+${ONT_ID}\s+${ONT_PORT_ID}\s+${ONT_PORT_TYPE}\s*(-|${SPEED_MBPS})\s*(-|${DUPLEX})\s*(-|${LINK_STATE})\s*(-|${RING_STATUS}) -> Record
  ^\s+-
  ^. -> Error

`