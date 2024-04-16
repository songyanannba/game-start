#rm *.go
#protoc *.proto --go_out=plugins=grpc:.
protoc --go_out=plugins=grpc:../ ./common/*.proto ./game/*.proto
