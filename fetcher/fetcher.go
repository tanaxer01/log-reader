package fetcher

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
)

type AwsFetcher struct {
	svc *cloudwatchlogs.Client
}

func NewAwsFetcher() *AwsFetcher {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithSharedConfigProfile("bicevida-dev"),
	)
	if err != nil {
		panic(err)
	}

	svc := cloudwatchlogs.NewFromConfig(cfg)
	return &AwsFetcher{
		svc: svc,
	}
}

func (f *AwsFetcher) FetchGroups(pattern string) []string {
	res, err := f.svc.DescribeLogGroups(context.TODO(), &cloudwatchlogs.DescribeLogGroupsInput{LogGroupNamePattern: &pattern})
	if err != nil {
		panic(err)
	}

	var arr []string
	for _, group := range res.LogGroups {
		arr = append(arr, *group.LogGroupName)
	}

	return arr
}

func (f *AwsFetcher) FetchStreams(groupName string) []string {
	res, err := f.svc.DescribeLogStreams(context.TODO(), &cloudwatchlogs.DescribeLogStreamsInput{LogGroupName: &groupName})
	if err != nil {
		panic(err)
	}

	var arr []string
	for _, stream := range res.LogStreams {
		arr = append(arr, *stream.LogStreamName)
	}

	return arr
}

func (f *AwsFetcher) FetchEvents(groupName, streamName string) []string {
	res, err := f.svc.GetLogEvents(context.TODO(), &cloudwatchlogs.GetLogEventsInput{
		LogGroupName:  &groupName,
		LogStreamName: &streamName,
	})
	if err != nil {
		panic(err)
	}

	var arr []string
	for _, event := range res.Events {
		arr = append(arr, *event.Message)
	}

	return arr
}
