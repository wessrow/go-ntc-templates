package cisco_xr

type ShowIsisNeighbors struct {
	Type      string `json:"TYPE"`
	Ietf_nsf  string `json:"IETF_NSF"`
	System_id string `json:"SYSTEM_ID"`
	Interface string `json:"INTERFACE"`
	Snpa      string `json:"SNPA"`
	State     string `json:"STATE"`
	Hold_time string `json:"HOLD_TIME"`
}

var ShowIsisNeighbors_Template string = `Value SYSTEM_ID ((\d+.\d+.\d+)|\S+)
Value INTERFACE (\S+)
Value SNPA ((\d+.\d+.\d+)|\S+)
Value STATE (\S+)
Value HOLD_TIME (\d+)
Value TYPE ((L\d)+)
Value IETF_NSF (\S+)

Start
  ^${SYSTEM_ID}\s+${INTERFACE}\s+${SNPA}\s+${STATE}\s+${HOLD_TIME}\s+${TYPE}\s+${IETF_NSF} -> Record
`
