package models

type CiscoNxosShowRouteMap struct {
	Name	string	`json:"NAME"`
	Action	string	`json:"ACTION"`
	Seq	string	`json:"SEQ"`
	Match_clauses	[]string	`json:"MATCH_CLAUSES"`
	Set_clauses	[]string	`json:"SET_CLAUSES"`
}