package DemoContext

import (
	"context"
	"fmt"
	"time"
)

func TestContextBackground() {
	// Tạo một context cha
	ctx := context.Background()

	// Tạo một context con từ context cha
	ctx, cancel := context.WithCancel(ctx)

	// Khởi chạy các goroutine
	go worker(ctx, "Worker 1")
	go worker(ctx, "Worker 2")

	// Hủy bỏ context sau 3 giây
	time.Sleep(3 * time.Second)
	cancel()

	// Đợi tất cả các goroutine hoàn thành
	time.Sleep(1 * time.Second)
	fmt.Println("Tất cả các goroutine đã được hủy bỏ.")
}

func worker(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "đã bị hủy bỏ.")
			return
		default:
			fmt.Println(name, "đang chạy.")
			time.Sleep(1 * time.Second)
		}
	}
}
