package models

type CiscoXrShowControllersHundredgigeAll struct {
	Interface	string	`json:"INTERFACE"`
	Admin_state	string	`json:"ADMIN_STATE"`
	Ops_state	string	`json:"OPS_STATE"`
	Vendor	string	`json:"VENDOR"`
	Part_num	string	`json:"PART_NUM"`
	Serial_num	string	`json:"SERIAL_NUM"`
	Led_state	string	`json:"LED_STATE"`
	Media_type	string	`json:"MEDIA_TYPE"`
	Temperature_value	string	`json:"TEMPERATURE_VALUE"`
	Voltage_value	string	`json:"VOLTAGE_VALUE"`
	Dom_alarms	string	`json:"DOM_ALARMS"`
	Lane	string	`json:"LANE"`
	Wavelength	string	`json:"WAVELENGTH"`
	Rx_value_dbm	string	`json:"RX_VALUE_DBM"`
	Rx_value_mw	string	`json:"RX_VALUE_MW"`
	Tx_value_dbm	string	`json:"TX_VALUE_DBM"`
	Tx_value_mw	string	`json:"TX_VALUE_MW"`
	Laser_bias	string	`json:"LASER_BIAS"`
	Temperature_alarm_high	string	`json:"TEMPERATURE_ALARM_HIGH"`
	Temperature_alarm_low	string	`json:"TEMPERATURE_ALARM_LOW"`
	Temperature_warn_high	string	`json:"TEMPERATURE_WARN_HIGH"`
	Temperature_warn_low	string	`json:"TEMPERATURE_WARN_LOW"`
	Voltage_alarm_high	string	`json:"VOLTAGE_ALARM_HIGH"`
	Voltage_alarm_low	string	`json:"VOLTAGE_ALARM_LOW"`
	Voltage_warn_high	string	`json:"VOLTAGE_WARN_HIGH"`
	Voltage_warn_low	string	`json:"VOLTAGE_WARN_LOW"`
	Amps_alarm_high	string	`json:"AMPS_ALARM_HIGH"`
	Amps_alarm_low	string	`json:"AMPS_ALARM_LOW"`
	Amps_warn_high	string	`json:"AMPS_WARN_HIGH"`
	Amps_warn_low	string	`json:"AMPS_WARN_LOW"`
	Rx_alarm_mw_high	string	`json:"RX_ALARM_MW_HIGH"`
	Rx_alarm_mw_low	string	`json:"RX_ALARM_MW_LOW"`
	Rx_warn_mw_high	string	`json:"RX_WARN_MW_HIGH"`
	Rx_warn_mw_low	string	`json:"RX_WARN_MW_LOW"`
	Tx_alarm_mw_high	string	`json:"TX_ALARM_MW_HIGH"`
	Tx_alarm_mw_low	string	`json:"TX_ALARM_MW_LOW"`
	Tx_warn_mw_high	string	`json:"TX_WARN_MW_HIGH"`
	Tx_warn_mw_low	string	`json:"TX_WARN_MW_LOW"`
	Rx_alarm_dbm_high	string	`json:"RX_ALARM_DBM_HIGH"`
	Rx_alarm_dbm_low	string	`json:"RX_ALARM_DBM_LOW"`
	Rx_warn_dbm_high	string	`json:"RX_WARN_DBM_HIGH"`
	Rx_warn_dbm_low	string	`json:"RX_WARN_DBM_LOW"`
	Tx_alarm_dbm_high	string	`json:"TX_ALARM_DBM_HIGH"`
	Tx_alarm_dbm_low	string	`json:"TX_ALARM_DBM_LOW"`
	Tx_warn_dbm_high	string	`json:"TX_WARN_DBM_HIGH"`
	Tx_warn_dbm_low	string	`json:"TX_WARN_DBM_LOW"`
	Fec_corrected	string	`json:"FEC_CORRECTED"`
	Fec_uncorrected	string	`json:"FEC_UNCORRECTED"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Bia	string	`json:"BIA"`
	Autonegotiation	string	`json:"AUTONEGOTIATION"`
	Speed	string	`json:"SPEED"`
	Duplex	string	`json:"DUPLEX"`
	Flow_control	string	`json:"FLOW_CONTROL"`
	Loopback	string	`json:"LOOPBACK"`
	Mtu	string	`json:"MTU"`
	Mru	string	`json:"MRU"`
	Fec	string	`json:"FEC"`
}

var CiscoXrShowControllersHundredgigeAll_Template = `Value Filldown INTERFACE (\S+)
Value Filldown ADMIN_STATE (\S+)
Value Filldown OPS_STATE (\S+)
Value Filldown VENDOR (\S+)
Value Filldown PART_NUM (\S+)
Value Filldown SERIAL_NUM (\S+)
Value LED_STATE (.*)
Value MEDIA_TYPE (.*)
Value TEMPERATURE_VALUE (\S+)
Value VOLTAGE_VALUE (\S+)
Value DOM_ALARMS (.*)
Value LANE (\d+)
Value WAVELENGTH (\S+)
Value RX_VALUE_DBM (\S+)
Value RX_VALUE_MW (\S+)
Value TX_VALUE_DBM (\S+)
Value TX_VALUE_MW (\S+)
Value LASER_BIAS (\S+)
Value TEMPERATURE_ALARM_HIGH (\S+)
Value TEMPERATURE_ALARM_LOW (\S+)
Value TEMPERATURE_WARN_HIGH (\S+)
Value TEMPERATURE_WARN_LOW (\S+)
Value VOLTAGE_ALARM_HIGH (\S+)
Value VOLTAGE_ALARM_LOW (\S+)
Value VOLTAGE_WARN_HIGH (\S+)
Value VOLTAGE_WARN_LOW (\S+)
Value AMPS_ALARM_HIGH (\S+)
Value AMPS_ALARM_LOW (\S+)
Value AMPS_WARN_HIGH (\S+)
Value AMPS_WARN_LOW (\S+)
Value RX_ALARM_MW_HIGH (\S+)
Value RX_ALARM_MW_LOW (\S+)
Value RX_WARN_MW_HIGH (\S+)
Value RX_WARN_MW_LOW (\S+)
Value TX_ALARM_MW_HIGH (\S+)
Value TX_ALARM_MW_LOW (\S+)
Value TX_WARN_MW_HIGH (\S+)
Value TX_WARN_MW_LOW (\S+)
Value RX_ALARM_DBM_HIGH (\S+)
Value RX_ALARM_DBM_LOW (\S+)
Value RX_WARN_DBM_HIGH (\S+)
Value RX_WARN_DBM_LOW (\S+)
Value TX_ALARM_DBM_HIGH (\S+)
Value TX_ALARM_DBM_LOW (\S+)
Value TX_WARN_DBM_HIGH (\S+)
Value TX_WARN_DBM_LOW (\S+)
Value FEC_CORRECTED (\d+)
Value FEC_UNCORRECTED (\d+)
Value MAC_ADDRESS (\S+)
Value BIA (\S+)
Value AUTONEGOTIATION (\S+)
Value SPEED (\S+)
Value DUPLEX (.*)
Value FLOW_CONTROL (\S+)
Value LOOPBACK (.*)
Value MTU (\d+)
Value MRU (\d+)
Value FEC (\S+)

Start
  ^Operational\s+data\s+for\s+interface\s+${INTERFACE}: -> StateChanger
  
StateChanger
  ^State: -> State
  ^Phy: -> Phy
  ^\s+Alarm.+Alarm -> AlarmHeader
  ^\s+Statistics: -> Stats
  ^Lane -> Lanes
  ^MAC\s+address\s+information: -> MacInfo
  ^Operational\s+data\s+for\s+interface\s+${INTERFACE}:
  ^Operational\s+values: -> OpsValues
  ^Autonegotiation\s+disabled.*
  ^Management\s+information\s+for\s+interface -> Ender
  ^$$
  ^\s+$$
  ^.* -> Error "LINE NOT FOUND"

State
  ^\s+Administrative\s+state:\s+${ADMIN_STATE} 
  ^\s+Operational\s+state:\s+${OPS_STATE}
  ^\s+LED\s+state:\s+${LED_STATE}
  ^Autonegotiation\s+${AUTONEGOTIATION}
  ^\s+$$ -> Record StateChanger
  ^$$ -> Record StateChanger
  ^.* -> Error "LINE NOT FOUND"
  
Phy
  ^\s+Media\s+type:\s+${MEDIA_TYPE}
  ^\s+Optics:
  ^\s+Vendor:\s+${VENDOR}
  ^\s+Part\s+number:\s+${PART_NUM}
  ^\s+Serial\s+number:\s+${SERIAL_NUM}
  ^\s+Wavelength:\s+${WAVELENGTH}\s+nm
  ^\s+Digital\s+Optical\s+Monitoring:
  ^\s+Transceiver\s+Temp:\s+${TEMPERATURE_VALUE}
  ^\s+Transceiver\s+Voltage:\s+${VOLTAGE_VALUE}\s+V
  ^\s+Alarms.*high
  ^\s+\(L\).+low
  ^\s+Wavelength.+Bias
  ^\s+Lane.+\(mA\) -> Record
  ^\s+$$
  ^$$ -> Record StateChanger
  ^.* -> Error "LINE NOT FOUND"
  
Lanes
  ^\s+${LANE}\s+${WAVELENGTH}\s+${TX_VALUE_DBM}\s+${TX_VALUE_MW}\s+${RX_VALUE_DBM}\s+${RX_VALUE_MW}\s+${LASER_BIAS}  -> Record
  ^\s+DOM\s+alarms: -> Dom
  ^\s*$$ -> StateChanger
  ^.* -> Error "LINE NOT FOUND"
  
Dom
  ^\s+${DOM_ALARMS} -> Record
  ^\s*$$ -> StateChanger
  ^.* -> Error "LINE NOT FOUND"

AlarmHeader
  ^\s+Thresholds.+Low\s+Low
  ^\s+--.+-- -> AlarmThresholds
  ^.* -> Error "LINE NOT FOUND"
  
AlarmThresholds
  ^\s+Transceiver\s+Temp\s+\(C\):\s+${TEMPERATURE_ALARM_HIGH}\s+${TEMPERATURE_WARN_HIGH}\s+${TEMPERATURE_WARN_LOW}\s+${TEMPERATURE_ALARM_LOW}
  ^\s+Transceiver\s+Voltage\s+\(V\):\s+${VOLTAGE_ALARM_HIGH}\s+${VOLTAGE_WARN_HIGH}\s+${VOLTAGE_WARN_LOW}\s+${VOLTAGE_ALARM_LOW}
  ^\s+Laser\s+Bias\s+\(mA\):\s+${AMPS_ALARM_HIGH}\s+${AMPS_WARN_HIGH}\s+${AMPS_WARN_LOW}\s+${AMPS_ALARM_LOW}
  ^\s+Transmit\s+Power\s+\(mW\):\s+${TX_ALARM_MW_HIGH}\s+${TX_WARN_MW_HIGH}\s+${TX_WARN_MW_LOW}\s+${TX_ALARM_MW_LOW}
  ^\s+Transmit\s+Power\s+\(dBm\):\s+${TX_ALARM_DBM_HIGH}\s+${TX_WARN_DBM_HIGH}\s+${TX_WARN_DBM_LOW}\s+${TX_ALARM_DBM_LOW}
  ^\s+Receive\s+Power\s+\(mW\):\s+${RX_ALARM_MW_HIGH}\s+${RX_WARN_MW_HIGH}\s+${RX_WARN_MW_LOW}\s+${RX_ALARM_MW_LOW}
  ^\s+Receive\s+Power\s+\(dBm\):\s+${RX_ALARM_DBM_HIGH}\s+${RX_WARN_DBM_HIGH}\s+${RX_WARN_DBM_LOW}\s+${RX_ALARM_DBM_LOW} -> Record StateChanger
  ^.* -> Error "LINE NOT FOUND"

Stats
  ^\s+FEC:
  ^\s+Corrected\s+Codeword\s+Count:\s+${FEC_CORRECTED}
  ^\s+Uncorrected\s+Codeword\s+Count:\s+${FEC_UNCORRECTED} -> Record
  ^\s+$$ -> StateChanger
  ^.* -> Error "LINE NOT FOUND"
  
MacInfo
  ^\s+Operational\s+address:\s+${MAC_ADDRESS}
  ^\s+Burnt-in\s+address:\s+${BIA}
  ^\s+No\s+unicast.*
  ^\s+Operating\s+in.*
  ^\s+$$ -> Record StateChanger
  ^$$ -> Record StateChanger
  ^.* -> Error "LINE NOT FOUND"
 
OpsValues
  ^\s+Speed:\s+${SPEED}
  ^\s+Duplex:\s+${DUPLEX}
  ^\s+Flowcontrol:\s+${FLOW_CONTROL}
  ^\s+Loopback:\s+${LOOPBACK}
  ^\s+MTU:\s+${MTU}
  ^\s+MRU:\s+${MRU}
  ^\s+Forward\s+error\s+correction:\s+${FEC}
  ^\s+Inter\-pack.*
  ^\s+BER\s+monitoring:.*
  ^\s+Signal.*
  ^\s+$$ -> Record StateChanger
  ^$$ -> Record StateChanger
  ^.* -> Error "LINE NOT FOUND"
  
Ender
  ^.*

EOF
`