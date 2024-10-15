package huawei_vrp

type DisplaySnLicense struct {
	Serial_number string `json:"SERIAL_NUMBER"`
}

var DisplaySnLicense_Template string = `Value SERIAL_NUMBER (\S+)


Start
  ^.*ESN of device:\s${SERIAL_NUMBER}$$
  ^. -> Error

`
