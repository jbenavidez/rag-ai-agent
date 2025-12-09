FROM python:3.10-slim

ENV PYTHONDONTWRITEBYTECODE=1
ENV PYTHONUNBUFFERED=1

RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
    git \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Do NOT upgrade pip. It breaks under buildx.
RUN pip install --no-cache-dir \
        torch==2.9.1 \
        sentence-transformers \
        flask

COPY app.py .

EXPOSE 8000

CMD ["python", "app.py"]
