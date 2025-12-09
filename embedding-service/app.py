from flask import Flask, request, jsonify
from sentence_transformers import SentenceTransformer
import numpy as np

# Load model on startup
model = SentenceTransformer("all-MiniLM-L6-v2")

app = Flask(__name__)

@app.route("/", methods=["GET"])
def root():
    return jsonify({
        "status": "running",
        "model": "all-MiniLM-L6-v2"
    })

@app.route("/embed", methods=["POST"])
def embed():
    data = request.get_json()

    if not data or "text" not in data:
        return jsonify({"error": "Missing 'text' field"}), 400

    text = data["text"]
    normalize = data.get("normalize", True)

    # Handle string vs list
    inputs = [text] if isinstance(text, str) else text

    embeddings = model.encode(inputs, normalize_embeddings=normalize)
    embeddings = embeddings.tolist()

    return jsonify({
        "model": "all-MiniLM-L6-v2",
        "count": len(embeddings),
        "embeddings": embeddings
    })


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8000)
