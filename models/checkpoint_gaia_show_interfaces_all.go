package models

type CheckpointGaiaShowInterfacesAll struct {
	Interface	string	`json:"INTERFACE"`
	State	string	`json:"STATE"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Type	string	`json:"TYPE"`
	Link_state	string	`json:"LINK_STATE"`
	Mtu	string	`json:"MTU"`
	Instance	string	`json:"INSTANCE"`
	Autoneg	string	`json:"AUTONEG"`
	Speed	string	`json:"SPEED"`
	Comment	string	`json:"COMMENT"`
	Ipv4_address	string	`json:"IPV4_ADDRESS"`
	Ipv6_address	string	`json:"IPV6_ADDRESS"`
	Ipv6_ll_address	string	`json:"IPV6_LL_ADDRESS"`
	Ipv6_ll_mask	string	`json:"IPV6_LL_MASK"`
}