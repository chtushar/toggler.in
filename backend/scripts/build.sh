echo "Running go generate"

go generate ./...


OUTPUT='bin/'
go build -o "$OUTPUT" ./app

echo "Build successful. Output -> ${OUTPUT}app"
