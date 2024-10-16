package huawei_vrp

type DisplayServiceSetIdId struct {
	Forward_mode          string `json:"FORWARD_MODE"`
	Id                    string `json:"ID"`
	User_isolate          string `json:"USER_ISOLATE"`
	Max_nb_user           string `json:"MAX_NB_USER"`
	Wlan_bss_interface    string `json:"WLAN_BSS_INTERFACE"`
	Igmp_mode             string `json:"IGMP_MODE"`
	Traffic_profile_name  string `json:"TRAFFIC_PROFILE_NAME"`
	Security_profile_name string `json:"SECURITY_PROFILE_NAME"`
	Name                  string `json:"NAME"`
	Ssid                  string `json:"SSID"`
	Hide_ssid             string `json:"HIDE_SSID"`
	Type                  string `json:"TYPE"`
	Assoc_timeout         string `json:"ASSOC_TIMEOUT"`
}

var DisplayServiceSetIdId_Template string = `Value ID (\d+)
Value NAME (\S+)
Value SSID (\S+)
Value HIDE_SSID (disable|enable)
Value USER_ISOLATE (disable|enable)
Value TYPE (service|ap-management)
Value MAX_NB_USER (\d+)
Value ASSOC_TIMEOUT (\d+)
Value TRAFFIC_PROFILE_NAME (\S+)
Value SECURITY_PROFILE_NAME (\S+)
Value WLAN_BSS_INTERFACE (\S+)
Value IGMP_MODE (snooping|off)
Value FORWARD_MODE (direct-forward|gre-tunnel)

Start
  ^\s*-+ -> Next
  ^\s*(Service-(S|s)et\s+)?ID\s+:\s${ID}\s*$$
  ^\s*(Service-(S|s)et\s+)?(N|n)ame\s+:\s${NAME}\s*$$
  ^\s*SSID\s+:\s${SSID}\s*$$
  ^\s*Hide\sSSID\s+:\s${HIDE_SSID}\s*$$
  ^\s*User\sisolate\s+:\s${USER_ISOLATE}\s*$$
  ^\s*Type\s+:\s${TYPE}\s*$$
  ^\s*Maximum\snumber\sof\suser\s+:\s${MAX_NB_USER}\s*$$
  ^\s*Association\stimeout\(min\)\s+:\s${ASSOC_TIMEOUT}\s*$$
  ^\s*Traffic\sprofile\sname\s+:\s${TRAFFIC_PROFILE_NAME}\s*$$
  ^\s*Security\sprofile\sname\s+:\s${SECURITY_PROFILE_NAME}\s*$$
  ^\s*Wlan-bss\sinterface\s+:\s${WLAN_BSS_INTERFACE}\s*$$
  ^\s*Igmp\smode\s+:\s${IGMP_MODE}\s*$$
  ^\s*Forward\smode\s+:\s${FORWARD_MODE}\s*$$
  ^.*$$ -> Error
`
