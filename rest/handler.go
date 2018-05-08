package rest

import (
	"encoding/json"
	"fmt"
	"github.com/lsavouillannxw/hockey-kids-lines/service"
	"net/http"
	"strconv"
	"strings"
	"os"
	"io/ioutil"
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
		return
	}
	numberOfPlayersPerLine, err := strconv.Atoi(r.URL.Query().Get("numberOfPlayersPerLine"))
	if err != nil {
		w.WriteHeader(400)
		return
	}
	numberOfLinesPerMatch, err := strconv.Atoi(r.URL.Query().Get("numberOfLinesPerMatch"))
	if err != nil {
		w.WriteHeader(400)
		return
	}
	if numberOfPlayers > 16 || numberOfPlayers < 7 || numberOfPlayersPerLine < 3 || numberOfPlayersPerLine > 5 || numberOfLinesPerMatch < 5 || numberOfLinesPerMatch > 16 || numberOfPlayers%numberOfPlayersPerLine == 0 {
		w.WriteHeader(400)
		return
	}
	filename := fmt.Sprintf("results/%dp-%dl-%ds.json", numberOfPlayers, numberOfLinesPerMatch, numberOfPlayersPerLine)
	var resAsBytes []byte
	if _, err := os.Stat(filename); err == nil {
		resAsBytes, err = ioutil.ReadFile(filename)
	} else {
		res := service.NewProcessingResult(*service.NewProcessingHandler(numberOfPlayers, numberOfLinesPerMatch, numberOfPlayersPerLine).Process())
		resAsBytes, err = json.Marshal(res)
	}
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, string(resAsBytes))
}
