FROM alpine:latest

WORKDIR /app

COPY clientApp /app/clientApp
 
COPY cmd /app/cmd

CMD ["/app/clientApp"]