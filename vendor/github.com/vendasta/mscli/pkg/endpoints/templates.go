package endpoints

//GRPCService represents the information Cloud Endpoints needs to know about a GRPC service
type GRPCService struct {
	Name            string
	EnvironmentName string
	ProjectName     string
	Apis            []string
	VerifyIdentity  bool
	PublicRoutes    []string
}

//EndpointsEnvironment is the template to generate the CloudEndpoints YAML input
const EndpointsEnvironment = `type: google.api.Service
config_version: 3
name: {{.Name}}
title: {{.ProjectName}} gRPC {{.EnvironmentName}}

apis:
{{ range .Apis}}
- name: {{ . }}
{{end}}

usage:
  rules:
    - selector: '*'
      allow_unregistered_calls: true
{{if .VerifyIdentity}}
authentication:
  rules:
    - selector: '*'
      requirements:
        - provider_id: google_id_token
          audiences: "https://{{.Name}}"
        - provider_id: iam_token
          audiences: "https://iam-{{.EnvironmentName}}.vendasta-internal.com"
  {{- range $i, $route := .PublicRoutes}}
    - selector: {{ $route }}
      requirements:
  {{ end -}}
  providers:
    - id: google_id_token
      issuer: accounts.google.com
    - id: iam_token
      jwks_uri: https://iam-{{.EnvironmentName}}.vendasta-internal.com/oauth2/v1/certs
      issuer: iam-{{.EnvironmentName}}.vendasta-internal.com
{{end}}
`

//LocalServiceAccountJSON is the service account to use when running our µs locally
const LocalServiceAccountJSON = `{
  "type": "service_account",
  "project_id": "repcore-prod",
  "private_key_id": "ea4e97e2cd377952a2d48660f15b1ad8178e8646",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQDI3e/Bkb8xtGSa\neve5wNICIz6V6nnADEGw4yqV2NCgmL1OnxfZR+iAIMy7cufpmcUuLXM19YF7ENea\nEGYgikdzvXLtclkb7g5z5tAvzz6qBg5uHzIe9R0YZ6sRhyFXAWdcph/dncp/it9L\nd2uabpGtw6xVYTXvN+ulVqbZjl6lJFOkdOYvJbjgPaz9Q8RdjdpBIJY3iFjnNgew\nU+ZOdrZQpEWeCrJdciEJfVPoqYYJrb2/jn76SZe5hxTew0z/BldGYb/hqygIeAZ5\n76bmd3O1nv2Vj+/ymQiJz7tXHMYuhtykePRf7HcycctQyE85wjPgx0FzLLyLD80z\ng8qvcZedAgMBAAECggEBALaYyjL1AxF0I2ArGLdLiZ1X3LT7ieIDQlRCrZu6lSzo\nrG6kQVHMBJc1OefM9PVuS8lGl80deK8fCF+NOMmr6nhUZ/1iTY8OQSUrVWWgyB5X\nkxaYVT+/IsvksaYlaTFmnUej3KPhpQt8erBnxvGEQfnsPuT+kcD83MJYmyR6P59u\nYjJDSRAKar38lShfvh8SWdhXcBXjvqPhsBRkgLT46d2dU9zhS0NRgZBedzpdSEPO\nrRpaqy6rl4knVEi7qphRwGQhQGlBP9Za96bfcFjK/Gs96VV+uz22rZinK4avNJoO\n39SdRHqo2k6vAFrDZomafY4yKxs39NNCw5Xra2gB+yECgYEA4uM2K5Vmhzw5AmFP\nUU2NojqTob+/4KZCYDA34hLRayNClntDnpm+1xz9BSdX6w1cjVMFpwhB/iFfMAt8\npGEYuBXv5C22ZSan5j4lA3IMiaGxQEEnOUs/ys4/7db2E7F2BtNvrhuK2hTo2LkT\nrVKTLbz3Ra7cjW7M+oO9Y52YWT8CgYEA4qQAbMIoi5E3ogc43Vvmz5Q1q/rX4p2P\nXJpRnyBGdF3JbOIgzHLCFX3ycES4RqftkBaysKXJ1WN/DHX6o3F+0vgPhckxiKaA\nfBQ6d8Lz/AeldgTLGF3GcLcs6y0wg8QJrUWeXA/dqAO/RCjt3QI7vy8xMaVK7Jzm\nFK+kS7EFnCMCgYB2Do4gTYPk+GQdpe60umrpMBujfXfk9/3vuQdK/kmzuswqKwd7\nXjqcCfxqExe/OdufucRmLnjQOCMkh8Wabt+C4f9KNrMCnQOeXVW7HjwB3X2ylnRH\nbq1J5NoE8uZmRt/IG3qwGKq+YUTriBoCBMaKdRohSyR9/1pvO99vFxCv1wKBgQCt\n2BqzUeMD22oPEPcfZIURKKBawNij7TOwVnbRVlJ0pBQDPxjleglBrxAt6ahoGhtx\ncNe0BBdjZSRpDH4qrL4ZWUme14r1RLI9es62WoYIBl12v4qBsDys7PS5mDokeFTr\n4gDsoQbm/6jQwYxjAHmP1sUnTaDBCkC6EVg27xxcfQKBgQCzA9PKB0p6OSlxSIt5\nB4+DbwBNTbtRW8hn+y7VwI8DKoGTmvw2G2tr2eW4ipNXJaPaI5QN1ZOKpHWryA9c\nVjBBVS/KRza6iOfmhAvbXVgEilvuxPEC+9So4b2GlkehZQKeGcKEIg4K5CfLHw0M\nSNZdxWWrOtAYXiouVuuXFql8GA==\n-----END PRIVATE KEY-----\n",
  "client_email": "vendasta-local@repcore-prod.iam.gserviceaccount.com",
  "client_id": "113285331591202673944",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://accounts.google.com/o/oauth2/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/vendasta-local%40repcore-prod.iam.gserviceaccount.com"
}`
