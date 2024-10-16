package cisco_s300

type ShowSystem struct {
	Mac_address string   `json:"MAC_ADDRESS"`
	Oid         string   `json:"OID"`
	Fans_status string   `json:"FANS_STATUS"`
	Description string   `json:"DESCRIPTION"`
	Up_time     string   `json:"UP_TIME"`
	Contact     string   `json:"CONTACT"`
	Hostname    string   `json:"HOSTNAME"`
	Location    []string `json:"LOCATION"`
}

var ShowSystem_Template string = `Value DESCRIPTION (.*)
Value UP_TIME (.*)
Value CONTACT (.*)
Value HOSTNAME (.*)
Value List LOCATION (.*)
Value MAC_ADDRESS (.*)
Value OID (.*)
Value FANS_STATUS (.*)

Start
  ^\s*System\s+Description:\s*${DESCRIPTION}\s*$$
  ^\s*System\s+Up\s+Time\s+\(days,hour:min:sec\):\s*${UP_TIME}\s*$$
  ^\s*System\s+Contact:\s*${CONTACT}\s*$$
  ^\s*System\s+Name:\s*${HOSTNAME}\s*$$
  ^\s*System\s+MAC\s+Address:\s*${MAC_ADDRESS}\s*$$
  ^\s*System\s+Object\s+ID:\s*${OID}\s*$$
  ^\s*Fans\s+Status:\s*${FANS_STATUS}\s*$$
  ^\s*System\s+Location:\s*${LOCATION}\s*$$
  ^\s*$$
  ^\s*${LOCATION}\s*$$
  ^. -> Error
`
