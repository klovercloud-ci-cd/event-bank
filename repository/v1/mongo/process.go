package mongo

import (
	"context"
	"github.com/klovercloud-ci/core/v1/repository"
	v1 "github.com/klovercloud-ci/core/v1"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

var (
	ProcessCollection="processCollection"
)

type processRepository struct {
	manager  *dmManager
	timeout  time.Duration
}

func (p processRepository) Store(process v1.Process) {
	coll := p.manager.Db.Collection(ProcessCollection)
	_, err := coll.InsertOne(p.manager.Ctx, process)
	if err != nil {
		log.Println("[ERROR] Insert document:", err.Error())
	}
}

func (p processRepository) GetByCompanyIdAndRepositoryIdAndAppName(companyId, repositoryId, appId string,option v1.ProcessQueryOption) []v1.Process {
	query := bson.M{
		"$and": []bson.M{
			{"company_id": companyId},
			{"app_id": appId},
			{"repository_id": repositoryId},
		},
	}
	coll := p.manager.Db.Collection(ProcessCollection)

	curser, err := coll.Find(p.manager.Ctx, query)
	if err!=nil{
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
		results= append(results, *elemValue)
	}
	return results
}

func NewProcessRepository(timeout int) repository.ProcessRepository {
	return &processRepository{
		manager: GetDmManager(),
		timeout: time.Duration(timeout),
	}

}
