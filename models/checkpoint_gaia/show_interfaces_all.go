package checkpoint_gaia

type ShowInterfacesAll struct {
	Link_state      string `json:"LINK_STATE"`
	Comment         string `json:"COMMENT"`
	Ipv6_ll_address string `json:"IPV6_LL_ADDRESS"`
	Type            string `json:"TYPE"`
	Instance        string `json:"INSTANCE"`
	Autoneg         string `json:"AUTONEG"`
	Ipv6_address    string `json:"IPV6_ADDRESS"`
	Mtu             string `json:"MTU"`
	State           string `json:"STATE"`
	Ipv4_address    string `json:"IPV4_ADDRESS"`
	Interface       string `json:"INTERFACE"`
	Speed           string `json:"SPEED"`
	Ipv6_ll_mask    string `json:"IPV6_LL_MASK"`
	Mac_address     string `json:"MAC_ADDRESS"`
}

var ShowInterfacesAll_Template string = `Value INTERFACE (\S+)
Value STATE (\w+)
Value MAC_ADDRESS (.*)
Value TYPE (\S+)
Value LINK_STATE (.*)
Value MTU (\d+)
Value INSTANCE (\d+)
Value AUTONEG (.*)
Value SPEED (.*)
Value COMMENT (\S.*)
Value IPV4_ADDRESS ((?:[0-9]{1,3}\.){3}[0-9]{1,3}(/\d{1,2})?|\w+)
Value IPV6_ADDRESS ([0-9a-f:]*)
Value IPV6_LL_ADDRESS ([0-9a-f:]*)
Value IPV6_LL_MASK (\d+)


Start
  ^Interface\s${INTERFACE}
  ^\s+state\s${STATE}
  ^\s+mac-addr\s${MAC_ADDRESS}
  ^\s+type\s${TYPE}
  ^\s+link-state\s${LINK_STATE}
  ^\s+instance\s${INSTANCE}
  ^\s+mtu\s${MTU}
  ^\s+auto-negotiation\s${AUTONEG}
  ^\s+speed\s${SPEED}
  # sometimes the comment is enclosed in braces
  ^\s+comments\s${COMMENT}\s*$$
  ^\s+ipv4-address\s${IPV4_ADDRESS}\s*$$
  ^\s+ipv6-address\s${IPV6_ADDRESS}\s*$$
  ^\s+ipv6-local-link-address\s${IPV6_LL_ADDRESS}/${IPV6_LL_MASK}
  ^Statistics: -> Record`
