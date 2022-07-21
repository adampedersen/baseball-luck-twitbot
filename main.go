package main
	//must set GOTWI_API_KEY and GOTWI_API_KEY_SECRET 
	//environment variables for gotwi API
import (
    "fmt"
	"os"
	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/managetweet"
	"github.com/michimani/gotwi/tweet/managetweet/types"

	"github.com/adampedersen/luck-factor-bot/pkg/stats"
	
)

func main() {
	playerName := "Cody Bellinger"
	str := stats.GetLuckRating(playerName)
	fmt.Println(str)
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("TWITTER_ACCESS_SECRET")
	
	in := &gotwi.NewClientInput{
		AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
		OAuthToken:           accessToken,
		OAuthTokenSecret:     accessTokenSecret,
	}

	oauth1Client, err := gotwi.NewClient(in)
	if err != nil {
		panic(err)
	}

	tweetID, err := StringOnlyTweet(oauth1Client, str)
	if err != nil {
		panic(err)
	}

	fmt.Println("Posted successfully. ID: ", tweetID)

}

func StringOnlyTweet(c *gotwi.Client, text string) (string, error) {
	p := &types.CreateInput{
		Text: gotwi.String(text),
	}

	res, err := managetweet.Create(context.Background(), c, p)
	if err != nil {
		return "", err
	}

	return gotwi.StringValue(res.Data.ID), nil
}



