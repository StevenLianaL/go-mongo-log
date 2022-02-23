package mongo_log

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
		function: GetFuncName(),
		created:  time.Now(),
		user:     user,
		project:  m.Project,
		app:      m.App,
	}
	_, _ = m.Collection.InsertOne(*m.Ctx, bson.D{
		{"log", recorder.log},
		{"level", recorder.level},
		{"func", recorder.function},
		{"created", recorder.created},
		{"user", recorder.user},
		{"project", recorder.project},
		{"app", recorder.app},
	})
}

func (m *Manager) Info(msg string, user int) {
	m.Log(msg, INFO, user)
}

func (m *Manager) Debug(msg string, user int) {
	m.Log(msg, DEBUG, user)
}

func (m *Manager) Warning(msg string, user int) {
	m.Log(msg, WARNING, user)
}

func (m *Manager) Error(msg string, user int) {
	m.Log(msg, ERROR, user)
}
