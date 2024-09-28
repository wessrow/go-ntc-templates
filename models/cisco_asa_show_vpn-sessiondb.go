package models

type CiscoAsaShowVpnSessiondb struct {
	Vpn_session_name	[]string	`json:"VPN_SESSION_NAME"`
	Vpn_session_active	[]string	`json:"VPN_SESSION_ACTIVE"`
	Vpn_session_cumulative	[]string	`json:"VPN_SESSION_CUMULATIVE"`
	Vpn_session_peak_concurrent	[]string	`json:"VPN_SESSION_PEAK_CONCURRENT"`
	Vpn_session_inactive	[]string	`json:"VPN_SESSION_INACTIVE"`
	Total_active_and_inactive	string	`json:"TOTAL_ACTIVE_AND_INACTIVE"`
	Total_cumulative	string	`json:"TOTAL_CUMULATIVE"`
	Device_total_vpn_capacity	string	`json:"DEVICE_TOTAL_VPN_CAPACITY"`
	Device_load_percent	string	`json:"DEVICE_LOAD_PERCENT"`
	Tunnels_summary_name	[]string	`json:"TUNNELS_SUMMARY_NAME"`
	Tunnels_summary_active	[]string	`json:"TUNNELS_SUMMARY_ACTIVE"`
	Tunnels_summary_cumulative	[]string	`json:"TUNNELS_SUMMARY_CUMULATIVE"`
	Tunnels_summary_peak_concurrent	[]string	`json:"TUNNELS_SUMMARY_PEAK_CONCURRENT"`
	Totals_active	string	`json:"TOTALS_ACTIVE"`
	Totals_cumulative	string	`json:"TOTALS_CUMULATIVE"`
}