package env_vars

import (
	"html"
	"strings"
)

type EnvVars interface {
	GenerateVarsFromFileBytes(file []byte, lowercase bool) []Envvar
	GenerateSecretsFromFileBytes(file []byte, lowercase bool) []Envvar
}

type Envvar struct {
	Name  string
	Value string
}

type env struct{}

func New() EnvVars {
	return &env{}
}

func (e *env) GenerateVarsFromFileBytes(file []byte, lowercase bool) []Envvar {
	result := make([]Envvar, 0)
	//Split the file into lines
	lines := strings.Split(string(file), "\n")
	//Loop through the lines and create a slice of envvars
	for _, line := range lines {
		if line != "" && !strings.HasPrefix(line, "#") {
			if !containsSecrets(line) {
				env := strings.Split(line, "=")
				name := strings.ReplaceAll(html.EscapeString(strings.ToLower(env[0])), "&#34;", "")
				value := strings.ReplaceAll(html.EscapeString(env[1]), "&#34;", "")
				if !lowercase {
					name = strings.ToUpper(env[0])
					value = strings.ToLower(env[0])
				}
				result = append(result, Envvar{Name: name, Value: value})
			}
		}
	}
	return result
}

func (e *env) GenerateSecretsFromFileBytes(file []byte, lowercase bool) []Envvar {
	result := make([]Envvar, 0)
	//Split the file into lines
	lines := strings.Split(string(file), "\n")
	//Loop through the lines and create a slice of envvars
	for _, line := range lines {
		if line != "" && !strings.HasPrefix(line, "#") {
			if containsSecrets(line) {
				env := strings.Split(line, "=")
				name := strings.ReplaceAll(html.EscapeString(strings.ToLower(env[0])), "&#34;", "")
				value := strings.ReplaceAll(html.EscapeString(env[1]), "&#34;", "")
				if !lowercase {
					name = strings.ToUpper(env[0])
					value = strings.ToLower(env[0])
				}
				result = append(result, Envvar{Name: name, Value: value})
			}
		}
	}
	return result
}

var listOfSecrets = []string{"secret", "api", "key", "pass", "user", "token"}

func containsSecrets(envvar string) bool {
	//Check if the envvar contains any of the secrets
	for _, secret := range listOfSecrets {
		if strings.Contains(strings.ToLower(envvar), secret) {
			return true
		}
	}
	return false
}
