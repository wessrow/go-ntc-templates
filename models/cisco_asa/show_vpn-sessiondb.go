package cisco_asa

type ShowVpnSessiondb struct {
	Device_total_vpn_capacity       string   `json:"DEVICE_TOTAL_VPN_CAPACITY"`
	Tunnels_summary_active          []string `json:"TUNNELS_SUMMARY_ACTIVE"`
	Tunnels_summary_cumulative      []string `json:"TUNNELS_SUMMARY_CUMULATIVE"`
	Vpn_session_inactive            []string `json:"VPN_SESSION_INACTIVE"`
	Total_cumulative                string   `json:"TOTAL_CUMULATIVE"`
	Device_load_percent             string   `json:"DEVICE_LOAD_PERCENT"`
	Tunnels_summary_peak_concurrent []string `json:"TUNNELS_SUMMARY_PEAK_CONCURRENT"`
	Totals_cumulative               string   `json:"TOTALS_CUMULATIVE"`
	Vpn_session_peak_concurrent     []string `json:"VPN_SESSION_PEAK_CONCURRENT"`
	Vpn_session_cumulative          []string `json:"VPN_SESSION_CUMULATIVE"`
	Tunnels_summary_name            []string `json:"TUNNELS_SUMMARY_NAME"`
	Vpn_session_active              []string `json:"VPN_SESSION_ACTIVE"`
	Total_active_and_inactive       string   `json:"TOTAL_ACTIVE_AND_INACTIVE"`
	Totals_active                   string   `json:"TOTALS_ACTIVE"`
	Vpn_session_name                []string `json:"VPN_SESSION_NAME"`
}

var ShowVpnSessiondb_Template string = `Value List VPN_SESSION_NAME (\S+?\s?\S+)
Value List VPN_SESSION_ACTIVE (\d+)
Value List VPN_SESSION_CUMULATIVE (\d+)
Value List VPN_SESSION_PEAK_CONCURRENT (\d+)
Value List VPN_SESSION_INACTIVE (\d+)
#
Value TOTAL_ACTIVE_AND_INACTIVE (\d+)
Value TOTAL_CUMULATIVE (\d+)
Value DEVICE_TOTAL_VPN_CAPACITY (\d+)
Value DEVICE_LOAD_PERCENT (\d+)
#
Value List TUNNELS_SUMMARY_NAME (\S+?\s?\S+)
Value List TUNNELS_SUMMARY_ACTIVE (\d+)
Value List TUNNELS_SUMMARY_CUMULATIVE (\d+)
Value List TUNNELS_SUMMARY_PEAK_CONCURRENT (\d+)
#
Value TOTALS_ACTIVE (\d+)
Value TOTALS_CUMULATIVE (\d+)

Start
  ^.+#\s+show\s+vpn-sessiondb\s*$$
  ^\s*\-+\s*$$
  ^\s*VPN\s+Session\s+Summary\s*$$ -> VPN_Session_Summary

VPN_Session_Summary
  ^\s*\-+\s*$$
  ^\s*Active\s+:\s+Cumulative\s+:\s+Peak\s+Concur\s+:\s+Inactive\s*$$
  ^\s*${VPN_SESSION_NAME}\s+:\s+${VPN_SESSION_ACTIVE}\s+:\s+${VPN_SESSION_CUMULATIVE}\s+:\s+${VPN_SESSION_PEAK_CONCURRENT}\s*(:\s+${VPN_SESSION_INACTIVE}\s*)?$$
  ^\s*Total\s+Active\s+and\s+Inactive\s+:\s+${TOTAL_ACTIVE_AND_INACTIVE}\s+Total\s+Cumulative\s+:\s+${TOTAL_CUMULATIVE}\s*$$
  ^\s*Device\s+Total\s+VPN\s+Capacity\s+:\s+${DEVICE_TOTAL_VPN_CAPACITY}
  ^\s*Device\s+Load\s+:\s+${DEVICE_LOAD_PERCENT}%\s*$$
  #
  ^\s*Tunnels\s+Summary\s* -> Tunnels_Summary
  #
  ^\s*$$
  ^. -> Error

Tunnels_Summary
  ^\s*\-+\s*$$
  ^\s*Active\s+:\s+Cumulative\s+:\s+Peak\s+Concurrent\s*$$
  ^\s*${TUNNELS_SUMMARY_NAME}\s+:\s+${TUNNELS_SUMMARY_ACTIVE}\s+:\s+${TUNNELS_SUMMARY_CUMULATIVE}\s+:\s+${TUNNELS_SUMMARY_PEAK_CONCURRENT}\s*$$
  ^\s*Totals\s+:\s+${TOTALS_ACTIVE}\s+:\s+${TOTALS_CUMULATIVE}\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`
