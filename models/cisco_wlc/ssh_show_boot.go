package cisco_wlc

type SshShowBoot struct {
	Image_status  string `json:"IMAGE_STATUS"`
	Image_type    string `json:"IMAGE_TYPE"`
	Image_version string `json:"IMAGE_VERSION"`
}

var SshShowBoot_Template string = `Value IMAGE_TYPE (\S+\s+\S+\s+[A-z]+)
Value IMAGE_VERSION (\S+)
Value IMAGE_STATUS ((?:\(\S+\))\s*(?:\(\S+\))?)


Start
  ^${IMAGE_TYPE}\.+\s+${IMAGE_VERSION}\s*(${IMAGE_STATUS})?$$ -> Record
  ^\s*$$
  ^. -> Error
`
