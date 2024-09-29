package paloalto_panos 

type ShowCounterGlobal struct {
	Name	string	`json:"NAME"`
	Value	string	`json:"VALUE"`
	Rate	string	`json:"RATE"`
	Severity	string	`json:"SEVERITY"`
	Category	string	`json:"CATEGORY"`
	Aspect	string	`json:"ASPECT"`
	Description	string	`json:"DESCRIPTION"`
}

var ShowCounterGlobal_Template = `Value NAME (\S+)
Value VALUE (\d+)
Value RATE (\d+)
Value SEVERITY (\S+)
Value CATEGORY (\S+)
Value ASPECT (\S+)
Value DESCRIPTION ([\S+\s+]+)

Start
  ^${NAME}\s+${VALUE}\s+${RATE}\s+${SEVERITY}\s+${CATEGORY}\s+${ASPECT}\s+${DESCRIPTION} -> Record
`