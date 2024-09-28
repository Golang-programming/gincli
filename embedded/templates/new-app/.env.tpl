{{- if .DBHost }}
DB_HOST={{ .DBHost }}
{{- end }}

{{- if .DBUsername }}
DB_USER={{ .DBUsername }}
{{- end }}

{{- if .DBPassword }}
DB_PASSWORD={{ .DBPassword }}
{{- end }}

{{- if .DBName }}
DB_NAME={{ .DBName }}
{{- end }}

{{- if .DBPort }}
DB_PORT={{ .DBPort }}
{{- end }}

{{- if .DBUri }}
MONGODB_URI={{ .DBUri }}
{{- end }}