package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"runtime/trace"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("failed to create trace output file: %v", err)
	}
	if err := trace.Start(f); err != nil {
		log.Fatalf("failed to start trace: %v", err)
	}
	defer trace.Stop()

	ctx := context.Background()

	decoder := json.NewDecoder(os.Stdin)

	trace.WithRegion(ctx, "MANTEP", func() {

		// sc := bufio.NewScanner(os.Stdin)
		// gg := trace.StartRegion(ctx, "LOOP")
		for {
			var mj map[string]interface{}

			e := trace.StartRegion(ctx, "DECODE")
			if err := decoder.Decode(&mj); err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}
			e.End()

			e = trace.StartRegion(ctx, "PRINT")
			// for sc.Scan() {
			// v := sc.Bytes()

			// err := json.Unmarshal(v, &mj)
			// if err != nil {
			// 	log.Println("Unmarshal err", err)
			// 	continue
			// }

			fmt.Println(mj)
			e.End()

		}
		// gg.End()
	})
}
