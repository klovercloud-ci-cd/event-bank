package logic

import (
	"fmt"
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/repository"
	"github.com/klovercloud-ci/core/v1/service"
	"github.com/klovercloud-ci/enums"
	"strings"
)

type processLifeCycleEventService struct {
	repo repository.ProcessLifeCycleEventRepository
}

func (p processLifeCycleEventService) PullNonInitializedAndAutoTriggerEnabledEventsByStepType(count int64, stepType string) []v1.ProcessLifeCycleEvent {
	return p.repo.PullNonInitializedAndAutoTriggerEnabledEventsByStepType(count,stepType)
}

func (p processLifeCycleEventService) PullPausedAndAutoTriggerEnabledResourcesByAgentName(count int64, agent string) []v1.Resource {
	resources:=[]v1.Resource{}
	events:=p.repo.PullPausedAndAutoTriggerEnabledResourcesByAgentName(count,agent)
	for _,event:=range events{

		var step *v1.Step
		for _,each:=range event.Pipeline.Steps{
			if each.Name==event.Step{
				step=&each
				break
			}
		}
		if step!=nil{
			resources=append(resources,v1.Resource{
				Step:        step.Name,
				ProcessId:   event.ProcessId,
				Descriptors: step.Descriptors,
				Type:        enums.PIPELINE_RESOURCE_TYPE(step.Params["type"]),
				Name:        step.Params["name"],
				Namespace:   step.Params["namespace"],
				Images:     strings.Split(fmt.Sprintf("%v", step.Params["images"]), ",") ,
			} )
		}
	}
	return resources
}

func (p processLifeCycleEventService) Store(events []v1.ProcessLifeCycleEvent) {
	p.repo.Store(events)
}

func NewProcessLifeCycleEventService(repo repository.ProcessLifeCycleEventRepository) service.ProcessLifeCycleEvent {
	return &processLifeCycleEventService{
		repo: repo,
	}
}
