package zte_zxros 

type ShowArp struct {
	Address	string	`json:"ADDRESS"`
	Age	string	`json:"AGE"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Interface	[]string	`json:"INTERFACE"`
	State	string	`json:"STATE"`
	External_vlan	string	`json:"EXTERNAL_VLAN"`
	Internal_vlan	string	`json:"INTERNAL_VLAN"`
	Sub_interface	[]string	`json:"SUB_INTERFACE"`
}

var ShowArp_Template = `Value ADDRESS (\d+\.\d+\.\d+\.\d+)
Value AGE (\S+)
Value MAC_ADDRESS (\S+)
Value List INTERFACE (\S+)
Value STATE (\S+)
Value EXTERNAL_VLAN (\d+|N/A)
Value INTERNAL_VLAN (\d+|N/A)
Value List SUB_INTERFACE (\S+)

Start
  ^\d+\.\d+\.\d+\.\d+ -> Continue.Record
  ^${ADDRESS}\s+${AGE}\s+${MAC_ADDRESS}\s+${INTERFACE}\s+${EXTERNAL_VLAN}\s+${INTERNAL_VLAN}\s+${SUB_INTERFACE}\s*$$
  ^\s+${INTERFACE}\s*$$
  ^\s+${INTERFACE}(\s+${SUB_INTERFACE})?\s*$$
  ^Arp\s+protect
  ^The\s+count\s+is\s+\d+
  ^IP\s+Hardware\s+Exter\s+Inter\s+Sub
  ^Address\s+Age\s+Address\s+Interface\s+VlanID\s+VlanID\s+Interface
  ^-+
  ^\s*$$
  ^. -> Error
`