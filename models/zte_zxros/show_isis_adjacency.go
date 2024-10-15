package zte_zxros

type ShowIsisAdjacency struct {
	Hold_time string `json:"HOLD_TIME"`
	Mt        string `json:"MT"`
	Context   string `json:"CONTEXT"`
	Interface string `json:"INTERFACE"`
	System_id string `json:"SYSTEM_ID"`
	Pri       string `json:"PRI"`
	Nsf       string `json:"NSF"`
	Af        string `json:"AF"`
	State     string `json:"STATE"`
	Lev       string `json:"LEV"`
	Snpa      string `json:"SNPA"`
}

var ShowIsisAdjacency_Template string = `Value Filldown CONTEXT (\d+)
Value Required INTERFACE (\S+)
Value SYSTEM_ID ((\d+.\d+.\d+)|\S+)
Value STATE (\S+)
Value LEV (\S+)
Value HOLD_TIME (\d+)
Value SNPA ((\d+.\d+.\d+)|\S+)
Value PRI (\S+)
Value MT (\S+)
Value NSF (\S+)
Value AF (\S+)

Start
  ^Process\s+ID:\s+${CONTEXT}\s*$$
  ^Interface\s+System\s+id\s+State\s+Lev\s+Holds\s+SNPA\S*\s+Pri\s+MT\s+NSF\s+AF\s*$$
  ^${INTERFACE}\s+${SYSTEM_ID}\s+${STATE}\s+${LEV}\s+${HOLD_TIME}\s+${SNPA}\s+${PRI}\s+${MT}?\s+${NSF}\s+${AF}\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`
