package huawei_vrp

type DisplayIpRoutingTableVerbose struct {
	Qos_info       string `json:"QOS_INFO"`
	Neighbour      string `json:"NEIGHBOUR"`
	Age            string `json:"AGE"`
	Protocol       string `json:"PROTOCOL"`
	Tag            string `json:"TAG"`
	Tunnel_id      string `json:"TUNNEL_ID"`
	Entry_id       string `json:"ENTRY_ID"`
	Entry_flags    string `json:"ENTRY_FLAGS"`
	Reference_cnt  string `json:"REFERENCE_CNT"`
	Process_id     string `json:"PROCESS_ID"`
	Prefix_length  string `json:"PREFIX_LENGTH"`
	Label          string `json:"LABEL"`
	Preference     string `json:"PREFERENCE"`
	Cost           string `json:"COST"`
	Indirect_id    string `json:"INDIRECT_ID"`
	Flags          string `json:"FLAGS"`
	Destination    string `json:"DESTINATION"`
	State          string `json:"STATE"`
	Priority       string `json:"PRIORITY"`
	Relay_next_hop string `json:"RELAY_NEXT_HOP"`
	Interface      string `json:"INTERFACE"`
	Next_hop       string `json:"NEXT_HOP"`
}

var DisplayIpRoutingTableVerbose_Template string = `Value DESTINATION (\S+)
Value PREFIX_LENGTH (\d+)
Value PROTOCOL (Direct|Static|EBGP|IBGP||ISIS|OSPF|RIP|UNR|O_ASE)
Value PROCESS_ID (\d+)
Value LABEL (\S+)
Value QOS_INFO (\S+)
Value NEIGHBOUR (\S+)
Value NEXT_HOP (\S+)
Value PREFERENCE (\d+)
Value STATE ((\s*(Active|Invalid|Inactive|NoAdv|Adv|Del|Relied|Stale))+)
Value COST (\d+)
Value AGE ([\d\w]+)
Value TAG (\d+)
Value PRIORITY (low|medium|high|critical)
Value INDIRECT_ID (\S+)
Value RELAY_NEXT_HOP (\S+)
Value INTERFACE (\S+)
Value TUNNEL_ID (\S+)
Value FLAGS ((D|R)*)
Value ENTRY_ID (\S+)
Value ENTRY_FLAGS (\S+)
Value REFERENCE_CNT (\d+)

Start
  ^\s*Route\sFlags:.*$$
  ^\s*-*\s*$$
  ^\s*Routing\sTable(s)?\s*:\sPublic\s*$$
  ^\s*Destinations\s*:\s+\d+\s+Routes\s*:\s+\d+\s*$$
  ^\s*Destination\s*:\s+.*$$ -> Continue.Record
  ^\s*Destination\s*:\s+${DESTINATION}(\s+PrefixLength\s*:\s+${PREFIX_LENGTH})?\s*$$
  ^\s*Protocol\s*:\s+${PROTOCOL}\s+Process\sID\s*:\s+${PROCESS_ID}\s*$$
  ^\s*Preference\s*:\s+${PREFERENCE}\s+Cost\s*:\s+${COST}\s*$$
  ^\s*NextHop\s*:\s+${NEXT_HOP}\s+Neighbour\s*:\s+${NEIGHBOUR}\s*$$
  ^\s*State\s*:\s+${STATE}\s+Age\s*:\s+${AGE}\s*$$
  ^\s*Tag\s*:\s+${TAG}\s+Priority\s*:\s+${PRIORITY}\s*$$
  ^\s*Label\s*:\s+${LABEL}\s+QoSInfo\s*:\s+${QOS_INFO}\s*$$
  ^\s*IndirectID\s*:\s+${INDIRECT_ID}\s*$$
  ^\s*RelayNextHop\s*:\s+${RELAY_NEXT_HOP}\s+Interface\s*:\s+${INTERFACE}\s*$$
  ^\s*TunnelID\s*:\s+${TUNNEL_ID}\s+Flags\s*:\s+${FLAGS}\s*$$
  # IPv6 only regex
  ^\s*NextHop\s*:\s+${NEXT_HOP}\s+Preference\s*:\s+${PREFERENCE}\s*$$
  ^\s*Neighbour\s*:\s+${NEIGHBOUR}\s+ProcessID\s*:\s+${PROCESS_ID}\s*$$
  ^\s*Label\s*:\s+${LABEL}\s+Protocol\s*:\s+${PROTOCOL}\s*$$
  ^\s*State\s*:\s+${STATE}\s+Cost\s*:\s+${COST}\s*$$
  ^\s*Entry\sID\s*:\s+${ENTRY_ID}\s+EntryFlags\s*:\s+${ENTRY_FLAGS}\s*$$
  ^\s*Reference\sCnt\s*:\s+${REFERENCE_CNT}\s+Tag\s*:\s+${TAG}\s*$$
  ^\s*Priority\s*:\s+${PRIORITY}\s+Age\s*:\s+${AGE}\s*$$
  ^\s*RelayNextHop\s*:\s+${RELAY_NEXT_HOP}\s+TunnelID\s*:\s+${TUNNEL_ID}\s*$$
  ^\s*Interface\s*:\s+${INTERFACE}\s+Flags\s*:\s+${FLAGS}\s*$$
  ^\s*$$
  ^. -> Error
`
