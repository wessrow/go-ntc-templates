package models

type AristaEosBashDfH struct {
	File_system	string	`json:"FILE_SYSTEM"`
	Size	string	`json:"SIZE"`
	Used	string	`json:"USED"`
	Avail	string	`json:"AVAIL"`
	Use_percent	string	`json:"USE_PERCENT"`
	Mounted_on	string	`json:"MOUNTED_ON"`
}