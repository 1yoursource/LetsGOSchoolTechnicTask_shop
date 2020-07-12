package helpers

import (
	"github.com/night-codes/mgo-ai"
	"gopkg.in/mgo.v2"
)

func Counter() {
	sess, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	ai.Connect(sess.DB("example-db").C("counters"))
}
