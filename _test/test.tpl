{{ with $values := file "./_test/config-values.json" | parseJSON -}}
{
    "Severity": {{indexOrDefault $values "DEV123"}},
}
{{- end -}}