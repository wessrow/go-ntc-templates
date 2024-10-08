package zyxel_os 

type CfgWlanGet struct {
	Index	string	`json:"INDEX"`
	Band	string	`json:"BAND"`
	Ssid	string	`json:"SSID"`
	Enabled	string	`json:"ENABLED"`
	Bandwidth	string	`json:"BANDWIDTH"`
	Channel	string	`json:"CHANNEL"`
	Max_devices	string	`json:"MAX_DEVICES"`
	Encryption_type	string	`json:"ENCRYPTION_TYPE"`
	Key	string	`json:"KEY"`
}

var CfgWlanGet_Template = `Value INDEX (\d+)
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