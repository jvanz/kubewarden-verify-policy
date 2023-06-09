// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson3fc85001DecodeGithubComKubewardenK8sObjectsApiBatchV1(in *jlexer.Lexer, out *PodFailurePolicyOnPodConditionsPattern) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "status":
			if in.IsNull() {
				in.Skip()
				out.Status = nil
			} else {
				if out.Status == nil {
					out.Status = new(string)
				}
				*out.Status = string(in.String())
			}
		case "type":
			if in.IsNull() {
				in.Skip()
				out.Type = nil
			} else {
				if out.Type == nil {
					out.Type = new(string)
				}
				*out.Type = string(in.String())
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson3fc85001EncodeGithubComKubewardenK8sObjectsApiBatchV1(out *jwriter.Writer, in PodFailurePolicyOnPodConditionsPattern) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix[1:])
		if in.Status == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Status))
		}
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		if in.Type == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Type))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PodFailurePolicyOnPodConditionsPattern) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3fc85001EncodeGithubComKubewardenK8sObjectsApiBatchV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PodFailurePolicyOnPodConditionsPattern) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3fc85001EncodeGithubComKubewardenK8sObjectsApiBatchV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PodFailurePolicyOnPodConditionsPattern) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3fc85001DecodeGithubComKubewardenK8sObjectsApiBatchV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PodFailurePolicyOnPodConditionsPattern) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3fc85001DecodeGithubComKubewardenK8sObjectsApiBatchV1(l, v)
}
