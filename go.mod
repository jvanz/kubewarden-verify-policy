module github.com/KhaledEmaraDev/kubewarden-verify-policy

go 1.19

replace github.com/go-openapi/strfmt => github.com/kubewarden/strfmt v0.1.2

replace github.com/kubewarden/policy-sdk-go => /home/kemara/Documents/workspaces/kubewarden/policy-sdk-go

require (
	github.com/francoispqt/onelog v0.0.0-20190306043706-8c2bb31b10a4
	github.com/kubewarden/k8s-objects v1.26.0-kw1
	github.com/kubewarden/policy-sdk-go v0.3.0
	github.com/mailru/easyjson v0.7.7
	github.com/wapc/wapc-guest-tinygo v0.3.3
)

require (
	github.com/francoispqt/gojay v0.0.0-20181220093123-f2cc13a668ca // indirect
	github.com/go-openapi/strfmt v0.21.3 // indirect
	github.com/josharian/intern v1.0.0 // indirect
)
