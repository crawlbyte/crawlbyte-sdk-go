package crawlbyte_test

import (
	"github.com/joho/godotenv"
	"os"
	"testing"

	"github.com/crawlbyte/crawlbyte-sdk-go"
)

func TestUniversalWalmartTask(t *testing.T) {
	_ = godotenv.Load(".env")

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		t.Fatal("Missing API_KEY")
	}

	sdk := crawlbyte.New(apiKey)

	task, err := sdk.CreateTask(crawlbyte.TaskPayload{
		Template: "universal",
		Input: []string{
			"https://www.walmart.com/",
		},
	})
	if err != nil {
		t.Fatalf("CreateTask failed: %v", err)
	}

	t.Logf("[✓] Task created: ID=%s, status=%s", task.ID, task.Status)

	if task.Status == "completed" || task.Status == "failed" {
		t.Logf("[✓] Task finished immediately: status=%s", task.Status)
		t.Logf("[→] Result: %+v", task.Result)
		return
	}

	t.Logf("[…] Polling for result...")

	result, err := sdk.PollTask(task.ID, crawlbyte.PollOptions{
		IntervalSeconds: 5,
		TimeoutSeconds:  60,
	})
	if err != nil {
		t.Fatalf("[✗] Polling failed: %v", err)
	}

	if result.Status != "completed" {
		t.Errorf("Expected status 'completed', got '%s'", result.Status)
	}

	t.Logf("[✓] Final task status: %s", result.Status)
	t.Logf("[→] Final result: %+v", result.Result)
}
