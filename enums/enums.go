package enums

type STEP_TYPE string


const (
	BUILD=STEP_TYPE("BUILD")
	DEPLOY=STEP_TYPE("DEPLOY")
)

const (
	Mongo = "MONGO"
	Inmemory = "INMEMORY"
)

const (
	DEFAULT_PAGE_LIMIT = 10
	DEFAULT_PAGE       = 0
)
type PIPELINE_PURGING string
const (
	PIPELINE_PURGING_ENABLE      = PIPELINE_PURGING("ENABLE")
	PIPELINE_PURGING_DISABLE =PIPELINE_PURGING("DISABLE")
)


type TRIGGER string
const (
	AUTO      = TRIGGER("AUTO")
	MANUAL =TRIGGER("MANUAL")
)

type PARAMS string
const (
	REPOSITORY_TYPE      = PARAMS("repository_type")
	REVISION             =PARAMS("revision")
	SERVICE_ACCOUNT      =PARAMS("service_account")
	IMAGES               =PARAMS("images")
	ARGS_FROM_CONFIGMAPS =PARAMS("args_from_configmaps")
	ARGS_FROM_SECRETS    =PARAMS("args_from_secrets")
	ENVS_FROM_CONFIGMAPS =PARAMS("envs_from_configmaps")
	ENVS_FROM_SECRETS    =PARAMS("envs_from_secrets")
	ARGS                 =PARAMS("arg")
	ENVS                 =PARAMS("envs")
	AGENT                =PARAMS("agent")
	NAME                 =PARAMS("name")
	NAMESPACE            =PARAMS("namespace")
	ENV                  =PARAMS("env")
	URL                  =PARAMS("url")
)

type PROCESS_STATUS string

const (
	ACTIVE   = PROCESS_STATUS("active")
	COMPLETED   = PROCESS_STATUS("completed")
	FAILED             =PROCESS_STATUS("failed")
	PAUSED             =PROCESS_STATUS("paused")
)
type PIPELINE_RESOURCE_TYPE string
const  (
	GIT=PIPELINE_RESOURCE_TYPE("git")
	IMAGE=PIPELINE_RESOURCE_TYPE("image")
	DEPLOYMENT=PIPELINE_RESOURCE_TYPE("deployment")
	STATEFULSET=PIPELINE_RESOURCE_TYPE("statefulset")
	DAEMONSET=PIPELINE_RESOURCE_TYPE("daemonset")
	POD=PIPELINE_RESOURCE_TYPE("pod")
	REPLICASET=PIPELINE_RESOURCE_TYPE("replicaset")
)
