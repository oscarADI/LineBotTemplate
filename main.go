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

	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	sess := session.Must(session.NewSession())

svc := sqs.New(sess)

params := &sqs.ReceiveMessageInput{
    QueueUrl: aws.String("String"), // Required
    AttributeNames: []*string{
        aws.String("QueueAttributeName"), // Required
        // More values...
    },
    MaxNumberOfMessages: aws.Int64(1),
    MessageAttributeNames: []*string{
        aws.String("MessageAttributeName"), // Required
        // More values...
    },
    ReceiveRequestAttemptId: aws.String("String"),
    VisibilityTimeout:       aws.Int64(1),
    WaitTimeSeconds:         aws.Int64(1),
}
resp, err := svc.ReceiveMessage(params)

if err != nil {
    // Print the error, cast err to awserr.Error to get the Code and
    // Message from an error.
    fmt.Println(err.Error())
    return
}

// Pretty-print the response data.
fmt.Println(resp)
	
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

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text+"HAHAHARVEY")).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
	/*var messages []linebot.Message

    // append some message to messages
	
	messages = "HI"
    	_, err := bot.PushMessage(oscar1229, messages...).Do()
    	if err != nil {
		log.Print(err)
        // Do something when some bad happened
    	}*/
}
