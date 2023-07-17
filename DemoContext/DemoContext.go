package DemoContext

import (
	"context"
	"fmt"
	"time"
)

func DemoContext() {
	// Tạo một context với hạn chế thời gian 2 giây
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go process(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("Quá thời gian!")
	}
}

func process(ctx context.Context) {
	// Mô phỏng công việc mất thời gian
	time.Sleep(3 * time.Second)

	// Kiểm tra xem có yêu cầu hủy bỏ không
	select {
	case <-ctx.Done():
		fmt.Println("Công việc đã bị hủy bỏ.")
	default:
		fmt.Println("Công việc đã hoàn thành.")
	}
}
