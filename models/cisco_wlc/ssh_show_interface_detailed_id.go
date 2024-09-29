package cisco_wlc 

type SshShowInterfaceDetailedId struct {
	Interface_name	string	`json:"INTERFACE_NAME"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Netmask	string	`json:"NETMASK"`
	Gateway	string	`json:"GATEWAY"`
	Primary_dhcp_server	string	`json:"PRIMARY_DHCP_SERVER"`
	Secondary_dhcp_server	string	`json:"SECONDARY_DHCP_SERVER"`
}

var SshShowInterfaceDetailedId_Template = `Value INTERFACE_NAME (\S+)
Value MAC_ADDRESS (\w+\.\w+\.\w+|\w+\:\w+\:\w+\:\w+\:\w+\:\w+)
Value IP_ADDRESS (\d+.\d+.\d+.\d+)
Value NETMASK (\d+.\d+.\d+.\d+)
Value GATEWAY (\d+.\d+.\d+.\d+)
Value PRIMARY_DHCP_SERVER (\d+.\d+.\d+.\d+|Unconfigured)
Value SECONDARY_DHCP_SERVER (\d+.\d+.\d+.\d+|Unconfigured)

Start
  ^\s*Interface\s+Name\.+\s+${INTERFACE_NAME}\s*$$
  ^\s*MAC\s+Address\.+\s+${MAC_ADDRESS}\s*$$
  ^\s*IP\s+Address\.+\s+${IP_ADDRESS}\s*$$
  ^\s*IP\s+Netmask\.+\s+${NETMASK}\s*$$
  ^\s*IP\s+Gateway\.+\s+${GATEWAY}\s*$$
  ^\s*Primary\s+DHCP\s+Server\.+\s+${PRIMARY_DHCP_SERVER}\s*$$
  ^\s*Secondary\s+DHCP\s+Server\.+\s+${SECONDARY_DHCP_SERVER}\s*$$ -> Record
  ^\s*External\s+NAT\s+IP\s+State.*$$
  ^\s*External\s+NAT\s+IP\s+Address.*$$
  ^\s*Link\s+Local\s+IPv6\s+Address.*$$
  ^\s*STATE.*$$
  ^\s*Primary\s+IPv6\s+Address.*$$
  ^\s*IPv6\s+Address.*$$
  ^\s*Primary\s+IPv6\s+Gateway.*$$
  ^\s*IPv6\s+Gateway.*$$
  ^\s*Primary\s+IPv6\s+Gateway\s+Mac\s+Address.*$$
  ^\s*IPv6\s+Gateway\s+Mac\s+Address.*$$
  ^\s*VLAN.*$$
  ^\s*Quarantine-vlan.*$$
  ^\s*NAS-Identifier.*$$
  ^\s*Active\s+Physical\s+Port.*$$
  ^\s*Primary\s+Physical\s+Port.*$$
  ^\s*Backup\s+Physical\s+Port.*$$
  ^\s*DHCP\s+Proxy\s+Mode.*$$
  ^\s*DHCP\s+Option\s+82.*$$
  ^\s*DHCP\s+Option\s+82\s+bridge.*$$
  ^\s*DHCP\s+Option\s+6\s+Opendns\s+Override.*$$
  ^\s*IPv4\s+ACL.*$$
  ^\s*URL\s+ACL.*$$
  ^\s*mDNS\s+Profile\s+Name.*$$
  ^\s*AP\s+Manager.*$$
  ^\s*Guest\s+Interface.*$$
  ^\s*3G\s+VLAN.*$$
  ^\s*L2\s+Multicast.*$$
  ^\s*IPv6\s+ACL.*$$
  ^\s*$$
  ^.* -> Error
`