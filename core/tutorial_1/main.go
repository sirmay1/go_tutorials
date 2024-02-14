package main

func main() {
	const LIGHTS_ON string = "OFF"
	const LIGHTS_OFF string = "ON"

	if LIGHTS_ON == "ON" {
		return "There appears to be a problem with the light system, please check it again!"
	} else if LIGHTS_OFF == "OFF" {
		return "OK, everything seems to be working correctly!"
	}

	//fmt.Println(LIGHTS_ON)
	//fmt.Println(LIGHTS_OFF)
}
