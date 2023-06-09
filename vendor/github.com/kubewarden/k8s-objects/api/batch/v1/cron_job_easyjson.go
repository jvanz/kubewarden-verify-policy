// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1

import (
	json "encoding/json"
	_v11 "github.com/kubewarden/k8s-objects/api/core/v1"
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

func easyjsonA56f4d58DecodeGithubComKubewardenK8sObjectsApiBatchV1(in *jlexer.Lexer, out *CronJob) {
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
		case "apiVersion":
			out.APIVersion = string(in.String())
		case "kind":
			out.Kind = string(in.String())
		case "metadata":
			if in.IsNull() {
				in.Skip()
				out.Metadata = nil
			} else {
				if out.Metadata == nil {
					out.Metadata = new(_v1.ObjectMeta)
				}
				(*out.Metadata).UnmarshalEasyJSON(in)
			}
		case "spec":
			if in.IsNull() {
				in.Skip()
				out.Spec = nil
			} else {
				if out.Spec == nil {
					out.Spec = new(CronJobSpec)
				}
				easyjsonA56f4d58DecodeGithubComKubewardenK8sObjectsApiBatchV11(in, out.Spec)
			}
		case "status":
			if in.IsNull() {
				in.Skip()
				out.Status = nil
			} else {
				if out.Status == nil {
					out.Status = new(CronJobStatus)
				}
				easyjsonA56f4d58DecodeGithubComKubewardenK8sObjectsApiBatchV12(in, out.Status)
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
func easyjsonA56f4d58EncodeGithubComKubewardenK8sObjectsApiBatchV1(out *jwriter.Writer, in CronJob) {
	out.RawByte('{')
	first := true
	_ = first
	if in.APIVersion != "" {
		const prefix string = ",\"apiVersion\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.APIVersion))
	}
	if in.Kind != "" {
		const prefix string = ",\"kind\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Kind))
	}
	if in.Metadata != nil {
		const prefix string = ",\"metadata\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Metadata).MarshalEasyJSON(out)
	}
	if in.Spec != nil {
		const prefix string = ",\"spec\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjsonA56f4d58EncodeGithubComKubewardenK8sObjectsApiBatchV11(out, *in.Spec)
	}
	if in.Status != nil {
		const prefix string = ",\"status\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjsonA56f4d58EncodeGithubComKubewardenK8sObjectsApiBatchV12(out, *in.Status)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CronJob) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA56f4d58EncodeGithubComKubewardenK8sObjectsApiBatchV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CronJob) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA56f4d58EncodeGithubComKubewardenK8sObjectsApiBatchV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CronJob) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA56f4d58DecodeGithubComKubewardenK8sObjectsApiBatchV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CronJob) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA56f4d58DecodeGithubComKubewardenK8sObjectsApiBatchV1(l, v)
}
func easyjsonA56f4d58DecodeGithubComKubewardenK8sObjectsApiBatchV12(in *jlexer.Lexer, out *CronJobStatus) {
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
			if in.IsNull() {
				in.Skip()
				out.Active = nil
			} else {
				in.Delim('[')
				if out.Active == nil {
					if !in.IsDelim(']') {
						out.Active = make([]*_v11.ObjectReference, 0, 8)
					} else {
						out.Active = []*_v11.ObjectReference{}
					}
				} else {
					out.Active = (out.Active)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *_v11.ObjectReference
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(_v11.ObjectReference)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Active = append(out.Active, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "lastScheduleTime":
			if in.IsNull() {
				in.Skip()
				out.LastScheduleTime = nil
			} else {
				if out.LastScheduleTime == nil {
					out.LastScheduleTime = new(_v1.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.LastScheduleTime).UnmarshalJSON(data))
				}
			}
		case "lastSuccessfulTime":
			if in.IsNull() {
				in.Skip()
				out.LastSuccessfulTime = nil
			} else {
				if out.LastSuccessfulTime == nil {
					out.LastSuccessfulTime = new(_v1.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.LastSuccessfulTime).UnmarshalJSON(data))
				}
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
func easyjsonA56f4d58EncodeGithubComKubewardenK8sObjectsApiBatchV12(out *jwriter.Writer, in CronJobStatus) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Active) != 0 {
		const prefix string = ",\"active\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('[')
			for v2, v3 := range in.Active {
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
	if in.LastScheduleTime != nil {
		const prefix string = ",\"lastScheduleTime\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((*in.LastScheduleTime).MarshalJSON())
	}
	if in.LastSuccessfulTime != nil {
		const prefix string = ",\"lastSuccessfulTime\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((*in.LastSuccessfulTime).MarshalJSON())
	}
	out.RawByte('}')
}
func easyjsonA56f4d58DecodeGithubComKubewardenK8sObjectsApiBatchV11(in *jlexer.Lexer, out *CronJobSpec) {
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
		case "concurrencyPolicy":
			out.ConcurrencyPolicy = string(in.String())
		case "failedJobsHistoryLimit":
			out.FailedJobsHistoryLimit = int32(in.Int32())
		case "jobTemplate":
			if in.IsNull() {
				in.Skip()
				out.JobTemplate = nil
			} else {
				if out.JobTemplate == nil {
					out.JobTemplate = new(JobTemplateSpec)
				}
				easyjsonA56f4d58DecodeGithubComKubewardenK8sObjectsApiBatchV13(in, out.JobTemplate)
			}
		case "schedule":
			if in.IsNull() {
				in.Skip()
				out.Schedule = nil
			} else {
				if out.Schedule == nil {
					out.Schedule = new(string)
				}
				*out.Schedule = string(in.String())
			}
		case "startingDeadlineSeconds":
			out.StartingDeadlineSeconds = int64(in.Int64())
		case "successfulJobsHistoryLimit":
			out.SuccessfulJobsHistoryLimit = int32(in.Int32())
		case "suspend":
			out.Suspend = bool(in.Bool())
		case "timeZone":
			out.TimeZone = string(in.String())
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
func easyjsonA56f4d58EncodeGithubComKubewardenK8sObjectsApiBatchV11(out *jwriter.Writer, in CronJobSpec) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ConcurrencyPolicy != "" {
		const prefix string = ",\"concurrencyPolicy\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.ConcurrencyPolicy))
	}
	if in.FailedJobsHistoryLimit != 0 {
		const prefix string = ",\"failedJobsHistoryLimit\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.FailedJobsHistoryLimit))
	}
	{
		const prefix string = ",\"jobTemplate\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.JobTemplate == nil {
			out.RawString("null")
		} else {
			easyjsonA56f4d58EncodeGithubComKubewardenK8sObjectsApiBatchV13(out, *in.JobTemplate)
		}
	}
	{
		const prefix string = ",\"schedule\":"
		out.RawString(prefix)
		if in.Schedule == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Schedule))
		}
	}
	if in.StartingDeadlineSeconds != 0 {
		const prefix string = ",\"startingDeadlineSeconds\":"
		out.RawString(prefix)
		out.Int64(int64(in.StartingDeadlineSeconds))
	}
	if in.SuccessfulJobsHistoryLimit != 0 {
		const prefix string = ",\"successfulJobsHistoryLimit\":"
		out.RawString(prefix)
		out.Int32(int32(in.SuccessfulJobsHistoryLimit))
	}
	if in.Suspend {
		const prefix string = ",\"suspend\":"
		out.RawString(prefix)
		out.Bool(bool(in.Suspend))
	}
	if in.TimeZone != "" {
		const prefix string = ",\"timeZone\":"
		out.RawString(prefix)
		out.String(string(in.TimeZone))
	}
	out.RawByte('}')
}
func easyjsonA56f4d58DecodeGithubComKubewardenK8sObjectsApiBatchV13(in *jlexer.Lexer, out *JobTemplateSpec) {
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
		case "metadata":
			if in.IsNull() {
				in.Skip()
				out.Metadata = nil
			} else {
				if out.Metadata == nil {
					out.Metadata = new(_v1.ObjectMeta)
				}
				(*out.Metadata).UnmarshalEasyJSON(in)
			}
		case "spec":
			if in.IsNull() {
				in.Skip()
				out.Spec = nil
			} else {
				if out.Spec == nil {
					out.Spec = new(JobSpec)
				}
				easyjsonA56f4d58DecodeGithubComKubewardenK8sObjectsApiBatchV14(in, out.Spec)
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
func easyjsonA56f4d58EncodeGithubComKubewardenK8sObjectsApiBatchV13(out *jwriter.Writer, in JobTemplateSpec) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Metadata != nil {
		const prefix string = ",\"metadata\":"
		first = false
		out.RawString(prefix[1:])
		(*in.Metadata).MarshalEasyJSON(out)
	}
	if in.Spec != nil {
		const prefix string = ",\"spec\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjsonA56f4d58EncodeGithubComKubewardenK8sObjectsApiBatchV14(out, *in.Spec)
	}
	out.RawByte('}')
}
func easyjsonA56f4d58DecodeGithubComKubewardenK8sObjectsApiBatchV14(in *jlexer.Lexer, out *JobSpec) {
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
		case "activeDeadlineSeconds":
			out.ActiveDeadlineSeconds = int64(in.Int64())
		case "backoffLimit":
			out.BackoffLimit = int32(in.Int32())
		case "completionMode":
			out.CompletionMode = string(in.String())
		case "completions":
			out.Completions = int32(in.Int32())
		case "manualSelector":
			out.ManualSelector = bool(in.Bool())
		case "parallelism":
			out.Parallelism = int32(in.Int32())
		case "podFailurePolicy":
			if in.IsNull() {
				in.Skip()
				out.PodFailurePolicy = nil
			} else {
				if out.PodFailurePolicy == nil {
					out.PodFailurePolicy = new(PodFailurePolicy)
				}
				easyjsonA56f4d58DecodeGithubComKubewardenK8sObjectsApiBatchV15(in, out.PodFailurePolicy)
			}
		case "selector":
			if in.IsNull() {
				in.Skip()
				out.Selector = nil
			} else {
				if out.Selector == nil {
					out.Selector = new(_v1.LabelSelector)
				}
				(*out.Selector).UnmarshalEasyJSON(in)
			}
		case "suspend":
			out.Suspend = bool(in.Bool())
		case "template":
			if in.IsNull() {
				in.Skip()
				out.Template = nil
			} else {
				if out.Template == nil {
					out.Template = new(_v11.PodTemplateSpec)
				}
				(*out.Template).UnmarshalEasyJSON(in)
			}
		case "ttlSecondsAfterFinished":
			out.TTLSecondsAfterFinished = int32(in.Int32())
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
func easyjsonA56f4d58EncodeGithubComKubewardenK8sObjectsApiBatchV14(out *jwriter.Writer, in JobSpec) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ActiveDeadlineSeconds != 0 {
		const prefix string = ",\"activeDeadlineSeconds\":"
		first = false
		out.RawString(prefix[1:])
		out.Int64(int64(in.ActiveDeadlineSeconds))
	}
	if in.BackoffLimit != 0 {
		const prefix string = ",\"backoffLimit\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.BackoffLimit))
	}
	if in.CompletionMode != "" {
		const prefix string = ",\"completionMode\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.CompletionMode))
	}
	if in.Completions != 0 {
		const prefix string = ",\"completions\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Completions))
	}
	if in.ManualSelector {
		const prefix string = ",\"manualSelector\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.ManualSelector))
	}
	if in.Parallelism != 0 {
		const prefix string = ",\"parallelism\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Parallelism))
	}
	if in.PodFailurePolicy != nil {
		const prefix string = ",\"podFailurePolicy\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjsonA56f4d58EncodeGithubComKubewardenK8sObjectsApiBatchV15(out, *in.PodFailurePolicy)
	}
	if in.Selector != nil {
		const prefix string = ",\"selector\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Selector).MarshalEasyJSON(out)
	}
	if in.Suspend {
		const prefix string = ",\"suspend\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Suspend))
	}
	{
		const prefix string = ",\"template\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Template == nil {
			out.RawString("null")
		} else {
			(*in.Template).MarshalEasyJSON(out)
		}
	}
	if in.TTLSecondsAfterFinished != 0 {
		const prefix string = ",\"ttlSecondsAfterFinished\":"
		out.RawString(prefix)
		out.Int32(int32(in.TTLSecondsAfterFinished))
	}
	out.RawByte('}')
}
func easyjsonA56f4d58DecodeGithubComKubewardenK8sObjectsApiBatchV15(in *jlexer.Lexer, out *PodFailurePolicy) {
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
		case "rules":
			if in.IsNull() {
				in.Skip()
				out.Rules = nil
			} else {
				in.Delim('[')
				if out.Rules == nil {
					if !in.IsDelim(']') {
						out.Rules = make([]*PodFailurePolicyRule, 0, 8)
					} else {
						out.Rules = []*PodFailurePolicyRule{}
					}
				} else {
					out.Rules = (out.Rules)[:0]
				}
				for !in.IsDelim(']') {
					var v4 *PodFailurePolicyRule
					if in.IsNull() {
						in.Skip()
						v4 = nil
					} else {
						if v4 == nil {
							v4 = new(PodFailurePolicyRule)
						}
						easyjsonA56f4d58DecodeGithubComKubewardenK8sObjectsApiBatchV16(in, v4)
					}
					out.Rules = append(out.Rules, v4)
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
func easyjsonA56f4d58EncodeGithubComKubewardenK8sObjectsApiBatchV15(out *jwriter.Writer, in PodFailurePolicy) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"rules\":"
		out.RawString(prefix[1:])
		if in.Rules == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Rules {
				if v5 > 0 {
					out.RawByte(',')
				}
				if v6 == nil {
					out.RawString("null")
				} else {
					easyjsonA56f4d58EncodeGithubComKubewardenK8sObjectsApiBatchV16(out, *v6)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjsonA56f4d58DecodeGithubComKubewardenK8sObjectsApiBatchV16(in *jlexer.Lexer, out *PodFailurePolicyRule) {
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
		case "action":
			if in.IsNull() {
				in.Skip()
				out.Action = nil
			} else {
				if out.Action == nil {
					out.Action = new(string)
				}
				*out.Action = string(in.String())
			}
		case "onExitCodes":
			if in.IsNull() {
				in.Skip()
				out.OnExitCodes = nil
			} else {
				if out.OnExitCodes == nil {
					out.OnExitCodes = new(PodFailurePolicyOnExitCodesRequirement)
				}
				easyjsonA56f4d58DecodeGithubComKubewardenK8sObjectsApiBatchV17(in, out.OnExitCodes)
			}
		case "onPodConditions":
			if in.IsNull() {
				in.Skip()
				out.OnPodConditions = nil
			} else {
				in.Delim('[')
				if out.OnPodConditions == nil {
					if !in.IsDelim(']') {
						out.OnPodConditions = make([]*PodFailurePolicyOnPodConditionsPattern, 0, 8)
					} else {
						out.OnPodConditions = []*PodFailurePolicyOnPodConditionsPattern{}
					}
				} else {
					out.OnPodConditions = (out.OnPodConditions)[:0]
				}
				for !in.IsDelim(']') {
					var v7 *PodFailurePolicyOnPodConditionsPattern
					if in.IsNull() {
						in.Skip()
						v7 = nil
					} else {
						if v7 == nil {
							v7 = new(PodFailurePolicyOnPodConditionsPattern)
						}
						easyjsonA56f4d58DecodeGithubComKubewardenK8sObjectsApiBatchV18(in, v7)
					}
					out.OnPodConditions = append(out.OnPodConditions, v7)
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
func easyjsonA56f4d58EncodeGithubComKubewardenK8sObjectsApiBatchV16(out *jwriter.Writer, in PodFailurePolicyRule) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"action\":"
		out.RawString(prefix[1:])
		if in.Action == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Action))
		}
	}
	if in.OnExitCodes != nil {
		const prefix string = ",\"onExitCodes\":"
		out.RawString(prefix)
		easyjsonA56f4d58EncodeGithubComKubewardenK8sObjectsApiBatchV17(out, *in.OnExitCodes)
	}
	{
		const prefix string = ",\"onPodConditions\":"
		out.RawString(prefix)
		if in.OnPodConditions == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.OnPodConditions {
				if v8 > 0 {
					out.RawByte(',')
				}
				if v9 == nil {
					out.RawString("null")
				} else {
					easyjsonA56f4d58EncodeGithubComKubewardenK8sObjectsApiBatchV18(out, *v9)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjsonA56f4d58DecodeGithubComKubewardenK8sObjectsApiBatchV18(in *jlexer.Lexer, out *PodFailurePolicyOnPodConditionsPattern) {
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
func easyjsonA56f4d58EncodeGithubComKubewardenK8sObjectsApiBatchV18(out *jwriter.Writer, in PodFailurePolicyOnPodConditionsPattern) {
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
func easyjsonA56f4d58DecodeGithubComKubewardenK8sObjectsApiBatchV17(in *jlexer.Lexer, out *PodFailurePolicyOnExitCodesRequirement) {
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
		case "containerName":
			out.ContainerName = string(in.String())
		case "operator":
			if in.IsNull() {
				in.Skip()
				out.Operator = nil
			} else {
				if out.Operator == nil {
					out.Operator = new(string)
				}
				*out.Operator = string(in.String())
			}
		case "values":
			if in.IsNull() {
				in.Skip()
				out.Values = nil
			} else {
				in.Delim('[')
				if out.Values == nil {
					if !in.IsDelim(']') {
						out.Values = make([]int32, 0, 16)
					} else {
						out.Values = []int32{}
					}
				} else {
					out.Values = (out.Values)[:0]
				}
				for !in.IsDelim(']') {
					var v10 int32
					v10 = int32(in.Int32())
					out.Values = append(out.Values, v10)
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
func easyjsonA56f4d58EncodeGithubComKubewardenK8sObjectsApiBatchV17(out *jwriter.Writer, in PodFailurePolicyOnExitCodesRequirement) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ContainerName != "" {
		const prefix string = ",\"containerName\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.ContainerName))
	}
	{
		const prefix string = ",\"operator\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Operator == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Operator))
		}
	}
	{
		const prefix string = ",\"values\":"
		out.RawString(prefix)
		if in.Values == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v11, v12 := range in.Values {
				if v11 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v12))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
