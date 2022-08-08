
import asyncio
import logging

import grpc
import logs_pb2
import logs_pb2_grpc


class Logs(logs_pb2_grpc.LogsDataServicer):

    async def WriteLog(
            self, request: logs_pb2.LogMsg,
            context: grpc.aio.ServicerContext):
        print(request.msg)
        return logs_pb2.google_dot_protobuf_dot_empty__pb2.Empty()


async def serve() -> None:
    server = grpc.aio.server()
    logs_pb2_grpc.add_LogsDataServicer_to_server(Logs(), server)
    listen_addr = '[::]:50053'
    server.add_insecure_port(listen_addr)
    logging.info("Starting server on %s", listen_addr)
    await server.start()
    await server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig(level=logging.INFO)
    asyncio.run(serve())
