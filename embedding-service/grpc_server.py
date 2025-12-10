import grpc
from concurrent import futures
from fastembed import TextEmbedding

import embedding_pb2, embedding_pb2_grpc


model = TextEmbedding("sentence-transformers/all-MiniLM-L6-v2")


class EmbeddingService(embedding_pb2_grpc.EmbeddingServiceServicer):
    """ gRPC service for generating text embeddings."""

    def TextToEmbedding(self, request, context):
        text = request.text
        print("Valinor_is_calling:", text)


        vectors_gen = model.embed([text])
        vectors_list = list(vectors_gen)
        embeddings = vectors_list[0]

        # Convert to list for RPC-safe format
        embedding_list = embeddings.tolist() if hasattr(embeddings, "tolist") else embeddings

        return embedding_pb2.EmbeddingsMessageResponse(
            text=str(embedding_list)
        )


def start_grpc_server():
    """ Starts the gRPC server on port 50001. """
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    embedding_pb2_grpc.add_EmbeddingServiceServicer_to_server(
        EmbeddingService(), server
    )

    server.add_insecure_port("[::]:50001")
    server.start()
    print("GRPC server running on port 50001")

    server.wait_for_termination()


if __name__ == "__main__":
    start_grpc_server()
