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

func easyjson24779c0DecodeGithubComKubewardenK8sObjectsApiCoreV1(in *jlexer.Lexer, out *TypedObjectReference) {
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
		case "apiGroup":
			out.APIGroup = string(in.String())
		case "kind":
			if in.IsNull() {
				in.Skip()
				out.Kind = nil
			} else {
				if out.Kind == nil {
					out.Kind = new(string)
				}
				*out.Kind = string(in.String())
			}
		case "name":
			if in.IsNull() {
				in.Skip()
				out.Name = nil
			} else {
				if out.Name == nil {
					out.Name = new(string)
				}
				*out.Name = string(in.String())
			}
		case "namespace":
			out.Namespace = string(in.String())
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
func easyjson24779c0EncodeGithubComKubewardenK8sObjectsApiCoreV1(out *jwriter.Writer, in TypedObjectReference) {
	out.RawByte('{')
	first := true
	_ = first
	if in.APIGroup != "" {
		const prefix string = ",\"apiGroup\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.APIGroup))
	}
	{
		const prefix string = ",\"kind\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Kind == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Kind))
		}
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		if in.Name == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Name))
		}
	}
	if in.Namespace != "" {
		const prefix string = ",\"namespace\":"
		out.RawString(prefix)
		out.String(string(in.Namespace))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v TypedObjectReference) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson24779c0EncodeGithubComKubewardenK8sObjectsApiCoreV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TypedObjectReference) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson24779c0EncodeGithubComKubewardenK8sObjectsApiCoreV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TypedObjectReference) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson24779c0DecodeGithubComKubewardenK8sObjectsApiCoreV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TypedObjectReference) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson24779c0DecodeGithubComKubewardenK8sObjectsApiCoreV1(l, v)
}
