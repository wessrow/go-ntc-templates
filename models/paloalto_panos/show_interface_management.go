package paloalto_panos

type ShowInterfaceManagement struct {
	Ipv6_address   string `json:"IPV6_ADDRESS"`
	Ipv6_linklocal string `json:"IPV6_LINKLOCAL"`
	Ipv6_gateway   string `json:"IPV6_GATEWAY"`
	Name           string `json:"NAME"`
	Mac_address    string `json:"MAC_ADDRESS"`
	Ipv4_address   string `json:"IPV4_ADDRESS"`
	Ipv4_netmask   string `json:"IPV4_NETMASK"`
	Ipv4_gateway   string `json:"IPV4_GATEWAY"`
}

var ShowInterfaceManagement_Template string = `Value NAME (.+)
Value MAC_ADDRESS ([a-fA-F0-9]{2}\:[a-fA-F0-9]{2}\:[a-fA-F0-9]{2}\:[a-fA-F0-9]{2}\:[a-fA-F0-9]{2}\:[a-fA-F0-9]{2})
Value IPV4_ADDRESS (\S+)
Value IPV4_NETMASK (\S+)
Value IPV4_GATEWAY (\S+)
Value IPV6_ADDRESS (\S+)
Value IPV6_LINKLOCAL (\S+)
Value IPV6_GATEWAY (\S+)

Start
  ^-+
  ^Name:\s+${NAME}$$
  ^Link\s+status
  ^\s+Runtime
  ^\s+Configured\s+link
  ^MAC\s+address
  ^\s+Port\s+MAC\s+address\s+${MAC_ADDRESS}
  ^Ip\s+address:\s+${IPV4_ADDRESS}
  ^Netmask:\s+${IPV4_NETMASK}
  ^Default\s+gateway:\s+${IPV4_GATEWAY}
  ^Ipv6\s+address:\s+${IPV6_ADDRESS}
  ^Ipv6\s+link\s+local\s+address:\s+${IPV6_LINKLOCAL}
  ^Ipv6\s+default\s+gateway:(\s+${IPV6_GATEWAY})?
  ^Logical\s+interface\s+counters
  ^bytes\s+received
  ^bytes\s+transmitted
  ^packets\s+received
  ^packets\s+transmitted
  ^receive\s+errors
  ^transmit\s+errors
  ^receive\s+packets\s+dropped
  ^transmit\s+packets\s+dropped
  ^multicast\s+packets\s+received
  ^\s*$$
  ^. -> Error
`
