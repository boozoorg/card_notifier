// in technical specifications wore wroten to use only standard library so that why i made my scheduler (if i had a choice i probable use gocron)
package scheduler

import (
	"card_notifier/config"
	"log"
	"time"
)

// own ticker to run job several times on set time
type timer struct {
	tickers []*time.Ticker
}

// set timer
func (t *timer) set() {
	var tmr *time.Ticker
	tmr = time.NewTicker(time.Second * time.Duration(5))
	t.tickers = append(t.tickers, tmr)
}

// Setup run jobs on goroutine on timers set time
func Setup() {
	log.Println("[.................jobs start..................]")
	var t timer
	t.set()

	// is to run job (for example if it is dev mode it will be false because of unneeded)
	toRun := make(map[string]bool)
	toRun["Notifier"] = config.Configs.ProdMode

	run := []bool{
		toRun["Notifier"],
	}

	for i, tt := range t.tickers {
		if run[i] {
			go func(i int, t *time.Ticker) {
				defer t.Stop()
				for {
					select {
					case <-t.C:
						go runJob(i)
					}
				}
			}(i, tt)
		}
	}
}

// run job by index of timers set time
func runJob(i int) {
	switch i {
	case 0:
		Notifier()
	}
}
