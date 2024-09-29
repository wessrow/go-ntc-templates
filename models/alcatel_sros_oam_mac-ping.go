package models

type AlcatelSrosOamMacPing struct {
	Router	string	`json:"ROUTER"`
	Port	string	`json:"PORT"`
	Tag	string	`json:"TAG"`
}

var AlcatelSrosOamMacPing_Template = `Value ROUTER (\d+\.\d+\.\d+\.\d+)
Value PORT ([0-9]0?\/[1-2]\/[0-9]+|lag-[0-9]{1,3})
Value TAG ([0-9]{1,4}|\d+.?\d+?)

Start
  ^1\s+${ROUTER}:sap${PORT}:${TAG} -> Record

`