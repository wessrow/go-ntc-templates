package cisco_ios

type ShowVtpStatus struct {
	Local_updater_iface  string `json:"LOCAL_UPDATER_IFACE"`
	Max_vlans            string `json:"MAX_VLANS"`
	Revision_number      string `json:"REVISION_NUMBER"`
	Device_id            string `json:"DEVICE_ID"`
	Last_modified_date   string `json:"LAST_MODIFIED_DATE"`
	Pruning              string `json:"PRUNING"`
	Traps                string `json:"TRAPS"`
	Last_modified_server string `json:"LAST_MODIFIED_SERVER"`
	Local_updater_addr   string `json:"LOCAL_UPDATER_ADDR"`
	Mode                 string `json:"MODE"`
	Existing_vlan_count  string `json:"EXISTING_VLAN_COUNT"`
	Version              string `json:"VERSION"`
	Domain               string `json:"DOMAIN"`
}

var ShowVtpStatus_Template string = `Value VERSION (\d)
Value DOMAIN (\S+)
Value PRUNING (\S+)
Value TRAPS (\S+)
Value DEVICE_ID (\S+)
Value LAST_MODIFIED_SERVER (\d+.\d+.\d+.\d+)
Value LAST_MODIFIED_DATE (\d+-\d+-\d+\s\d+:\d+:\d+)
Value LOCAL_UPDATER_ADDR (\d+.\d+.\d+.\d+)
Value LOCAL_UPDATER_IFACE (\S+)
Value MODE (\w+)
Value MAX_VLANS (\d+)
Value EXISTING_VLAN_COUNT (\d+)
Value REVISION_NUMBER (\d+)

Start
  ^VTP\s+[v,V]ersion\s+[r,R]unning\s+:\s${VERSION}
  ^VTP\s+[d,D]omain\s+[n,N]ame\s+:\s${DOMAIN}
  ^VTP\s+[p,P]runing\s+[m,M]ode\s+:\s${PRUNING}
  ^VTP\s+[t,T]raps\s+[g,G]eneration\s+:\s${TRAPS}
  ^[d,D]evice\s+ID\s+:\s${DEVICE_ID}
  ^\w+\s\w+\s\w+\s\w+\s${LAST_MODIFIED_SERVER}\s\w+\s${LAST_MODIFIED_DATE}
  ^\w+\s\w+\s\w+\s\w+\s${LOCAL_UPDATER_ADDR}\s\w+\s\w+\s${LOCAL_UPDATER_IFACE}\s.*
  ^VTP\s[o,O]perating\s[m,M]ode\s+:\s${MODE}
  ^Maximum\sVLANs\s[s,S]upported\s[l,L]ocally\s+:\s${MAX_VLANS}
  ^Number\sof\sexisting\sVLANs\s+:\s${EXISTING_VLAN_COUNT}
  ^Configuration\s[r,R]evision\s+:\s${REVISION_NUMBER} -> Record
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is
`
