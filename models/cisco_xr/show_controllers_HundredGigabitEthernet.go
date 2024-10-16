package cisco_xr

type ShowControllersHundredGigabitEthernet struct {
	Ops_state         string   `json:"OPS_STATE"`
	Rx_value_mw       []string `json:"RX_VALUE_MW"`
	Laser_bias        []string `json:"LASER_BIAS"`
	Media_type        string   `json:"MEDIA_TYPE"`
	Lane              []string `json:"LANE"`
	Wavelength        []string `json:"WAVELENGTH"`
	Interface         string   `json:"INTERFACE"`
	Admin_state       string   `json:"ADMIN_STATE"`
	Led_state         string   `json:"LED_STATE"`
	Temperature_value string   `json:"TEMPERATURE_VALUE"`
	Voltage_value     string   `json:"VOLTAGE_VALUE"`
	Rx_value_dbm      []string `json:"RX_VALUE_DBM"`
	Tx_value_dbm      []string `json:"TX_VALUE_DBM"`
	Tx_value_mw       []string `json:"TX_VALUE_MW"`
}

var ShowControllersHundredGigabitEthernet_Template string = `Value Required INTERFACE (\S+)
Value ADMIN_STATE (\S+)
Value OPS_STATE (\S+)
Value LED_STATE (\S+)
Value MEDIA_TYPE (.+)
Value TEMPERATURE_VALUE (\S+)
Value VOLTAGE_VALUE (\S+)
Value List LANE (\d+)
Value List WAVELENGTH (\d+|\S+|n/a)
Value List RX_VALUE_DBM ([-\+]?\d+\.\d+|n/a)
Value List RX_VALUE_MW ([-\+]?\d+\.\d+)
Value List TX_VALUE_DBM ([-\+]?\d+\.\d+|n/a)
Value List TX_VALUE_MW ([-\+]?\d+\.\d+)
Value List LASER_BIAS ([-\+]?\d+\.\d+)

Start
  ^$$
  ^\s+$$
  ^\w+\s+\w+\s+\d{1,2}\s+\d{2}:\d{2}:\d{2}.\d{3}\s+\w+\s*$$
  ^Operational\s+data\s+for\s+interface.*$$ -> Continue.Record
  ^Operational\s+data\s+for\s+interface\s+${INTERFACE}:\s*$$
  ^State:\s*$$
  ^\s+Administrative\s+state:\s+${ADMIN_STATE}\s*$$
  ^\s+Operational\s+state:\s+${OPS_STATE}\s*$$
  ^\s+LED\s+state:\s+${LED_STATE}.*$$
  ^Phy:\s*$$
  ^\s+Media\s+type:\s+${MEDIA_TYPE}\s*$$
  ^\s+Optics:\s*$$
  ^\s+Vendor:.*$$
  ^\s+Part\s+number:.*$$
  ^\s+Serial\s+number:.*$$
  ^\s+Wavelength:.*$$
  ^\s+Digital\s+Optical\s+Monitoring:\s*$$
  ^\s+Transceiver\s+Temp:\s+${TEMPERATURE_VALUE}\s+C\s*$$
  ^\s+Transceiver\s+Voltage:\s+${VOLTAGE_VALUE}\s+V\s*$$
  ^\s+Alarms\s+key:.*$$
  ^\s+(\(\w+\).*(,|))+$$
  ^\s+Wavelength\s+Tx\s+Power\s+Rx\s+Power\s+Laser\s+Bias\s*$$
  ^\s+Lane\s+\(\w+\).+$$
  # NEED TO ACCOUNT FOR ALARM STATUS IN OPTICAL READINGS
  ^\s+${LANE}\s+${WAVELENGTH}\s+${TX_VALUE_DBM}\s+${TX_VALUE_MW}\s+${RX_VALUE_DBM}\s+${RX_VALUE_MW}\s+${LASER_BIAS}\s*$$
  ^\s+DOM\s+alarms:\s*$$
  # NEED TO ADD SUPPORT FOR ALARM MESSAGES
  ^\s+No\s+alarms\s*$$
  ^\s+Alarm\s+Alarm.*$$
  ^\s+Thresholds.*$$
  ^(\s+-+)+\s*$$
  ^\s+Transceiver\s+Temp.+$$
  ^\s+Transceiver\s+Voltage.+$$
  ^\s+Laser\s+Bias.+$$
  ^\s+(Transmit|Receive)\s+Power.+$$
  ^\s+Statistics:\s*$$
  ^\s+FEC:\s*$$
  ^\s+Corrected.+$$
  ^\s+Uncorrected.+$$
  ^MAC.+:\s*$$
  ^\s+Operational\s+address:\s+([a-f0-9]{4}(\.|))+\s*$$
  ^\s+Burnt-in\s+address:\s+([a-f0-9]{4}(\.|))+\s*$$
  ^Autonegotiation\s+\w+.\s*$$
  ^Operational\s+values:\s*$$
  ^\s+Speed:\s+\w+\s*$$
  ^\s+Duplex:\s+\w+.*$$
  ^\s+Flowcontrol:\s+.+?\s*$$
  ^\s+Loopback:\s+.+?\s*$$
  ^\s+MTU:\s+\d+\s*$$
  ^\s+MRU:\s+\d+\s*$$
  ^\s+Forward\s+error\s+correction:\s+\w+\s*$$
  ^.* -> Error "LINE NOT FOUND"
`
