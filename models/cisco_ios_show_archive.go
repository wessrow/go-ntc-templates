package models

type CiscoIosShowArchive struct {
	State	string	`json:"STATE"`
	Next_filename	string	`json:"NEXT_FILENAME"`
	Filenames	[]string	`json:"FILENAMES"`
	Current_index	string	`json:"CURRENT_INDEX"`
}