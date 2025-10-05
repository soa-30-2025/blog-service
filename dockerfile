# 1. Koristi zvanični Go image
FROM golang:1.25

# 2. Postavi radni direktorijum
WORKDIR /app

# 3. Kopiraj sve fajlove u kontejner
COPY . .

# 4. Preuzmi dependency-e
RUN go mod download

# 5. Izgradi aplikaciju
RUN go build -o blog-service .

# 6. Pokreni aplikaciju
CMD ["./blog-service"]
