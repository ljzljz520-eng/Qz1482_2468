package main

import (
	"aerialfarm/storage"
	"aerialfarm/workflow"
	"context"
	"fmt"
	"os"
)

func main() {
	path := "aerialfarm.db"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	s, e := storage.Open(path)
	if e != nil {
		fmt.Println(e)
		return
	}
	defer s.Close()
	svc := workflow.NewService(s)
	r, e := svc.Register(context.Background(), "demo-33", "field-33", "pilot")
	if e == nil {
		_, _ = svc.Begin(context.Background(), r.ID)
		r, _ = svc.Process(context.Background(), r.ID)
		fmt.Println(r.Status)
	}
}
