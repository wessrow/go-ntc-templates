package cisco_ios

type ShowAdjacency struct {
	Endpoint            string   `json:"ENDPOINT"`
	Rewrite_headers     []string `json:"REWRITE_HEADERS"`
	Recursive_interface string   `json:"RECURSIVE_INTERFACE"`
	Recursive_nexthop   string   `json:"RECURSIVE_NEXTHOP"`
	Interface           string   `json:"INTERFACE"`
}

var ShowAdjacency_Template string = `Value Required INTERFACE (\S+)
Value Required ENDPOINT (\S+)
Value List REWRITE_HEADERS ([A-F0-9]+)
Value RECURSIVE_INTERFACE (\S+)
Value RECURSIVE_NEXTHOP (\S+)

Start
  ^Protocol\s+Interface\s+Address\s*$$ -> Entries

Entries
  ^IP.*?\(\d+\).*?$$ -> Continue.Record
  ^IP\s+${INTERFACE}\s+${ENDPOINT}\(\d+\).*?$$
  ^\s+connectionid
  ^\s+\d+\s+packets,\s+\d+\s+bytes
  ^\s+epoch\s+\d+
  ^\s+sourced\s+in\s+sev-epoch
  ^\s+punt\s+\(rate-limited\)\s+packets
  ^\s+empty\s+encap\s+string
  ^\s+P2P-ADJ
  ^\s+Encap\s+length\s+\d+
  ^\s+${REWRITE_HEADERS}$$
  ^\s+L2\s+destination\s+address
  ^\s+Link-type\s+after\s+encap:\s+
  ^\s+Inject\s+p2mp\s+Multicast
  ^\s+ARP
  ^\s+Tun\s+endpt
  ^\s+Next\s+chain\s+element: -> Chain
  ^\s+CEF\s+expires:
  ^\s+refresh:
  ^\s+Epoch:
  ^\s*$$
  ^. -> Error

Chain
  ^\s+IP\sadj\sout\sof\s${RECURSIVE_INTERFACE},\saddr\s${RECURSIVE_NEXTHOP}$$ -> Entries
  ^\s+${RECURSIVE_INTERFACE} -> Entries
  ^\s*$$
  ^. -> Error`
