package huawei_vrp 

type DisplaySnmpAgentCommunityRead struct {
	Name	string	`json:"NAME"`
	Group	string	`json:"GROUP"`
	Storage_type	string	`json:"STORAGE_TYPE"`
	View	string	`json:"VIEW"`
	Acl	string	`json:"ACL"`
}

var DisplaySnmpAgentCommunityRead_Template = `Value NAME (\S+)
Value GROUP (\S+)
Value STORAGE_TYPE (volatile|nonVolatile|permanent|readOnly|other)
Value VIEW (\S+)
Value ACL (\d+)

Start
  ^\s*Community\sname:.*$$ -> Continue.Record
  ^\s*Community\sname:\s+${NAME}\s*$$
  ^\s*Group\sname:\s+${GROUP}\s*$$
  ^\s*Storage(\s|-)type:\s+${STORAGE_TYPE}\s*$$
  ^\s*View\sname:\s+${VIEW}\s*$$
  ^\s*Acl:\s+${ACL}\s*$$
  ^\s*Total\snumber\sis\s\d+\s*$$
  ^\s*$$
  ^. -> Error
`