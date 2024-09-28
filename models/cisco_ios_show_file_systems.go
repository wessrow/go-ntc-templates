package models

type CiscoIosShowFileSystems struct {
	Size	string	`json:"SIZE"`
	Free	string	`json:"FREE"`
	Type	string	`json:"TYPE"`
	Flags	string	`json:"FLAGS"`
	Prefixes	string	`json:"PREFIXES"`
}