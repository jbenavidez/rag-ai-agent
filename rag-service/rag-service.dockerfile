FROM alpine:latest

WORKDIR /app

COPY RagAgent /app/RagAgent
 
COPY cmd /app/cmd

CMD ["/app/RagAgent"]