package main

import (
	"fmt"

	"github.com/alecthomas/kingpin"
)

type Account struct {
	AvatarURL  interface{} `json:"avatar_url"`
	ID         int         `json:"id"`
	Login      string      `json:"login"`
	Name       string      `json:"name"`
	ReposCount int         `json:"repos_count"`
	Type       string      `json:"type"`
	Subscribed bool        `json:"subscribed"`
}

var accountsCommand = kingpin.Command("accounts", "displays accounts and their subscription status").Action(func(ctx *kingpin.ParseContext) error {
	err := auth()
	if err != nil {
		return err
	}

	u, _, err := client.Users.GetAuthenticated()
	if err != nil {
		return err
	}

	req, err := client.NewRequest("GET", "/accounts?all=true", nil, nil)
	if err != nil {
		return err
	}

	var res struct {
		Accounts []Account `json:"accounts"`
	}
	resp, err := client.Do(req, &res)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	s := "not subscribed"
	for _, account := range res.Accounts {
		if account.Name != u.Name {
			s = "subscribed"
			break
		}
	}

	for _, account := range res.Accounts {
		fmt.Printf("%s (%s): %s, %d repositories\n", account.Login, account.Name, s, account.ReposCount)
	}

	return nil
})
