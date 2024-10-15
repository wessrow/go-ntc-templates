package cisco_nxos

type ShowLicenseUsage struct {
	Expiry_date   string `json:"EXPIRY_DATE"`
	Comments      string `json:"COMMENTS"`
	Feature       string `json:"FEATURE"`
	Installed     string `json:"INSTALLED"`
	License_count string `json:"LICENSE_COUNT"`
	Status        string `json:"STATUS"`
}

var ShowLicenseUsage_Template string = `Value FEATURE (\S+)
Value INSTALLED (Yes|No)
Value LICENSE_COUNT (\d+|-)
Value STATUS (Unused|In\s+use)
Value EXPIRY_DATE (\S+|\s*?)
Value COMMENTS (.+)

Start
  ^Feature\s+Ins\s+Lic\s+Status\s+Expiry\s+Date\s+Comments -> Begin
  ^\s*$$
  ^. -> Error

Begin
  ^\s+Count
  ^${FEATURE}\s+${INSTALLED}\s+${LICENSE_COUNT}\s+${STATUS}\s+?${EXPIRY_DATE}\s+${COMMENTS}\s*$$ -> Record
  ^-+\s*$$
  ^\s*$$
  ^. -> Error
`
