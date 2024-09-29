package cisco_ios 

type ShowIsisNeighbors struct {
	System_id	string	`json:"SYSTEM_ID"`
	Type	string	`json:"TYPE"`
	Interface	string	`json:"INTERFACE"`
	Ip_address	string	`json:"IP_ADDRESS"`
	State	string	`json:"STATE"`
	Hold_time	string	`json:"HOLD_TIME"`
	Circuit_id	string	`json:"CIRCUIT_ID"`
}

var ShowIsisNeighbors_Template = `Value SYSTEM_ID ((\d+.\d+.\d+)|\S+)
Value TYPE (L\d)
Value INTERFACE (\S+)
Value IP_ADDRESS (\d+.\d+.\d+.\d+)
Value STATE (\S+)
Value HOLD_TIME (\d+)
Value CIRCUIT_ID (\S+)

Start
  ^${SYSTEM_ID}\s+${TYPE}\s+${INTERFACE}\s+${IP_ADDRESS}\s+${STATE}\s+${HOLD_TIME}\s+${CIRCUIT_ID} -> Record
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is
`