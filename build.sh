../bin/tinygo/bin/tinygo build -o policy.wasm -target=wasi -no-debug .
../bin/spdx-sbom-generator -f json
mv bom-go-mod.json policy-sbom.spdx.json
../bin/kwctl-linux-x86_64/kwctl annotate -m metadata.yml -u README.md -o annotated-policy.wasm policy.wasm
COSIGN_EXPERIMENTAL=1 cosign sign-blob --output-certificate policy-sbom.spdx.cert --output-signature policy-sbom.spdx.sig policy-sbom.spdx.json
PATH="/home/kemara/Documents/workspaces/kubewarden/bin/kwctl-linux-x86_64:$PATH" bats e2e.bats
IMMUTABLE_REF=$(../bin/kwctl-linux-x86_64/kwctl push -o json annotated-policy.wasm khaledemaradev/hello-world-signed:latest | jq -r .immutable_ref)
COSIGN_EXPERIMENTAL=1 cosign sign ${IMMUTABLE_REF}
echo $IMMUTABLE_REF
