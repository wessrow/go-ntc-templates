package oneaccess_oneos 

type ShowIpAsPathAccessList struct {
	Name	string	`json:"NAME"`
	Entries	[]string	`json:"ENTRIES"`
}

var ShowIpAsPathAccessList_Template = `Value Required NAME (\S+)
Value Required,List ENTRIES (.*)

Start
  ^AS\spath\saccess\slist\s${NAME}$$ -> ASPATH
  ^. -> Error

ASPATH
  ^[\s\t]+${ENTRIES}.*
  ^\s+$$ -> Record Start
  ^AS\spath -> Continue.Record
  ^AS\spath\saccess\slist\s${NAME}$$ -> ASPATH
  ^. -> Error
`