package cisco_nxos

type ShowForwardingAdjacency struct {
	Interface string `json:"INTERFACE"`
	Origin_as string `json:"ORIGIN_AS"`
	Peer_as   string `json:"PEER_AS"`
	Neighbor  string `json:"NEIGHBOR"`
	Slot      string `json:"SLOT"`
	Nexthop   string `json:"NEXTHOP"`
	Rewrite   string `json:"REWRITE"`
}

var ShowForwardingAdjacency_Template string = `# Nexus 7k have multiple slots
Value Filldown SLOT (\d+)
Value Required NEXTHOP ((?:\d{1,3}\.){3}\d{1,3})
Value Required REWRITE (\S+)
Value Required INTERFACE (\S+)
Value ORIGIN_AS (\S+)
Value PEER_AS (\S+)
Value NEIGHBOR (\S+)

Start
  # Extract slot #, if available (Nexus 7k)
  ^slot\s+${SLOT}
  ^====
  # Match headers
  ^IPv4\s+adjacency\s+information
  ^next-hop\s+rewrite\s+info\s+interface\s*$$ -> Entries
  ^next-hop\s+rewrite\s+info\s+interface\s+Origin\s+AS\s+Peer\s+AS\s+Neighbor -> EntriesDetail
  # Empty and unknowns lines
  ^\s*$$
  ^. -> Error

Entries
  # Ignore delimiter
  ^--------
  # Match single adjacency entry (brief view)
  ^${NEXTHOP}\s+${REWRITE}\s+${INTERFACE} -> Record
  # New slot?
  ^slot\s+${SLOT} -> Start
  # Empty and unknowns lines
  ^\s*$$
  ^. -> Error

EntriesDetail
  # Ignore delimiter
  ^--------
  # Match single adjacency entry (detailed view)
  ^${NEXTHOP}\s+${REWRITE}\s+${INTERFACE}\s*(?:${ORIGIN_AS}\s+${PEER_AS}\s+${NEIGHBOR})? -> Record
  # New slot?
  ^slot\s+${SLOT} -> Start
  # Empty and unknowns lines
  ^\s*$$
  ^. -> Error`
