package oneaccess_oneos 

type ShowSoftwareImage struct {
	Bank	string	`json:"BANK"`
	Software	string	`json:"SOFTWARE"`
	Created	string	`json:"CREATED"`
	Checksum	string	`json:"CHECKSUM"`
}

var ShowSoftwareImage_Template = `Value Required,Key BANK (Active|Alternate)
Value SOFTWARE (\S+)
Value CREATED (.*)
Value CHECKSUM (\S+)

Start
  ^-+\s+\S+\s+bank -> Continue.Record
  ^-+\s${BANK}\sbank\s-+
  ^Software\sversion\s+:\s${SOFTWARE}
  ^Creation\sdate\s+:\s${CREATED}
  ^Header\schecksum\s+:\s${CHECKSUM}
  ^Installation\sstatus.*
  ^\s*$$
  ^. -> Error
`