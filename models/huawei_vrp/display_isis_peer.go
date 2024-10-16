package huawei_vrp

type DisplayIsisPeer struct {
	Type       string   `json:"TYPE"`
	Pri        string   `json:"PRI"`
	System_id  []string `json:"SYSTEM_ID"`
	Interface  string   `json:"INTERFACE"`
	Circuit_id string   `json:"CIRCUIT_ID"`
	State      string   `json:"STATE"`
	Hold_time  string   `json:"HOLD_TIME"`
}

var DisplayIsisPeer_Template string = `Value List SYSTEM_ID (\S+)
Value INTERFACE (\S+)
Value CIRCUIT_ID (\d+)
Value STATE (\S+)
Value HOLD_TIME (\S+)
Value TYPE ((L\d)+)
Value PRI (\S+)

Start
  ^\s+Peer\s+information\s+for\s+\S+
  ^\s+System\s+Id\s+Interface\s+Circuit\s+Id\s+State\s+HoldTime\s+Type\s+PRI
  ^-+
  ^\S+\s+\S+\s+\d+ -> Continue.Record
  ^${SYSTEM_ID}\s+${INTERFACE}\s+${CIRCUIT_ID}\s+${STATE}\s+${HOLD_TIME}\s+${TYPE}\s+${PRI}\s*$$
  ^${SYSTEM_ID}\s*$$
  ^Total\s+Peer\S+:\s+\d+
  ^\s*$$
  ^. -> Error
`
