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

## 참고자료

* [Introduction to gRPC](https://grpc.io/docs/what-is-grpc/introduction/)
* [Protocol Buffers Language Guide](https://developers.google.com/protocol-buffers/docs/overview#simple)