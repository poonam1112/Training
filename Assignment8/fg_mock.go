package mux

import (
	"leap/exception"
	"leap/protocol"
	"leap/resource"
	"leap/token"
)

// mock registers to "/mock". When read, it returns the details of the given mock(server)
type mock struct {
	functionGroupBase
}

func init() {
	tokenizedUrlToFunctionGroupMap["/mock"] = mock{}

}

func (mock) validate(t *token.TokenURL, connection *ConnectionValueObject) *exception.Exception {
	return nil
}

func (mock) read(t *token.TokenURL, cvo *ConnectionValueObject) (body interface{}, messageBodyType string, status protocol.ActionResponse, exc *exception.Exception) {

	b, m, e := resource.BodyAndMessageTypeForMockDefinition()
	return b, m, protocol.NewActionResponseComplete(), e
}
