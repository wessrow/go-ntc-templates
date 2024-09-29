package models

type ExtremeExosShowIpconfig struct {
	Use_redirects	string	`json:"USE_REDIRECTS"`
	Ip_option_lsrr	string	`json:"IP_OPTION_LSRR"`
	Ip_option_ssrr	string	`json:"IP_OPTION_SSRR"`
	Ip_option_rr	string	`json:"IP_OPTION_RR"`
	Ip_option_ts	string	`json:"IP_OPTION_TS"`
	Ip_option_ra	string	`json:"IP_OPTION_RA"`
	Route_sharing	string	`json:"ROUTE_SHARING"`
	Route_compression	string	`json:"ROUTE_COMPRESSION"`
	Originated_packets	string	`json:"ORIGINATED_PACKETS"`
	Max_shared_gateways_current	string	`json:"MAX_SHARED_GATEWAYS_CURRENT"`
	Max_shared_gateways_configured	string	`json:"MAX_SHARED_GATEWAYS_CONFIGURED"`
	Route_sharing_hash	string	`json:"ROUTE_SHARING_HASH"`
	Irdp_advertisement_address	string	`json:"IRDP_ADVERTISEMENT_ADDRESS"`
	Irdp_max_interval	string	`json:"IRDP_MAX_INTERVAL"`
	Irdp_min_interval	string	`json:"IRDP_MIN_INTERVAL"`
	Irdp_lifetime	string	`json:"IRDP_LIFETIME"`
	Irdp_preference	string	`json:"IRDP_PREFERENCE"`
	Interface	[]string	`json:"INTERFACE"`
	Ip	[]string	`json:"IP"`
	Subnet	[]string	`json:"SUBNET"`
	Flags	[]string	`json:"FLAGS"`
	Nsia	[]string	`json:"NSIA"`
}

var ExtremeExosShowIpconfig_Template = `Value USE_REDIRECTS (.*)
Value IP_OPTION_LSRR (.*)
Value IP_OPTION_SSRR (.*)
Value IP_OPTION_RR (.*)
Value IP_OPTION_TS (.*)
Value IP_OPTION_RA (.*)
Value ROUTE_SHARING (.*)
Value ROUTE_COMPRESSION (.*)
Value ORIGINATED_PACKETS (.*)
Value MAX_SHARED_GATEWAYS_CURRENT (\d+)
Value MAX_SHARED_GATEWAYS_CONFIGURED (\d+)
Value ROUTE_SHARING_HASH (.*)
Value IRDP_ADVERTISEMENT_ADDRESS (\S+)
Value IRDP_MAX_INTERVAL (\d+)
Value IRDP_MIN_INTERVAL (\d+)
Value IRDP_LIFETIME (\d+)
Value IRDP_PREFERENCE (\d+)
Value List INTERFACE (\S+)
Value List IP (\S+)
Value List SUBNET (/\d+)
Value List FLAGS (\S+)
Value List NSIA (\d+)

Start
  ^\s*Use\s+Redirects\s*:\s*${USE_REDIRECTS}\s*$$
  ^\s*IpOption\s+LSRR\s*:\s*${IP_OPTION_LSRR}\s*$$
  ^\s*IpOption\s+SSRR\s*:\s*${IP_OPTION_SSRR}\s*$$
  ^\s*IpOption\s+RR\s*:\s*${IP_OPTION_RR}\s*$$
  ^\s*IpOption\s+TS\s*:\s*${IP_OPTION_TS}\s*$$
  ^\s*IpOption\s+RA\s*:\s*${IP_OPTION_RA}\s*$$
  ^\s*Route\s+Sharing\s*:\s*${ROUTE_SHARING}\s*$$
  ^\s*Route\s+Compression\s*:\s*${ROUTE_COMPRESSION}\s*$$
  ^\s*Originated\s+Packets\s*:\s*${ORIGINATED_PACKETS}\s*$$
  ^\s*Max\s+Shared\s+Gateways\s*:\s*Current:\s*${MAX_SHARED_GATEWAYS_CURRENT}\s+Configured:\s*${MAX_SHARED_GATEWAYS_CONFIGURED}\s*$$
  ^\s*Route\s+Sharing\s+Hash\s*:\s*${ROUTE_SHARING_HASH}\s*$$
  ^\s*IRDP:\s*$$ -> Irdp
  ^\s*Interface\s+IP\s+Address\s+Flags\s+nSIA\s*$$ -> IpAdresses
  ^\s*Flags:\s+\(A\)\s+Address\s+Mask\s+Reply\s+Enabled\s+\(B\)\s+BOOTP\s+Enabled\s*$$
  ^\s*\(b\)\s+Broadcast\s+Forwarding\s+Enabled\s*$$
  ^\s*\(D\)\s+Duplicate\s+address\s+detected\s+on\s+VLAN,\s+\(E\)\s+Interface\s+Enabled\s*$$
  ^\s*\(f\)\s+Forwarding\s+Enabled\s+\(g\)\s+Ignore\s+IP\s+Broadcast\s+Enabled\s*$$
  ^\s*\(h\)\s+Directed\s+Broadcast\s+Forwarding\s+by\s+Hardware\s+Enabled\s*$$
  ^\s*\(I\)\s+IRDP\s+Advertisement\s+Enabled,\s+\(M\)\s+Send\s+Parameter\s+Problem\s+Enabled\s*$$
  ^\s*\(m\)\s+Multicast\s+forwarding\s+Enabled,\s+\(n\)\s+Multinetted\s+VLAN\s*$$
  ^\s*\(nSIA\s+\)\s+Number\s+of\s+Secondary\s+IP\s+Addresses\s*$$
  ^\s*\(P\)\s+Send\s+Port\s+Unreachables\s+Enabled,\s+\(R\)\s+Send\s+Redirects\s+Enabled\s*$$
  ^\s*\(r\)\s+Unicast\s+Reverse\s+Path\s+Enabled\s+on\s+at\s+least\s+one\s+port\s+of\s+the\s+VLAN\s*$$
  ^\s*\(t\)\s+Tentative\s+address,\s+\(T\)\s+Time\s+Stamp\s+Reply\s+Enabled\s*$$
  ^\s*\(u\)\s+Send\s+Unreachables\s+Enabled,\s+\(U\)\s+Interface\s+Up\s*$$
  ^\s*\(v\)\s+VRRP\s+Enabled,\s+\(X\)\s+Send\s+Time\s+Exceeded\s+Enabled\s*$$
  ^\s*
  ^. -> Error

Irdp
  ^\s*Advertisement\s+Address:\s+${IRDP_ADVERTISEMENT_ADDRESS}\s+Maximum\s+Interval:\s+${IRDP_MAX_INTERVAL}\s*$$
  ^\s*Minimum\s+Interval:\s+${IRDP_MIN_INTERVAL}\s+Lifetime:\s+${IRDP_LIFETIME}\s+Preference:\s+${IRDP_PREFERENCE}\s*$$
  ^\s* -> Start
  ^. -> Error

IpAdresses
  ^\s*${INTERFACE}\s+${IP}\s+${SUBNET}\s+${FLAGS}\s+${NSIA}\s*$$
  ^\s* -> Start
  ^. -> Error

`