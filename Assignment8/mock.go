package resource

import (
	"leap/exception"
)

// mockResourceTitle is a capitalized version of the resource name, used to generate
// MessageBodyType values in the header response, such as OnemockDefintion/*
const mockResourceTitle = "Mock"

type singleMockDefinition struct {
	Mock string
}

func BodyAndMessageTypeForMockDefinition() (*singleMockDefinition, string, *exception.Exception) {

	mock := "This is a Mock Data"
	return &singleMockDefinition{
		Mock: mock,
	}, GenerateMessageBodyType(mockResourceTitle, MESSAGE_BODY_DEFINITION, SingleObject), nil
}
