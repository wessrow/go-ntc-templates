package cisco_nxos

type ShowIpOspfDatabase struct {
	Tag        string `json:"TAG"`
	Router_id  string `json:"ROUTER_ID"`
	Process_id string `json:"PROCESS_ID"`
	Area       string `json:"AREA"`
	Link_id    string `json:"LINK_ID"`
	Adv_router string `json:"ADV_ROUTER"`
	Age        string `json:"AGE"`
	Link_count string `json:"LINK_COUNT"`
}

var ShowIpOspfDatabase_Template string = `Value Filldown ROUTER_ID (\d+\.\d+\.\d+\.\d+)
Value Filldown PROCESS_ID (\w+)
Value Filldown AREA (\d+\.\d+\.\d+\.\d+|\d+)
Value LINK_ID (\d+\.\d+\.\d+\.\d+)
Value ADV_ROUTER (\d+\.\d+\.\d+\.\d+)
Value AGE (\d+)
Value LINK_COUNT (\d+)
Value TAG (\d+)

Start
  ^.*\(${ROUTER_ID}\)\s+\(.+\s${PROCESS_ID}\)$$
  ^.*\(Area\s${AREA}\)$$
  ^${LINK_ID}\s+${ADV_ROUTER}\s+${AGE}\s+\S+\s+\S+\s+${LINK_COUNT}$$ -> Record
  ^${LINK_ID}\s+${ADV_ROUTER}\s+${AGE}\s+\S+\s+\S+$$ -> Record
  ^\s+Type-5 AS External Link States -> Tag

Tag
  ^Link ID\s+ADV Router\s+Age\s+Seq#\s+Checksum\s+Tag -> Next.Clearall
  ^${LINK_ID}\s+${ADV_ROUTER}\s+${AGE}\s+\S+\s+\S+\s+${TAG} -> Next
  ^\s -> Start

EOF`
