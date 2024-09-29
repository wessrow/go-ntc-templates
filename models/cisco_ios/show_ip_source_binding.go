package cisco_ios 

type ShowIpSourceBinding struct {
	Mac_address	string	`json:"MAC_ADDRESS"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Lease	string	`json:"LEASE"`
	Type	string	`json:"TYPE"`
	Vlan_id	string	`json:"VLAN_ID"`
	Interface	string	`json:"INTERFACE"`
}

var ShowIpSourceBinding_Template = `Value Required MAC_ADDRESS ((?:[0-9a-fA-F]{2}[:-]){5}(?:[0-9a-fA-F]{2}))
Value Required IP_ADDRESS ((?:[0-9]{1,3}\.){3}[0-9]{1,3})
Value Required LEASE ([0-9]+)
Value Required TYPE (\S+)
Value Required VLAN_ID ([0-9]+)
Value Required INTERFACE (\S+)

Start
  ^\s*MacAddress\s*IpAddress\s*Lease\(sec\)\s*Type\s*VLAN\s*Interface\s*
  ^-+
  ^${MAC_ADDRESS}\s+${IP_ADDRESS}\s+${LEASE}\s+${TYPE}\s+${VLAN_ID}\s+${INTERFACE} -> Record
  ^Total number of bindings
  ^\s*$$
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is
  ^.*$$ -> Error
`