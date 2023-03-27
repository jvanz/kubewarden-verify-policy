#!/usr/bin/env bats

@test "reject because name is on deny list" {
  run kwctl run annotated-policy.wasm -r test_data/pod.json --settings-json '{"image": "khaledemaradev/hello-world-signed@sha256:f54a58bc1aac5ea1a25d796ae155dc228b3f0e11d046ae276b39c4bf2f13d8c4", "pub_keys": ["-----BEGIN PUBLIC KEY-----\nMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE6WQpyouAQN7ZPmc5F50O17xkPq/T\np1aXg2yEVIxcZxcaGjHqPnUVbEnGaFFSgd/DjpEa1CYpn85h9WE4kYKFpA==\n-----END PUBLIC KEY-----\n"]}'

  # this prints the output when one the checks below fails
  echo "output = ${output}"

  # request rejected
  [ "$status" -eq 0 ]
  [ $(expr "$output" : '.*allowed.*false') -ne 0 ]
  [ $(expr "$output" : ".*The 'test-pod' name is on the deny list.*") -ne 0 ]
}

@test "accept because name is not on the deny list" {
  run kwctl run annotated-policy.wasm -r test_data/pod.json --settings-json '{"denied_names": ["foo"]}'
  # this prints the output when one the checks below fails
  echo "output = ${output}"

  # request accepted
  [ "$status" -eq 0 ]
  [ $(expr "$output" : '.*allowed.*true') -ne 0 ]
}

@test "accept because the deny list is empty" {
  run kwctl run annotated-policy.wasm -r test_data/pod.json
  # this prints the output when one the checks below fails
  echo "output = ${output}"

  # request accepted
  [ "$status" -eq 0 ]
  [ $(expr "$output" : '.*allowed.*true') -ne 0 ]
}
