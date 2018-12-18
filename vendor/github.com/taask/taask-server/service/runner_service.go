package service

import (
	"fmt"
	"net"
	"time"

	log "github.com/cohix/simplog"
	"github.com/pkg/errors"
	"github.com/taask/taask-server/brain"
	"github.com/taask/taask-server/model"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

const runnerServicePort = ":3687"

// StartRunnerService starts the runner service
func StartRunnerService(brain *brain.Manager, errChan chan error) {
	lis, err := net.Listen("tcp", runnerServicePort)
	if err != nil {
		errChan <- err
		return
	}

	grpcServer := grpc.NewServer()

	RegisterRunnerServiceServer(grpcServer, &RunnerService{Manager: brain})

	log.LogInfo("starting taask-server runner service on :3687")
	if err := grpcServer.Serve(lis); err != nil {
		errChan <- err
	}
}

// RunnerService describes the service available to taask runners
type RunnerService struct {
	Manager *brain.Manager
}

// AuthRunner allows a runner to advertise itself and perform auth with the server
func (rs *RunnerService) AuthRunner(ctx context.Context, req *AuthRunnerRequest) (*AuthRunnerResponse, error) {
	defer log.LogTrace("AuthRunner")()

	encRunnerChallenge, err := rs.Manager.AuthRunner(req.PubKey, req.JoinCodeSignature)
	if err != nil {
		return nil, err
	}

	resp := &AuthRunnerResponse{
		EncChallenge:    encRunnerChallenge.EncChallenge,
		EncChallengeKey: encRunnerChallenge.EncChallengeKey,
	}

	return resp, nil
}

// RegisterRunner allows a runner to connect (with a valid session) get a stream of tasks to execute
func (rs *RunnerService) RegisterRunner(req *RegisterRunnerRequest, stream RunnerService_RegisterRunnerServer) error {
	defer log.LogTrace(fmt.Sprintf("RegisterRunner kind %s", req.Kind))()

	tasksChan := make(chan *model.Task, 128)

	runner := &model.Runner{
		UUID:        model.NewRunnerUUID(),
		Kind:        req.Kind,
		Tags:        req.Tags,
		TaskChannel: tasksChan,
	}

	if err := rs.Manager.RegisterRunner(runner, req.ChallengeSignature); err != nil {
		log.LogError(errors.Wrap(err, "failed to RegisterRunner"))
		return err
	}

	defer rs.Manager.UnregisterRunner(runner)
	go startRunnerHeartbeat(tasksChan)

	log.LogInfo(fmt.Sprintf("runner %s ready to receive tasks", runner.UUID))

	for {
		task := <-tasksChan

		update := model.TaskUpdate{
			Status:     model.TaskStatusQueued,
			RunnerUUID: runner.UUID,
		}

		// if uuid is "", then it's a heartbeat
		if task.UUID != "" {
			runnerEncKey, err := rs.Manager.EncryptTaskKeyForRunner(runner.UUID, task.Meta.MasterEncTaskKey)
			if err != nil {
				log.LogError(errors.Wrap(err, "failed to ReEncryptTaskKey"))
				continue
			}

			update.RunnerEncTaskKey = runnerEncKey

			var updateErr error
			update, updateErr = task.Update(update)

			if updateErr != nil {
				log.LogError(errors.Wrap(updateErr, "RegisterRunner failed to task.Update"))
				continue
			}

			log.LogInfo(fmt.Sprintf("sending task %s to runner %s", task.UUID, runner.UUID))
		} else {
			log.LogInfo(fmt.Sprintf("sending runner %s heartbeat", runner.UUID))
		}

		if err := stream.Send(task); err != nil {
			log.LogError(errors.Wrap(err, "failed to stream.Send"))

			if task.UUID != "" {
				rs.Manager.Updater.UpdateTask(update) // persist the queued update so that the task goes waiting -> queued -> retrying
				log.LogInfo(fmt.Sprintf("task %s is dead, a retry worker should be started for it", task.UUID))
			}

			break
		}

		if task.UUID != "" {
			rs.Manager.Updater.UpdateTask(update)
		}
	}

	return nil
}

// UpdateTask handles update task calls
func (rs *RunnerService) UpdateTask(ctx context.Context, req *model.TaskUpdate) (*Empty, error) {
	defer log.LogTrace(fmt.Sprintf("UpdateTask task %s", req.UUID))()

	if err := rs.checkTaskVersion(req); err != nil {
		log.LogError(errors.Wrap(err, "failed to checkTaskVersion"))
		return &Empty{}, nil
	}

	rs.Manager.Updater.UpdateTask(*req)

	return &Empty{}, nil
}

// TODO: find a way to terminate this
func startRunnerHeartbeat(taskChan chan *model.Task) {
	for {
		<-time.After(time.Second * time.Duration(10))

		taskChan <- &model.Task{}
	}
}

func (rs *RunnerService) checkTaskVersion(update *model.TaskUpdate) error {
	task, err := rs.Manager.GetTask(update.UUID)
	if err != nil {
		return errors.Wrap(err, "failed to storage.Get")
	}

	if update.Version != task.Meta.Version+1 {
		return fmt.Errorf("runner tried to apply update with version %d to task %s with version %d", update.Version, task.UUID, task.Meta.Version)
	}

	return nil
}
