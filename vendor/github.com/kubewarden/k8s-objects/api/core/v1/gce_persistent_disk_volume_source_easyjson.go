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

func easyjson2cae620eDecodeGithubComKubewardenK8sObjectsApiCoreV1(in *jlexer.Lexer, out *GCEPersistentDiskVolumeSource) {
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
		case "fsType":
			out.FSType = string(in.String())
		case "partition":
			out.Partition = int32(in.Int32())
		case "pdName":
			if in.IsNull() {
				in.Skip()
				out.PdName = nil
			} else {
				if out.PdName == nil {
					out.PdName = new(string)
				}
				*out.PdName = string(in.String())
			}
		case "readOnly":
			out.ReadOnly = bool(in.Bool())
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
func easyjson2cae620eEncodeGithubComKubewardenK8sObjectsApiCoreV1(out *jwriter.Writer, in GCEPersistentDiskVolumeSource) {
	out.RawByte('{')
	first := true
	_ = first
	if in.FSType != "" {
		const prefix string = ",\"fsType\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.FSType))
	}
	if in.Partition != 0 {
		const prefix string = ",\"partition\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Partition))
	}
	{
		const prefix string = ",\"pdName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.PdName == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.PdName))
		}
	}
	if in.ReadOnly {
		const prefix string = ",\"readOnly\":"
		out.RawString(prefix)
		out.Bool(bool(in.ReadOnly))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GCEPersistentDiskVolumeSource) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson2cae620eEncodeGithubComKubewardenK8sObjectsApiCoreV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GCEPersistentDiskVolumeSource) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson2cae620eEncodeGithubComKubewardenK8sObjectsApiCoreV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GCEPersistentDiskVolumeSource) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson2cae620eDecodeGithubComKubewardenK8sObjectsApiCoreV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GCEPersistentDiskVolumeSource) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson2cae620eDecodeGithubComKubewardenK8sObjectsApiCoreV1(l, v)
}
