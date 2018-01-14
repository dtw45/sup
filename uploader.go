package main

import (
		"os"
	    "fmt"
	    "strconv"
		"github.com/strava/go.strava"
)

func main() {

	API_TOKEN := os.Getenv("STRAVA_UP_API_TOKEN")
	ATHLETE_ID := str2int(os.Getenv("STRAVA_UP_ATHELETE_ID"))

	// Decleare client used for all strava api calls
	client := strava.NewClient(API_TOKEN)

	// Get the datetime of the last activity uploaded
	service := strava.NewAthletesService(client)
	activities, err := service.ListActivities(ATHLETE_ID).PerPage(1).Do()
	check(err)

	date_of_last_uploaded_activity := activities[0].StartDate

	fmt.Println("last uploaded activity: ", date_of_last_uploaded_activity)
	// Check garmin for activities since last upload


	// Upload each activity to strava and check response code
	// service := strava.NewUploadsService(client)

	// for _, activity := queue {
	// 	file_reader, err := os.Open(activity)
	// 	upload, err := service.Create(FILE_DATA_TYPE, activity, file_reader).Private(1).Do()
	// 	if e, ok := err.(*strava.Error); ok {
	// 		fmt.Println("strava error:", e)
	// 	} else {
	// 		fmt.Println("os level error", e)
	// 	}

	// 	upload_ids = append(upload.Id)

	// }

	// // Poll strava for upload status
	// service := strava.
	// for _, id := upload_ids {
	// 	upload, err := service.Get(id)
	// }

	// Launch webserver to display new activities and their metadata

	// Detach and Exit

}

func check(err error) {
	// Raise error if not none
	if err != nil {
	// Raise and report error
		fmt.Println("ERROR: ", err)
	}

}

func str2int(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic("Not a number")
	}
	return i
}
