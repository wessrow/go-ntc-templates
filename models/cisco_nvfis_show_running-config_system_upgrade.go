package models

type CiscoNvfisShowRunningConfigSystemUpgrade struct {
	Apply_image	string	`json:"APPLY_IMAGE"`
	Schedule_time	string	`json:"SCHEDULE_TIME"`
	Image_name	string	`json:"IMAGE_NAME"`
	Image_location	string	`json:"IMAGE_LOCATION"`
}