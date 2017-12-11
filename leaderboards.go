package main

import (
	"fmt"
	"log"
	"net/http"
	_ "io/ioutil"
	"encoding/json"
)

type Leaderboard struct {
	TimePosted int `json:"time_posted"`
	NextPostTime int `json:"next_scheduled_post_time"`
	ServerTime int `json:"server_time"`
	Players []Player `json:"leaderboard"`
}

//{"rank":1,"name":"SmAsH","team_id":5154535,"team_tag":"\u2122PERU","country":"pe","sponsor":"\u2764q","solo_mmr":null}
type Player struct {
	Name string `json:"name"`
	Rank int `json:"rank"`
	Team int `json:"team_id"`
	TeamTag string `json:"team_tag"`
	Country string `json:"country"`
	Sponsor string `json:"sponsor"`
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	var lb Leaderboard

	division := r.URL.Query().Get("division")
	res, err := http.Get("http://www.dota2.com/webapi/ILeaderboard/GetDivisionLeaderboard/v0001?division="+division)
	if err != nil {
		panic(err)
	}
	log.Println("Scrapeo correcto.")
	defer res.Body.Close()

	//"<h1>Dota2 Leaderboards</h1><p>A mi web.</p>"
	//robots, _ := ioutil.ReadAll(res.Body)
	if json.NewDecoder(res.Body).Decode(&lb) != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	ranking := ""
	for index, player := range lb.Players {
		ranking += fmt.Sprintf("Player: [%d] %s.%s.%s \n", index + 1, player.TeamTag, player.Name, player.Sponsor)
	}
	fmt.Fprintf(w, "Fecha del Servidor es: [%d]\n%s", lb.ServerTime, string(ranking))
}

func main() {
	log.Println("Iniciando el scrapeo...")


	http.HandleFunc("/", mainHandler)
	http.ListenAndServe(":8080", nil)
}