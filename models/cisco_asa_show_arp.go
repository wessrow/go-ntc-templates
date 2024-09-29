package models

type CiscoAsaShowArp struct {
	Interface	string	`json:"INTERFACE"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Age	string	`json:"AGE"`
	Mac_address	string	`json:"MAC_ADDRESS"`
}

var CiscoAsaShowArp_Template = `Value Required INTERFACE (\S+)
Value Required IP_ADDRESS (\d+\.\d+\.\d+\.\d+)
Value Required AGE (\S+)
Value Required MAC_ADDRESS (\S+)

Start
  ^\s*${INTERFACE}\s+${IP_ADDRESS}\s+${MAC_ADDRESS}\s+${AGE} -> Record
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is
  ^\s*$$
  ^. -> Error

`