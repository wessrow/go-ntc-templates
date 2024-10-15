package cisco_asa

type ShowModuleStatus struct {
	Status string `json:"STATUS"`
	Module string `json:"MODULE"`
}

var ShowModuleStatus_Template string = `Value Key MODULE (\d+|sfr|ips|cxsc)
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

End`
