package db

type notification struct {
	ID        int
	SessionID string
	Message   string
	Send      bool
}

// cache like Redis or go-cache
type storage []notification

var db storage

func GetDB() storage {
	return db
}

func setupDB() {
	db = []notification{}
}

func closeDB() {
	// close db here
}

// my implementation of db request like 'select * from table_name where send <> false'
func (s storage) ReadAllUnSendMessages() []notification {
	var resp []notification
	for id, n := range s {
		if !n.Send {
			n.ID = id
			resp = append(resp, n)
		}
	}
	return resp
}

func (s storage) Create(mes, sID string) {
	db = append(db, notification{Message: mes, SessionID: sID})
}

func (s storage) Read() []notification {
	return db
}

func (s storage) Update(id int, ok bool) {
	db[id] = notification{Message: s[id].Message, Send: ok}
}

func (s storage) Delete(id int) {
	db = append(db[:id], db[id+1:]...)
}
