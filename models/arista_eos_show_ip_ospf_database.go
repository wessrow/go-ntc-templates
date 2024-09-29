package models

type AristaEosShowIpOspfDatabase struct {
	Router_id	string	`json:"ROUTER_ID"`
	Process_id	string	`json:"PROCESS_ID"`
	Vrf	string	`json:"VRF"`
	Area	string	`json:"AREA"`
	Link_id	string	`json:"LINK_ID"`
	Adv_router	string	`json:"ADV_ROUTER"`
	Age	string	`json:"AGE"`
	Link_count	string	`json:"LINK_COUNT"`
	Tag	string	`json:"TAG"`
}

var AristaEosShowIpOspfDatabase_Template = `Value Filldown ROUTER_ID (\d+\.\d+\.\d+\.\d+)
Value Filldown PROCESS_ID (\d+)
Value Filldown VRF (\S+)
Value Filldown AREA (\d+\.\d+\.\d+\.\d+|\d+)
Value LINK_ID (\d+\.\d+\.\d+\.\d+)
Value ADV_ROUTER (\d+\.\d+\.\d+\.\d+)
Value AGE (\d+:\d+:\d+|\d+)
Value LINK_COUNT (\d+)
Value TAG (\d+)

Start
  ^.*\(${ROUTER_ID}\).*?${PROCESS_ID}.*VRF\s+${VRF}\)
  ^.*\(Area\s+${AREA}\)
  ^${LINK_ID}\s+${ADV_ROUTER}\s+${AGE}\s+\S+\s+\S+\s+${LINK_COUNT} -> Record
  ^${LINK_ID}\s+${ADV_ROUTER}\s+${AGE}\s+\S+\s+\S+ -> Record
  ^Link ID\s+ADV Router\s+Age\s+Seq#\s+Checksum\s+Tag -> Tag

Tag
  ^Link ID\s+ADV Router\s+Age\s+Seq#\s+Checksum\s+Tag -> Next.Clearall
  ^${LINK_ID}\s+${ADV_ROUTER}\s+${AGE}\s+\S+\s+\S+\s+${TAG} -> Next
  ^\s -> Start

EOF

`