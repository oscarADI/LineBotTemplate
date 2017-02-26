// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}
	var t0,t1 int
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			//for n:=0;;n++{
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if(message.Text == "start") {t0 := time.Now()}
				if(message.Text == "end") {
					t1 := time.Now()
					//var n int:= t1-t0
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(strconv.Itoa(t1-t0))).Do(); err != nil {
					log.Print(message.Text)}
					break
				}
			}
		}
	}
}


