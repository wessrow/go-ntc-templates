package cisco_wlc

type SshShowClientDetail struct {
	Ip_address                      string `json:"IP_ADDRESS"`
	Ap_name                         string `json:"AP_NAME"`
	Client_state                    string `json:"CLIENT_STATE"`
	Wireless_lan_profile_name       string `json:"WIRELESS_LAN_PROFILE_NAME"`
	Connected_for_measurement       string `json:"CONNECTED_FOR_MEASUREMENT"`
	Interface                       string `json:"INTERFACE"`
	Radio_signal_strength_indicator string `json:"RADIO_SIGNAL_STRENGTH_INDICATOR"`
	Signal_to_noise_ratio           string `json:"SIGNAL_TO_NOISE_RATIO"`
	Vlan_id                         string `json:"VLAN_ID"`
	Client_username                 string `json:"CLIENT_USERNAME"`
	Wireless_lan_network_name       string `json:"WIRELESS_LAN_NETWORK_NAME"`
	Connected_for                   string `json:"CONNECTED_FOR"`
}

var SshShowClientDetail_Template string = `Value CLIENT_USERNAME (.+?)
Value AP_NAME (.+?)
Value CLIENT_STATE (\S+)
Value WIRELESS_LAN_NETWORK_NAME (.+?)
Value WIRELESS_LAN_PROFILE_NAME (.+?)
Value CONNECTED_FOR (\d+)
Value CONNECTED_FOR_MEASUREMENT (\S+)
Value IP_ADDRESS (\d.*)
Value INTERFACE (.*)
Value RADIO_SIGNAL_STRENGTH_INDICATOR (\S+)
Value SIGNAL_TO_NOISE_RATIO (\d+)
Value VLAN_ID (.*)


Start
  ^Client\s+Username\s*\.+\s*${CLIENT_USERNAME}\s*$$
  ^AP\s+Name\s*\.+\s*${AP_NAME}\s*$$
  ^Client\s+State\s*\.+\s*${CLIENT_STATE}\s*$$
  ^Wireless\s+LAN\s+Network\s+Name\s+\(\S+?\)\s*\.+\s*${WIRELESS_LAN_NETWORK_NAME}\s*$$
  ^Wireless\s+LAN\s+Profile\s+Name\s*\.+\s*${WIRELESS_LAN_PROFILE_NAME}\s*$$
  ^Connected\s+For\s*\.+\s*${CONNECTED_FOR}\s+${CONNECTED_FOR_MEASUREMENT}\s*$$
  ^IP\s+Address\s*\.+\s*${IP_ADDRESS}\s*$$
  ^\s+Radio\s+Signal\s+Strength\s+Indicator\s*\.+\s*${RADIO_SIGNAL_STRENGTH_INDICATOR}\s+dBm\s*$$
  ^Interface\s*\.+\s*${INTERFACE}\s*$$
  ^\s+Signal\s+to\s+Noise\s+Ratio\s*\.+\s*${SIGNAL_TO_NOISE_RATIO}\s+dB\s*$$
  ^VLAN\s*\.+\s*${VLAN_ID}\s*$$
  ^.+\.+
  ^\s+AP.+\(.+\)
  ^Client\s+Capabilities:
  ^Client\s+Wifi\s+Direct Capabilities:
  ^Fast\s+BSS\s+Transition\s+Details:
  ^Client\s+Statistics:
  ^Client\s+Rate\s+Limiting\s+Statistics:
  ^Nearby\s+AP\s+Statistics:
  ^DNS\s+Server\s+details:
  ^Assisted\s+Roaming\s+Prediction\s+List\s+details:
  ^\s*Client\s+Dhcp\s+Required:
  ^Allowed\s+\(URL\)IP\s+Addresses
  ^-+
  ^\s*$$
  ^. -> Error
`
