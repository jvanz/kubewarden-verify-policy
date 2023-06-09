// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1

import (
	json "encoding/json"
	_v1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
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

func easyjsonCe8ae4eDecodeGithubComKubewardenK8sObjectsApiBatchV1(in *jlexer.Lexer, out *JobStatus) {
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
		case "active":
			out.Active = int32(in.Int32())
		case "completedIndexes":
			out.CompletedIndexes = string(in.String())
		case "completionTime":
			if in.IsNull() {
				in.Skip()
				out.CompletionTime = nil
			} else {
				if out.CompletionTime == nil {
					out.CompletionTime = new(_v1.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.CompletionTime).UnmarshalJSON(data))
				}
			}
		case "conditions":
			if in.IsNull() {
				in.Skip()
				out.Conditions = nil
			} else {
				in.Delim('[')
				if out.Conditions == nil {
					if !in.IsDelim(']') {
						out.Conditions = make([]*JobCondition, 0, 8)
					} else {
						out.Conditions = []*JobCondition{}
					}
				} else {
					out.Conditions = (out.Conditions)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *JobCondition
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(JobCondition)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Conditions = append(out.Conditions, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "failed":
			out.Failed = int32(in.Int32())
		case "ready":
			out.Ready = int32(in.Int32())
		case "startTime":
			if in.IsNull() {
				in.Skip()
				out.StartTime = nil
			} else {
				if out.StartTime == nil {
					out.StartTime = new(_v1.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.StartTime).UnmarshalJSON(data))
				}
			}
		case "succeeded":
			out.Succeeded = int32(in.Int32())
		case "uncountedTerminatedPods":
			if in.IsNull() {
				in.Skip()
				out.UncountedTerminatedPods = nil
			} else {
				if out.UncountedTerminatedPods == nil {
					out.UncountedTerminatedPods = new(UncountedTerminatedPods)
				}
				easyjsonCe8ae4eDecodeGithubComKubewardenK8sObjectsApiBatchV11(in, out.UncountedTerminatedPods)
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
func easyjsonCe8ae4eEncodeGithubComKubewardenK8sObjectsApiBatchV1(out *jwriter.Writer, in JobStatus) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Active != 0 {
		const prefix string = ",\"active\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.Active))
	}
	if in.CompletedIndexes != "" {
		const prefix string = ",\"completedIndexes\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.CompletedIndexes))
	}
	if in.CompletionTime != nil {
		const prefix string = ",\"completionTime\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((*in.CompletionTime).MarshalJSON())
	}
	if len(in.Conditions) != 0 {
		const prefix string = ",\"conditions\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v2, v3 := range in.Conditions {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					(*v3).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if in.Failed != 0 {
		const prefix string = ",\"failed\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Failed))
	}
	if in.Ready != 0 {
		const prefix string = ",\"ready\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Ready))
	}
	if in.StartTime != nil {
		const prefix string = ",\"startTime\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((*in.StartTime).MarshalJSON())
	}
	if in.Succeeded != 0 {
		const prefix string = ",\"succeeded\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Succeeded))
	}
	if in.UncountedTerminatedPods != nil {
		const prefix string = ",\"uncountedTerminatedPods\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjsonCe8ae4eEncodeGithubComKubewardenK8sObjectsApiBatchV11(out, *in.UncountedTerminatedPods)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v JobStatus) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCe8ae4eEncodeGithubComKubewardenK8sObjectsApiBatchV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v JobStatus) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCe8ae4eEncodeGithubComKubewardenK8sObjectsApiBatchV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *JobStatus) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCe8ae4eDecodeGithubComKubewardenK8sObjectsApiBatchV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *JobStatus) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCe8ae4eDecodeGithubComKubewardenK8sObjectsApiBatchV1(l, v)
}
func easyjsonCe8ae4eDecodeGithubComKubewardenK8sObjectsApiBatchV11(in *jlexer.Lexer, out *UncountedTerminatedPods) {
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
		case "failed":
			if in.IsNull() {
				in.Skip()
				out.Failed = nil
			} else {
				in.Delim('[')
				if out.Failed == nil {
					if !in.IsDelim(']') {
						out.Failed = make([]string, 0, 4)
					} else {
						out.Failed = []string{}
					}
				} else {
					out.Failed = (out.Failed)[:0]
				}
				for !in.IsDelim(']') {
					var v4 string
					v4 = string(in.String())
					out.Failed = append(out.Failed, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "succeeded":
			if in.IsNull() {
				in.Skip()
				out.Succeeded = nil
			} else {
				in.Delim('[')
				if out.Succeeded == nil {
					if !in.IsDelim(']') {
						out.Succeeded = make([]string, 0, 4)
					} else {
						out.Succeeded = []string{}
					}
				} else {
					out.Succeeded = (out.Succeeded)[:0]
				}
				for !in.IsDelim(']') {
					var v5 string
					v5 = string(in.String())
					out.Succeeded = append(out.Succeeded, v5)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjsonCe8ae4eEncodeGithubComKubewardenK8sObjectsApiBatchV11(out *jwriter.Writer, in UncountedTerminatedPods) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Failed) != 0 {
		const prefix string = ",\"failed\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('[')
			for v6, v7 := range in.Failed {
				if v6 > 0 {
					out.RawByte(',')
				}
				out.String(string(v7))
			}
			out.RawByte(']')
		}
	}
	if len(in.Succeeded) != 0 {
		const prefix string = ",\"succeeded\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v8, v9 := range in.Succeeded {
				if v8 > 0 {
					out.RawByte(',')
				}
				out.String(string(v9))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
