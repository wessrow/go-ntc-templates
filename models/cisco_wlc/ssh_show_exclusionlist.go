package cisco_wlc 

type SshShowExclusionlist struct {
	Mac_address	string	`json:"MAC_ADDRESS"`
	Comment	string	`json:"COMMENT"`
}

var SshShowExclusionlist_Template = `Value MAC_ADDRESS (([\da-fA-F]{2}[-:]){5}[\da-fA-F]{2})
Value COMMENT (.*)

Start
  ^Manually\s*Disabled\s*Clients -> ManuallyDisabled
  ^No\s*dynamically
  ^\s*$$
  ^. -> Error

ManuallyDisabled
  ^---
  ^MAC\s+Address\s+Description
  ^${MAC_ADDRESS}\s+${COMMENT} -> Record
  ^\s*$$ -> Start`