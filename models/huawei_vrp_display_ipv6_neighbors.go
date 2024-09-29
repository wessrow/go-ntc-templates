package models

type HuaweiVrpDisplayIpv6Neighbors struct {
	Ipv6_address	string	`json:"IPV6_ADDRESS"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	State	string	`json:"STATE"`
	Interface	string	`json:"INTERFACE"`
	Age	string	`json:"AGE"`
	Vlan_id	string	`json:"VLAN_ID"`
	Ce_vlan	string	`json:"CE_VLAN"`
	Vpn_name	string	`json:"VPN_NAME"`
	Is_router	string	`json:"IS_ROUTER"`
	Secure_flag	string	`json:"SECURE_FLAG"`
	Nickname	string	`json:"NICKNAME"`
}

var HuaweiVrpDisplayIpv6Neighbors_Template = `Value IPV6_ADDRESS (\S+)
Value MAC_ADDRESS (([\d\w]{4}-){2}[\d\w]{4})
Value STATE (INCMP|REACH|STALE|DELAY|PROBE)
Value INTERFACE (\S+)
Value AGE (\S+)
Value VLAN_ID ([-\d]+)
Value CE_VLAN ([-\d]+)
Value VPN_NAME (\S+|\s{0})
Value IS_ROUTER (TRUE|FALSE)
Value SECURE_FLAG (SECURE|UNSECURE|UN-SECURE)
Value NICKNAME (\S+)

Start
  ^\s*$$
  ^\s*-+\s*$$
  ^\s*IPv6\sAddress\s*:\s*\S+\s*$$ -> Continue.Record
  ^\s*IPv6\sAddress\s*:\s*${IPV6_ADDRESS}\s*$$
  ^\s*Link-layer\s*:\s*${MAC_ADDRESS}\s+State\s*:\s*${STATE}\s*$$
  ^\s*Interface\s*:\s*${INTERFACE}\s+Age\s*:\s*${AGE}\s*$$
  ^\s*VLAN\s*:\s*${VLAN_ID}\s+CEVLAN\s*:\s*${CE_VLAN}\s*$$
  ^\s*(VPN\sname\s*:\s*${VPN_NAME}\s+)?Is\sRouter\s*:\s*${IS_ROUTER}\s*$$
  ^\s*Secure FLAG\s*:\s*${SECURE_FLAG}\s*(Nickname\s*:\s*${NICKNAME}\s*)?$$
  ^\s*Total:\s+\d+\s+Dynamic:\s+\d+\s+Static:\s+\d+\s*$$
  ^. -> Error

`