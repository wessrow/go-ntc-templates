package models

type BrocadeNetironShowRunningConfigInterface struct {
	Interface_type	string	`json:"INTERFACE_TYPE"`
	Interface	string	`json:"INTERFACE"`
	Portname	string	`json:"PORTNAME"`
	Qostostrust	string	`json:"QOSTOSTRUST"`
	Qostosmark	string	`json:"QOSTOSMARK"`
	Ospfarea	string	`json:"OSPFAREA"`
	Ospfpassive	string	`json:"OSPFPASSIVE"`
	Ospfcost	string	`json:"OSPFCOST"`
	Vrf	string	`json:"VRF"`
	Ippimsparse	string	`json:"IPPIMSPARSE"`
	Iprouterisis	string	`json:"IPROUTERISIS"`
	Isismetric	string	`json:"ISISMETRIC"`
	Ip_helper	[]string	`json:"IP_HELPER"`
	Ipredirect	string	`json:"IPREDIRECT"`
	Ip_address	[]string	`json:"IP_ADDRESS"`
	Ip_addr_cidr	[]string	`json:"IP_ADDR_CIDR"`
	Ipv6_addr	[]string	`json:"IPV6_ADDR"`
	Ipv6_addr_cidr	[]string	`json:"IPV6_ADDR_CIDR"`
	Admin_state	string	`json:"ADMIN_STATE"`
	Acl_in	string	`json:"ACL_IN"`
	Acl_out	string	`json:"ACL_OUT"`
	Vrid	[]string	`json:"VRID"`
	Ip_addr_vrrpe	[]string	`json:"IP_ADDR_VRRPE"`
}

var BrocadeNetironShowRunningConfigInterface_Template = `Value INTERFACE_TYPE (.+?)
Value Required INTERFACE (.+?)
Value PORTNAME (.+?)
Value QOSTOSTRUST ([a-z\-]+)
Value QOSTOSMARK ([a-z\-]+)
Value OSPFAREA (\d+)
Value OSPFPASSIVE (.+?)
Value OSPFCOST (\d+)
Value VRF (.+?)
Value IPPIMSPARSE (pim-sparse)
Value IPROUTERISIS (isis)
Value ISISMETRIC (\d+)
Value List IP_HELPER (.+?)
Value IPREDIRECT ([no]*)
Value List IP_ADDRESS (\S+?)
Value List IP_ADDR_CIDR (\d*)
Value List IPV6_ADDR (\S+?)
Value List IPV6_ADDR_CIDR (\d*)
Value ADMIN_STATE (disable)
Value ACL_IN (\S+)
Value ACL_OUT (\S+)
Value List VRID (\d+)
Value List IP_ADDR_VRRPE (\S+)


Start
  ^interface -> Continue.Record
  ^interface\s+${INTERFACE_TYPE}\s+${INTERFACE}\s*$$
  ^\s+${ADMIN_STATE}\s*$$
  ^\s+port-name\s+${PORTNAME}\s*$$
  ^\s+qos-tos\s+trust\s+${QOSTOSTRUST}
  ^\s+qos-tos\s+mark\s+${QOSTOSMARK}
  ^\s+ip\s+ospf\s+area\s+${OSPFAREA}\s*$$
  ^\s+ip\s+ospf\s+${OSPFPASSIVE}
  ^\s+ip\s+ospf\s+cost\s+${OSPFCOST}
  ^\s+vrf\s+forwarding\s+${VRF}\s*$$
  ^\s+ip\s+${IPPIMSPARSE}\s*$$
  ^\s+ip\s+router\s+${IPROUTERISIS}
  ^\s+isis\s+metric\s+${ISISMETRIC}\s*$$
  ^\s+${IPREDIRECT}\s*ip\s+redirect
  ^\s+ip\s+helper-address\s+${IP_HELPER}\s*$$
  ^\s+ip\s+access-group\s+${ACL_IN}\s+in
  ^\s+ip\s+access-group\s+${ACL_OUT}\s+out
  ^\s+ip\s+address\s+${IP_ADDRESS}/*${IP_ADDR_CIDR}\s*$$
  ^\s+ipv6\s+address\s+${IPV6_ADDR}/*${IPV6_ADDR_CIDR}\s*$$
  ^\s+ip(?:v6|)\s+vrrp-extended\s+vrid\s+${VRID} -> VRRPEXTENDED

VRRPEXTENDED
  ^\s+ip(?:v6|)-address\s+${IP_ADDR_VRRPE}
  ^\s+ip(?:v6|)\s+vrrp-extended\s+vrid\s+${VRID}
  ^interface -> Continue.Record
  ^interface\s+${INTERFACE_TYPE}\s+${INTERFACE}\s*$$ -> Start

`