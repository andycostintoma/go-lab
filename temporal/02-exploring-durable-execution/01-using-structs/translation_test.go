package translation

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
)

func TestSayHelloGoodbye(t *testing.T) {
	var suite testsuite.WorkflowTestSuite
	env := suite.NewTestWorkflowEnvironment()
	env.OnActivity(TranslateTerm, mock.Anything, TranslationActivityInput{
		Term: "Hello", LanguageCode: "fr",
	}).Return(TranslationActivityOutput{Translation: "Bonjour"}, nil).Once()
	env.OnActivity(TranslateTerm, mock.Anything, TranslationActivityInput{
		Term: "Goodbye", LanguageCode: "fr",
	}).Return(TranslationActivityOutput{Translation: "Au revoir"}, nil).Once()

	env.ExecuteWorkflow(SayHelloGoodbye, TranslationWorkflowInput{Name: "Andy", LanguageCode: "fr"})
	require.NoError(t, env.GetWorkflowError())
	var result TranslationWorkflowOutput
	require.NoError(t, env.GetWorkflowResult(&result))
	require.Equal(t, TranslationWorkflowOutput{
		HelloMessage: "Bonjour, Andy", GoodbyeMessage: "Au revoir, Andy",
	}, result)
	env.AssertExpectations(t)
}
