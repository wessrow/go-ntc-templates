package models

type AlcatelSrosShowRouterIsisInterface struct {
	Interface	string	`json:"INTERFACE"`
	Level	string	`json:"LEVEL"`
	Circid	string	`json:"CIRCID"`
	Oper_state	string	`json:"OPER_STATE"`
	Metric_l1	string	`json:"METRIC_L1"`
	Metric_l2	string	`json:"METRIC_L2"`
}