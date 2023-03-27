// This package provides access to the structs and functions offered by the Kubewarden host.
// This allows policies to perform operations that are not doable inside of the WebAssembly
// runtime. Such as, policy verification, reverse DNS lookups, interacting with OCI registries,...
package capabilities

import (
	"encoding/json"
	"fmt"

	"github.com/mailru/easyjson"
)

// Host makes possible to interact with the policy host from inside of a
// policy.
//
// Use the `NewHost` function to create an instance of `Host`.
type Host struct {
	Client WapcClient
}

type WapcClient interface {
	HostCall(binding, namespace, operation string, payload []byte) (response []byte, err error)
}

// GetOCIManifestDigest computes the digest of the OCI object referenced by image
// Arguments:
// * image: image to be verified (e.g.: `registry.testing.lan/busybox:1.0.0`)
func (h *Host) GetOCIManifestDigest(image string) (string, error) {
	// build request payload, e.g: `"ghcr.io/kubewarden/policies/pod-privileged:v0.1.10"`
	payload, err := json.Marshal(image)
	if err != nil {
		return "", fmt.Errorf("cannot serialize image to JSON: %w", err)
	}

	// perform host callback
	responsePayload, err := h.Client.HostCall("kubewarden", "oci", "v1/manifest_digest", payload)
	if err != nil {
		return "", err
	}

	response := OciManifestResponse{}
	if err := easyjson.Unmarshal(responsePayload, &response); err != nil {
		return "", fmt.Errorf("cannot unmarshall response: %w", err)
	}

	return response.Digest, nil
}

// LookupHost looks up the addresses for a given hostname via DNS
func (h *Host) LookupHost(host string) ([]string, error) {
	// build request, e.g: `"localhost"`
	payload, err := json.Marshal(host)
	if err != nil {
		return []string{}, fmt.Errorf("cannot serialize host to JSON: %w", err)
	}

	// perform host callback
	responsePayload, err := h.Client.HostCall("kubewarden", "net", "v1/dns_lookup_host", payload)
	if err != nil {
		return []string{}, err
	}

	response := LookupHostResponse{}
	if err := easyjson.Unmarshal(responsePayload, &response); err != nil {
		return []string{}, fmt.Errorf("cannot unmarshall response: %w", err)
	}

	return response.Ips, nil
}

// VerifyPubKeys verifies sigstore signatures of an image using public keys
// Arguments
// * image: image to be verified (e.g.: `registry.testing.lan/busybox:1.0.0`)
// * pubKeys: list of PEM encoded keys that must have been used to sign the OCI object
// * annotations: annotations that must have been provided by all signers when they signed the OCI artifact
func (h *Host) VerifyPubKeys(image string, pubKeys []string, annotations map[string]string) (VerificationResponse, error) {
	requestObj := sigstorePubKeysVerify{
		Image:       image,
		PubKeys:     pubKeys,
		Annotations: annotations,
	}

	return h.verify(requestObj, "v1/verify")
}

// VerifyKeyless verifies sigstore signatures of an image using keyless signing
// Arguments
// * image: image to be verified (e.g.: `registry.testing.lan/busybox:1.0.0`)
// * keyless: list of KeylessInfo pairs, containing Issuer and Subject info from OIDC providers
// * annotations: annotations that must have been provided by all signers when they signed the OCI artifact
func (h *Host) VerifyKeyless(image string, keyless []KeylessInfo, annotations map[string]string) (VerificationResponse, error) {
	requestObj := sigstoreKeylessVerify{
		Image:       image,
		Keyless:     keyless,
		Annotations: annotations,
	}

	return h.verify(requestObj, "v1/verify")
}

// VerifyPubKeysImageV2 verifies sigstore signatures of an image using public keys
// Arguments
// * image: image to be verified (e.g.: `registry.testing.lan/busybox:1.0.0`)
// * pubKeys: list of PEM encoded keys that must have been used to sign the OCI object
// * annotations: annotations that must have been provided by all signers when they signed the OCI artifact
func (h *Host) VerifyPubKeysImageV2(image string, pubKeys []string, annotations map[string]string) (VerificationResponse, error) {
	requestObj := sigstorePubKeysVerify{
		Image:       image,
		PubKeys:     pubKeys,
		Annotations: annotations,
	}

	return h.verify(requestObj)
}

// VerifyKeylessExactMatchV2 verifies sigstore signatures of an image using keyless signing
// Arguments
// * image: image to be verified (e.g.: `registry.testing.lan/busybox:1.0.0`)
// * keyless: list of KeylessInfo pairs, containing Issuer and Subject info from OIDC providers
// * annotations: annotations that must have been provided by all signers when they signed the OCI artifact
func (h *Host) VerifyKeylessExactMatchV2(image string, keyless []KeylessInfo, annotations map[string]string) (VerificationResponse, error) {
	requestObj := sigstoreKeylessVerify{
		Image:       image,
		Keyless:     keyless,
		Annotations: annotations,
	}

	return h.verify(requestObj)
}

// verify sigstore signatures of an image using keyless. Here, the provided
// subject string is treated as a URL prefix, and sanitized to a valid URL on
// itself by appending `/` to prevent typosquatting. Then, the provided subject
// will satisfy the signature only if it is a prefix of the signature subject.
// # Arguments
// * `image` -  image to be verified
// * `keyless`  -  list of issuers and subjects
// * `annotations` - annotations that must have been provided by all signers when they signed the OCI artifact
func (h *Host) VerifyKeylessPrefixMatchV2(image string, keylessPrefix []KeylessPrefixInfo, annotations map[string]string) (VerificationResponse, error) {
	requestObj := sigstoreKeylessPrefixVerify{
		Image:         image,
		KeylessPrefix: keylessPrefix,
		Annotations:   annotations,
	}

	return h.verify(requestObj)
}

// verify sigstore signatures of an image using keyless signatures made via
// Github Actions.
// # Arguments
// * `image` -  image to be verified
// * `owner` - owner of the repository. E.g: octocat
// * `repo` - Optional. repo of the GH Action workflow that signed the artifact. E.g: example-repo. Optional.
// * `annotations` - annotations that must have been provided by all signers when they signed the OCI artifact
func (h *Host) VerifyKeylessGithubActionsV2(image string, owner string, repo string, annotations map[string]string) (VerificationResponse, error) {
	requestObj := sigstoreGithubActionsVerify{
		Image:       image,
		Owner:       owner,
		Repo:        repo,
		Annotations: annotations,
	}

	return h.verify(requestObj)
}

// verify sigstore signatures of an image using a user provided certificate
// # Arguments
//   - `image` -  image to be verified
//   - `certificate` - PEM encoded certificate used to verify the signature
//   - `certificate_chain` - Optional. PEM encoded certificates used to verify `certificate`.
//     When not specified, the certificate is assumed to be trusted
//   - `require_rekor_bundle` - require the  signature layer to have a Rekor bundle.
//     Having a Rekor bundle allows further checks to be performed,
//     like ensuring the signature has been produced during the validity
//     time frame of the certificate.
//     It is recommended to set this value to `true` to have a more secure
//     verification process.
//   - `annotations` - annotations that must have been provided by all signers when they signed the OCI artifact
func (h *Host) VerifyCertificateV2(image string, certificate string, certificateChain []string, requireRekorBundle bool, annotations map[string]string) (VerificationResponse, error) {
	chain := make([][]rune, len(certificateChain))
	for i, c := range certificateChain {
		chain[i] = []rune(c)
	}

	requestObj := sigstoreCertificateVerify{
		Image:              image,
		Certificate:        []rune(certificate),
		CertificateChain:   chain,
		RequireRekorBundle: requireRekorBundle,
		Annotations:        annotations,
	}

	return h.verify(requestObj)
}

func (h *Host) verify(requestObj easyjson.Marshaler, option_operation ...string) (VerificationResponse, error) {
	// failsafe return response
	vr := VerificationResponse{
		IsTrusted: false,
		Digest:    "",
	}

	payload, err := easyjson.Marshal(requestObj)
	if err != nil {
		return vr, fmt.Errorf("cannot serialize request object: %w", err)
	}

	operation := "v2/verify"
	if len(option_operation) == 1 {
		operation = option_operation[0]
	}

	// perform callback
	responsePayload, err := h.Client.HostCall("kubewarden", "oci", operation, payload)
	if err != nil {
		return vr, fmt.Errorf("%s: %w", operation, err)
	}

	responseObj := VerificationResponse{}
	if err := easyjson.Unmarshal(responsePayload, &responseObj); err != nil {
		return vr, fmt.Errorf("cannot unmarshall response object: %w", err)
	}

	return responseObj, nil
}

// ListResourcesByNamespace gets all the Kubernetes resources defined inside of
// the given namespace
// Note: cannot be used for cluster-wide resources
func (h *Host) ListResourcesByNamespace(req ListResourcesByNamespaceRequest) ([]byte, error) {
	payload, err := req.MarshalJSON()
	if err != nil {
		return []byte{}, fmt.Errorf("cannot serialize request object: %w", err)
	}

	// perform callback
	responsePayload, err := h.Client.HostCall("kubewarden", "kubernetes", "list_resources_by_namespace", payload)
	if err != nil {
		return []byte{}, err
	}

	return responsePayload, nil
}

// ListResources gets all the Kubernetes resources defined inside of the cluster.
// Note: this has be used for cluster-wide resources
func (h *Host) ListResources(req ListAllResourcesRequest) ([]byte, error) {
	payload, err := req.MarshalJSON()
	if err != nil {
		return []byte{}, fmt.Errorf("cannot serialize request object: %w", err)
	}

	// perform callback
	responsePayload, err := h.Client.HostCall("kubewarden", "kubernetes", "list_all_resources", payload)
	if err != nil {
		return []byte{}, err
	}

	return responsePayload, nil
}

// GetResource gets a specific Kubernetes resource.
func (h *Host) GetResource(req GetResourceRequest) ([]byte, error) {
	payload, err := req.MarshalJSON()
	if err != nil {
		return []byte{}, fmt.Errorf("cannot serialize request object: %w", err)
	}

	// perform callback
	responsePayload, err := h.Client.HostCall("kubewarden", "kubernetes", "get_resource", payload)
	if err != nil {
		return []byte{}, err
	}

	return responsePayload, nil
}

// Verify_cert verifies cert's trust against the passed cert_chain, and
// expiration and validation time of the certificate.
// Accepts 3 arguments:
//   - cert: PEM-encoded certificate to verify.
//   - cert_chain: list of PEM-encoded certs, ordered by trust usage
//     (intermediates first, root last). If empty, certificate is assumed trusted.
//   - not_after: string in RFC 3339 time format, to check expiration against.
//     If None, certificate is assumed never expired.
func (h *Host) VerifyCert(cert Certificate, certChain []Certificate, notAfter string) (bool, error) {
	requestObj := CertificateVerificationRequest{
		Cert:      cert,
		CertChain: certChain,
		NotAfter:  notAfter,
	}

	payload, err := easyjson.Marshal(requestObj)
	if err != nil {
		return false, fmt.Errorf("cannot serialize request object: %w", err)
	}

	// perform callback
	responsePayload, err := h.Client.HostCall("kubewarden", "crypto", "v1/is_certificate_trusted", payload)
	if err != nil {
		return false, err
	}

	responseObj := CertificateVerificationResponse{}
	if err := easyjson.Unmarshal(responsePayload, &responseObj); err != nil {
		return false, fmt.Errorf("cannot unmarshall response object: %w", err)
	}

	switch responseObj.Trusted {
	case true:
		return true, nil
	case false:
		return false, fmt.Errorf(responseObj.Reason)
	default:
		return false, fmt.Errorf("unknown response from the host")
	}
}
