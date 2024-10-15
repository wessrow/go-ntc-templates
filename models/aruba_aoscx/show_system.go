package aruba_aoscx

type ShowSystem struct {
	Uptime_days    string `json:"UPTIME_DAYS"`
	Hostname       string `json:"HOSTNAME"`
	Location       string `json:"LOCATION"`
	Vendor         string `json:"VENDOR"`
	Uptime_weeks   string `json:"UPTIME_WEEKS"`
	Version        string `json:"VERSION"`
	Time_zone      string `json:"TIME_ZONE"`
	Uptime_hours   string `json:"UPTIME_HOURS"`
	Uptime_minutes string `json:"UPTIME_MINUTES"`
	Contact        string `json:"CONTACT"`
	Product        string `json:"PRODUCT"`
	Serial         string `json:"SERIAL"`
	Base_mac       string `json:"BASE_MAC"`
}

var ShowSystem_Template string = `Value Required HOSTNAME (\S+)
Value CONTACT (\S+)
Value LOCATION (\S+)
Value VENDOR (\S+)
Value PRODUCT (.+)
Value SERIAL (\S+)
Value BASE_MAC (\S+)
Value VERSION (\S+)
Value TIME_ZONE (\S+)
Value UPTIME_WEEKS (\d+)
Value UPTIME_DAYS (\d+)
Value UPTIME_HOURS (\d+)
Value UPTIME_MINUTES (\d+)

Start
  ^Hostname\s+:\s+${HOSTNAME}
  ^System Description\s+:\s*\S+
  ^System Contact\s+:\s+${CONTACT}		
  ^System Contact\s+:
  ^System Location\s+:\s+${LOCATION}
  ^System Location\s+:
  ^Vendor\s+:\s+${VENDOR}
  ^Product Name\s+:\s+${PRODUCT}
  ^Chassis Serial Nbr\s+:\s+${SERIAL}
  ^Base MAC Address\s+:\s+${BASE_MAC}	
  ^ArubaOS-CX Version\s+:\s+${VERSION}
  ^Time Zone\s+:\s+${TIME_ZONE}
  ^Up Time\s+:(\s+${UPTIME_WEEKS}\s+weeks?,?)?(\s+${UPTIME_DAYS}\s+days?,?)?(\s+${UPTIME_HOURS}\s+hours?,?)?(\s+(${UPTIME_MINUTES}|under\s+a)\s+minutes?)? -> Record
  ^CPU
  ^Memory
  ^-+$$
  ^\*+$$
  ^\s*$$
  ^. -> Error "LINE NOT FOUND"
`
