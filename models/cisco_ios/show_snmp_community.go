package cisco_ios

type ShowSnmpCommunity struct {
	Index              string `json:"INDEX"`
	Security_name      string `json:"SECURITY_NAME"`
	Storage_type       string `json:"STORAGE_TYPE"`
	Access_list_number string `json:"ACCESS_LIST_NUMBER"`
	Name               string `json:"NAME"`
}

var ShowSnmpCommunity_Template string = `Value NAME (\S+)
Value INDEX (\S+)
Value SECURITY_NAME (\S+)
Value STORAGE_TYPE (\S+)
Value ACCESS_LIST_NUMBER (\S+)

Start
  ^Community\s+name: -> Continue.Record
  ^Community\s+name:\s+${NAME}\s*$$
  ^Community\sIndex:\s+${INDEX}
  ^Community\sSecurityName:\s+${SECURITY_NAME}
  ^storage-type:\s${STORAGE_TYPE}\s+active(\s+access-list:\s${ACCESS_LIST_NUMBER})?
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is
  ^\s*$$
  ^. -> Error
`
