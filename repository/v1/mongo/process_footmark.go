package mongo

import (
	"context"
	v1 "github.com/klovercloud-ci-cd/event-bank/core/v1"
	"github.com/klovercloud-ci-cd/event-bank/core/v1/repository"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

// ProcessFootmarkCollection process footmark collection name
var (
	ProcessFootmarkCollection = "processFootmarkCollection"
)


type processFootmarkRepository struct {
	manager *dmManager
	timeout time.Duration
}

func (l processFootmarkRepository) GetByProcessIdAndStep(processId, step string) []v1.ProcessFootmark {
	query := bson.M{
		"$and": []bson.M{
			{"process_id": processId},
			{"step": step},
		},
	}
	coll := l.manager.Db.Collection(ProcessFootmarkCollection)
	curser, err := coll.Find(l.manager.Ctx, query)
	if err != nil {
		log.Println(err.Error())
	}
	var results []v1.ProcessFootmark
	for curser.Next(context.TODO()) {
		elemValue := new(v1.ProcessFootmark)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		results = append(results, *elemValue)
	}
	return results
}

func (l processFootmarkRepository) GetByProcessId(processId string) []v1.ProcessFootmark {
	query := bson.M{
		"$and": []bson.M{
			{"process_id": processId},
		},
	}
	coll := l.manager.Db.Collection(ProcessFootmarkCollection)
	curser, err := coll.Find(l.manager.Ctx, query)
	if err != nil {
		log.Println(err.Error())
	}
	var results []v1.ProcessFootmark
	for curser.Next(context.TODO()) {
		elemValue := new(v1.ProcessFootmark)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		results = append(results, *elemValue)
	}
	return results
}

func (l processFootmarkRepository) GetFootmarkByProcessIdAndStepAndFootmark(processId,step,footmark string) *v1.ProcessFootmark{
	query := bson.M{
		"$and": []bson.M{
			{"process_id": processId},
			{"step": step},
			{"footmark":footmark},
		},
	}
	coll := l.manager.Db.Collection(ProcessFootmarkCollection)
	curser, err := coll.Find(l.manager.Ctx, query)
	if err != nil {
		log.Println(err.Error())
	}
	for curser.Next(context.TODO()) {
		elemValue := new(v1.ProcessFootmark)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		return elemValue
	}
	return nil
}
func (l processFootmarkRepository) Store(data v1.ProcessFootmark) {
	coll := l.manager.Db.Collection(ProcessFootmarkCollection)
	_, err := coll.InsertOne(l.manager.Ctx, data)
	if err != nil {
		log.Println("[ERROR] Insert document:", err.Error())
	}

}
// NewProcessFootmarkRepository returns ProcessLifeCycleEventRepository type object
func NewProcessFootmarkRepository(timeout int) repository.ProcessFootmarkRepository {
	return &processFootmarkRepository{
		manager: GetDmManager(),
		timeout: time.Duration(timeout),
	}

}