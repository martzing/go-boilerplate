package job

import (
	"time"

	utilityHelper "github.com/martzing/go-boilerplate/helpers/utility"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

type Schedule struct {
	Id       string
	Schedule string
	Body     any
	Func     func(any)
}

type JobManeger struct {
	job *cron.Cron
	log *zap.SugaredLogger
}

var Manager *JobManeger

func New() *JobManeger {
	logger := utilityHelper.NewLogger()
	// Start Create init
	logger.Infoln("###Starting intial Jobs###")

	tz, err := time.LoadLocation("Asia/Bangkok")

	if err != nil {
		panic(err)
	}

	_cron := cron.New(cron.WithSeconds(), cron.WithLocation(tz))

	Manager = &JobManeger{
		job: _cron,
		log: logger,
	}

	return Manager
}

func (j *JobManeger) CreateJob(schedule Schedule) {
	entryId, err := j.job.AddFunc(schedule.Schedule, func() {
		schedule.Func(schedule.Body)
	})

	if err != nil {
		j.log.Errorf("job:error:%v", err.Error())
	}

	j.log.Infof("job:added:%v:%v", schedule.Id, entryId)
}

func (j *JobManeger) Start() {
	j.job.Start()
}
