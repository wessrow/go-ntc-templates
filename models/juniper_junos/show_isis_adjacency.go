package juniper_junos

type ShowIsisAdjacency struct {
	Interface string `json:"INTERFACE"`
	System_id string `json:"SYSTEM_ID"`
	Type      string `json:"TYPE"`
	State     string `json:"STATE"`
	Hold_time string `json:"HOLD_TIME"`
	Snpa      string `json:"SNPA"`
}

var ShowIsisAdjacency_Template string = `Value INTERFACE (\S+)
Value SYSTEM_ID ((\d+.\d+.\d+)|\S+)
Value TYPE (\d)
Value STATE (\S+)
Value HOLD_TIME (\d+)
Value SNPA ((\S+)?)

Start
  ^${INTERFACE}\s+${SYSTEM_ID}\s+${TYPE}\s+${STATE}\s+${HOLD_TIME}(\s+)?${SNPA} -> Record
  ^\s*$$
  ^{master:\d+}
`
