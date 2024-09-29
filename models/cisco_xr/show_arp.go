package cisco_xr 

type ShowArp struct {
	Ip_address	string	`json:"IP_ADDRESS"`
	Age	string	`json:"AGE"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	State	string	`json:"STATE"`
	Type	string	`json:"TYPE"`
	Interface	string	`json:"INTERFACE"`
	Cpu	string	`json:"CPU"`
}

var ShowArp_Template = `Value Required IP_ADDRESS (\d+\.\d+\.\d+\.\d+)
Value Required AGE (\S+)
Value Required MAC_ADDRESS (\S+)
Value Required STATE (\S+)
Value Required TYPE (\S+)
Value Required INTERFACE (\S+)
Value Filldown CPU (\d+/\S*\d+/CPU\d+$)

Start
  # Match the timestamp at beginning of command output
  ^\S+\s+\S+\s+\d+\s+\d+:\d+:\d+\.\d+\s+\S+\s*$$
  ^Address\s+Age\s+Hardware\s+Addr\s+State\s+Type\s+Interface$$
  ^-+
  ^${CPU}
  ^\s*$$
  ^${IP_ADDRESS}\s+${AGE}\s+${MAC_ADDRESS}\s+${STATE}\s+${TYPE}\s+${INTERFACE}\s*$$ -> Record
  ^. -> Error
`