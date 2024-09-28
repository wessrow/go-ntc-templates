package models

type CiscoIosShowIpRouteSummary struct {
	Route_source	string	`json:"ROUTE_SOURCE"`
	Name	string	`json:"NAME"`
	Networks	string	`json:"NETWORKS"`
	Subnets	string	`json:"SUBNETS"`
	Replicates	string	`json:"REPLICATES"`
	Overhead	string	`json:"OVERHEAD"`
	Memory	string	`json:"MEMORY"`
}