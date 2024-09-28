package models

type AlcatelSrosShowRouterBgpSummaryFamily struct {
	Neighbor string `json:"NEIGHBOR"`
	As       string `json:"AS"`
	Pktrcvd  string `json:"PKTRCVD"`
	Pktsent  string `json:"PKTSENT"`
	Inq      string `json:"INQ"`
	Outq     string `json:"OUTQ"`
	Up_down  string `json:"UP_DOWN"`
	State    string `json:"STATE"`
	Rcv      string `json:"RCV"`
	Act      string `json:"ACT"`
	Sent     string `json:"SENT"`
}
