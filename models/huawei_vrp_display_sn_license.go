package models

type HuaweiVrpDisplaySnLicense struct {
	Serial_number	string	`json:"SERIAL_NUMBER"`
}

var HuaweiVrpDisplaySnLicense_Template = `Value SERIAL_NUMBER (\S+)


Start
  ^.*ESN of device:\s${SERIAL_NUMBER}$$
  ^. -> Error


`