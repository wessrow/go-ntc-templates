package cisco_nxos

type ShowModule struct {
	Module string `json:"MODULE"`
	Ports  string `json:"PORTS"`
	Type   string `json:"TYPE"`
	Model  string `json:"MODEL"`
	Status string `json:"STATUS"`
}

var ShowModule_Template string = `Value MODULE (\d+)
Value PORTS (\d+)
Value TYPE (\S+(\s+\S+)+)
Value MODEL (\S+)
Value STATUS (ok|active \*|ha-standby|powered-dn|powered-up)

Start
  ^Xbar\s+Ports\s+Module-Type\s+Model\s+Status -> Fail
  ^${MODULE}\s+${PORTS}\s+${TYPE}\s+${MODEL}\s+${STATUS} -> Record
  ^${MODULE}\s+${PORTS}\s+${TYPE}\s+${STATUS} -> Record

Fail
  ^.* -> NoRecord
`
