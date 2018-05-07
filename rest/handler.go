package rest

import (
	"encoding/json"
	"fmt"
	"github.com/lsavouillannxw/hockey-kids-lines/service"
	"net/http"
	"strconv"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, "/api") {
		http.Redirect(w, r, strings.Split(r.URL.String(), r.URL.Path)[0]+"/web/hockeyKidsLinesPage.html", http.StatusPermanentRedirect)
		return
	}

	if r.Method != http.MethodGet {
		w.WriteHeader(405)
		return
	}

	numberOfPlayers, err := strconv.Atoi(r.URL.Query().Get("numberOfPlayers"))
	if err != nil {
		w.WriteHeader(400)
	}
	numberOfPlayersPerLine, err := strconv.Atoi(r.URL.Query().Get("numberOfPlayersPerLine"))
	if err != nil {
		w.WriteHeader(400)
	}
	numberOfLinesPerMatch, err := strconv.Atoi(r.URL.Query().Get("numberOfLinesPerMatch"))
	if err != nil {
		w.WriteHeader(400)
	}
	if numberOfPlayers > 16 || numberOfPlayers < 7 || numberOfPlayersPerLine < 3 || numberOfPlayersPerLine > 5 || numberOfLinesPerMatch < 5 || numberOfLinesPerMatch > 16 || numberOfPlayers%numberOfPlayersPerLine == 0 {
		w.WriteHeader(400)
		return
	}
	res := service.NewProcessingResult(*service.NewProcessingHandler().Process(numberOfPlayers, numberOfLinesPerMatch, numberOfPlayersPerLine), numberOfPlayers, numberOfLinesPerMatch)
	resAsBytes, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, string(resAsBytes))
}
