package cisco_asa

type ShowObjectGroupNetwork struct {
	Network     string `json:"NETWORK"`
	Netmask     string `json:"NETMASK"`
	Grp_object  string `json:"GRP_OBJECT"`
	Name        string `json:"NAME"`
	Description string `json:"DESCRIPTION"`
	Type        string `json:"TYPE"`
	Host        string `json:"HOST"`
	Net_object  string `json:"NET_OBJECT"`
}

var ShowObjectGroupNetwork_Template string = `Value Filldown NAME (\S+)
Value Filldown DESCRIPTION (.+)
Value TYPE ([newrkguphsobjct]+)
Value HOST (\d+.\d+.\d+.\d+)
Value NET_OBJECT (\S+)
Value NETWORK (\d+.\d+.\d+.\d+)
Value NETMASK (\d+.\d+.\d+.\d+)
Value GRP_OBJECT (\S+)

Start
  ^object-group -> Continue.Clearall
  ^object-group\s+network\s+${NAME}\s*
  ^\s+description:?\s+${DESCRIPTION}\s*
  ^\s+network-object\s+${TYPE}\s+${HOST}\s* -> Record
  ^\s+network-object\s+${TYPE}\s+${NET_OBJECT}\s* -> Record
  ^\s+${TYPE}-object\s+${NETWORK}\s+${NETMASK}\s* -> Record
  ^\s+${TYPE}-object\s+${GRP_OBJECT}\s* -> Record

EOF
`
