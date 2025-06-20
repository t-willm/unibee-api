# UniBee Recurring Billing API

# Introduction

The UniBee API is an essential service for the UniBee Billing system, designed to empower our customers to effectively manage their subscription plans. It provides seamless access to multiple payment channels, subscription management, invoicing services, and more, simplifying and streamlining the billing process.

# Local Environment

## Prerequisites

- Docker Desktop 
- Make

## Description

The local environment is based on Docker Desktop and Make, which makes it easy to run and manage the UniBee API locally.

It uses the following Docker images:
- chdotworld/dotworld:golang-air-atlas-ubuntu
- redis:alpine
- mysql:8.0.37

All necessary ports are exposed, including the Air port, Redis port, and MySQL port.

Configured using `.env` in `.devcontainer/docker` directory and `config.yaml.docker`.

**So you need to set up the config.yaml first.**

```bash
cp config.yaml.docker config.yaml
```

Then we can execute `make run` to start the container.

## Usage As Local Environment (no codespace)

### How To Run

```bash
make run
```

### How to enter the container
```bash
make bash
or 
make zsh
```

### Serve

```bash
make serve
```

### Share

```bash
make expose
```

## Usage As Codespace Environment

### How To Serve

```bash
dotdev serve
```

### How to share 

```bash
dotdev share --port 8088
```

# Infra
### Gateway
- Stripe
- Paypal
- Changelly

### VatGateway
- VatSense

### Email
- SendGrid


#### Mysql Integration

- \_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

#### Redis Integration

- \_ "github.com/gogf/gf/contrib/nosql/redis/v2"

## How To Run

Manual setting `config.yaml` in project dir before run unibee-api, the template `config.yaml` is `manifest/config/config.yaml.template` 
```bash
go run main.go
```

#### OpenAPI V3 Doc：http://127.0.0.1:8088/swagger

#### Swagger TryOut API：http://127.0.0.1:8088/swagger-ui.html

#### OpenAPI V3 Json：http://127.0.0.1:8088/api.json

## How To build Docker Image
```bash
docker build -f manifest/docker/Dockfile . 
```

## Develop Tools
#### Generate API Controller Shell：
```bash
gf gen ctrl
```
#### Generate Dao Code Shell：
```bash
gf gen dao
```

### GoFrame Quick Start:
- Github : https://github.com/gogf/gf
- Doc : https://goframe.org/pages/viewpage.action?pageId=1114399
- API Generate Shell: gf gen ctrl (should have gf install)
  - Generate Controller After API Definition
  - Doc https://goframe.org/pages/viewpage.action?pageId=93880327
- Dao Generate Shell: gf gen dao (should have gf install)
  - Generate Dao|Entity|DO After Database Table Change 
  - Doc https://goframe.org/pages/viewpage.action?pageId=3673173 (need delete config.yaml under root dic)
  - 

Basic Directory Structure
The basic directory structure of the GoFrame business project is as follows (using the Single Repo as an example) :


├── api\
├── hack\
├── internal\
│   ├── cmd\
│   ├── consts\
│   ├── controller\
│   ├── dao\
│   ├── logic\
│   ├── model\
│   |   ├── do\
│   │   └── entity\
│   └── service\
├── manifest\
├── resource\
├── utility\
├── go.mod\
└── main.go\
🔥 Important note 🔥 : The engineering catalog of the framework adopts a universal design to meet the needs of different complex business projects, but the actual project can be appropriate to increase or decrease the default catalog according to the needs of the project. For example, if no i18n/template/protobuf requirement is required, delete the corresponding directory. For example, for very simple business projects (such as verification/demonstration projects), do not consider the use of rigorous dao/logic/model directory and features, then directly delete the corresponding directory, you can directly implement the business logic in the controller. Everything can be flexibly assembled by the developer!

API External interface Defines the input/output data structure of the service provided externally. Considering the need for version management, api/xxx/v1... Exists.
hack tool scripts Store project development tools and scripts. For example, the configuration of CLI tools, various shell/bat scripts and other files.
internal Internal logic Directory for storing service logic. Hide visibility from the outside through the Golang internal feature.
- cmd CLI manages the directory. Multiple command lines can be managed and maintained.
- consts
  Constant definition

All constant definitions for the project.

- controller Indicates the interface layer that receives and parses user input parameters.
- dao

Data Access Data access objects, a layer of abstract objects used to interact with the underlying database, contain only the most basic CURD methods
- logic Service encapsulation Service logic encapsulation management, specific service logic implementation and encapsulation. Is often the most complex part of a project.
- redismq Message System Core Implementation Base On Redis Stream
- consumer Message Topic Consumer For Redis Stream 
- cronjob CronJobs
- model structure Model data structure management module, manage data entity objects, and input and output data structure definitions.
  -do domain objects are used to transform the service model and instance model in dao data operations. They are maintained by the tool and cannot be modified by users.
- entity Data model A data model is a one-to-one relationship between a model and a data set. It is maintained by the tool and cannot be modified by users.
- [Deprecate]service Service interface Indicates the interface definition layer for decoupling service modules. The specific interface implementation is injected in logic.
  The manifest delivery list contains the files that compile, deploy, run, and configure the program. Common contents are as follows:
- config Indicates the directory for storing configuration files.
- docker image file Docker image related dependency files, script files, and so on.
- deploy deployment file Specifies the deployment file. Yaml templates for clustered deployment of Kubernetes are provided by default, managed through kustomize.
- protobuf protocol file Specifies the protocol definition file for the GRPC protocol. After the protocol file is compiled, a go file is generated and stored in the api directory.
  resource Static resource Static resource file. These files can often be injected into a release file in the form of resource packaging/image compilation.
  Go.mod Dependency Management A dependency description file managed using the Go Module package.
  main.go entry file Program entry file.