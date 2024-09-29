package mikrotik_routeros 

type SnmpCommunityPrintWithoutPaging struct {
	Index	string	`json:"INDEX"`
	Flags	string	`json:"FLAGS"`
	Name	string	`json:"NAME"`
	Addresses	string	`json:"ADDRESSES"`
	Security	string	`json:"SECURITY"`
	Read_access	string	`json:"READ_ACCESS"`
	Write_access	string	`json:"WRITE_ACCESS"`
}

var SnmpCommunityPrintWithoutPaging_Template = `Value Key INDEX (\d+)
Value FLAGS ([X*]+)
Value NAME (\S+?)
Value ADDRESSES (\S+?)
Value SECURITY (authorized|none|private)
Value READ_ACCESS (yes|no)
Value WRITE_ACCESS (yes|no)

Start
  ^\s#\s+NAME\s+ADDRESSES\s+SECURITY\s+READ-ACCESS\s+WRITE-ACCESS\s*$$ -> SNMPCommunitiesTable

SNMPCommunitiesTable
  ^\s?${INDEX}\s+(${FLAGS})?\s+${NAME}\s+${ADDRESSES}\s+${SECURITY}\s+${READ_ACCESS}\s+${WRITE_ACCESS}\s*$$ -> Record
  ^\s*(?:\d{2}:){2}\d{2}\s+echo:\s*.*$$ -> Next
  ^. -> Error
`