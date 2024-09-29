package models

type CiscoWlcSshShowBoot struct {
	Image_type	string	`json:"IMAGE_TYPE"`
	Image_version	string	`json:"IMAGE_VERSION"`
	Image_status	string	`json:"IMAGE_STATUS"`
}

var CiscoWlcSshShowBoot_Template = `Value IMAGE_TYPE (\S+\s+\S+\s+[A-z]+)
Value IMAGE_VERSION (\S+)
Value IMAGE_STATUS ((?:\(\S+\))\s*(?:\(\S+\))?)


Start
  ^${IMAGE_TYPE}\.+\s+${IMAGE_VERSION}\s*(${IMAGE_STATUS})?$$ -> Record
  ^\s*$$
  ^. -> Error

`