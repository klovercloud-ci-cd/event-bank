package mongo

import (
	v1 "github.com/klovercloud-ci-cd/event-bank/core/v1"
	"github.com/klovercloud-ci-cd/event-bank/core/v1/repository"
	"github.com/twinj/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// ProcessEventCollection process event collection name
var (
	ProcessEventCollection = "processEventCollection"
)

type processEventRepository struct {
	manager *dmManager
	timeout time.Duration
}

func (p processEventRepository) Store(data v1.PipelineProcessEvent) {
	data.Id = uuid.NewV4().String()
	coll := p.manager.Db.Collection(ProcessEventCollection)
	_, err := coll.InsertOne(p.manager.Ctx, data)
	if err != nil {
		log.Println("[ERROR] Insert document:", err.Error())
	}
}

func (p processEventRepository) GetByCompanyId(companyId string) map[string]interface{} {
	var processEvent = new(v1.PipelineProcessEvent)
	query := bson.M{
		"$and": []bson.M{
			{"company_id": companyId},
		},
	}
	coll := p.manager.Db.Collection(ProcessEventCollection)
	result := coll.FindOne(p.manager.Ctx, query)
	err := result.Decode(&processEvent)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return processEvent.Data
}

func (p processEventRepository) DequeueByCompanyId(companyId string) map[string]interface{} {
	var processEvent = new(v1.PipelineProcessEvent)
	query := bson.M{
		"$and": []bson.M{
			{"company_id": companyId},
			{"read": false},
		},
	}
	coll := p.manager.Db.Collection(ProcessEventCollection)
	result := coll.FindOne(p.manager.Ctx, query)
	err := result.Decode(&processEvent)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	filter := bson.M{
		"$and": []bson.M{
			{"company_id": companyId},
			{"id": processEvent.Id},
		},
	}
	updatedProcessEvent := v1.PipelineProcessEvent{
		Id:        processEvent.Id,
		ProcessId: processEvent.ProcessId,
		CompanyId: processEvent.CompanyId,
		Read:      true,
		Data:      processEvent.Data,
	}
	update := bson.M{
		"$set": updatedProcessEvent,
	}
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	result = coll.FindOneAndUpdate(p.manager.Ctx, filter, update, &opt)
	if result.Err() != nil {
		log.Println("[ERROR]", result.Err().Error())
	}
	return processEvent.Data
}

// NewProcessEventRepository returns ProcessEventRepository type object
func NewProcessEventRepository(timeout int) repository.ProcessEventRepository {
	return &processEventRepository{
		manager: GetDmManager(),
		timeout: time.Duration(timeout),
	}
}
