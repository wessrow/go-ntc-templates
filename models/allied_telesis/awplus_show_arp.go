package allied_telesis 

type AwplusShowArp struct {
	Ip_address	string	`json:"IP_ADDRESS"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Type	string	`json:"TYPE"`
	Interface	string	`json:"INTERFACE"`
	Port	string	`json:"PORT"`
}

var AwplusShowArp_Template = `Value Required IP_ADDRESS (\d+\.\d+\.\d+\.\d+)
Value Required MAC_ADDRESS (\S+)
Value TYPE (\S+)
Value INTERFACE (\S+)
Value PORT (\S+)

Start
  ^.*IP\s+Address\s+MAC\s+Address\s+Interface\s+Port\s+Type
  ^.*IP\s+Address\s+LL\s+Address\s+Interface\s+Port\s+Type
  ^${IP_ADDRESS}\s+${MAC_ADDRESS}\s+${INTERFACE}\s+${PORT}\s+${TYPE} -> Record
  ^. -> Error
`