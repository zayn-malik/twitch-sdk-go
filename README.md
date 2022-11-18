# twitch-sdk-go

https://github.com/deepmap/oapi-codegen

# Example

```go
func main() {
   	var clientID = ""
	var clientSecret = ""

    bearerTokenProvider, err := twitch.NewBearerToken(clientID, clientSecret)
	if err != nil {
		log.Fatal(err)
	}

	// add Client-id in the HEADER
	client, err := twitch.NewClientWithResponses("https://api.twitch.tv/helix", twitch.WithRequestEditorFn(bearerTokenProvider.Intercept), twitch.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Add("Client-Id", clientID)
		return nil
	}))
	if err != nil {
		log.Fatal(err)
	}

    var login = "some login"
	resp, err := client.GetUserWithResponse(context.Background(), &twitch.GetUserParams{Login: &login})
	if err != nil {
		log.Fatal(err)
	}

    for _, u := range *resp.JSON200.Data {
		fmt.Printf("%s\n", *u.DisplayName)
		r, err := client.GetVideoWithResponse(context.Background(), &twitch.GetVideoParams{UserId: u.Id})
		if err != nil {
			log.Fatal(err)
		}

		for _, v := range *r.JSON200.Data {
			fmt.Println(*v.Title)
		}
	}

}
```
