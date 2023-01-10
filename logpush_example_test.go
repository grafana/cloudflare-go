package cloudflare_test

import (
	"context"
	"fmt"
	"log"

	cloudflare "github.com/grafana/cloudflare-go"
)

var exampleNewLogpushJob = cloudflare.LogpushJob{
	Enabled:         false,
	Name:            "example.com",
	LogpullOptions:  "fields=RayID,ClientIP,EdgeStartTimestamp&timestamps=rfc3339",
	DestinationConf: "s3://mybucket/logs?region=us-west-2",
}

var exampleUpdatedLogpushJob = cloudflare.LogpushJob{
	Enabled:         true,
	Name:            "updated.com",
	LogpullOptions:  "fields=RayID,ClientIP,EdgeStartTimestamp",
	DestinationConf: "gs://mybucket/logs",
}

func ExampleAPI_CreateZoneLogpushJob() {
	api, err := cloudflare.New(apiKey, user)
	if err != nil {
		log.Fatal(err)
	}

	zoneID, err := api.ZoneIDByName(domain)
	if err != nil {
		log.Fatal(err)
	}

	job, err := api.CreateZoneLogpushJob(context.Background(), zoneID, exampleNewLogpushJob)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", job)
}

func ExampleAPI_UpdateZoneLogpushJob() {
	api, err := cloudflare.New(apiKey, user)
	if err != nil {
		log.Fatal(err)
	}

	zoneID, err := api.ZoneIDByName(domain)
	if err != nil {
		log.Fatal(err)
	}

	err = api.UpdateZoneLogpushJob(context.Background(), zoneID, 1, exampleUpdatedLogpushJob)
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleAPI_ListZoneLogpushJobs() {
	api, err := cloudflare.New(apiKey, user)
	if err != nil {
		log.Fatal(err)
	}

	zoneID, err := api.ZoneIDByName(domain)
	if err != nil {
		log.Fatal(err)
	}

	jobs, err := api.ListZoneLogpushJobs(context.Background(), zoneID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", jobs)
	for _, r := range jobs {
		fmt.Printf("%+v\n", r)
	}
}

func ExampleAPI_GetZoneLogpushJob() {
	api, err := cloudflare.New(apiKey, user)
	if err != nil {
		log.Fatal(err)
	}

	zoneID, err := api.ZoneIDByName(domain)
	if err != nil {
		log.Fatal(err)
	}

	job, err := api.GetZoneLogpushJob(context.Background(), zoneID, 1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", job)
}

func ExampleAPI_DeleteZoneLogpushJob() {
	api, err := cloudflare.New(apiKey, user)
	if err != nil {
		log.Fatal(err)
	}

	zoneID, err := api.ZoneIDByName(domain)
	if err != nil {
		log.Fatal(err)
	}

	err = api.DeleteZoneLogpushJob(context.Background(), zoneID, 1)
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleAPI_GetZoneLogpushOwnershipChallenge() {
	api, err := cloudflare.New(apiKey, user)
	if err != nil {
		log.Fatal(err)
	}

	zoneID, err := api.ZoneIDByName(domain)
	if err != nil {
		log.Fatal(err)
	}

	ownershipChallenge, err := api.GetZoneLogpushOwnershipChallenge(context.Background(), zoneID, "destination_conf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", ownershipChallenge)
}

func ExampleAPI_ValidateZoneLogpushOwnershipChallenge() {
	api, err := cloudflare.New(apiKey, user)
	if err != nil {
		log.Fatal(err)
	}

	zoneID, err := api.ZoneIDByName(domain)
	if err != nil {
		log.Fatal(err)
	}

	isValid, err := api.ValidateZoneLogpushOwnershipChallenge(context.Background(), zoneID, "destination_conf", "ownership_challenge")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", isValid)
}

func ExampleAPI_CheckZoneLogpushDestinationExists() {
	api, err := cloudflare.New(apiKey, user)
	if err != nil {
		log.Fatal(err)
	}

	zoneID, err := api.ZoneIDByName(domain)
	if err != nil {
		log.Fatal(err)
	}

	exists, err := api.CheckZoneLogpushDestinationExists(context.Background(), zoneID, "destination_conf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", exists)
}
