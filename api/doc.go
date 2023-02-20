package api

//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=config.yaml openapi.yaml
//generate API code.

//go:generate go run github.com/vektra/mockery/v2@v2.18.0 --name ClientWithResponsesInterface --output ../mocks --with-expecter --filename api.gen.go --structname Api
//generate client interface mocks

//go:generate go run github.com/twitchtv/circuitgen --name ClientWithResponsesInterface --pkg ./ --debug --out ./circuit.gen.go --alias ClientWithResponsesCircuit --circuit-major-version 3
//generate circuits for interface

/*
func GenerateREPLACEMEClient(
	server string,
	subjectString string,
	client api.HttpRequestDoer,
	manager *circuit.Manager,
	conf api.CircuitWrapperClientWithResponsesCircuitConfig,
) (api.ClientWithResponsesInterface, error) {
	e, err := api.NewClientWithResponses(server, api.ClientOption(func(cl *api.Client) error {
		cl.RequestEditors = []api.RequestEditorFn{
			func(ctx context.Context, req *http.Request) error {
				req.Header.Add("subject-uuid", subjectString)
				return nil
			},
		}
		cl.Client = client
		return nil
	}))
	if err != nil {
		return nil, err
	}

	return api.NewCircuitWrapperClientWithResponsesCircuit(manager, e, conf)
}
*/
