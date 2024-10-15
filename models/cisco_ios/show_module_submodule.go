package cisco_ios

type ShowModuleSubmodule struct {
	Submodule_serial string `json:"SUBMODULE_SERIAL"`
	Submodule_hw     string `json:"SUBMODULE_HW"`
	Submodule_status string `json:"SUBMODULE_STATUS"`
	Module           string `json:"MODULE"`
	Submodule        string `json:"SUBMODULE"`
	Submodule_model  string `json:"SUBMODULE_MODEL"`
}

var ShowModuleSubmodule_Template string = `Value Key MODULE (\d+)
Value SUBMODULE (.+?)
Value SUBMODULE_MODEL (\S+)
Value SUBMODULE_SERIAL (\S+)
Value SUBMODULE_HW (\S+)
Value SUBMODULE_STATUS (\w+)


Start
  ^\s*Mod\s+Sub-Module\s+Model\s+Serial\s+Hw\s+Status\s*$$ -> SubModule

SubModule
  ^-+
  ^\s*${MODULE}\s+${SUBMODULE}\s+${SUBMODULE_MODEL}\s+${SUBMODULE_SERIAL}\s+${SUBMODULE_HW}\s+${SUBMODULE_STATUS}\s*$$ -> Record
  ^Mod\s+Ports\s+Card\s+Type\s+Model\s+Serial -> End
  ^(?:Mod|\s+M)\s+MAC\s+addresses\s+Hw\s+Fw\s+Sw\s+Status -> End
  ^Mod\s+Online\s+Diag\s+Status -> End
  ^Mod\s+Redundancy\s+Role\s+Operating\s+Redundancy\s+Mode\s+Configured\s+Redundancy\s+Mode -> End
  ^Mod\s+Redundancy\s+Role\s+Operating\s+Mode\s+Configured\s+Mode\s+Redundancy\s+Status -> End
  ^\s*$$
  ^. -> Error

End`
