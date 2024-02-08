package scheduler

import (
	"card_notifier/pkg/db"
	"fmt"
	"log"
)

// job print client notification to terminal
func Notifier() {
	notifications := db.GetDB().ReadAllUnSendMessages()
	for _, n := range notifications {
		log.Println(fmt.Sprintf(`{client:"%s", notification:"%s"}`, n.SessionID, n.Message))
		db.GetDB().Update(n.ID, true)
	}
}
