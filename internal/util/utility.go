package util

import (
	"log"
	"net/http"
	"strconv"
)

func CheckErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func ExtractCampaignId(r *http.Request) (int, error) {
	return strconv.Atoi(r.PathValue("campaignId"))
}
func ExtractAdventureId(r *http.Request) (int, error) {
	return strconv.Atoi(r.PathValue("adventureId"))
}
func ExtractUserId(r *http.Request) (int, error) {
	return strconv.Atoi(r.PathValue("userId"))
}
