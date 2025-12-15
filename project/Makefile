RAG_SERVICE=RagAgent
FRONT_END=frontEndApp

up_build:  build_rag_agent build_frontend
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Starting Docker images..."
	docker-compose up 
	@echo "Docker images started!"


build_rag_agent:
	@echo "Building RAG IA Agent  binary..."
	cd ../rag-ai-agent && env GOOS=linux CGO_ENABLED=0 go build -o ${RAG_SERVICE} ./cmd/api
	@echo "Done!"


build_frontend:
	@echo "Building front end  binary..."
	cd ../front-end && env GOOS=linux CGO_ENABLED=0 go build -o ${FRONT_END} ./cmd/web
	@echo "Done!"
