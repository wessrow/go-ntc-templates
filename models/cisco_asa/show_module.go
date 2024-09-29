package cisco_asa 

type ShowModule struct {
	Module	string	`json:"MODULE"`
	Cardtype	string	`json:"CARDTYPE"`
	Model	string	`json:"MODEL"`
	Serial	string	`json:"SERIAL"`
}

var ShowModule_Template = `Value Key MODULE (\d+|sfr|ips|cxsc)
Value CARDTYPE (\S.+?)
Value MODEL (\S+)
Value SERIAL (\w+)

Start
  ^Mod\s+Card\s+Type\s+Model\s+Serial\s+No\. -> Status
  ^. -> Error

Status
  ^-+
  ^\s*${MODULE}\s+${CARDTYPE}(\s+${MODEL})?(\s+${SERIAL})?\s*$$ -> Record
  ^\s*Mod\s+MAC\s+Address\s+Range\s+Hw\s+Version\s+Fw\s+Version\s+Sw\s+Version -> End
  ^\s*$$
  ^. -> Error Status

End
`