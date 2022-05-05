[![Go Reference](https://pkg.go.dev/badge/github.com/joakimofv/restartpb.svg)](https://pkg.go.dev/github.com/joakimofv/restartpb)

restartpb
=========

Protobuf and gRPC interface for restart/terminate/crash.
Useful for checking (in tests) that microservices (in Kubernetes) recover correctly after restart, whether clean or violent.

# Server Side

## Register

```go
grpcServer := grpc.NewServer(...)
restartpb.RegisterRestartServer(grpcServer, myService)
```

## Implement

```go
type MyService struct {
	restartpb.UnimplementedRestartServer
	// ...
}

func (myService *MyService) Restart(ctx context.Context, req *restartpb.RestartRequest) (*restartpb.RestartResponse, error) {
	// TODO: Make myService restart without exiting the program
	return &restartpb.RestartResponse{}, nil
}

func (myService *MyService) Terminate(ctx context.Context, req *restartpb.TerminateRequest) (*restartpb.TerminateResponse, error) {
	// TODO: Make myService stop and the program exit in a clean way
	return &restartpb.TerminateResponse{}, nil
}

func (myService *MyService) Crash(ctx context.Context, req *restartpb.CrashRequest) (*restartpb.CrashResponse, error) {
	// TODO: Crash the program, e.g. with panic()
	return &restartpb.CrashResponse{}, nil
}
```

# Client Side

## New Client

```go
conn, err := grpc.Dial(...)
restartClient := restartpb.NewRestartClient(conn)
```

## Call

```go
_, err = restartClient.Restart(ctx, &restartpb.RestartRequest{})
// Expecting nil error

_, err = restartClient.Terminate(ctx, &restartpb.TerminateRequest{})
// Expecting nil error

_, err = restartClient.Crash(ctx, &restartpb.CrashRequest{})
// Expecting non-nil error
```
