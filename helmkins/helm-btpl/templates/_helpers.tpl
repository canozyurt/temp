{{- define "resourcename" -}}
{{- ternary ( .Release.Name | lower) ( list (regexReplaceAll ".*/" .Values.branch "" | default "development") .Release.Name | join "-" | lower) (ne .Release.Namespace "test") -}}
{{- end -}}