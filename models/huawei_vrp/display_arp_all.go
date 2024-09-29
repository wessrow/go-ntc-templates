package huawei_vrp 

type DisplayArpAll struct {
	Ip_address	string	`json:"IP_ADDRESS"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Expire	string	`json:"EXPIRE"`
	Type	string	`json:"TYPE"`
	Interface	string	`json:"INTERFACE"`
	Vpn_instance	string	`json:"VPN_INSTANCE"`
}

var DisplayArpAll_Template = `Value Required IP_ADDRESS (\d+\.\d+\.\d+\.\d+)
Value Required MAC_ADDRESS (\S+)
Value EXPIRE (\d+)
Value Required TYPE (\S+\s\S+|\S+)
Value Required INTERFACE (\S+)
Value VPN_INSTANCE (\S+)

Start
  ^IP\s+ADDRESS\s+MAC\s+ADDRESS\s+EXPIRE\S+\s+TYPE\s+INTERFACE\s+VPN-INSTANCE
  ^\s+VLAN\/CEVLAN
  ^-+
  ^${IP_ADDRESS}\s+${MAC_ADDRESS}\s+${EXPIRE}?\s+${TYPE}\s+${INTERFACE}(\s+)?${VPN_INSTANCE}?$$ -> Record
  ^\s+(\d+)\/(\S+)
  ^Total:(\d+)\s+Dynamic:(\d+)\s+Static:(\d+)\s+Interface:(\d+)
  ^Redirect:(\d+)
  ^\s*$$
  ^. -> Error
`