package server

import (
	"reflect"
	"testing"

	"github.com/prateekgogia/echoserver/api"
	"golang.org/x/net/context"
)

func TestServer_EchoRequest(t *testing.T) {
	type args struct {
		in0 context.Context
		in  *api.EchoMessage
	}
	tests := []struct {
		name string
		// s       *Server
		args    args
		want    *api.EchoMessage
		wantErr bool
	}{
		{
			name:    "Test public API in package",
			args:    args{context.TODO(), &api.EchoMessage{Message: "Hello"}},
			want:    &api.EchoMessage{Message: "Hello"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{}
			got, err := s.EchoRequest(tt.args.in0, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.EchoRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.EchoRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_echoRequest(t *testing.T) {
	type args struct {
		in *api.EchoMessage
	}
	tests := []struct {
		name    string
		args    args
		want    *api.EchoMessage
		wantErr bool
	}{
		{
			name:    "Send and expect hello message",
			args:    args{in: &api.EchoMessage{Message: "Hello"}},
			want:    &api.EchoMessage{Message: "Hello"},
			wantErr: false,
		},
		{
			name:    "Empty message from client",
			args:    args{in: &api.EchoMessage{Message: ""}},
			want:    &api.EchoMessage{Message: ""},
			wantErr: false,
		},
		{
			name:    "Alphnumeric characters from client",
			args:    args{in: &api.EchoMessage{Message: "abcAA332#*&332"}},
			want:    &api.EchoMessage{Message: "abcAA332#*&332"},
			wantErr: false,
		},
		{
			name:    "non-ascii characters",
			args:    args{in: &api.EchoMessage{Message: "日本語"}},
			want:    &api.EchoMessage{Message: "日本語"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := echoRequest(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("echoRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("echoRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
