package mongo

import (
	"context"
	v1 "github.com/klovercloud-ci-cd/event-bank/core/v1"
	"github.com/klovercloud-ci-cd/event-bank/core/v1/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// ProcessCollection process collection name
var (
	ProcessCollection = "processCollection"
)

type processRepository struct {
	manager *dmManager
	timeout time.Duration
}

func (p processRepository) GetById(companyId, processId string) v1.Process {
	var process = new(v1.Process)
	query := bson.M{
		"$and": []bson.M{
			{"process_id": processId},
			{"company_id": companyId},
		},
	}
	coll := p.manager.Db.Collection(ProcessCollection)
	result := coll.FindOne(p.manager.Ctx, query)
	err := result.Decode(&process)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return *process
}

func (p processRepository) CountTodaysRanProcessByCompanyId(companyId string) int64 {
	time.Local = time.UTC
	t := time.Now()
	fromDate := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	toDate := fromDate.AddDate(0, 0, 0).Add(time.Hour * 23).Add(time.Minute * 59).Add(time.Second * 59)
	query := bson.M{
		"$and": []bson.M{
			{"company_id": companyId},
			{
				"created_at": bson.M{
					"$gte": fromDate,
					"$lte": toDate,
				},
			},
		},
	}
	total, err := p.manager.Db.Collection(ProcessCollection).CountDocuments(p.manager.Ctx, query)
	if err != nil {
		log.Println(err.Error())
	}
	return total
}

func (p processRepository) Store(process v1.Process) {
	coll := p.manager.Db.Collection(ProcessCollection)
	_, err := coll.InsertOne(p.manager.Ctx, process)
	if err != nil {
		log.Println("[ERROR] Insert document:", err.Error())
	}
}

func (p processRepository) GetByCompanyIdAndRepositoryIdAndAppName(companyId, repositoryId, appId string, option v1.ProcessQueryOption) ([]v1.Process, int64) {
	query := bson.M{
		"$and": []bson.M{},
	}
	and := []bson.M{
		{"company_id": companyId},
		{"app_id": appId},
		{"repository_id": repositoryId},
	}
	if option.Step != "" {
		and = append(and, map[string]interface{}{"step": option.Step})
	}
	query["$and"] = and
	coll := p.manager.Db.Collection(ProcessCollection)
	count, err := coll.CountDocuments(p.manager.Ctx, query)
	if err != nil {
		log.Println(err.Error())
	}
	skip := option.Pagination.Page * option.Pagination.Limit

	curser, err := coll.Find(p.manager.Ctx, query, &options.FindOptions{
		Limit: &option.Pagination.Limit,
		Skip:  &skip,
		Sort:  bson.M{"created_at": -1},
	})
	if err != nil {
		log.Println(err.Error())
	}
	var results []v1.Process
	for curser.Next(context.TODO()) {
		elemValue := new(v1.Process)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		results = append(results, *elemValue)
	}
	return results, count
}

func (p processRepository) GetByCompanyIdAndCommitId(companyId, commitId string, option v1.ProcessQueryOption) ([]v1.Process, int64) {
	query := bson.M{
		"$and": []bson.M{},
	}
	and := []bson.M{
		{"company_id": companyId},
		{"commit_id": commitId},
	}
	if option.Step != "" {
		and = append(and, map[string]interface{}{"step": option.Step})
	}
	query["$and"] = and
	coll := p.manager.Db.Collection(ProcessCollection)
	count, err := coll.CountDocuments(p.manager.Ctx, query)
	if err != nil {
		log.Println(err.Error())
	}
	skip := option.Pagination.Page * option.Pagination.Limit

	curser, err := coll.Find(p.manager.Ctx, query, &options.FindOptions{
		Limit: &option.Pagination.Limit,
		Skip:  &skip,
		Sort:  bson.M{"created_at": -1},
	})
	if err != nil {
		log.Println(err.Error())
	}
	var results []v1.Process
	for curser.Next(context.TODO()) {
		elemValue := new(v1.Process)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		results = append(results, *elemValue)
	}
	return results, count
}

// NewProcessRepository returns ProcessRepository type object
func NewProcessRepository(timeout int) repository.ProcessRepository {
	return &processRepository{
		manager: GetDmManager(),
		timeout: time.Duration(timeout),
	}

}
