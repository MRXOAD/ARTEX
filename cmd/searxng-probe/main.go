// searxng-probe 是部署自检小工具：用打过补丁的 norma 搜索模块，对
// ARTEX_SEARX_URL 指向的 SearXNG 实例跑一次真实查询，验证 searxng 后端可用。
// 用法：ARTEX_SEARX_URL=http://192.168.31.1:18088 go run ./cmd/searxng-probe
package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Autumn-27/norma/tool"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	res, err := tool.WebSearchProbe(ctx, tool.WebSearchConfig{Backend: "searxng"}, "nginx 稳定版", 5)
	if err != nil {
		fmt.Println("PROBE FAIL:", err)
		os.Exit(1)
	}
	fmt.Println("PROBE OK, results:", len(res))
	for _, r := range res {
		fmt.Printf("- %s | %s\n", r.Title, r.URL)
	}
}
