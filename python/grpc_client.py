import logging

import grpc

import message_definition_pb2
import message_definition_pb2_grpc


def run():
    with grpc.insecure_channel('localhost:50051') as channel:
        stub = message_definition_pb2_grpc.YGTestServiceStub(channel)
        response = stub.SayHello(message_definition_pb2.YGRequest(req_message='Test Message'))
    print(f"Response is '{response.res_message}'")


if __name__ == "__main__":
    logging.basicConfig()
    run()