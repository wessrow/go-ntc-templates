package cisco_nxos 

type ShowFabricpathRoute struct {
	Ftag	string	`json:"FTAG"`
	Switch_id	string	`json:"SWITCH_ID"`
	Subswitch_id	string	`json:"SUBSWITCH_ID"`
	Ports	[]string	`json:"PORTS"`
	Distance	string	`json:"DISTANCE"`
	Metric	string	`json:"METRIC"`
}

var ShowFabricpathRoute_Template = `Value Filldown FTAG (\d+)
Value Filldown SWITCH_ID (\d+)
Value SUBSWITCH_ID (\d+)
Value Required,List PORTS (\S+)
Value DISTANCE (\d+)
Value METRIC (\d+)

Start
  ^FabricPath\s+Unicast\s+Route\s+Table\s+for\s+Topology-Default -> Start_record

Start_record
  ^\d+/\d+/\d+, -> Continue.Record
  ^${FTAG}/${SWITCH_ID}/${SUBSWITCH_ID},\s+number\s+of\s+next-hops
  ^.+via\s+${PORTS}\s*,\s\[${DISTANCE}/${METRIC}\],\s
`