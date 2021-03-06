// {{ .FuncName }} {{ if eq .Comment "" -}}
generated by sqlbrick, update data in database.
{{- else -}}
{{ .Comment }}
{{- end }}
{{ if .IsTx -}}
func (b *{{ .BrickName }}BrickTx){{ .FuncName }}(args *{{ .BrickName }}) (int64, error) {
{{- else -}}
func (b *{{ .BrickName }}Brick){{ .FuncName }}(args *{{ .BrickName }}) (int64, error) {
{{- end -}}
    {{- $conditionLen := len .Conditions -}}
    {{- if gt $conditionLen 0 }}
    conditionQuery := `{{ index .Segments 0 }}`
    {{- $segments := .Segments -}}
    {{- $segmentLen := len .Segments -}}
    {{- CacheSet "segment" 1 -}}
    {{- $removeComma := .RemoveComma -}}
    {{- $indexOfWhere := .IndexOfWhere -}}
    {{- range $k, $v := .Conditions }}
    if args.{{ $v.Expression }} {
        conditionQuery += `{{ $v.Query }}`
    }
    {{- if not $v.AppendNext }}
    {{- $segmentIndex := CacheGet "segment" -}}
    {{- $segment := index $segments $segmentIndex -}}
    {{- if and $removeComma (eq $segmentIndex $indexOfWhere) }}
    if strings.HasSuffix(conditionQuery, ",") {
        conditionQuery = strings.TrimSuffix(conditionQuery, ",")
    }
    {{- end }}
    conditionQuery += `{{ $segment }}`
    {{- $segmentIndex := Add $segmentIndex 1 -}}
    {{- CacheSet "segment" $segmentIndex -}}
    {{- end -}}
    {{- end }}

    {{ if .IsTx -}}
    if err := b.checkTx(); err != nil {
        return 0, err
    }

    stmt, err := b.tx.PrepareNamed(conditionQuery)
    {{- else -}}
    stmt, err := b.db.PrepareNamed(conditionQuery)
    {{- end -}}
    {{- else -}}
    {{- if .IsTx -}}
    stmt, err := b.tx.PrepareNamed(
        `{{ index .Segments 0 }}`)
    {{- else -}}
    stmt, err := b.db.PrepareNamed(
            `{{ index .Segments 0 }}`)
    {{- end -}}
    {{- end }}
    if err != nil {
        return 0, err
    }

    result, err := stmt.Exec(args)
    if err != nil {
    {{- if .IsTx }}
        if rbe := b.tx.Rollback(); rbe != nil {
            return 0, rbe
        }
    {{- end }}
        return 0, err
    }

    return result.RowsAffected()
}
