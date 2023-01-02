package qcdn

import (
	"context"
	"strings"

	cdn "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdn/v20180606"
)

var (
	preset = []string{
		"https://tangx.in/index.xml",
		"https://tangx.in/sitemap.xml",
	}
)

func Do(flag *Flag) {

	client := mustClient()
	urls := mustURLs(flag)

	urls = append(urls, preset...)

	if !flag.Purge && !flag.Push {
		flag.Purge = true
		flag.Push = true
	}

	if flag.Purge {
		urls := exclude(urls, "posts")
		// 刷新 URL
		_ = PurgeSite(client, urls)

	}

	if flag.Push {
		// 预热 URL
		PushSite(client, urls)
	}
}

func isPurgeOK(ctx context.Context, client *cdn.Client, purgeId string) bool {
	select {
	case <-ctx.Done():
		// timeout, treat as OK
		return true
	default:
		ok := IsPurgeTaskDone(client, purgeId)
		if ok {
			return true
		}
		return false
	}
}

func exclude(urls []string, str string) []string {
	result := []string{}
	for _, u := range urls {
		if strings.Contains(u, "posts") {
			continue
		}
		result = append(result, u)
	}

	return result
}
