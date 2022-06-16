package mongo

import (
	"context"
	v1 "github.com/klovercloud-ci-cd/event-bank/core/v1"
	"github.com/klovercloud-ci-cd/event-bank/core/v1/repository"
	"github.com/klovercloud-ci-cd/event-bank/enums"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// ProcessLifeCycleCollection process life cycle event collection name
var (
	ProcessLifeCycleCollection = "processLifeCycleEventCollection"
)

type processLifeCycleRepository struct {
	manager *dmManager
	timeout time.Duration
}

func (p processLifeCycleRepository) GetByTime(time time.Time) ([]v1.ProcessLifeCycleEvent, error) {
	var data []v1.ProcessLifeCycleEvent
	query := bson.M{
		"$and": []bson.M{
			{"updated_at": bson.M{"$lte": time}},
		},
	}
	coll := p.manager.Db.Collection(ProcessLifeCycleCollection)
	result, err := coll.Find(p.manager.Ctx, query, nil)
	if err != nil {
		log.Println(err.Error())
		return []v1.ProcessLifeCycleEvent{}, err
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.ProcessLifeCycleEvent)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		data = append(data, *elemValue)
	}
	return data, nil
}

func (p processLifeCycleRepository) GetByCompanyId(companyId string, fromDate, toDate time.Time) []v1.ProcessLifeCycleEvent {
	var data []v1.ProcessLifeCycleEvent
	query := bson.M{
		"$and": []bson.M{
			{"updated_at": bson.M{"$gte": fromDate, "$lte": toDate}},
			{"pipeline.metadata.companyid": companyId},
		},
	}
	coll := p.manager.Db.Collection(ProcessLifeCycleCollection)
	result, err := coll.Find(p.manager.Ctx, query, nil)
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.ProcessLifeCycleEvent)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		data = append(data, *elemValue)
	}
	return data
}

func (p processLifeCycleRepository) UpdateStatusesByTime(time time.Time) error {
	query := bson.M{
		"$and": []bson.M{
			{"status": enums.ACTIVE},
			{"updated_at": bson.M{"$lte": time}},
		},
	}
	update := bson.M{
		"$set": bson.M{
			"status": enums.FAILED,
		},
	}
	coll := p.manager.Db.Collection(ProcessLifeCycleCollection)
	_, err := coll.UpdateMany(p.manager.Ctx, query, update)
	if err != nil {
		log.Println("[ERROR]", err)
		return err
	}
	return nil
}

func (p processLifeCycleRepository) UpdateClaim(processId, step, status string) error {
	data := p.GetByProcessIdAndStep(processId, step)
	data.Claim = data.Claim + 1
	data.Status = enums.PROCESS_STATUS(status)
	data.UpdatedAt = time.Now().UTC()
	data.ClaimedAt = time.Now().UTC()
	filter := bson.M{
		"$and": []bson.M{
			{"process_id": processId},
			{"step": step},
		},
	}
	update := bson.M{
		"$set": data,
	}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	coll := p.manager.Db.Collection(ProcessLifeCycleCollection)
	err := coll.FindOneAndUpdate(p.manager.Ctx, filter, update, &opt)
	if err.Err() != nil {
		log.Println("[ERROR]", err.Err())
		return err.Err()
	}
	return nil
}

func (p processLifeCycleRepository) PullNonInitializedAndAutoTriggerEnabledEventsByStepType(count int64, stepType string) []v1.ProcessLifeCycleEvent {
	var data []v1.ProcessLifeCycleEvent
	var query bson.M
	if stepType == string(enums.BUILD) {
		query = bson.M{
			"$and": []bson.M{
				{"status": enums.NON_INITIALIZED},
				{"trigger": enums.AUTO},
				{"step_type": stepType},
			},
		}
	} else {
		query = bson.M{
			"$and": []bson.M{
				{"status": enums.PAUSED},
				{"trigger": enums.AUTO},
				{"step_type": stepType},
			},
		}
	}

	coll := p.manager.Db.Collection(ProcessLifeCycleCollection)
	result, err := coll.Find(p.manager.Ctx, query, &options.FindOptions{
		Limit: &count,
	})
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.ProcessLifeCycleEvent)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		data = append(data, *elemValue)
	}
	for _, each := range data {
		each.ClaimedAt = time.Now().UTC()
		each.UpdatedAt = time.Now().UTC()
		go p.updateStatus(each, string(enums.ACTIVE))
	}
	return data
}

func (p processLifeCycleRepository) PullPausedAndAutoTriggerEnabledResourcesByAgentName(count int64, agent string) []v1.ProcessLifeCycleEvent {
	var data []v1.ProcessLifeCycleEvent
	query := bson.M{
		"$and": []bson.M{
			{"agent": agent},
			{"status": enums.PAUSED},
			{"trigger": enums.AUTO},
			{"step_type": enums.DEPLOY},
		},
	}
	coll := p.manager.Db.Collection(ProcessLifeCycleCollection)
	result, err := coll.Find(p.manager.Ctx, query, &options.FindOptions{
		Limit: &count,
	})
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.ProcessLifeCycleEvent)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		data = append(data, *elemValue)
	}
	for _, each := range data {
		each.ClaimedAt = time.Now().UTC()
		each.UpdatedAt = time.Now().UTC()
		go p.updateStatus(each, string(enums.ACTIVE))
	}
	return data
}

func (p processLifeCycleRepository) Get() []v1.ProcessLifeCycleEvent {
	var data []v1.ProcessLifeCycleEvent
	query := bson.M{}
	coll := p.manager.Db.Collection(ProcessLifeCycleCollection)
	result, err := coll.Find(p.manager.Ctx, query, nil)
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.ProcessLifeCycleEvent)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		data = append(data, *elemValue)
	}
	return data
}
func (p processLifeCycleRepository) Store(events []v1.ProcessLifeCycleEvent) {
	for i, each := range events {
		events[i].UpdatedAt = each.CreatedAt
	}
	coll := p.manager.Db.Collection(ProcessLifeCycleCollection)
	var pipeline *v1.Pipeline
	if len(events) > 0 {
		if events[0].StepType == enums.BUILD {
			pipeline = events[0].Pipeline
		}
	}
	for _, each := range events {
		existing := p.GetByProcessIdAndStep(each.ProcessId, each.Step)
		if existing == nil {
			each.CreatedAt = time.Now().UTC()
			each.ClaimedAt = time.Now().UTC()
			each.UpdatedAt = time.Now().UTC()
			each.Pipeline = pipeline
			_, err := coll.InsertOne(p.manager.Ctx, each)
			if err != nil {
				log.Println(err.Error())
			}
		} else {
			existing.ClaimedAt = time.Now().UTC()
			existing.UpdatedAt = time.Now().UTC()
			existing.Status = each.Status
			err := p.update(*existing)
			if err != nil {
				log.Println(err.Error())
			}
		}
	}
}
func (p processLifeCycleRepository) updateStatus(data v1.ProcessLifeCycleEvent, status string) error {
	data.UpdatedAt = time.Now().UTC()
	filter := bson.M{
		"$and": []bson.M{
			{"process_id": data.ProcessId},
			{"step": data.Step},
		},
	}
	data.Status = enums.PROCESS_STATUS(status)
	update := bson.M{
		"$set": data,
	}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	coll := p.manager.Db.Collection(ProcessLifeCycleCollection)
	err := coll.FindOneAndUpdate(p.manager.Ctx, filter, update, &opt)
	if err.Err() != nil {
		log.Println("[ERROR]", err.Err())
		return err.Err()
	}

	return nil
}
func (p processLifeCycleRepository) update(data v1.ProcessLifeCycleEvent) error {
	filter := bson.M{
		"$and": []bson.M{
			{"process_id": data.ProcessId},
			{"step": data.Step},
		},
	}
	update := bson.M{
		"$set": data,
	}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	coll := p.manager.Db.Collection(ProcessLifeCycleCollection)
	err := coll.FindOneAndUpdate(p.manager.Ctx, filter, update, &opt)
	if err.Err() != nil {
		log.Println("[ERROR]", err.Err())
		return err.Err()
	}

	return nil
}
func (p processLifeCycleRepository) GetByProcessIdAndStep(processId, step string) *v1.ProcessLifeCycleEvent {
	query := bson.M{
		"$and": []bson.M{
			{"process_id": processId},
			{"step": step},
		},
	}

	temp := new(v1.ProcessLifeCycleEvent)
	coll := p.manager.Db.Collection(ProcessLifeCycleCollection)
	result := coll.FindOne(p.manager.Ctx, query)
	err := result.Decode(temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	if temp.ProcessId == "" {
		return nil
	}
	return temp

}
func (p processLifeCycleRepository) GetByProcessId(processId string) []v1.ProcessLifeCycleEvent {
	query := bson.M{
		"$and": []bson.M{
			{"process_id": processId},
		},
	}
	coll := p.manager.Db.Collection(ProcessLifeCycleCollection)

	curser, err := coll.Find(p.manager.Ctx, query)
	if err != nil {
		log.Println(err.Error())
	}
	var results []v1.ProcessLifeCycleEvent
	for curser.Next(context.TODO()) {
		elemValue := new(v1.ProcessLifeCycleEvent)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		results = append(results, *elemValue)
	}
	return results
}

// NewProcessLifeCycleRepository returns ProcessLifeCycleEventRepository type object
func NewProcessLifeCycleRepository(timeout int) repository.ProcessLifeCycleEventRepository {
	return &processLifeCycleRepository{
		manager: GetDmManager(),
		timeout: time.Duration(timeout),
	}

}
