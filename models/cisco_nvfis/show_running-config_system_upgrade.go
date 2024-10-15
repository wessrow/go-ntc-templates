package cisco_nvfis

type ShowRunningConfigSystemUpgrade struct {
	Image_name     string `json:"IMAGE_NAME"`
	Image_location string `json:"IMAGE_LOCATION"`
	Apply_image    string `json:"APPLY_IMAGE"`
	Schedule_time  string `json:"SCHEDULE_TIME"`
}

var ShowRunningConfigSystemUpgrade_Template string = `Value APPLY_IMAGE (\S+iso)
Value SCHEDULE_TIME (\d+)
Value IMAGE_NAME (\S+iso)
Value IMAGE_LOCATION (\S+iso)


Start
  ^system\supgrade\sapply-image\s${APPLY_IMAGE}\sscheduled-time\s${SCHEDULE_TIME}
  ^system\supgrade\simage-name\s${IMAGE_NAME}
  ^\slocation\s${IMAGE_LOCATION}
  ^! -> Record

EOF`
