package jira

import (
	"fmt"
	"github.com/gypsydave5/kickoff"
	"github.com/gypsydave5/kickoff/question"
)

type InitHandler struct {
	client *Client
}

func NewInitHandler(client *Client) *InitHandler {
	return &InitHandler{client}
}

func (t InitHandler) Handle(questioner question.Input) (*kickoff.Kickoff, error) {
	storyID, err := questioner.Ask(whatIsYourJiraStoryId{})
	if err != nil {
		return &kickoff.Kickoff{}, err
	}

	story, err := t.client.GetStory(storyID.String())
	if err != nil {
		return &kickoff.Kickoff{}, err
	}

	ko := kickoff.NewKickoff(jiraStoryTitle(story))
	ko.AddSection(jiraLinkLine(story))
	return ko, nil
}

type whatIsYourJiraStoryId struct{}

func (w whatIsYourJiraStoryId) String() string {
	return "What is your Jira story ID?"
}

func jiraStoryTitle(story Story) string {
	return story.ID + ": " + story.Summary
}

func jiraLinkLine(story Story) string {
	return fmt.Sprintf("[Link To Jira](%s)", story.Link)
}
