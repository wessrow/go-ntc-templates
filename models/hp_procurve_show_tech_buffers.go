package models

type HpProcurveShowTechBuffers struct {
	Product	string	`json:"PRODUCT"`
	Name	string	`json:"NAME"`
	Date	string	`json:"DATE"`
	Build	string	`json:"BUILD"`
	Version	string	`json:"VERSION"`
	Cpu	string	`json:"CPU"`
}

var HpProcurveShowTechBuffers_Template = `Value PRODUCT (.+)
Value NAME (.+)
Value DATE (.+)
Value BUILD (.+)
Value VERSION (.+)
Value CPU (\S+)

Start
  ^.*[cC]rash Log File Header -> HEADER

HEADER
  ^\s*Product:\s+${PRODUCT}
  ^\s*Name:\s+${NAME}
  ^\s*Date:\s+${DATE}
  ^\s*Build:\s+${BUILD}
  ^\s*Version:\s+${VERSION}
  ^\s*Directory:.+
  ^\s*CPU:\s+${CPU} -> Record SLOTS

SLOTS
  ^.*[cC]rash Log File Header -> NoRecord
`