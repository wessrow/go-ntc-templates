package zte_zxros

type ShowMplsTrafficEngTunnels struct {
	Signal      string   `json:"SIGNAL"`
	Inlabel     string   `json:"INLABEL"`
	Dst         string   `json:"DST"`
	Instance    string   `json:"INSTANCE"`
	Admin       string   `json:"ADMIN"`
	Oper        string   `json:"OPER"`
	Path        string   `json:"PATH"`
	Outlabel    string   `json:"OUTLABEL"`
	Outlabel_id string   `json:"OUTLABEL_ID"`
	Id          string   `json:"ID"`
	Dest        string   `json:"DEST"`
	Desc        string   `json:"DESC"`
	Inlabel_id  string   `json:"INLABEL_ID"`
	Ip_route    []string `json:"IP_ROUTE"`
	Name        string   `json:"NAME"`
	Src         string   `json:"SRC"`
}

var ShowMplsTrafficEngTunnels_Template string = `Value Required NAME (\S+)
Value DESC (\S+)
Value DEST (\S+)
Value ADMIN (\S+)
Value OPER (\S+)
Value PATH (\S+)
Value SIGNAL (\S+)
Value INLABEL (\S+)
Value INLABEL_ID (\d+)
Value OUTLABEL (\S+)
Value OUTLABEL_ID (\d+)
Value SRC (\S+)
Value DST (\S+)
Value ID (\d+)
Value INSTANCE (\d+)
Value List IP_ROUTE (\d+\.\d+\.\d+\.\d+|NONE)


Start
  ^Name: -> Continue.Record
  ^Name:\s+${NAME}\s*$$
  ^\s+${DESC}\s+Destination:\s+${DEST}\s*$$
  ^\s+Admin:\s+${ADMIN}\s+Oper:\s+${OPER}\s+Path:\s+${PATH}\s+Signalling:\s+${SIGNAL}\s*$$
  ^\s+Signalling:\s+${SIGNAL}\s*$$
  ^\s+InLabel:\s+${INLABEL}(,\s+${INLABEL_ID})?\s*$$
  ^\s+OutLabel:\s+${OUTLABEL}(,\s+${OUTLABEL_ID})?\s*$$
  ^\s+Src\s+${SRC},\s+Dst\s+${DST},\s+Tun-ID\s+${ID},\s+Tun-Instance\s+${INSTANCE}\s*$$
  ^\s+Explicit\s+Route:\s+${IP_ROUTE},*\s*$$
  ^\s+Explicit\s+Route:\s+${IP_ROUTE},* -> Continue
  ^\s+Explicit\s+(?:\S+\s+){2}${IP_ROUTE},*\s*$$
  ^\s+Explicit\s+(?:\S+\s+){2}${IP_ROUTE},* -> Continue
  ^\s+Explicit\s+(?:\S+\s+){3}${IP_ROUTE},*\s*$$
  ^\s+Explicit\s+(?:\S+\s+){3}${IP_ROUTE},* -> Continue
  ^\s+Explicit\s+(?:\S+\s+){4}${IP_ROUTE},*\s*$$
  ^\s+Explicit\s+(?:\S+\s+){4}${IP_ROUTE},* -> Continue
  ^\s+Explicit\s+(?:\S+\s+){5}${IP_ROUTE},*\s*$$
  ^\s+${IP_ROUTE},*\s*$$
  ^\s+${IP_ROUTE},* -> Continue
  ^\s+\d+\.\d+\.\d+\.\d+\s+${IP_ROUTE},*\s*$$
  ^\s+\d+\.\d+\.\d+\.\d+\s+${IP_ROUTE},* -> Continue
  ^\s+(?:\d+\.\d+\.\d+\.\d+\s+){2}${IP_ROUTE},*\s*$$
  ^\s+(?:\d+\.\d+\.\d+\.\d+\s+){2}${IP_ROUTE},* -> Continue
  ^\s+(?:\d+\.\d+\.\d+\.\d+\s+){3}${IP_ROUTE},*\s*$$
  ^\s+(?:\d+\.\d+\.\d+\.\d+\s+){3}${IP_ROUTE},* -> Continue
  ^\s+(?:\d+\.\d+\.\d+\.\d+\s+){4}${IP_ROUTE},*\s*$$
  ^\s+Status:
  ^\s+Path\s+option:
  ^\s+Pre-setup\s+Path
  ^\s+Actual\s+Bandwidth
  ^\s+Hot-standby
  ^\s+protect\s+option
  ^\s+PCE
  ^\s+Active-MPLS-binding-SID
  ^\s+Config\s+Parameters:
  ^\s+Resv-Style
  ^\s+Metric\s+Type
  ^\s+Hop\s+Prior
  ^\s+Hot\s+Hop\s+Limit
  ^\s+Record-Route
  ^\s+(Facility|Detour)\s+Fast-reroute
  ^\s+Protect\s+(Coexist|Nest)
  ^\s+Main\s+LSP\s+Fast-reroute
  ^\s+Bandwidth
  ^\s+Hot-standby-lsp
  ^\s+E2E
  ^\s+BFD
  ^\s+Policy\s+Class
  ^\s+Track\s+Name 
  ^\s+Auto-reoptimize
  ^\s+Reference
  ^\s+Tunnel-Status
  ^\s+CBS
  ^\s+(Main|HSB|FRR)\s+affinity:
  ^\s+Exclude-any
  ^\s+Include-any
  ^\s+Include-all
  ^\s+AutoRoute
  ^\s+AUTO-BW
  ^\s+Forwarding-adjacency
  ^\s+Co-routed\s+Bidirect
  ^\s+Associated\s+Bidirect
  ^\s+Rate-limit
  ^\s+Crankback
  ^\s+Soft\s+Preemption
  ^\s+Addresses\s+of\s+preempting\s+links
  ^\s+Graceful\s+shutdown
  ^\s+Without-CSPF
  ^\s+Ultralimit
  ^\s+Advertise
  ^\s+LSP\s+recoverd
  ^\s+RSVP\s+(Signalling|Path|Resv)
  ^\s+(Exclude|Record)\s+Route: -> IgnoreRouting
  ^\s
  ^\s+History: -> Logs
  ^\s*$$
  ^. -> Error

IgnoreRouting
  ^\s+(Exclude|Record)\s+Route:
  ^\s+(\d+\.){3}\d+(?:\s*\(\d+\))?
  ^\s+(F|T)spec -> Start
  ^\s*$$
  ^. -> Error

Logs
  # Log data is unpredictable, and we can be confident important data has been captured by this point
  ^Name: -> Continue.Record
  ^Name:\s+${NAME}\s*$$ -> Start
`
