from concurrent import futures
import logging

import grpc

import message_definition_pb2
import message_definition_pb2_grpc


class YGServicer(message_definition_pb2_grpc.YGTestServiceServicer):

    def SayHello(self, request, context):
        return message_definition_pb2.YGResponse(
            res_message=f'Hello! Your request is {request.req_message}'
        )


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    message_definition_pb2_grpc.add_YGTestServiceServicer_to_server(YGServicer(), server)
    server.add_insecure_port('localhost:50051')
    server.start()
    server.wait_for_termination()


if __name__ == "__main__":
    logging.basicConfig()
    serve()