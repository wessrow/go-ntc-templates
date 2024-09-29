package models

type PaloaltoPanosRequestLicenseInfo struct {
	Feature	string	`json:"FEATURE"`
	Description	string	`json:"DESCRIPTION"`
	Serial	string	`json:"SERIAL"`
	Auth_code	string	`json:"AUTH_CODE"`
	Issued	string	`json:"ISSUED"`
	Expires	string	`json:"EXPIRES"`
	Expired	string	`json:"EXPIRED"`
	Base_license	string	`json:"BASE_LICENSE"`
}

var PaloaltoPanosRequestLicenseInfo_Template = `Value FEATURE (.+)
Value DESCRIPTION (.+)
Value SERIAL (.+)
Value AUTH_CODE (.+)
Value ISSUED (.+)
Value EXPIRES (.+)
Value EXPIRED (yes|no)
Value BASE_LICENSE (.+?)

Start
  ^License\s+entry: -> Record
  ^Feature: ${FEATURE}
  ^Description: ${DESCRIPTION}
  ^Serial: ${SERIAL}
  ^Authcode: ${AUTH_CODE}
  ^Issued: ${ISSUED}
  ^Expires: ${EXPIRES}
  ^Expired\?: ${EXPIRED}
  ^Base\s+[Ll]icense:\s+${BASE_LICENSE}\s*$$
  ^\s*Current\s+\S+\s+[Dd]ate
  ^\s*$$
  ^. -> Error

`