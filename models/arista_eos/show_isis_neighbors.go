package arista_eos 

type ShowIsisNeighbors struct {
	Instance	string	`json:"INSTANCE"`
	Vrf	string	`json:"VRF"`
	System_id	string	`json:"SYSTEM_ID"`
	Type	string	`json:"TYPE"`
	Interface	string	`json:"INTERFACE"`
	Snpa	string	`json:"SNPA"`
	State	string	`json:"STATE"`
	Hold_time	string	`json:"HOLD_TIME"`
	Circuit_id	string	`json:"CIRCUIT_ID"`
}

var ShowIsisNeighbors_Template = `Value INSTANCE (\S+)
Value VRF (\S+)
Value SYSTEM_ID ((\d+.\d+.\d+)|\S+)
Value TYPE ((L\d)+)
Value INTERFACE (\S+)
Value SNPA (\S+)
Value STATE (\S+)
Value HOLD_TIME (\d+)
Value CIRCUIT_ID ((\d+.\d+.\d+.\S+)|\S+)

Start
  ^${INSTANCE}\s+${VRF}\s+${SYSTEM_ID}\s+${TYPE}\s+${INTERFACE}\s+${SNPA}\s+${STATE}\s+${HOLD_TIME}\s+${CIRCUIT_ID} -> Record
`