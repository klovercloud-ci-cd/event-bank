package mongo

import (
	"context"
	v1 "github.com/klovercloud-ci-cd/event-bank/core/v1"
	"github.com/klovercloud-ci-cd/event-bank/core/v1/repository"
	"github.com/klovercloud-ci-cd/event-bank/enums"
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

func (p processEventRepository) GetByCompanyIdAndProcessId(companyId, processId string, option v1.ProcessQueryOption) ([]v1.PipelineProcessEvent, int64) {
	var results []v1.PipelineProcessEvent
	var query bson.M
	if processId == "" {
		query = bson.M{
			"$and": []bson.M{
				{"company_id": companyId},
			},
			"$or": []bson.M{
				{"data.status": string(enums.PROCESS_EVENT_INITIALIZING)},
				{"data.status": string(enums.PROCESS_EVENT_FAILED)},
				{"data.status": string(enums.PROCESS_EVENT_SUCCESSFUL)},
			},
		}
	} else {
		query = bson.M{
			"$and": []bson.M{
				{"company_id": companyId},
				{"process_id": processId},
			},
			"$or": []bson.M{
				{"data.status": string(enums.PROCESS_EVENT_INITIALIZING)},
				{"data.status": string(enums.PROCESS_EVENT_FAILED)},
				{"data.status": string(enums.PROCESS_EVENT_SUCCESSFUL)},
			},
		}
	}
	coll := p.manager.Db.Collection(ProcessEventCollection)
	skip := option.Pagination.Page * option.Pagination.Limit
	curser, err := coll.Find(p.manager.Ctx, query, &options.FindOptions{
		Limit: &option.Pagination.Limit,
		Skip:  &skip,
		Sort:  bson.M{"created_at": -1},
	})
	if err != nil {
		log.Println(err.Error())
	}
	for curser.Next(context.TODO()) {
		elemValue := new(v1.PipelineProcessEvent)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		results = append(results, *elemValue)
	}
	count, err := coll.CountDocuments(p.manager.Ctx, query)
	if err != nil {
		log.Println(err.Error())
	}
	return results, count
}

func (p processEventRepository) Store(data v1.PipelineProcessEvent) {
	data.Id = uuid.NewV4().String()
	data.CreatedAt = time.Now().UTC()
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
	findOneOptions := options.FindOneOptions{
		Sort: bson.M{"created_at": -1},
	}
	result := coll.FindOne(p.manager.Ctx, query, &findOneOptions)
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
	findOneOptions := options.FindOneOptions{
		Sort: bson.M{"created_at": -1},
	}
	result := coll.FindOne(p.manager.Ctx, query, &findOneOptions)
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
		CreatedAt: processEvent.CreatedAt,
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
