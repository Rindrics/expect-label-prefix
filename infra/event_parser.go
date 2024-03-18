package infra

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log/slog"
	"os"

	"github.com/Rindrics/require-label-prefix-single/domain"
	"github.com/google/go-github/github"
)

func LoadEventFromEnv() (interface{}, error) {
	eventPath := os.Getenv("GITHUB_EVENT_PATH")
	if eventPath == "" {
		return nil, fmt.Errorf("GITHUB_EVENT_PATH environment variable not set")
	}

	data, err := ioutil.ReadFile(eventPath)
	if err != nil {
		return nil, err
	}

	// Unmarshal the data into a generic map
	var genericData map[string]interface{}
	if err := json.Unmarshal(data, &genericData); err != nil {
		return nil, err
	}

	// Judge the type of the event
	if _, ok := genericData["issue"]; ok {
		var event github.IssuesEvent
		if err := json.Unmarshal(data, &event); err != nil {
			return nil, err
		}
		return &event, nil
	} else if _, ok := genericData["pull_request"]; ok {
		var event github.PullRequestEvent
		if err := json.Unmarshal(data, &event); err != nil {
			return nil, err
		}
		return &event, nil
	} else {
		return nil, fmt.Errorf("unsupported or unknown event type in the provided data")
	}
}

func ParseEvent(event interface{}, l *slog.Logger) domain.EventInfo {
	switch e := event.(type) {
	case *github.IssuesEvent:
		title := e.GetIssue().GetTitle()
		number := e.GetIssue().GetNumber()
		var labels []string
		for _, label := range e.GetIssue().Labels {
			labels = append(labels, *label.Name)
		}
		l.Debug("parsed", "title", title, "number", number, "labels", labels)
		return domain.EventInfo{Number: number, Labels: labels}
	case *github.PullRequestEvent:
		title := e.GetPullRequest().GetTitle()
		number := e.GetNumber()
		var labels []string
		for _, label := range e.GetPullRequest().Labels {
			labels = append(labels, label.GetName())
		}
		l.Debug("parsed", "title", title, "number", number, "labels", labels)
		return domain.EventInfo{Number: number, Labels: labels}
	default:
		l.Error("Unsupported event type")
		return domain.EventInfo{}
	}
}
