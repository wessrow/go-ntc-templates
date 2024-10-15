package paloalto_panos

type ShowJobsAll struct {
	Dequeued  string `json:"DEQUEUED"`
	Id        string `json:"ID"`
	Type      string `json:"TYPE"`
	Status    string `json:"STATUS"`
	Result    string `json:"RESULT"`
	Completed string `json:"COMPLETED"`
	Enqueued  string `json:"ENQUEUED"`
}

var ShowJobsAll_Template string = `Value ENQUEUED (\S+\s+\S+?)
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
