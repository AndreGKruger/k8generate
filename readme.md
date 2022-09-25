# k8generate
## ⚡️ Quick start

First, [download](https://golang.org/dl/) and install **Go**. Version `1.19` or higher is required.

To install use the [`go install`](https://golang.org/cmd/go/#hdr-Compile_and_install_packages_and_dependencies) command and rename installed binary in `$GOPATH/bin`:

```bash
go install github.com/AndreGKruger/k8generate@latest
```
---

## 📖 Documentation
run `k8generate --help` for more information.

---
### Usage:
  - k8generate [command]

---

### Available Commands:
    - all         Generates all the k8 files (configmap, secret, deployment, service)    
    - configmap   Generates a k8_configmap.yaml file
    - deployment  Generates a k8_deployment.yaml file
    - help        Help about any command
    - secret      Generates a k8_secrets.yaml file
    - service     Generates a k8_service.yaml file
---

### Flags:
    -h, --help   help for k8generate
---

### Common Flags for ConfigMap, Secret, Service and Deployment
    -h, --help                                 help for configmap
    -a, --appname string        required       name of the application
    -e, --env string            required       name of the environment IE:production, staging, development    
    -n, --namespace string      optional       namespace of the application, defaults to appname-env IE myapp-production
Note:
    uses .env in the root of the folder file to generate the configmap, secret, and deployment file env entries

---

### Deployment Flags
    -h, --help                                 help for deployment
    -r, --repoendpoint string   required       endpoint of the repository IE: xyz.dkr.ecr.eu-west-1.amazonaws.com
    -p, --reponame string       required       name of the repository IE: myrepo/myapp
    -v, --repoversion string    required       version of the repository IE: 1.0.0
---

### Examples
    k8generate deployment -a myapp -e production -r xyz.dkr.ecr.eu-west-1.amazonaws.com -p myrepo/myapp -v 1.0.0
    k8generate configmap -a myapp -e production
    k8generate secret -a myapp -e production
    k8generate service -a myapp -e production
---

### Files Generated
    k8_configmap.yaml
    k8_deployment.yaml
    k8_secret.yaml
    k8_service.yaml
Files are generated in ./kubernetes/{env}/ IE : ./kubernetes/production/k8_configmap.yaml

### Secrets 
    Any env key that contains one of the following strings will be considered a secret and will be generated in the k8_secret.yaml file
    "secret", "api", "key", "pass", "user", "token"

---