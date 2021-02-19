package youtube

import (
	"strings"
	"time"
	"fmt"
	"context"
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type YouTube struct {
	Channels	[]string
	apikey    string
	youtubeService *youtube.Service
}

const sampleConfig = `
  ## List of channels to monitor.
  channels = [
    "UCBR8-60-B28hp2BmDPdntcQ",
    "UCnrgOD6G0y0_rcubQuICpTQ"
  ]
`

func (y *YouTube) SampleConfig() string {
	return sampleConfig
}

func (y *YouTube) Description() string {
	return "Gets channel subs,views,VideoCount from YouTube channels."
}

//To use an API key for authentication (note: some APIs do not support API keys), use option.WithAPIKey:
func (y *YouTube) createYouTubeService(ctx context.Context) (*youtube.Service, error) {
	//youtubeService, err := youtube.NewService(ctx, option.WithAPIKey())
	//return youtube.NewService(ctx, option.WithAPIKey(y.apikey))
	return youtube.NewService(ctx, option.WithAPIKey(y.apikey))
}

func (y *YouTube) Gather(acc telegraf.Accumulator) error {

	ctx := context.Background()
	youtubeService, err := youtube.NewService(ctx)
	if err != nil {
		return err
	}
	//https://developers.google.com/youtube/v3/docs/channels/list
	//call := youtubeService.Channels.List([]string{"snippet", "statistics"})
	list := []string{"snippet", "statistics"}
	call := youtubeService.Channels.List(strings.Join(list,",")).Id(strings.Join(y.Channels, ",")).MaxResults(50)
	resp, err := call.Do()
	if err != nil {
		return err
	}
	now := time.Now()


	for _, item := range resp.Items {
		tags := getTags(item)
		fields := getFields(item)
		acc.AddFields("youtube_channel", fields, tags, now)

	}

	return nil
}

func getTags(channelInfo *youtube.Channel) map[string]string {

	fmt.Println("id: ", channelInfo.Id, "title: ",channelInfo.Snippet.Title)
	tags := make(map[string]string)
	tags["id"] = channelInfo.Id
	tags["title"] = channelInfo.Snippet.Title
	return tags
	//return map[string]string{
		//"id":    channelInfo.Id,
		//"title": channelInfo.Snippet.Title,
		//}
}

func getFields(channelInfo *youtube.Channel) map[string]interface{} {

	fmt.Println("subscount:", channelInfo.Statistics.SubscriberCount)
	fields := make(map[string]interface{})
	fields["subscribers"] = channelInfo.Statistics.SubscriberCount
	fields["videos"] = channelInfo.Statistics.VideoCount
	fields["views"] = channelInfo.Statistics.ViewCount
	return fields
	//return map[string]interface{}{
		//"subscribers": ,
		//"videos":      channelInfo.Statistics.VideoCount,
		//"views":       channelInfo.Statistics.ViewCount,
		//}
}

func init() {
	inputs.Add("youtube", func() telegraf.Input {
		return &YouTube{apikey: "AIzaSyBBqZX1S8ZfJ6jFkzY9eU7JiQijJLZ1-Ow"}
	})
}
