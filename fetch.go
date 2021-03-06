package main

import (
	"fmt"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"github.com/json-iterator/go"
	"log"
	"net/http"
)

type FetchMeetup struct {
}

func (f FetchMeetup) Name() string {
	return "fetchMeetup"
}

func (f FetchMeetup) Version() string {
	return "1.0"
}

type FetchMeetupInput struct {
	TopicCategory string
	EndDate string
	Radius string
	Lat string
	Lon string
	Page string
}

type FetchMeetupOutput struct {
	Events []Event
}

func (f FetchMeetup) Execute(in step.Context) (interface{}, error) {

	input := FetchMeetupInput{}
	err := in.BindInputs(&input)
	if err != nil {
		log.Fatal(err, "BindInputs: ")
		return nil, err
	}

	output, err := f.execute(input)
	if err != nil {
		log.Fatal(err, "Execute: ")
		return nil, err
	}
	return output, nil
}

func (f FetchMeetup) execute(input FetchMeetupInput) (
	*FetchMeetupOutput, error) {
	requestUrl := fmt.Sprintf("https://api.meetup.com/find/upcoming_events?&key=443bc10506c5c6a5d376a5f35d2246&lon=%s&end_date=%s&topic_category=%s&page=%s&radius=%s&lat=%s", input.Lon, input.EndDate, input.TopicCategory, input.Page, input.Radius, input.Lat)
	response, err := http.Get(requestUrl)
	if err != nil {
		log.Fatal(err, "Request: ")
		return nil, err
	}
	defer response.Body.Close()

	dec := jsoniter.NewDecoder(response.Body)
	fmt.Println(dec)
	resp := Response{}
	fmt.Println(resp)

	err = dec.Decode(&resp)
	if err != nil {
		return nil, err
	}

	return &FetchMeetupOutput{
		Events: resp.Events,
	}, nil
}