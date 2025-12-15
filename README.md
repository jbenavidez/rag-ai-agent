# RAG-AI-Agent(WIP)
A lightweight RAG-based AI agent that ingests data from the NYC Capital Projects dataset and answers user questions with relevant, context-aware responses grounded in the ingested data. The AI agent is designed to operate independently and can be integrated into other applications or services. A command-line interface (CLI) is provided as a testing and interaction tool, allowing users to easily query and validate the agent’s capabilities.

 ## Stack
- **Go** — Core language for building the **RAG AI agent**
- **Weaviate** — Vector database used by the **RAG AI agent** for storing and retrieving embeddings
- **gRPC** — High-performance communication between microservices
- **Ollama** — Local LLM runtime for serving LLaMA 2
- **LLaMA 2** — Large Language Model for context-aware response generation
- **OpenAI Embeddings** — Used for text vectorization via `text2vec-openai`
- **Docker & Docker Compose** — Containerization and service orchestration

 ## Features
- **Retrieval-Augmented Generation (RAG)** for accurate, data-grounded responses
- **Semantic search** powered by vector embeddings
- **Context-aware responses** generated using **LLaMA 2**
- **Microservices-based architecture** for modularity and scalability
- **CLI tool** for testing and interacting with the AI agent

 ## Microservices Detais
The system is built using a **microservices architecture**:

- **RAG-service (Go)** — Handles queries, coordinates embedding retrieval, and orchestrates the AI agent  
- **Weaviate Database** — Stores and retrieves vector embeddings for semantic search  
- **Ollama (LLaMA 2)** — Generates natural language responses based on context  
- **CLI-service** — Command-line interface for testing and interacting with the AI agent

 ## Data source
The agent’s knowledge comes from   **data.cityofnewyork.us**, with data currently imported via a CSV file. This ingestion method is fully flexible—CSV loading can easily be replaced with direct API-based retrieval, since  **data.cityofnewyork.us** provides both CSV exports and a public API for accessing the capital projects information.



 ## Setup and Running the Project
### 1. Start services

Start all services using Docker Compose:

```bash
make up_build
```

### 2. Access the CLI(wip)

SSH into the CLI container::

```bash
docker exec -it project-cli-service-1 /bin/sh

```
You will see a shell prompt inside the container:
```bash
/ #
```

Start the CLI service inside the container:
```bash
/app/CliService

```
You will see a welcome message:

```bash
Welcome to the NYC Capital Project RAG AI CLI
Ask any question related to NYC capital projects from 2023-2025.
The AI agent will provide answers based on the data provided by data.cityofnewyork.us
Type your question and press Enter. Type 'exit' to quit.
```

Type your question and press Enter. Example:
```bash
Which boroughs have interceptor sewer projects?
```

 Display response( still on progress..)