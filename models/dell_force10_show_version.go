package models

type DellForce10ShowVersion struct {
	Os_version	string	`json:"OS_VERSION"`
	Uptime	string	`json:"UPTIME"`
	Device_type	string	`json:"DEVICE_TYPE"`
}

var DellForce10ShowVersion_Template = `Value OS_VERSION (\S+)
Value UPTIME (.+)
Value DEVICE_TYPE (\S+)

Start
  ^Dell Application Software Version:\s+${OS_VERSION}
  ^Dell Networking OS uptime is\s${UPTIME}
  ^System Type:\s${DEVICE_TYPE} -> Record

`