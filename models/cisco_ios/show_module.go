package cisco_ios

type ShowModule struct {
	Module        string `json:"MODULE"`
	Switch_number string `json:"SWITCH_NUMBER"`
	Port          string `json:"PORT"`
	Cardtype      string `json:"CARDTYPE"`
	Model         string `json:"MODEL"`
	Serial        string `json:"SERIAL"`
}

var ShowModule_Template string = `Value Key MODULE (\d+)
Value Filldown SWITCH_NUMBER (\d+)
Value PORT (\d+)
Value CARDTYPE (\S.+?)
Value MODEL (\S+)
Value SERIAL (\w+)


Start
  ^Switch Number:? ${SWITCH_NUMBER}
  ^Mod\s+Ports\s+Card\s+Type\s+Model\s+Serial -> Status
  ^Chassis\sType\s?:\s
  ^Power\sconsumed\sby\sbackplane\s:\s
  # #1083 C9200L-24T
  ^Switch\s+Ports\s+Model\s+Serial\sNo\.\s+MAC\saddress\s+Hw\sVer\.\s+Sw\sVer\. -> SwitchStack
  ^. -> Error NoMatchInStart

Status
  ^---+
  ^\s*${MODULE}\s+(${PORT}\s+)?${CARDTYPE}(\s+${MODEL})?(\s+${SERIAL})?\s*$$ -> Record
  ^(?:Mod|\s+M)\s+MAC\s+addresses\s+Hw\s+Fw\s+Sw\s+Status -> End
  ^Mod\s+Sub-Module\s+Model\s+Serial\s+Hw\s+Status -> End
  ^Mod\s+Redundancy\s+Role\s+Operating\s+Redundancy\s+Mode\s+Configured\s+Redundancy\s+Mode -> End
  ^Mod\s+Redundancy\s+Role\s+Operating\s+Mode\s+Configured\s+Mode\s+Redundancy\s+Status -> End
  ^Mod\s+Online\s+Diag\s+Status -> End
  ^\s*$$
  ^. -> Error NoMatchInStatus

SwitchStack
  ^---+
  # #1083 C9200L-24T
  ^\s*${MODULE}\s+${PORT}\s+${MODEL}\s+${SERIAL}\s+ -> Record
  ^\s*$$
  ^. -> Error NoMatchInSwitch

End
`
