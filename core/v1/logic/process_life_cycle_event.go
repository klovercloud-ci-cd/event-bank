package logic

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/klovercloud-ci-cd/event-bank/api/common"
	"github.com/klovercloud-ci-cd/event-bank/config"
	v1 "github.com/klovercloud-ci-cd/event-bank/core/v1"
	"github.com/klovercloud-ci-cd/event-bank/core/v1/repository"
	"github.com/klovercloud-ci-cd/event-bank/core/v1/service"
	"github.com/klovercloud-ci-cd/event-bank/enums"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type processLifeCycleEventService struct {
	repo         repository.ProcessLifeCycleEventRepository
	processEvent service.ProcessEvent
	httpClient   service.HttpClient
}

func (p processLifeCycleEventService) GetByProcessIdAndStep(processId, step string) v1.ProcessLifeCycleEvent {
	return p.repo.GetByProcessIdAndStep(processId, step)
}

func (p processLifeCycleEventService) UpdateStatusesByTime(time time.Time) {
	err := p.repo.UpdateStatusesByTime(time)
	if err != nil {
		log.Println(err.Error())
	}
	processLifeCycleEvents, err := p.repo.GetByTime(time)
	for _, each := range processLifeCycleEvents {
		if each.Pipeline == nil {
			continue
		}
		processEvent := v1.PipelineProcessEvent{
			ProcessId: each.ProcessId,
			CompanyId: each.Pipeline.MetaData.CompanyId,
			Data: map[string]interface{}{
				"log":       "forcefully terminated",
				"type":      each.Step,
				"claim":     each.Claim,
				"companyId": each.Pipeline.MetaData.CompanyId,
				"processId": each.ProcessId,
				"status":    enums.FAILED,
			},
			CreatedAt: time,
		}
		if each.StepType == enums.BUILD {
			processEvent.Data["footmark"] = enums.POST_BUILD_JOB
		} else if each.StepType == enums.INTERMEDIARY {
			processEvent.Data["footmark"] = enums.POST_INTERMEDIARY_JOB
		} else if each.StepType == enums.JENKIN {
			processEvent.Data["footmark"] = enums.POST_JENKINS_JOB
		} else {
			processEvent.Data["footmark"] = enums.POST_DEPLOY_JOB
		}
		p.processEvent.Store(processEvent)
	}
}

func (p processLifeCycleEventService) UpdateClaim(companyId, processId, step, status string) error {
	if processId == "" {
		return errors.New("processId cannot be empty")
	}
	process := p.repo.GetByProcessIdAndStep(processId, step)
	if process.StepType != enums.DEPLOY {
		if process.Status == enums.ACTIVE {
			response := common.ResponseDTO{}
			reclaimAbility := v1.ReclaimAbleStatus{}
			url := config.CiCoreBaseUrl + "/pipelines/" + processId + "/steps" + "?question=IfStepIsClaimable&name=" + step
			header := make(map[string]string)
			header["Content-Type"] = "application/json"
			bytes, err := p.httpClient.Get(url, header)
			if err != nil {
				log.Println(err.Error())
				return err
			}
			err = json.Unmarshal(bytes, &response)
			if err != nil {
				log.Println(err.Error())
				return err
			}
			reclaimAbility.ClaimAble = reflect.ValueOf(response.Data).Bool()
			if reclaimAbility.ClaimAble == true {
				return p.repo.UpdateClaim(companyId, processId, step, status)
			} else {
				return errors.New("process is not reclaimable")
			}
		} else {
			return p.repo.UpdateClaim(companyId, processId, step, status)
		}
	}
	return p.repo.UpdateClaim(companyId, processId, step, status)
}

func (p processLifeCycleEventService) GetByProcessId(processId string) []v1.ProcessLifeCycleEvent {
	return p.repo.GetByProcessId(processId)
}

func (p processLifeCycleEventService) PullQueuedAndAutoTriggerEnabledEventsByStepType(count int64, stepType string) []v1.ProcessLifeCycleEvent {
	return p.repo.PullNonInitializedAndAutoTriggerEnabledEventsByStepType(count, stepType)
}

func (p processLifeCycleEventService) PullPausedAndAutoTriggerEnabledResourcesByAgentName(count int64, agent string) []v1.DeployableResource {
	resources := []v1.DeployableResource{}
	events := p.repo.PullPausedAndAutoTriggerEnabledResourcesByAgentName(count, agent)
	for _, event := range events {

		var step *v1.Step
		for _, each := range event.Pipeline.Steps {
			if each.Name == event.Step {
				step = &each
				break
			}
		}
		if step != nil {
			log.Println(step.Params["rollout_restart"])
			rolloutRestart, _ := strconv.ParseBool(step.Params["rollout_restart"])
			resources = append(resources, v1.DeployableResource{
				Step:           step.Name,
				ProcessId:      event.ProcessId,
				Descriptors:    step.Descriptors,
				Type:           enums.PIPELINE_RESOURCE_TYPE(step.Params["type"]),
				Name:           step.Params["name"],
				Namespace:      step.Params["namespace"],
				Images:         strings.Split(fmt.Sprintf("%v", step.Params["images"]), ","),
				Pipeline:       event.Pipeline,
				Claim:          event.Claim,
				RolloutRestart: rolloutRestart,
			})
		}
	}
	return resources
}

func (p processLifeCycleEventService) Store(events []v1.ProcessLifeCycleEvent) {
	p.repo.Store(events)
}

// NewProcessLifeCycleEventService returns ProcessLifeCycleEvent type service
func NewProcessLifeCycleEventService(repo repository.ProcessLifeCycleEventRepository, processEvent service.ProcessEvent, httpClient service.HttpClient) service.ProcessLifeCycleEvent {
	return &processLifeCycleEventService{
		repo:         repo,
		processEvent: processEvent,
		httpClient:   httpClient,
	}
}
