package models

type CiscoAsaShowModuleStatus struct {
	Module	string	`json:"MODULE"`
	Status	string	`json:"STATUS"`
}

var CiscoAsaShowModuleStatus_Template = `Value Key MODULE (\d+|sfr|ips|cxsc)
Value STATUS (\S.+?)

Start
  ^\s*Mod\s+Status\s+Data\s+Plane\s+Status\s+Compatibility -> ModStatus

ModStatus
  ^-+
  ^\s*${MODULE}\s+${STATUS}\s*Not Applicable\s*$$ -> Record
  ^\s+${MODULE}\s+${STATUS}\s+(Up|Down) -> Record
  ^\s*Mod\s+License\s+Name\s+License\s+Status\s+Time\s+Remaining -> End
  ^\s*$$
  ^. -> Error ModStatus

End
`