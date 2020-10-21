package messenger

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func GetMessengerList(w http.ResponseWriter, r *http.Request) {
	filepath, err := filepath.Abs("./internal/pkg/datasources/messenger.json")
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(file))
}
