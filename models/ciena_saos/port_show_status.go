package ciena_saos

type PortShowStatus struct {
	Description  string `json:"DESCRIPTION"`
	Speed_duplex string `json:"SPEED_DUPLEX"`
	Flow_control string `json:"FLOW_CONTROL"`
	Stp          string `json:"STP"`
	Mtu          string `json:"MTU"`
	Name         string `json:"NAME"`
	Link         string `json:"LINK"`
	Duration     string `json:"DURATION"`
	Transceiver  string `json:"TRANSCEIVER"`
}

var PortShowStatus_Template string = `Value NAME (\S+)
Value DESCRIPTION (\S*)
Value LINK (Up|Down)
Value DURATION ([^\|]+)
Value TRANSCEIVER (\S+)
Value STP (\w+)
Value SPEED_DUPLEX (\S+)
Value MTU (\d+)
Value FLOW_CONTROL (\S+)


Start
  ^\+-+\s*PORT\s*OPERATIONAL\s*STATUS\s*-+\+ 
  ^\+-+
  ^\|(\s+\|){3}\s*Link\s*State\s*\|(\s+\|){2}\s*Speed/\s*\|\s*MTU \s*\|\s*Flow\s*\|\s*$$
  ^\|\s*##\s*\|\s*Description\s*\|\s*Link\s*\|\s*Duration\s*\|\s*XCVR\s*\|\s*STP\s*\|\s*Duplex\s*\|\s*Size\s*\|\s*Ctrl\s*\|\s*$$
  ^\|${NAME}\s*\|${DESCRIPTION}\s*\|\s*${LINK}\s*\|\s*${DURATION}\s*\|\s*(${TRANSCEIVER}|\s+)\s*\|\s*${STP}\s*\|\s*(${SPEED_DUPLEX}|\s+)\s*\|\s*${MTU}\s*\|\s*(${FLOW_CONTROL}|\s+)\s*\| -> Record
  ^\s*$$
  ^. -> Error
`
