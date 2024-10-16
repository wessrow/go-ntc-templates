package cisco_wlc

type SshShow80211aCleanairConfig struct {
	Bluetooth_link_alarm          string `json:"BLUETOOTH_LINK_ALARM"`
	Jammer_alarm                  string `json:"JAMMER_ALARM"`
	Dect_like_phone_alarm         string `json:"DECT_LIKE_PHONE_ALARM"`
	Wifi_invalid_channel_alarm    string `json:"WIFI_INVALID_CHANNEL_ALARM"`
	Cleanair                      string `json:"CLEANAIR"`
	Dot15_4                       string `json:"DOT15_4"`
	Wimax_fixed                   string `json:"WIMAX_FIXED"`
	Wimax_mobile                  string `json:"WIMAX_MOBILE"`
	Ble_beacon                    string `json:"BLE_BEACON"`
	Tdd_transmit_alarm            string `json:"TDD_TRANSMIT_ALARM"`
	Rogue_duty_cycle_thresh       string `json:"ROGUE_DUTY_CYCLE_THRESH"`
	Air_quality_report            string `json:"AIR_QUALITY_REPORT"`
	Dect_like_phone               string `json:"DECT_LIKE_PHONE"`
	Canopy                        string `json:"CANOPY"`
	Ed_rrm_sensitivity            string `json:"ED_RRM_SENSITIVITY"`
	Ed_rrm_thresh                 string `json:"ED_RRM_THRESH"`
	Rogue_contribution            string `json:"ROGUE_CONTRIBUTION"`
	Persistant_device_propagation string `json:"PERSISTANT_DEVICE_PROPAGATION"`
	Int_device_alarm              string `json:"INT_DEVICE_ALARM"`
	Bluetooth_discovery_alarm     string `json:"BLUETOOTH_DISCOVERY_ALARM"`
	Wimax_fixed_alarm             string `json:"WIMAX_FIXED_ALARM"`
	Dot11_fh_alarm                string `json:"DOT11_FH_ALARM"`
	Canopy_alarm                  string `json:"CANOPY_ALARM"`
	Air_quality_alarm             string `json:"AIR_QUALITY_ALARM"`
	Dot11_fh                      string `json:"DOT11_FH"`
	Wifi_inverted                 string `json:"WIFI_INVERTED"`
	Wifi_invalid_channel          string `json:"WIFI_INVALID_CHANNEL"`
	Ed_rrm_state                  string `json:"ED_RRM_STATE"`
	Aq_alarm_thresh               string `json:"AQ_ALARM_THRESH"`
	Interfere_device_report       string `json:"INTERFERE_DEVICE_REPORT"`
	Continuous_transmit           string `json:"CONTINUOUS_TRANSMIT"`
	Tdd_transmit                  string `json:"TDD_TRANSMIT"`
	Jammer                        string `json:"JAMMER"`
	Dot15_4_alarm                 string `json:"DOT15_4_ALARM"`
	Wifi_inverted_alarm           string `json:"WIFI_INVERTED_ALARM"`
	Superag_alarm                 string `json:"SUPERAG_ALARM"`
	Aq_unclassified_thresh        string `json:"AQ_UNCLASSIFIED_THRESH"`
	Bluetooth_link                string `json:"BLUETOOTH_LINK"`
	Bluetooth_discovery           string `json:"BLUETOOTH_DISCOVERY"`
	Wimax_mobile_alarm            string `json:"WIMAX_MOBILE_ALARM"`
	Ble_beacon_alarm              string `json:"BLE_BEACON_ALARM"`
	Microsoft_device              string `json:"MICROSOFT_DEVICE"`
	Microwave_alarm               string `json:"MICROWAVE_ALARM"`
	Microsoft_device_alarm        string `json:"MICROSOFT_DEVICE_ALARM"`
	Microwave                     string `json:"MICROWAVE"`
	Video_camera                  string `json:"VIDEO_CAMERA"`
	Superag                       string `json:"SUPERAG"`
	Video_camera_alarm            string `json:"VIDEO_CAMERA_ALARM"`
	Persistant_device_state       string `json:"PERSISTANT_DEVICE_STATE"`
	Air_quality_period            string `json:"AIR_QUALITY_PERIOD"`
	Aq_unclassified_inter         string `json:"AQ_UNCLASSIFIED_INTER"`
	Continuous_transmit_alarm     string `json:"CONTINUOUS_TRANSMIT_ALARM"`
}

var SshShow80211aCleanairConfig_Template string = `Value CLEANAIR (.+?)
Value AIR_QUALITY_REPORT (\w+)
Value AIR_QUALITY_PERIOD (\d+)
Value AIR_QUALITY_ALARM (.+?)
Value AQ_ALARM_THRESH (\d+)
Value AQ_UNCLASSIFIED_INTER (.+?)
Value AQ_UNCLASSIFIED_THRESH (\d+)
Value INTERFERE_DEVICE_REPORT (\w+)
Value BLUETOOTH_LINK (\w+)
Value MICROWAVE (\w+)
Value DOT11_FH (\w+)
Value BLUETOOTH_DISCOVERY (\w+)
Value TDD_TRANSMIT (\w+)
Value JAMMER (\w+)
Value CONTINUOUS_TRANSMIT (\w+)
Value DECT_LIKE_PHONE (\w+)
Value VIDEO_CAMERA (\w+)
Value DOT15_4 (\w+)
Value WIFI_INVERTED (\w+)
Value WIFI_INVALID_CHANNEL (\w+)
Value SUPERAG (\w+)
Value CANOPY (\w+)
Value MICROSOFT_DEVICE (\w+)
Value WIMAX_MOBILE (\w+)
Value WIMAX_FIXED (\w+)
Value BLE_BEACON (\w+)
Value INT_DEVICE_ALARM (\w+)
Value BLUETOOTH_LINK_ALARM (\w+)
Value MICROWAVE_ALARM (\w+)
Value DOT11_FH_ALARM (\w+)
Value BLUETOOTH_DISCOVERY_ALARM (\w+)
Value TDD_TRANSMIT_ALARM (\w+)
Value JAMMER_ALARM (\w+)
Value CONTINUOUS_TRANSMIT_ALARM (\w+)
Value DECT_LIKE_PHONE_ALARM (\w+)
Value VIDEO_CAMERA_ALARM (\w+)
Value DOT15_4_ALARM (\w+)
Value WIFI_INVERTED_ALARM (\w+)
Value WIFI_INVALID_CHANNEL_ALARM (\w+)
Value SUPERAG_ALARM (\w+)
Value CANOPY_ALARM (\w+)
Value MICROSOFT_DEVICE_ALARM (\w+)
Value WIMAX_MOBILE_ALARM (\w+)
Value WIMAX_FIXED_ALARM (\w+)
Value BLE_BEACON_ALARM (\w+)
Value ED_RRM_STATE (\w+)
Value ED_RRM_SENSITIVITY (\w+)
Value ED_RRM_THRESH (\d+)
Value ROGUE_CONTRIBUTION (\w+)
Value ROGUE_DUTY_CYCLE_THRESH (\d+)
Value PERSISTANT_DEVICE_STATE (\w+)
Value PERSISTANT_DEVICE_PROPAGATION (\w+)


Start
  ^Clean\s+Air\s+Solution\.*\s${CLEANAIR}s*$$
  ^Air\s+Quality\s+Settings\: -> AirQuality
  ^\s+CleanAir\s+ED-RRM\s+State\.*\s${ED_RRM_STATE}s*$$
  ^\s+CleanAir\s+ED-RRM\s+Sensitivity\.*\s${ED_RRM_SENSITIVITY}s*$$
  ^\s+CleanAir\s+ED-RRM\s+Custom\s+Threshold\.*\s${ED_RRM_THRESH}s*$$
  ^\s+CleanAir\s+Rogue\s+Contribution\.*\s${ROGUE_CONTRIBUTION}s*$$
  ^\s+CleanAir\s+Rogue\s+Duty-Cycle\s+Threshold\.*\s${ROGUE_DUTY_CYCLE_THRESH}s*$$
  ^\s+CleanAir\s+Persistent\s+Devices\s+state\.*\s${PERSISTANT_DEVICE_STATE}s*$$
  ^\s+CleanAir\s+Persistent\s+Device\s+Propagation\.*\s${PERSISTANT_DEVICE_PROPAGATION}s*$$
  ^\s*$$
  ^. -> Error


AirQuality
  ^\s+Air\s+Quality\s+Reporting\.*\s${AIR_QUALITY_REPORT}s*$$
  ^\s+Air\s+Quality\s+Reporting\s+Period\s+\(min\)\.*\s${AIR_QUALITY_PERIOD}s*$$
  ^\s+Air\s+Quality\s+Alarms\.*\s${AIR_QUALITY_ALARM}s*$$
  ^\s+Air\s+Quality\s+Alarm\s+Threshold\.*\s${AQ_ALARM_THRESH}s*$$
  ^\s+Unclassified\s+Interference\.*\s${AQ_UNCLASSIFIED_INTER}s*$$
  ^\s+Unclassified\s+Severity\s+Threshold\.*\s${AQ_UNCLASSIFIED_THRESH}s*$$
  ^Interference\s+Device\s+Settings\: -> Interference
  ^\s*$$
  ^. -> Error

Interference
  ^\s+Interference\s+Device\s+Reporting\.*\s${INTERFERE_DEVICE_REPORT}S*$$
  ^\s+Interference\s+Device\s+Types\: -> Device_type
  ^\s*$$
  ^. -> Error

Device_type
  ^\s+Bluetooth\s+Link\.*\s${BLUETOOTH_LINK}
  ^\s+Microwave\s+Oven\.*\s${MICROWAVE}s*$$
  ^\s+802.11\s+FH\.*\s${DOT11_FH}s*$$
  ^\s+Bluetooth\s+Discovery\.*\s${BLUETOOTH_DISCOVERY}s*$$
  ^\s+TDD\s+Transmitter\.*\s${TDD_TRANSMIT}s*$$
  ^\s+Jammer\.*\s${JAMMER}s*$$
  ^\s+Continuous\s+Transmitter\.*\s${CONTINUOUS_TRANSMIT}s*$$
  ^\s+DECT-like\s+Phone\.*\s${DECT_LIKE_PHONE}s*$$
  ^\s+Video\s+Camera\.*\s${VIDEO_CAMERA}s*$$
  ^\s+802.15.4\.*\s${DOT15_4}s*$$
  ^\s+WiFi\s+Inverted\.*\s${WIFI_INVERTED}s*$$
  ^\s+WiFi\s+Invalid\s+Channel\.*\s${WIFI_INVALID_CHANNEL}s*$$
  ^\s+SuperAG\.*\s${SUPERAG}s*$$
  ^\s+Canopy\.*\s${CANOPY}s*$$
  ^\s+Microsoft\s+Device\.*\s${MICROSOFT_DEVICE}s*$$
  ^\s+WiMax\s+Mobile\.*\s${WIMAX_MOBILE}s*$$
  ^\s+WiMax\s+Fixed\.*\s${WIMAX_FIXED}s*$$
  ^\s+BLE\s+Beacon\.*\s${BLE_BEACON}s*$$
  ^\s+Interference\s+Device\s+Alarms\.*\s${INT_DEVICE_ALARM}s*$$
  ^\s+Interference\s+Device\s+Types\s+Triggering\s+Alarms\: -> Device_alarms
  ^\s*$$
  ^. -> Error

Device_alarms
  ^\s+Bluetooth\s+Link\.*\s${BLUETOOTH_LINK_ALARM}s*$$
  ^\s+Microwave\s+Oven\.*\s${MICROWAVE_ALARM}s*$$
  ^\s+802.11\s+FH\.*\s${DOT11_FH_ALARM}s*$$
  ^\s+Bluetooth\s+Discovery\.*\s${BLUETOOTH_DISCOVERY_ALARM}s*$$
  ^\s+TDD\s+Transmitter\.*\s${TDD_TRANSMIT_ALARM}s*$$
  ^\s+Jammer\.*\s${JAMMER_ALARM}s*$$
  ^\s+Continuous\s+Transmitter\.*\s${CONTINUOUS_TRANSMIT_ALARM}s*$$
  ^\s+DECT-like\s+Phone\.*\s${DECT_LIKE_PHONE_ALARM}s*$$
  ^\s+Video\s+Camera\.*\s${VIDEO_CAMERA_ALARM}s*$$
  ^\s+802.15.4\.*\s${DOT15_4_ALARM}s*$$
  ^\s+WiFi\s+Inverted\.*\s${WIFI_INVERTED_ALARM}s*$$
  ^\s+WiFi\s+Invalid\s+Channel\.*\s${WIFI_INVALID_CHANNEL_ALARM}s*$$
  ^\s+SuperAG\.*\s${SUPERAG_ALARM}s*$$
  ^\s+Canopy\.*\s${CANOPY_ALARM}s*$$
  ^\s+Microsoft\s+Device\.*\s${MICROSOFT_DEVICE_ALARM}s*$$
  ^\s+WiMax\s+Mobile\.*\s${WIMAX_MOBILE_ALARM}s*$$
  ^\s+WiMax\s+Fixed\.*\s${WIMAX_FIXED_ALARM}s*$$
  ^\s+BLE\s+Beacon\.*\s${BLE_BEACON_ALARM}s*$$
  ^Additional\s+Clean\s+Air\s+Settings\: -> Start`
