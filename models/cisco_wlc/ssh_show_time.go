package cisco_wlc

type SshShowTime struct {
	Ntp_msg_auth_status  string `json:"NTP_MSG_AUTH_STATUS"`
	Time                 string `json:"TIME"`
	Ntp_key_index        string `json:"NTP_KEY_INDEX"`
	Ntp_server           string `json:"NTP_SERVER"`
	Ntp_version          string `json:"NTP_VERSION"`
	Ntp_polling_interval string `json:"NTP_POLLING_INTERVAL"`
	Index                string `json:"INDEX"`
	Status               string `json:"STATUS"`
	Tz_delta             string `json:"TZ_DELTA"`
	Timezone             string `json:"TIMEZONE"`
	Tz_name              string `json:"TZ_NAME"`
}

var SshShowTime_Template string = `Value Filldown TIME (.+)
Value Filldown TZ_DELTA (\S+)
Value Filldown TIMEZONE (.+)
Value Filldown TZ_NAME (.+)
Value Filldown NTP_VERSION (\d)
Value Filldown NTP_POLLING_INTERVAL (\d+)
Value Required INDEX (\d+)
Value NTP_KEY_INDEX (\d+)
Value NTP_SERVER (\S+)
Value STATUS (\S+\s?\S+)
Value NTP_MSG_AUTH_STATUS (\S+\s?\S+)

Start
  ^\s*Time\s*\.+\s+${TIME}\s*$$
  ^\s*Timezone [dD]elta\s*\.+\s+${TZ_DELTA}\s*$$
  ^\s*Timezone [lL]ocation\s*\.+\s+\(${TIMEZONE}\)\s+${TZ_NAME}\s*$$
  ^\s*NTP [sS]ervers\s*$$
  ^\s*NTP [vV]ersion\s*\.+\s+${NTP_VERSION}\s*$$
  ^\s*NTP [pP]olling [iI]nterval\s*\.+\s+${NTP_POLLING_INTERVAL}\s*$$
  #
  # NTP Server Table
  ^\s+Index\s+NTP Key Index\s+NTP Server\s+Status\s+NTP Msg Auth Status\s*$$ -> NTP_Servers
  ^\s*$$

NTP_Servers
  ^\s+${INDEX}\s+${NTP_KEY_INDEX}\s+${NTP_SERVER}\s+${STATUS}\s+${NTP_MSG_AUTH_STATUS}\s*$$ -> Record
  ^\s+[-\s]+$$
  ^\s*$$
  ^. -> Error
`
