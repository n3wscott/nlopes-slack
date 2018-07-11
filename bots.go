package slack

import (
	"context"
<<<<<<< HEAD
=======
	"errors"
	"log"
>>>>>>> testing.
	"net/url"
)

// Bot contains information about a bot
type Bot struct {
<<<<<<< HEAD
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Deleted bool     `json:"deleted"`
	UserID  string   `json:"user_id"`
	AppID   string   `json:"app_id"`
	Updated JSONTime `json:"updated"`
	Icons   Icons    `json:"icons"`
=======
	ID      string `json:"id"`
	AppID   string `json:"app_id"`
	Name    string `json:"name"`
	Deleted bool   `json:"deleted"`
	Icons   Icons  `json:"icons"`
>>>>>>> testing.
}

type botResponseFull struct {
	Bot `json:"bot,omitempty"` // GetBotInfo
	SlackResponse
}

func (api *Client) botRequest(ctx context.Context, path string, values url.Values) (*botResponseFull, error) {
	response := &botResponseFull{}
	err := api.postMethod(ctx, path, values, response)
	if err != nil {
		return nil, err
	}

	if err := response.Err(); err != nil {
		return nil, err
	}

	return response, nil
}

// GetBotInfo will retrieve the complete bot information
func (api *Client) GetBotInfo(bot *string) (*Bot, error) {
	return api.GetBotInfoContext(context.Background(), bot)
}

// GetBotInfoContext will retrieve the complete bot information using a custom context
func (api *Client) GetBotInfoContext(ctx context.Context, bot *string) (*Bot, error) {
	values := url.Values{
		"token": {api.token},
	}
	if bot != nil {
		values.Set("bot", *bot)
<<<<<<< HEAD
=======
	}
	log.Printf("values: %+v", values)

	if bot != "" {
		values.Add("bot", bot)
>>>>>>> testing.
	}

	response, err := api.botRequest(ctx, "bots.info", values)
	if err != nil {
		return nil, err
	}
	log.Printf("BotResponse: %+v", response)
	return &response.Bot, nil
}
