package cisco_ios

type ShowModuleOnlineDiag struct {
	Module      string `json:"MODULE"`
	Online_diag string `json:"ONLINE_DIAG"`
}

var ShowModuleOnlineDiag_Template string = `Value Key MODULE (\d+)
Value ONLINE_DIAG (\w+)

Start
  ^Mod\s+Online\s+Diag\s+Status -> Diag

Diag
  ^-+
  ^\s*${MODULE}\s+${ONLINE_DIAG} -> Record
  ^Mod\s+Ports\s+Card\s+Type\s+Model\s+Serial -> End
  ^(?:Mod|\s+M)\s+MAC\s+addresses\s+Hw\s+Fw\s+Sw\s+Status -> End
  ^Mod\s+Sub-Module\s+Model\s+Serial\s+Hw\s+Status -> End
  ^Mod\s+Redundancy\s+Role\s+Operating\s+Redundancy\s+Mode\s+Configured\s+Redundancy\s+Mode -> End
  ^Mod\s+Redundancy\s+Role\s+Operating\s+Mode\s+Configured\s+Mode\s+Redundancy\s+Status -> End
  ^Mod\s+Online\s+Diag\s+Status -> End
  ^\s*$$
  ^. -> Error

End`
