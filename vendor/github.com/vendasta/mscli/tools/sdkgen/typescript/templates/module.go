package templates

//ModuleTemplate is used to generate the <name>.module.ts
const ModuleTemplate = `{{- $template := . -}}
import {NgModule} from '@angular/core';
import {SessionServiceModule} from '../session';
import {EnvironmentServiceModule} from '../environment';
import {
  {{- range $template.Services -}}{{- $service := . }}
  {{ $service.TsName }}ApiService,
  {{- end }}
} from './_internal';
import {HostService} from './_generated';

@NgModule({
  imports: [SessionServiceModule, EnvironmentServiceModule],
  providers: [
    {{- range $template.Services -}}{{- $service := . }}
    {{ $service.TsName }}ApiService,
    {{- end }}
	HostService
  ],
  declarations: [],
  exports: [],
  entryComponents: []
})
export class {{ $template.TsName }}Module {}
`
