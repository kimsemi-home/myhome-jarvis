package financeconnector

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

const myDataCardsPath = "/v2/card/cards"

type myDataCardListResponse struct {
	RspCode  string       `json:"rsp_code"`
	RspMsg   string       `json:"rsp_msg"`
	NextPage string       `json:"next_page"`
	CardList []myDataCard `json:"card_list"`
}

type myDataCard struct {
	CardID    string `json:"card_id"`
	IsConsent bool   `json:"is_consent"`
}

func (client LiveClient) discoverCards(
	ctx context.Context,
	connection ConnectionConfig,
	token string,
) ([]string, error) {
	var cards []string
	nextPage := ""
	for page := 0; page < 100; page++ {
		query := url.Values{
			"org_code": {connection.OrgCode}, "search_timestamp": {"0"},
			"limit": {"500"},
		}
		if nextPage != "" {
			query.Del("search_timestamp")
			query.Set("next_page", nextPage)
		}
		data, err := client.doMyDataJSON(ctx, http.MethodGet, myDataCardsPath, query, token, nil)
		if err != nil {
			return nil, err
		}
		var response myDataCardListResponse
		if err := json.Unmarshal(data, &response); err != nil {
			return nil, err
		}
		if err := validateMyDataResponse(response.RspCode, response.RspMsg); err != nil {
			return nil, err
		}
		for _, card := range response.CardList {
			if card.IsConsent && card.CardID != "" {
				cards = append(cards, card.CardID)
			}
		}
		if response.NextPage == "" {
			break
		}
		nextPage = response.NextPage
	}
	return cards, nil
}
