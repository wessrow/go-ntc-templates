package oneaccess_oneos

type ShowTransceiverEquipment struct {
	Temp_high_warning         string   `json:"TEMP_HIGH_WARNING"`
	Tx_power_high_alarm       string   `json:"TX_POWER_HIGH_ALARM"`
	Connector                 string   `json:"CONNECTOR"`
	Linklen_625_mm            string   `json:"LINKLEN_625_MM"`
	Temp_high_alarm           string   `json:"TEMP_HIGH_ALARM"`
	Linklen_sm                string   `json:"LINKLEN_SM"`
	Linklen_copper            string   `json:"LINKLEN_COPPER"`
	Temp_low_alarm            string   `json:"TEMP_LOW_ALARM"`
	Rx_power_high_alarm       string   `json:"RX_POWER_HIGH_ALARM"`
	Vendor_name               string   `json:"VENDOR_NAME"`
	Tcvr_ether_compliance     string   `json:"TCVR_ETHER_COMPLIANCE"`
	Tcvr_fiber_linklen        string   `json:"TCVR_FIBER_LINKLEN"`
	Wavelength                string   `json:"WAVELENGTH"`
	Encoding                  string   `json:"ENCODING"`
	Vendor_partner            string   `json:"VENDOR_PARTNER"`
	Tcvr_fiber_tech           string   `json:"TCVR_FIBER_TECH"`
	Tcvr_fiber_tx_media       string   `json:"TCVR_FIBER_TX_MEDIA"`
	Voltage_low_alarm         string   `json:"VOLTAGE_LOW_ALARM"`
	Voltage_low_warning       string   `json:"VOLTAGE_LOW_WARNING"`
	Bias_low_warning          string   `json:"BIAS_LOW_WARNING"`
	Measured_supply_voltage   string   `json:"MEASURED_SUPPLY_VOLTAGE"`
	Measured_rx_input_power   string   `json:"MEASURED_RX_INPUT_POWER"`
	Physical_device           string   `json:"PHYSICAL_DEVICE"`
	Tcvr_10g_ether_compliance string   `json:"TCVR_10G_ETHER_COMPLIANCE"`
	Tcvr_fiber_speed          string   `json:"TCVR_FIBER_SPEED"`
	Voltage_high_alarm        string   `json:"VOLTAGE_HIGH_ALARM"`
	Bias_high_alarm           string   `json:"BIAS_HIGH_ALARM"`
	Rx_power_high_warning     string   `json:"RX_POWER_HIGH_WARNING"`
	Vendor_revision           string   `json:"VENDOR_REVISION"`
	Vendor_datecode           string   `json:"VENDOR_DATECODE"`
	Temp_low_warning          string   `json:"TEMP_LOW_WARNING"`
	Voltage_high_warning      string   `json:"VOLTAGE_HIGH_WARNING"`
	Bias_low_alarm            string   `json:"BIAS_LOW_ALARM"`
	Tx_power_low_warning      string   `json:"TX_POWER_LOW_WARNING"`
	Rx_power_low_alarm        string   `json:"RX_POWER_LOW_ALARM"`
	Measured_tx_output_power  string   `json:"MEASURED_TX_OUTPUT_POWER"`
	Port                      string   `json:"PORT"`
	Vendor_id                 string   `json:"VENDOR_ID"`
	Linklen_50_mm             string   `json:"LINKLEN_50_MM"`
	Min_bit_rate              string   `json:"MIN_BIT_RATE"`
	Bias_high_warning         string   `json:"BIAS_HIGH_WARNING"`
	Rx_power_low_warning      string   `json:"RX_POWER_LOW_WARNING"`
	Measured_tx_bias_current  string   `json:"MEASURED_TX_BIAS_CURRENT"`
	Nominal_bitrate           string   `json:"NOMINAL_BITRATE"`
	Options                   []string `json:"OPTIONS"`
	Max_bit_rate              string   `json:"MAX_BIT_RATE"`
	Tx_power_low_alarm        string   `json:"TX_POWER_LOW_ALARM"`
	Tx_power_high_warning     string   `json:"TX_POWER_HIGH_WARNING"`
	Measured_module_temp      string   `json:"MEASURED_MODULE_TEMP"`
}

var ShowTransceiverEquipment_Template string = `Value Filldown PORT (\S+)
Value Required PHYSICAL_DEVICE (.*\S)
Value CONNECTOR (\S+)
Value VENDOR_NAME (.*\S)
Value VENDOR_ID (.*\S)
Value VENDOR_PARTNER (.*\S)
Value VENDOR_REVISION (.*\S)
Value VENDOR_DATECODE (.*\S)
Value TCVR_ETHER_COMPLIANCE (.*\S)
Value TCVR_10G_ETHER_COMPLIANCE (.*\S)
Value TCVR_FIBER_LINKLEN (.*\S)
Value TCVR_FIBER_TECH (.*\S)
Value TCVR_FIBER_TX_MEDIA (.*\S)
Value TCVR_FIBER_SPEED (.*\S)
Value WAVELENGTH (.*\S)
Value ENCODING (.*\S)
Value NOMINAL_BITRATE (.*\S)
Value LINKLEN_SM (.*\S)
Value LINKLEN_50_MM (.*\S)
Value LINKLEN_625_MM (.*\S)
Value LINKLEN_COPPER (.*\S)
Value List OPTIONS (\S.*\S)
Value MAX_BIT_RATE (\S.*\S)
Value MIN_BIT_RATE (\S.*\S)
Value TEMP_HIGH_ALARM (\S.*\S)
Value TEMP_LOW_ALARM (\S.*\S)
Value TEMP_HIGH_WARNING (\S.*\S)
Value TEMP_LOW_WARNING (\S.*\S)
Value VOLTAGE_HIGH_ALARM (\S.*\S)
Value VOLTAGE_LOW_ALARM (\S.*\S)
Value VOLTAGE_HIGH_WARNING (\S.*\S)
Value VOLTAGE_LOW_WARNING (\S.*\S)
Value BIAS_HIGH_ALARM (\S.*\S)
Value BIAS_LOW_ALARM (\S.*\S)
Value BIAS_HIGH_WARNING (\S.*\S)
Value BIAS_LOW_WARNING (\S.*\S)
Value TX_POWER_HIGH_ALARM (\S.*\S)
Value TX_POWER_LOW_ALARM (\S.*\S)
Value TX_POWER_HIGH_WARNING (\S.*\S)
Value TX_POWER_LOW_WARNING (\S.*\S)
Value RX_POWER_HIGH_ALARM (\S.*\S)
Value RX_POWER_LOW_ALARM (\S.*\S)
Value RX_POWER_HIGH_WARNING (\S.*\S)
Value RX_POWER_LOW_WARNING (\S.*\S)
Value MEASURED_MODULE_TEMP (\S.*\S)
Value MEASURED_SUPPLY_VOLTAGE (\S.*\S)
Value MEASURED_TX_BIAS_CURRENT (\S.*\S)
Value MEASURED_TX_OUTPUT_POWER (\S.*\S)
Value MEASURED_RX_INPUT_POWER (\S.*\S)

Start
  ^SFP -> Continue.Record
  ^SFP\s${PORT}:
  ^\s*No\s+(Inventory\s+data\s+available|SFP\s+module\s+present)
  ^\s*Physical\sdevice\s+=\s+${PHYSICAL_DEVICE}
  ^\s*connector\s+=\s+${CONNECTOR}
  ^\s*vendor: -> VENDOR
  ^\s*wavelength\s+=\s+${WAVELENGTH}
  ^\s*encoding\s+=\s+${ENCODING}
  ^\s*nominalBitRate\s+=\s+${NOMINAL_BITRATE}
  ^\s*Link\slength\sin: -> LINK_LENGTH
  ^\s*[mM]in\sbit\srate\s+=\s+${MIN_BIT_RATE}
  ^\s*Diagnostics\scalibration\sis\sinternal -> DIAGNOSTICS
  ^\s*SFP\s+module\s+inventory\s+information\s*:
  ^\s*$$
  ^. -> Error

VENDOR
  ^\s*name\s+=\s+${VENDOR_NAME}
  ^\s*id\s+=\s+${VENDOR_ID}
  ^\s*partNumber\s+=\s+${VENDOR_PARTNER}
  ^\s*revision\s+=\s+${VENDOR_REVISION}
  ^\s*dateCode\s+=\s+${VENDOR_DATECODE}
  ^\s*transceiver: -> TRANSCEIVER
  ^\s+serialNumber\s+=
  ^\s*$$
  ^. -> Error

TRANSCEIVER
  ^\s*ethernetComplianceCode\s+=\s+${TCVR_ETHER_COMPLIANCE}
  ^\s*10G\sethernetComplianceCode\s+=\s+${TCVR_10G_ETHER_COMPLIANCE}
  ^\s*fiberLinkLen\s+=\s+${TCVR_FIBER_LINKLEN}
  ^\s*fiberTech\s+=\s+${TCVR_FIBER_TECH}
  ^\s*fiberTxMedia\s+=\s+${TCVR_FIBER_TX_MEDIA}
  ^\s*fiberSpeed\s+=\s+${TCVR_FIBER_SPEED} -> Continue
  ^\s*fiberSpeed -> Start
  ^\s+sonet
  ^\s*$$
  ^. -> Error

LINK_LENGTH
  ^\s*single\sfiber\smode\s+=\s+${LINKLEN_SM}
  ^\s*50u\smulti-mode\sfiber\s+=\s+${LINKLEN_50_MM}
  ^\s*62\.5u\smulti-mode\sfiber\s+=\s+${LINKLEN_625_MM}
  ^\s*copper\scable\s+=\s+${LINKLEN_COPPER} -> Continue
  ^\s*copper\scable -> OPTIONS
  ^\s*$$
  ^. -> Error

OPTIONS
  ^\s*Max\sbit\srate\s+=\s+${MAX_BIT_RATE} -> Start
  ^\s*options\s+=\s${OPTIONS}
  ^\s+${OPTIONS}
  ^\s*$$
  ^. -> Error

DIAGNOSTICS
  ^\s*Temp\sHigh\sAlarm\s+=\s+${TEMP_HIGH_ALARM}
  ^\s*Temp\sLow\sAlarm\s+=\s+${TEMP_LOW_ALARM}
  ^\s*Temp\sHigh\sWarning\s+=\s+${TEMP_HIGH_WARNING}
  ^\s*Temp\sLow\sWarning\s+=\s+${TEMP_LOW_WARNING}
  ^\s*Voltage\sHigh\sAlarm\s+=\s+${VOLTAGE_HIGH_ALARM}
  ^\s*Voltage\sLow\sAlarm\s+=\s+${VOLTAGE_LOW_ALARM}
  ^\s*Voltage\sHigh\sWarning\s+=\s+${VOLTAGE_HIGH_WARNING}
  ^\s*Voltage\sLow\sWarning\s+=\s+${VOLTAGE_LOW_WARNING}
  ^\s*Bias\sHigh\sAlarm\s+=\s+${BIAS_HIGH_ALARM}
  ^\s*Bias\sLow\sAlarm\s+=\s+${BIAS_LOW_ALARM}
  ^\s*Bias\sHigh\sWarning\s+=\s+${BIAS_HIGH_WARNING}
  ^\s*Bias\sLow\sWarning\s+=\s+${BIAS_LOW_WARNING}
  ^\s*TX\sPower\sHigh\sAlarm\s+=\s+${TX_POWER_HIGH_ALARM}
  ^\s*TX\sPower\sLow\sAlarm\s+=\s+${TX_POWER_LOW_ALARM}
  ^\s*TX\sPower\sHigh\sWarning\s+=\s+${TX_POWER_HIGH_WARNING}
  ^\s*TX\sPower\sLow\sWarning\s+=\s+${TX_POWER_LOW_WARNING}
  ^\s*RX\sPower\sHigh\sAlarm\s+=\s+${RX_POWER_HIGH_ALARM}
  ^\s*RX\sPower\sLow\sAlarm\s+=\s+${RX_POWER_LOW_ALARM}
  ^\s*RX\sPower\sHigh\sWarning\s+=\s+${RX_POWER_HIGH_WARNING}
  ^\s*RX\sPower\sLow\sWarning\s+=\s+${RX_POWER_LOW_WARNING}
  ^\s*Internally\sMeasured\sModule\sTemperature\s+=\s+${MEASURED_MODULE_TEMP}
  ^\s*Internally\sMeasured\sSupply\sVoltage\s+=\s+${MEASURED_SUPPLY_VOLTAGE}
  ^\s*Internally\sMeasured\sTx\sBias\sCurrent\s+=\s+${MEASURED_TX_BIAS_CURRENT}
  ^\s*Measured\sTx\sOutput\sPower\s+=\s+${MEASURED_TX_OUTPUT_POWER}
  ^\s*Measured\sRx\sInput\sPower\s+=\s+${MEASURED_RX_INPUT_POWER}
  ^\s*Optional\s+Status
  ^\s*Alarm
  ^SFP -> Continue.Record
  ^SFP\s+${PORT}: -> Start
  ^\s*$$
  ^. -> Error
`
