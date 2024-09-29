package models

type PaloaltoPanosShowJobsAll struct {
	Enqueued	string	`json:"ENQUEUED"`
	Dequeued	string	`json:"DEQUEUED"`
	Id	string	`json:"ID"`
	Type	string	`json:"TYPE"`
	Status	string	`json:"STATUS"`
	Result	string	`json:"RESULT"`
	Completed	string	`json:"COMPLETED"`
}

var PaloaltoPanosShowJobsAll_Template = `Value ENQUEUED (\S+\s+\S+?)
Value DEQUEUED (\S+)
Value ID (\d+)
Value TYPE (\w+)
Value STATUS (\w+)
Value RESULT (\w+)
Value COMPLETED (\S+)

Start
#  ^${ENQUEUED}\s+\S+\s+${ID}\s+${TYPE}\s+${STATUS}\s+${RESULT}\s+${COMPLETED} -> Record
  ^${ENQUEUED}\s+${ID}\s+${TYPE}\s+${STATUS}\s+${RESULT}\s+${COMPLETED} -> Record
  ^${ENQUEUED}\s+${DEQUEUED}\s+${ID}\s+${TYPE}\s+${STATUS}\s+${RESULT}\s+${COMPLETED} -> Record

`