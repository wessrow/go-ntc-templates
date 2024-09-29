package hp_procurve 

type ShowTrunks struct {
	Local_port	string	`json:"LOCAL_PORT"`
	Int_name	string	`json:"INT_NAME"`
	Int_type	string	`json:"INT_TYPE"`
	Trunk	string	`json:"TRUNK"`
	Trunk_type	string	`json:"TRUNK_TYPE"`
}

var ShowTrunks_Template = `Value Required LOCAL_PORT (\S+)
Value INT_NAME (.*?)
Value INT_TYPE (.*?)
Value TRUNK (\S+)
Value TRUNK_TYPE (\S+)

Start
  ^\s*Load\sBalancing\sMethod.*$$
  ^\s+Port\s+|\s+Name\s+Type\s+|\s+Group\s+Type.*$$
  ^\s+-+\s+\+\s+\-+.* -> TRUNKS
  ^\s*$$
  ^. -> Error

TRUNKS
  ^\s+${LOCAL_PORT}\s+\|\s${INT_NAME}\s+ ${INT_TYPE}\s+ \|\s${TRUNK}\s+${TRUNK_TYPE}.*$$ -> Record
  ^\S+\#\s*$$
  ^\s*$$
  ^. -> Error
`