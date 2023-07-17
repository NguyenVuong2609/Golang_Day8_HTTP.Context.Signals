package DemoContext

import (
	"context"
	"fmt"
	"time"
)

func DemoDeadlinesContext() {
	// Tạo một context với thời hạn 2 giây
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	defer cancel()

	// Khởi chạy một goroutine và truyền context
	go performTask(ctx)

	// Đợi một chút để cho goroutine chạy
	time.Sleep(3 * time.Second)
}

func performTask(ctx context.Context) {
	// Kiểm tra xem deadline của context
	deadline, ok := ctx.Deadline()
	if ok {
		// In ra thời hạn của context
		fmt.Println("Thời hạn:", deadline)
	}

	// Thực hiện công việc...
	time.Sleep(1 * time.Second)

	// Kiểm tra xem context có hết hạn không
	select {
	case <-ctx.Done():
		// Đã hết hạn, hủy bỏ công việc
		fmt.Println("Hết hạn! Hủy bỏ công việc.")
		return
	default:
		// Công việc được hoàn thành trong thời hạn
		fmt.Println("Công việc đã hoàn thành.")
	}
}
