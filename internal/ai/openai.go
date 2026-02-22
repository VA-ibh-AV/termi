package ai

import (
	"context"
	"errors"
	"os"

	"github.com/VA-ibh-AV/termi/utils"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/openai/openai-go/v3/responses"
	"github.com/openai/openai-go/v3/shared"

	"encoding/json"

	"github.com/invopop/jsonschema"
)

func GenerateSchema[T any]() map[string]any {
	reflector := jsonschema.Reflector{
		AllowAdditionalProperties: false,
		DoNotReference:            true,
	}
	var v T
	schema := reflector.Reflect(v)

	data, _ := json.Marshal(schema)
	var result map[string]any
	json.Unmarshal(data, &result)
	return result
}

type CommandResponse struct {
	Command string `json:"command" jsonschema_description:"The command to execute"`
}

var CommandResposneSchema = GenerateSchema[CommandResponse]()

func Generate(ctx context.Context, prompt string) (*CommandResponse, error) {

	openAIKey := os.Getenv("OPENAI_KEY")

	if len(openAIKey) == 0 {
		return nil, errors.New("SET OPENAI_KEY in env")
	}

	client := openai.NewClient(
		option.WithAPIKey(os.Getenv("OPENAI_KEY")), // defaults to
	)

	hostInfo := utils.GetHostInfo()
	systemPrompt := utils.BuildCommandSystemPrompt(hostInfo.Os, hostInfo.Distro, hostInfo.Arch)

	resp, err := client.Responses.New(ctx, responses.ResponseNewParams{
		Reasoning: shared.ReasoningParam{Effort: openai.ReasoningEffortLow},
		Input: responses.ResponseNewParamsInputUnion{
			OfString: openai.String(systemPrompt + "\n\nUser Prompt:\n" + prompt),
		},
		Model: openai.ChatModelGPT5Nano2025_08_07,
		Text: responses.ResponseTextConfigParam{
			Format: responses.ResponseFormatTextConfigParamOfJSONSchema(
				"command_response",
				CommandResposneSchema,
			),
			Verbosity: responses.ResponseTextConfigVerbosityLow,
		},
	})

	if err != nil {
		return nil, err
	}

	var commandResponse CommandResponse
	err = json.Unmarshal([]byte(resp.OutputText()), &commandResponse)

	if err != nil {
		return nil, err
	}

	return &commandResponse, nil
}
