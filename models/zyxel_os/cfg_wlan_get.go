package zyxel_os

type CfgWlanGet struct {
	Channel         string `json:"CHANNEL"`
	Encryption_type string `json:"ENCRYPTION_TYPE"`
	Band            string `json:"BAND"`
	Ssid            string `json:"SSID"`
	Bandwidth       string `json:"BANDWIDTH"`
	Max_devices     string `json:"MAX_DEVICES"`
	Key             string `json:"KEY"`
	Index           string `json:"INDEX"`
	Enabled         string `json:"ENABLED"`
}

var CfgWlanGet_Template string = `Value INDEX (\d+)
Value BAND (2\.4GHz|5GHz)
Value SSID (.+\S)
Value ENABLED (0|1)
Value BANDWIDTH (\d+(M|G|)Hz)
Value CHANNEL (\S+)
Value MAX_DEVICES (\d+)
Value ENCRYPTION_TYPE (\S+)
Value KEY (\S*)

Start
  ^Index\s+Band\s+SSID\s+Enable\s+Bandwidth\s+Channel\s+MaxDevices\s+SecurityMode\s+PskValue\s*$$ -> WLANTable
  ^\s*$$
  ^. -> Error

WLANTable
  ^${INDEX}\s+${BAND}\s+${SSID}\s+${ENABLED}\s+${BANDWIDTH}\s+${CHANNEL}\s+${MAX_DEVICES}\s+${ENCRYPTION_TYPE}\s+${KEY}\s*$$ -> Record
  ^Command\sSuccessful.\s*$$
  ^\s*$$
  ^. -> Error
`
