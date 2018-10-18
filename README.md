# Serverless for Golang with Legacy Code

### Purpose:
  This repository contains an example of a modern GoLang Serverless project which can also support legacy code in Node.js, Java, Scala, Python or C#

### Dependencies:
  This project depends on the Serverless.js library, Node.js, Go, dep, AWS and supporting tools.

  - Go: It is recommended you install Go `1.10.3` via `goenv` Instructions for installing can be found [here](https://github.com/syndbg/goenv) or by running the following commands:
    ```
    git clone https://github.com/syndbg/goenv.git ~/.goenv

    echo `export GOENV_ROOT="$HOME/.goenv"
          export PATH="$GOENV_ROOT/bin:$PATH"
          export PATH="$GOENV_ROOT/shims:$PATH"
          eval "$(goenv init -)"
          export GOPATH="$HOME/workspace/go"` > ~/.bash_profile
    source ~/.bash_profile

    goenv install 1.10.3
    goenv global  1.10.3
    ```
    
  - dep: It is recommended you install the most recent version of dep manually via the built-in Go toolchain. Instructions for installing can be found [here](https://golang.github.io/dep/docs/installation.html) or by running the following commands:
    ```
    go get -d -u github.com/golang/dep
    cd $(go env GOPATH)/src/github.com/golang/dep
    DEP_LATEST=$(git describe --abbrev=0 --tags)
    git checkout $DEP_LATEST
    go install -ldflags="-X main.version=$DEP_LATEST" ./cmd/dep
    git checkout master
    ```
    
  - Node.js: It is recommended you install Node.js `10.5.0` via `nodenv` Instructions for installing can be found [here](https://github.com/nodenv/nodenv) or by running the following commands:
    ```
    git clone https://github.com/nodenv/nodenv.git ~/.nodenv

    echo `export NODENV_ROOT="$HOME/.nodenv"
          export PATH="$NODENV_ROOT/bin:$PATH"
          export PATH="$NODENV_ROOT/shims:$PATH"
          eval "$(nodenv init -)"` > ~/.bash_profile
    source ~/.bash_profile

    nodenv install 10.5.0
    nodenv global  10.5.0
    ```

  - Serverless.js: It is recommended you install the most recent version of Serverless.js manually via the built-in Node.js `npm` toolchain. Instructions for installing can be found [here](https://serverless.com/framework/docs/getting-started/) or by running the following commands:
    ```
    npm install -g serverless
    ```

  - AWS Credentials: It is recommended you have your AWS credentials configured in the default manner documented [here](https://docs.aws.amazon.com/cli/latest/userguide/cli-config-files.html) however the AWS CLI is not required.

### Configuration:
  Minimal configuration is required for this project, just run the following command in the project directory to get started:

  ```
  serverless config
  ```

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
