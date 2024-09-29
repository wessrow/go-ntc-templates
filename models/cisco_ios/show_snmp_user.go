package cisco_ios 

type ShowSnmpUser struct {
	User_name	string	`json:"USER_NAME"`
	Engine_id	string	`json:"ENGINE_ID"`
	Storage_type	string	`json:"STORAGE_TYPE"`
	Access_list	string	`json:"ACCESS_LIST"`
	Authentication_protocol	string	`json:"AUTHENTICATION_PROTOCOL"`
	Privacy_protocol	string	`json:"PRIVACY_PROTOCOL"`
	Group_name	string	`json:"GROUP_NAME"`
}

var ShowSnmpUser_Template = `Value Required USER_NAME (\S+)
Value ENGINE_ID (\S+)
Value STORAGE_TYPE (\S+)
Value ACCESS_LIST (.*)
Value AUTHENTICATION_PROTOCOL (\S+)
Value PRIVACY_PROTOCOL (\S+)
Value GROUP_NAME (\S+)

Start
  ^User\s+name:\s+${USER_NAME}$$
  ^Engine\s+ID:\s+${ENGINE_ID}$$
  ^storage-type:\s${STORAGE_TYPE}\s+active\s+access-list:\s+${ACCESS_LIST}$$
  ^storage-type:\s${STORAGE_TYPE}\s+active$$
  ^Authentication\s+Protocol:\s+${AUTHENTICATION_PROTOCOL}$$
  ^Privacy\s+Protocol:\s+${PRIVACY_PROTOCOL}$$
  ^Group-name:\s+${GROUP_NAME}$$ -> Record
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is
  ^\s*$$
  ^. -> Error
`