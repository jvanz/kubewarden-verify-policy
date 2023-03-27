package main

import (
	"fmt"

	onelog "github.com/francoispqt/onelog"
	corev1 "github.com/kubewarden/k8s-objects/api/core/v1"
	kubewarden "github.com/kubewarden/policy-sdk-go"
	kubewarden_capabilities "github.com/kubewarden/policy-sdk-go/capabilities"
	kubewarden_protocol "github.com/kubewarden/policy-sdk-go/protocol"
	"github.com/mailru/easyjson"
)

func validate(payload []byte) ([]byte, error) {
	// Create a ValidationRequest instance from the incoming payload
	validationRequest := kubewarden_protocol.ValidationRequest{}
	err := easyjson.Unmarshal(payload, &validationRequest)
	if err != nil {
		return kubewarden.RejectRequest(
			kubewarden.Message(err.Error()),
			kubewarden.Code(400))
	}

	// Create a Settings instance from the ValidationRequest object
	settings, err := NewSettingsFromValidationReq(&validationRequest)
	if err != nil {
		return kubewarden.RejectRequest(
			kubewarden.Message(err.Error()),
			kubewarden.Code(400))
	}

	// Access the **raw** JSON that describes the object
	podJSON := validationRequest.Request.Object

	// Try to create a Pod instance using the RAW JSON we got from the
	// ValidationRequest.
	pod := &corev1.Pod{}
	if err := easyjson.Unmarshal([]byte(podJSON), pod); err != nil {
		return kubewarden.RejectRequest(
			kubewarden.Message(
				fmt.Sprintf("Cannot decode Pod object: %s", err.Error())),
			kubewarden.Code(400))
	}

	logger.DebugWithFields("validating pod object", func(e onelog.Entry) {
		e.String("name", pod.Metadata.Name)
		e.String("namespace", pod.Metadata.Namespace)
	})

	host := kubewarden_capabilities.NewHost()

	verificatopnResponse, err := host.VerifyPubKeys(settings.Image, settings.PubKeys, make(map[string]string))
	if err != nil {
		logger.InfoWithFields("rejecting pod object", func(e onelog.Entry) {
			e.String("name", pod.Metadata.Name)
			e.String("image", settings.Image)
		})

		return kubewarden.RejectRequest(
			kubewarden.Message(
				fmt.Sprintf("The '%s' image cannot be verified: %s", settings.Image, err)),
			kubewarden.NoCode)
	}

	if !verificatopnResponse.IsTrusted {
		logger.InfoWithFields("rejecting pod object", func(e onelog.Entry) {
			e.String("name", pod.Metadata.Name)
			e.String("image", settings.Image)
		})

		return kubewarden.RejectRequest(
			kubewarden.Message(
				fmt.Sprintf("The '%s' image is not trusted", settings.Image)),
			kubewarden.NoCode)
	}

	return kubewarden.AcceptRequest()
}
