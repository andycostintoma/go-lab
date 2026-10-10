package translation

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"go.temporal.io/sdk/workflow"
)

const TaskQueueName = "using-structs-tasks"

type TranslationWorkflowInput struct {
	Name         string
	LanguageCode string
}

type TranslationWorkflowOutput struct {
	HelloMessage   string
	GoodbyeMessage string
}

type TranslationActivityInput struct {
	Term         string
	LanguageCode string
}

type TranslationActivityOutput struct {
	Translation string
}

func SayHelloGoodbye(ctx workflow.Context, input TranslationWorkflowInput) (TranslationWorkflowOutput, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 45 * time.Second,
	})

	var hello, goodbye TranslationActivityOutput
	err := workflow.ExecuteActivity(ctx, TranslateTerm, TranslationActivityInput{
		Term: "Hello", LanguageCode: input.LanguageCode,
	}).Get(ctx, &hello)
	if err != nil {
		return TranslationWorkflowOutput{}, err
	}
	err = workflow.ExecuteActivity(ctx, TranslateTerm, TranslationActivityInput{
		Term: "Goodbye", LanguageCode: input.LanguageCode,
	}).Get(ctx, &goodbye)
	if err != nil {
		return TranslationWorkflowOutput{}, err
	}

	return TranslationWorkflowOutput{
		HelloMessage:   fmt.Sprintf("%s, %s", hello.Translation, input.Name),
		GoodbyeMessage: fmt.Sprintf("%s, %s", goodbye.Translation, input.Name),
	}, nil
}

func TranslateTerm(ctx context.Context, input TranslationActivityInput) (TranslationActivityOutput, error) {
	query := url.Values{"lang": {input.LanguageCode}, "term": {input.Term}}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:9998/translate?"+query.Encode(), nil)
	if err != nil {
		return TranslationActivityOutput{}, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return TranslationActivityOutput{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return TranslationActivityOutput{}, err
	}
	if resp.StatusCode >= 400 {
		return TranslationActivityOutput{}, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, body)
	}
	return TranslationActivityOutput{Translation: string(body)}, nil
}
