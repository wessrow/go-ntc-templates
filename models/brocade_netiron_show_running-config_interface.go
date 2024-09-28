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