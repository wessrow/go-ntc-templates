package ruckus_fastiron 

type ShowMediaValidation struct {
	Interface	string	`json:"INTERFACE"`
	Vendor	string	`json:"VENDOR"`
	Optic_description	string	`json:"OPTIC_DESCRIPTION"`
	Optic_type	string	`json:"OPTIC_TYPE"`
}

var ShowMediaValidation_Template = `Value Required INTERFACE (\d\S+)
Value VENDOR (\S+)
Value OPTIC_DESCRIPTION (.*)
Value OPTIC_TYPE (\S+)

Start
  ^\s*Port\s+Supported\.*
  ^\s*-+
  ^${INTERFACE}\s+\w+\s+${VENDOR}\s+Type\s+:\s+${OPTIC_DESCRIPTION}\s+\(${OPTIC_TYPE}\)\s*$$ -> Record
  ^. -> Error
`