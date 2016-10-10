FROM golang:latest

RUN apt-get update && apt-get install -y --no-install-recommends \
        texlive-full \
        librsvg2-bin \
    && rm -rf /var/lib/apt/lists/*

# docker build -t go-mathtex .
# docker run -it --rm -v "$PWD":/go/src/github.com/dotzero/go-mathtex --name go-mathtex go-mathtex bash
