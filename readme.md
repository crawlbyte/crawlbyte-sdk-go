# Crawlbyte Go SDK

Official Go SDK for interacting with the [Crawlbyte API](https://crawlbyte.ai).

## Features

- Create scraping tasks with various templates and configurations
- Poll task status until completion
- Retrieve task results
- Support for all Crawlbyte API parameters and options

## Installation

```bash
go get github.com/crawlbyte/crawlbyte-sdk-go
```

## Quick Start

### Initialize SDK

```go
import "github.com/crawlbyte/crawlbyte-sdk-go"

sdk := crawlbyte.New("your-api-key")
```

### Create and Handle Tasks

```go
// Create a task
task, err := sdk.CreateTask(crawlbyte.TaskPayload{
    Template: "universal",
    Input: []string{
        "https://www.walmart.com/",
    },
})
if err != nil {
    // handle error
}

fmt.Println("Task ID:", task.ID)

// Check if task completed immediately (under 20 seconds)
if task.Status == "completed" || task.Status == "failed" {
    fmt.Println("Task completed immediately")
    fmt.Println("Result:", task.Result)
    return
}

// If still processing, poll until completion
result, err := sdk.PollTask(task.ID, crawlbyte.PollOptions{
    IntervalSeconds: 5,
    TimeoutSeconds:  60,
})
if err != nil {
    // handle error
}

fmt.Println("Final Status:", result.Status)
fmt.Println("Result:", result.Result)
```

### Get Task Status

```go
task, err := sdk.GetTask("task-id")
if err != nil {
    // handle error
}
fmt.Println("Status:", task.Status)
```

## Available Methods

- `CreateTask(payload TaskPayload) (*Task, error)` - Create a new scraping task. Returns results immediately if completed within 20 seconds, otherwise returns task details for polling.
- `GetTask(id string) (*Task, error)` - Get task status and results.
- `PollTask(id string, opts PollOptions) (*Task, error)` - Poll task until completion (only needed if task takes longer than 20 seconds).

## Task Configuration

The `TaskPayload` struct supports all Crawlbyte API parameters. Common fields include:

- `Template` - The scraping template to use (e.g., "walmart")
- `Input` - Array of URLs/Product IDs to scrape
- `Fields` - Specific data fields to extract (sites that support this)
- `DataType` - Type of data to extract
- `JSRendering` - Enable JavaScript rendering
- `Proxy` - Proxy configuration
- `CustomHeaders` - Custom HTTP headers

For a complete list of all available fields, configuration options, required parameters for each template, and detailed API documentation, visit: **[https://developers.crawlbyte.ai/](https://developers.crawlbyte.ai/)**

## Task Statuses

- `queued` - Task is waiting to be processed
- `processing` - Task is currently being executed
- `completed` - Task finished successfully
- `failed` - Task encountered an error

## Error Handling

The SDK returns detailed error messages for HTTP errors and API failures:

```go
task, err := sdk.CreateTask(payload)
if err != nil {
fmt.Printf("Error creating task: %v\n", err)
return
}
```

## Examples

### Universal Template with Multiple URLs

```go
task, err := sdk.CreateTask(crawlbyte.TaskPayload{
Template: "universal",
Input: []string{
"https://example1.com",
"https://example2.com",
},
JSRendering: true,
})
```

## Documentation

For comprehensive API documentation, template specifications, and field requirements, please visit:
**[https://developers.crawlbyte.ai/](https://developers.crawlbyte.ai/)**

## License

MIT