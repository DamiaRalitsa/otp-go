# Gunakan image Ubuntu sebagai dasar
FROM ubuntu:22.04

# Set environment variables
ENV DEBIAN_FRONTEND=noninteractive
ENV TZ=Asia/Jakarta

# Install dependencies yang dibutuhkan
RUN apt-get update -qq \
    && apt-get upgrade -y -qq \
    && apt-get install -y --no-install-recommends -qq \
        wget ca-certificates ffmpeg make curl git gcc g++ \
        libwebp-dev \
        tesseract-ocr \
        tesseract-ocr-eng \
        libtesseract-dev \
        libleptonica-dev \
    && rm -rf /var/lib/apt/lists/*

# Install Go versi terbaru
RUN wget https://go.dev/dl/go1.22.12.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go1.22.12.linux-amd64.tar.gz \
    && rm -f go1.22.12.linux-amd64.tar.gz

# Set environment variables untuk Go
ENV GOROOT=/usr/local/go
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:/usr/local/go/bin:$PATH
ENV CGO_ENABLED=1  

# Set working directory
WORKDIR /opt

# Copy go.mod dan go.sum lalu install dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy seluruh kode aplikasi ke dalam container
COPY . .

# Build aplikasi Go
RUN go build -o /opt/app cmd/app/main.go

# Expose port aplikasi
EXPOSE 8333

# Jalankan aplikasi
CMD ["./app"]
