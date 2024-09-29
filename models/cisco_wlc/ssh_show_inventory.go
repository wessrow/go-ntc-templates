package cisco_wlc 

type SshShowInventory struct {
	Bia	string	`json:"BIA"`
	Max_ap_num	string	`json:"MAX_AP_NUM"`
	Name	string	`json:"NAME"`
	Description	string	`json:"DESCRIPTION"`
	Pid	string	`json:"PID"`
	Vid	string	`json:"VID"`
	Sn	string	`json:"SN"`
}

var SshShowInventory_Template = `Value BIA ([\da-zA-Z:]{17})
Value MAX_AP_NUM (\d+)
Value NAME (.*)
Value DESCRIPTION (.*)
Value PID ([\w\-]+)
Value VID (\w+)
Value SN (\w+)

Start
  ^Burned-in\s+MAC\s+Address\.+\s+${BIA}
  ^Power\sSupply\s\d\.*\s(Present|Absent)
  ^Maximum\s+number\s+of\s+APs\s+supported\.+\s+${MAX_AP_NUM}
  ^NAME:\s+"${NAME}"\s+,\s*DESCR:\s+"${DESCRIPTION}"
  ^PID:\s+${PID},\s*VID:\s+${VID},\s*SN:\s+${SN} -> Record
  ^\s*$$
  ^. -> Error
`