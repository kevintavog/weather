package main

func dataCollector() {

	initDataDirectories()

	go soundersScheduleCollector()
	go weatherDataCollector()
	go stopInfoDataCollector()
}
