package service

import (
	"context"
	"fmt"

	"connectrpc.com/connect"

	sampleStreamv1 "grpc_client_streaming_sample/gen/sample_stream/v1" // generated by protoc-gen-go
	// generated by protoc-gen-connect-go
)

type SampleStreamServer struct{}

func (s *SampleStreamServer) SampleClientStream(ctx context.Context, stream *connect.ClientStream[sampleStreamv1.SampleClientStreamRequest]) (*connect.Response[sampleStreamv1.SampleClientStreamResponse], error) {
	var fileInfo sampleStreamv1.SampleClientStreamRequest_FileInfo
	var data []byte
	fmt.Println("--------------------------stream")
	for stream.Receive() {
		if stream.Msg().GetFileInfo() != nil {
			fileInfo = *stream.Msg().GetFileInfo()
		}
		if stream.Msg().GetData() != nil {
			data = append(data, stream.Msg().GetData()...)
		}
	}

	fmt.Printf("fileinfo:%+v", fileInfo)
	fmt.Printf("data:%s\n", data)

	return connect.NewResponse(&sampleStreamv1.SampleClientStreamResponse{}), nil
}
