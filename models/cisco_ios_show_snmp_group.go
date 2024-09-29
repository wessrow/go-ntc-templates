package models

type CiscoIosShowSnmpGroup struct {
	Groupname	string	`json:"GROUPNAME"`
	Security_model	string	`json:"SECURITY_MODEL"`
	Contextname	string	`json:"CONTEXTNAME"`
	Storage_type	string	`json:"STORAGE_TYPE"`
	Readview	string	`json:"READVIEW"`
	Writeview	string	`json:"WRITEVIEW"`
	Notifyview	string	`json:"NOTIFYVIEW"`
	Row_status	string	`json:"ROW_STATUS"`
	Access_list	string	`json:"ACCESS_LIST"`
}

var CiscoIosShowSnmpGroup_Template = `Value Required GROUPNAME (\S+)
Value SECURITY_MODEL ([\s\S]+?)
Value CONTEXTNAME ([\s\S]+?)
Value STORAGE_TYPE (\S+)
Value READVIEW ([\s\S]+?)
Value WRITEVIEW ([\s\S]+?)
Value NOTIFYVIEW ([\s\S]+?)
Value ROW_STATUS ([\s\S]+?)
Value ACCESS_LIST (\S+)

Start
  ^\s*groupname:\s*\S+\s+security\s+model:\s*[\s\S]+\s*$$ -> Continue.Record
  ^\s*groupname:\s*${GROUPNAME}\s+security\s+model:\s*${SECURITY_MODEL}\s*$$
  ^\s*contextname:\s+${CONTEXTNAME}\s+storage-type:\s+${STORAGE_TYPE}\s*$$
  ^\s*readview\s*:\s+${READVIEW}\s+writeview:\s+${WRITEVIEW}\s*$$
  ^\s*notifyview:\s+${NOTIFYVIEW}\s*$$
  ^\s*row\s+status:\s+${ROW_STATUS}(?:\s+access-list:\s+${ACCESS_LIST})*\s*$$
  ^\s*$$
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is
  ^. -> Error

`