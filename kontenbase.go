package kontenbase

import "github.com/drScripts/kontenbase-go-test/client"

type Client struct {
	apiKey   string
	QueryUrl string
	headers  map[string]string

	Auth     *client.AuthClient
	Storage  *client.StorageClient
	Realtime *client.RealtimeClient
}

const defaultDomain = "https://api.kontenbase.com"

func DefaultURL() string {
	return defaultDomain + "/query/api/v1"
}

func NewClient(apiKey string, url string) *Client {

	if url == "" {
		url = DefaultURL()
	}

	c := &Client{
		apiKey:   apiKey,
		headers:  map[string]string{},
		QueryUrl: url + "/" + apiKey,
	}

	c.Auth = client.NewAuthClient(c.QueryUrl+"/auth", c.headers)
	c.Storage = client.NewStorageClient(c.QueryUrl+"/storage", c.Auth)
	c.Realtime = client.NewRealtimeClient(defaultDomain+"/stream", c.apiKey, c.Auth)

	return c
}

func (c *Client) getHeaders() map[string]string {
	headers := c.headers
	authBearer := c.Auth.Token()

	if authBearer != "" {
		headers["Authorization"] = "Bearer " + authBearer
	}

	return headers
}

func (c *Client) Service(name string) *client.QueryClient {
	return client.NewQueryClient(c.QueryUrl+"/"+name, c.getHeaders())
}
