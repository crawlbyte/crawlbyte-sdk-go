package crawlbyte

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/crawlbyte/crawlbyte-sdk-go/client"
)

type SDK struct {
	client *client.Client
}

func New(apiKey string) *SDK {
	return &SDK{
		client: client.New("https://api.crawlbyte.ai/api", apiKey),
	}
}

func (sdk *SDK) CreateTask(payload TaskPayload) (*Task, error) {
	respBody, err := sdk.client.DoRequest("POST", "/tasks", payload)
	if err != nil {
		return nil, err
	}

	var task Task
	if taskErr := json.Unmarshal(respBody, &task); taskErr != nil {
		return nil, taskErr
	}

	return &task, nil
}

func (sdk *SDK) GetTask(id string) (*Task, error) {
	respBody, err := sdk.client.DoRequest("GET", fmt.Sprintf("/tasks/%s", id), nil)
	if err != nil {
		return nil, err
	}

	var task Task
	if taskErr := json.Unmarshal(respBody, &task); taskErr != nil {
		return nil, taskErr
	}

	return &task, nil
}

func (sdk *SDK) PollTask(id string, opts PollOptions) (*Task, error) {
	interval := time.Duration(opts.IntervalSeconds) * time.Second
	timeout := time.Duration(opts.TimeoutSeconds) * time.Second

	deadline := time.Now().Add(timeout)

	for {
		task, err := sdk.GetTask(id)
		if err != nil {
			return nil, err
		}

		switch task.Status {
		case "completed", "failed":
			return task, nil
		case "queued", "processing":
			// continue polling
		default:
			return nil, fmt.Errorf("unexpected task status: %s", task.Status)
		}

		if time.Now().After(deadline) {
			return nil, fmt.Errorf("timeout reached while polling task %s", id)
		}

		time.Sleep(interval)
	}
}
