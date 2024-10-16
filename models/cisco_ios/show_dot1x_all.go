package cisco_ios

type ShowDot1xAll struct {
	Dot1xversion            string `json:"DOT1XVERSION"`
	Hostmode                string `json:"HOSTMODE"`
	Reauthentication        string `json:"REAUTHENTICATION"`
	Masreq                  string `json:"MASREQ"`
	Critical_recovery_delay string `json:"CRITICAL_RECOVERY_DELAY"`
	Interface               string `json:"INTERFACE"`
	Pae                     string `json:"PAE"`
	Critical_eapol          string `json:"CRITICAL_EAPOL"`
	Quietperiod             string `json:"QUIETPERIOD"`
	Ratelimitperiod         string `json:"RATELIMITPERIOD"`
	Reauthmax               string `json:"REAUTHMAX"`
	Txperiod                string `json:"TXPERIOD"`
	Sysauthcontrol          string `json:"SYSAUTHCONTROL"`
	Portcontrol             string `json:"PORTCONTROL"`
	Controldirection        string `json:"CONTROLDIRECTION"`
	Servertimeout           string `json:"SERVERTIMEOUT"`
	Supptimeout             string `json:"SUPPTIMEOUT"`
	Reauthperiod            string `json:"REAUTHPERIOD"`
}

var ShowDot1xAll_Template string = `Value Filldown SYSAUTHCONTROL (\w+)
Value Filldown DOT1XVERSION (\d)
Value Filldown CRITICAL_RECOVERY_DELAY (\d+)
Value Filldown CRITICAL_EAPOL (\w+)
Value Required INTERFACE (\S+)
Value PAE (\w+)
Value PORTCONTROL (\w+)
Value CONTROLDIRECTION (\w+)
Value HOSTMODE (\w+)
Value REAUTHENTICATION (\w+)
Value QUIETPERIOD (\d+)
Value SERVERTIMEOUT (\d+)
Value SUPPTIMEOUT (\d+)
Value REAUTHPERIOD (\d+)
Value REAUTHMAX (\d+)
Value MASREQ (\d+)
Value TXPERIOD (\d+)
Value RATELIMITPERIOD (\d+)

Start
  ^Sysauthcontrol\s+${SYSAUTHCONTROL}
  ^Dot1x Protocol Version\s+${DOT1XVERSION}
  ^Critical Recovery Delay\s+${CRITICAL_RECOVERY_DELAY}
  ^Critical EAPOL\s+${CRITICAL_EAPOL}
  ^$$ -> Interfaces
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is

Interfaces
  ^Dot1x Info for\s -> Continue.Record
  ^Dot1x Info for\s${INTERFACE}
  ^PAE\s+=\s${PAE}
  ^PortControl\s+=\s${PORTCONTROL}
  ^ControlDirection\s+=\s${CONTROLDIRECTION}
  ^HostMode\s+=\s${HOSTMODE}
  ^ReAuthentication\s+=\s${REAUTHENTICATION}
  ^QuietPeriod\s+=\s${QUIETPERIOD}
  ^ServerTimeout\s+=\s${SERVERTIMEOUT}
  ^SuppTimeout\s+=\s${SUPPTIMEOUT}
  ^ReAuthPeriod\s+=\s${REAUTHPERIOD}
  ^ReAuthMax\s+=\s${REAUTHMAX}
  ^MaxReq\s+=\s${MASREQ}
  ^TxPeriod\s+=\s${TXPERIOD}
  ^RateLimitPeriod\s+=\s+${RATELIMITPERIOD}
  ^-+\s*$$
  ^\s*$$
  ^. -> Error
`
