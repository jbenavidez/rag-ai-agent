RAG_SERVICE=RagAgent

up_build:  build_client
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Starting Docker images..."
	docker-compose up 
	@echo "Docker images started!"


build_client:
	@echo "Building client  binary..."
	cd ../rag-service && env GOOS=linux CGO_ENABLED=0 go build -o ${RAG_SERVICE} ./cmd/api
	@echo "Done!"
