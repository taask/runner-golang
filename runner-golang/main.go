package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	log "github.com/cohix/simplog"
	"github.com/pkg/errors"
	"github.com/taask/runner-golang"
	"github.com/taask/taask-server/model"
)

type addition struct {
	First  int
	Second int
}

type answer struct {
	Answer int
}

func main() {
	runner, err := taask.NewRunner("com.taask.dummy", []string{}, func(task *model.Task) (interface{}, error) {
		var problem addition
		if err := json.Unmarshal(task.Body, &problem); err != nil {
			return nil, errors.Wrap(err, "failed to Unmarshal")
		}

		seconds := rand.Intn(5)

		log.LogInfo(fmt.Sprintf("solving %d + %d in %d seconds", problem.First, problem.Second, seconds))

		<-time.After(time.Second * time.Duration(seconds))

		return answer{Answer: problem.First + problem.Second}, nil
	})

	if err != nil {
		log.LogError(errors.Wrap(err, "failed to NewRunner"))
		os.Exit(1)
	}

	if len(os.Args) == 0 {
		log.LogError(errors.New("missing argument: join code"))
		os.Exit(1)
	}

	if err := runner.ConnectAndRun(os.Args[1], "localhost", "3687"); err != nil {
		log.LogError(errors.Wrap(err, "failed to ConnectAndRun"))
		os.Exit(1)
	}
}
