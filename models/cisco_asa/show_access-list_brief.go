package cisco_asa

type ShowAccessListBrief struct {
	Acl_name_hash      string `json:"ACL_NAME_HASH"`
	Line_hash          string `json:"LINE_HASH"`
	Group_hash         string `json:"GROUP_HASH"`
	Counter            string `json:"COUNTER"`
	Last_hit           string `json:"LAST_HIT"`
	Acl_name           string `json:"ACL_NAME"`
	Acl_total_elements string `json:"ACL_TOTAL_ELEMENTS"`
}

var ShowAccessListBrief_Template string = `Value Filldown ACL_NAME (\S+)
Value Filldown ACL_TOTAL_ELEMENTS (\d+)
Value Filldown ACL_NAME_HASH (0x\w+)
Value Required LINE_HASH (\w{8})
Value GROUP_HASH (\w{8})
Value COUNTER (\w{8})
Value LAST_HIT (\w{8})

Start
  ^access\-list\s+${ACL_NAME};\s+${ACL_TOTAL_ELEMENTS}\s+elements;\s+name\s+hash:\s+${ACL_NAME_HASH}\s*
  ^${LINE_HASH}\s+${GROUP_HASH}\s+${COUNTER}\s+${LAST_HIT}\s* -> Record
  ^\s*$$
  ^.* -> Error 

EOF
`
