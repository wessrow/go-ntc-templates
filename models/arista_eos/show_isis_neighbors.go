package arista_eos

type ShowIsisNeighbors struct {
	System_id  string `json:"SYSTEM_ID"`
	Snpa       string `json:"SNPA"`
	Hold_time  string `json:"HOLD_TIME"`
	Instance   string `json:"INSTANCE"`
	Vrf        string `json:"VRF"`
	Type       string `json:"TYPE"`
	Interface  string `json:"INTERFACE"`
	State      string `json:"STATE"`
	Circuit_id string `json:"CIRCUIT_ID"`
}

var ShowIsisNeighbors_Template string = `Value INSTANCE (\S+)
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
