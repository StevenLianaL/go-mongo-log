package mongo_log

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"mongo-log/helper"
	"time"
)

type Manager struct {
	Project    string
	App        string
	Collection *mongo.Collection
	Ctx        *context.Context
}
type Recorder struct {
	log      string
	level    string
	created  time.Time
	user     int
	project  string
	app      string
	function string
}

func (m *Manager) Log(msg string, level string, user int) {
	var recorder = Recorder{
		log:      msg,
		level:    level,
		function: helper.GetFuncName(),
		created:  time.Now(),
		user:     user,
		project:  m.Project,
		app:      m.App,
	}
	_, _ = m.Collection.InsertOne(*m.Ctx, bson.D{
		{"Log", recorder.log},
		{"level", recorder.level},
		{"function", recorder.function},
		{"created", recorder.created},
		{"user", recorder.user},
		{"Project", recorder.project},
		{"App", recorder.app},
	})
}
