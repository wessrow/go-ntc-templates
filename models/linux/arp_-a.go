package linux 

type ArpA struct {
	Rev_dns	string	`json:"REV_DNS"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Hw_type	string	`json:"HW_TYPE"`
	Interface	string	`json:"INTERFACE"`
}

var ArpA_Template = `Value REV_DNS (\S+)
Value IP_ADDRESS (\S+)
Value MAC_ADDRESS (\S+)
Value HW_TYPE (\S+)
Value INTERFACE (\S+)

Start
  ^${REV_DNS}\s+\(${IP_ADDRESS}\)\s+\S+\s+${MAC_ADDRESS}\s+\[${HW_TYPE}\]\s+\S+\s+${INTERFACE}$$ -> Record
  ^${REV_DNS}\s+\(${IP_ADDRESS}\)\s+\S+\s+<${MAC_ADDRESS}>\s+\S+\s+${INTERFACE}$$ -> Record
  ^\s*$$
  ^. -> Error
`