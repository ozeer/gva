package global

{{- if .HasGlobal }}

import "github.com/ozeer/gva/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}