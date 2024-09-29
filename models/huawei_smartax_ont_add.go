package models

type HuaweiSmartaxOntAdd struct {
	Onts	string	`json:"ONTS"`
	Success	string	`json:"SUCCESS"`
	Port_id	string	`json:"PORT_ID"`
	Ont_id	string	`json:"ONT_ID"`
}

var HuaweiSmartaxOntAdd_Template = `Value ONTS (\d+)
Value SUCCESS (\d+)
Value PORT_ID (\d+)
Value ONT_ID (\d+)

Start
  ^\s*Number\s*of\s*ONTs\s*that\s*can\s*be\s*added:\s*${ONTS},\s*success:\s*${SUCCESS} -> OntDetails

OntDetails
  ^\s*PortID\s*:${PORT_ID},\s*ONTID\s*:${ONT_ID} -> Record
  ^\s*$$
  ^. -> Error
`