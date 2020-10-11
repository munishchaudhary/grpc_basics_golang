# run this command in project root directory to generate code for defined protobuf
protoc  firstpb/first.proto --go_out=plugins=grpc:.
protoc personpb/person.proto --go_out=plugins=grpc:.