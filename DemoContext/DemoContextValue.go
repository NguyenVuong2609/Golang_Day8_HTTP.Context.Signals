package DemoContext

import (
	"context"
	"fmt"
	"time"
)

type key string

func DemoValueContext() {
	// Tạo một context cha
	ctx := context.Background()

	// Đặt giá trị vào context cha
	ctx = context.WithValue(ctx, key("username"), "john_doe")

	// Khởi chạy goroutine và truy cập giá trị từ context
	go printUsername(ctx)

	// Đợi một chút để cho goroutine chạy
	time.Sleep(1 * time.Second)
}

func printUsername(ctx context.Context) {
	// Truy cập giá trị từ context
	if username, ok := ctx.Value(key("username")).(string); ok {
		fmt.Println("Username từ context:", username)
	} else {
		fmt.Println("Không tìm thấy giá trị username trong context.")
	}
}
