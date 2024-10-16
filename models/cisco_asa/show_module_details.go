package cisco_asa

type ShowModuleDetails struct {
	Mod_fw      string `json:"MOD_FW"`
	Mod_sw      string `json:"MOD_SW"`
	Module      string `json:"MODULE"`
	Mac_address string `json:"MAC_ADDRESS"`
	Mod_hw      string `json:"MOD_HW"`
}

var ShowModuleDetails_Template string = `Value Key MODULE (\d+|sfr|ips|cxsc)
Value MAC_ADDRESS (\S+\s+to\s+\S+)
Value MOD_HW (\S+)
Value MOD_FW (\S+(\s+\[\w+\])?)
Value MOD_SW (\S+)

Start
  ^\s*Mod\s+MAC\s+Address\s+Range\s+Hw\s+Version\s+Fw\s+Version\s+Sw\s+Version -> ModDetails

ModDetails
  ^-+
  ^\s*${MODULE}\s+${MAC_ADDRESS}\s+${MOD_HW}\s+${MOD_FW}(\s+${MOD_SW})?\s*$$ -> Record
#  ^\s*${MODULE}\s+${MAC_ADDRESS}\s+${MOD_HW}\s+${MOD_FW}\s*$$ -> Record
  ^\s*Mod\s+SSM\s+Application\s+Name\s+Status\s+SSM\s+Application\s+Version -> End
  ^\s*$$
  ^. -> Error ModDetails

End`
