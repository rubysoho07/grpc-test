# gRPC Test

gRPC를 어떻게 사용할 수 있는지 테스트 해 봅니다.

## 폴더 구성

* `golang`: Golang으로 만든 클라이언트, 서버
* `python`: Python으로 만든 클라이언트, 서버
* `protocol_buffers`: Protocol Buffers 설정

## Python으로 서버/클라이언트 만들기

먼저 필요한 패키지를 설치합니다.

```
cd python/
pip install -r requirements.txt 
```

그리고 다음과 같이 프로토콜 버퍼를 컴파일 합니다.

```
python -m grpc_tools.protoc -I../protocol_buffers --python_out=. --grpc_python_out=. ../protocol_buffers/message_definition.proto
```

생성된 Python 코드를 가지고 서버와 클라이언트를 만들어야 합니다. `grpc_client.py` 파일과 `grpc_server.py` 파일을 참고하세요.

다음 명령으로 서버를 실행합니다. 

```
python grpc_server.py
```

그리고 다른 터미널 창에서 클라이언트를 실행하면, 다음과 같은 메시지가 나옵니다. 

```
python grpc_client.py
Response is 'Hello! Your request is Test Message'
```

서버를 끝내고 싶으면 Ctrl+C를 눌러줍니다. 

## Go로 서버/클라이언트 만들기

참고: [Quick Start](https://grpc.io/docs/languages/go/quickstart/)

프로토콜 버퍼 컴파일러를 설치해야 합니다. [이 문서](https://grpc.io/docs/protoc-installation/)를 참고하세요.

코드 생성에 사용되는 플러그인을 컴파일합니다. 

```
export GO111MODULE=on  # Enable module mode
go get github.com/golang/protobuf/protoc-gen-go
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
export PATH=$PATH:$(go env GOPATH)/bin
```

그리고 프로토콜 버퍼를 컴파일 합니다. 

```
mkdir -p $(go env GOPATH)/src/yg_test_service/
protoc -I../protocol_buffers --go_out=$(go env GOPATH)/src/yg_test_service/ --go-grpc_out=$(go env GOPATH)/src/yg_test_service/ \
    --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative \
    ../protocol_buffers/message_definition.proto
```

(GOPATH로 지정된 디렉터리 내 src/yg_test_service 디렉토리에 go 파일이 생성됩니다. 내용은 yg_test_service 폴더에 있는 내용과 같습니다.)

서버와 클라이언트 소스는 `server/main.go` 파일과 `client/main.go` 파일을 참고하세요.

서버를 실행하려면 다음과 같이 실행합니다. 

```
go run server/main.go
```

다른 터미널 창에서 클라이언트를 실행하면 다음과 같이 실행됩니다. 

```
go run client/main.go
2020/09/30 17:21:27 Greeting: Hello world
```

서버를 끝내고 싶으면 Ctrl+C를 눌러줍니다. 

### 겪었던 문제들

* 프로토콜 버퍼 import가 안 되는 경우: `GOPATH` 환경변수를 확인합니다. 
* gRPC import가 안 되는 경우: `go get google.golang.org/grpc`을 실행합니다. (먼저 `GOPATH` 환경변수를 확인해 보면 좋습니다)

## 참고자료

* [Introduction to gRPC](https://grpc.io/docs/what-is-grpc/introduction/)
* [Protocol Buffers Language Guide](https://developers.google.com/protocol-buffers/docs/overview#simple)