package models

type HpComwareDisplayLinkAggregationVerbose struct {
	Interface	string	`json:"INTERFACE"`
	Creation_mode	string	`json:"CREATION_MODE"`
	Aggregation_mode	string	`json:"AGGREGATION_MODE"`
	Loadsharing	string	`json:"LOADSHARING"`
	Local_interfaces	[]string	`json:"LOCAL_INTERFACES"`
	Remote_interfaces	[]string	`json:"REMOTE_INTERFACES"`
}

var HpComwareDisplayLinkAggregationVerbose_Template = `Value Required INTERFACE (\S+)
Value CREATION_MODE (.+)
Value AGGREGATION_MODE (.+)
Value LOADSHARING (.+)
Value List LOCAL_INTERFACES (\S+)
Value List REMOTE_INTERFACES (\S+)

Start
  ^Aggregat\S+\s+Interface:\s+${INTERFACE}
  ^Creation\s+Mode:\s+${CREATION_MODE}
  ^Aggregation\s+Mode:\s+${AGGREGATION_MODE}
  ^Loadsharing\s+Type:\s+${LOADSHARING}
  ^Local: -> LocalPorts
  ^\s*$$ ^. -> Error

LocalPorts
  ^\s+Port
  ^\s*---
  ^\s+${LOCAL_INTERFACES}\s+.*
  # Start of remote block port
  ^Remote: -> RemotePorts
  ^. -> Error
  ^\s*$$ ^. -> Error

RemotePorts
  ^\s+Actor
  ^\s*---
  ^\s+${REMOTE_INTERFACES}\s+.*
  # New item, restart from the beginning
  ^Aggregat -> Continue.Record
  ^Aggregat\S+\s+Interface:\s+${INTERFACE} -> Start
  ^. -> Error
  ^\s*$$ ^. -> Error

`