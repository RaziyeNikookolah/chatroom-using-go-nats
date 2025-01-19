gen-proto:
		protoc -I=./auth/api/pb --go_out=./auth/api/pb --go_opt=paths=source_relative  --go-grpc_out=./auth/api/pb --go-grpc_opt=paths=source_relative ./auth/api/pb/*.proto