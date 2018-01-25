# Serverless for Golang with Legacy Code

This repository contains an example of a modern GoLang Serverless project which can also support legacy code in Node.js, Java, Scala, Python or C#

## Depencies
1. Node.js 6.3
2. Serverless.js with a configured cloud provider
3. GoLang 1.9.x

## Quick Start
1. Compile the Go code (assuming you have setup your GOPATH and GOROOT)

```
cd serverless-go
GOOS=linux go build -o bin/main
```

2. Test Go Code
```
go test
```

3. Deploy

```
serverless deploy
```
