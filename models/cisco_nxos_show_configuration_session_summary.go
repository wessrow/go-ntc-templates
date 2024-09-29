package models

type CiscoNxosShowConfigurationSessionSummary struct {
	Session_name	string	`json:"SESSION_NAME"`
	Session_owner	string	`json:"SESSION_OWNER"`
	Creation_time	string	`json:"CREATION_TIME"`
}

var CiscoNxosShowConfigurationSessionSummary_Template = `Value SESSION_NAME (\S+)
Value SESSION_OWNER (\S+)
Value CREATION_TIME (\S+\s\S+\s\S+\s\S+\s\S+)

Start
  ^Name -> Session

Session
  ^Number
  ^${SESSION_NAME}\s+${SESSION_OWNER}\s+${CREATION_TIME} -> Record


`