package alcatel_sros 

type ShowRouterIsisInterface struct {
	Interface	string	`json:"INTERFACE"`
	Level	string	`json:"LEVEL"`
	Circid	string	`json:"CIRCID"`
	Oper_state	string	`json:"OPER_STATE"`
	Metric_l1	string	`json:"METRIC_L1"`
	Metric_l2	string	`json:"METRIC_L2"`
}

var ShowRouterIsisInterface_Template = `Value INTERFACE (\S+)
Value LEVEL (\S+)
Value CIRCID (\d+)
Value OPER_STATE (Up|Down)
Value METRIC_L1 (\d+)
Value METRIC_L2 (\d+)

Start
  ^----------- -> Interface

Interface
  ^${INTERFACE}\s+${LEVEL}\s+${CIRCID}\s+${OPER_STATE}\s+${METRIC_L1}/${METRIC_L2} -> Record
  ^\s*$$
  ^-----------
  ^===========
  ^Interfaces
  ^. -> Error
`