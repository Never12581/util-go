package {{.PackageName}};

import lombok.Getter;
import lombok.Setter;
{{range .Imports}}import {{.ImportName}};
{{end}}
@Getter
@Setter
public class {{.ClassName}} {
    {{if .Fields}}{{range .Fields}}{{if .FieldNote}}
    /*
     * {{.FieldNote}}
     */{{end}}
    private {{.FieldClass}} {{.FieldName}};
    {{end}}{{end}}
}
