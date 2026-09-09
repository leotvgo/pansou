package util

import (
	"fmt"
	"strings"
	"testing"
)

func TestParseSearchResultsSkipsTelegramDateHeader(t *testing.T) {
	html := telegramMessageHTML(
		"testchannel/1",
		`<b>📅 9月9日</b><br/><br/><b>🎬 【库库光影】【<mark class="highlight">亡灵之村</mark>】【2026】 夸克网盘</b><br/><br/>类型：恐怖片<br/>分享：雷欧-Leo<br/>📝 简介：<br/>正文内容<br/><br/>🔍 查看资源：<a href="https://pan.quark.cn/s/cd3562338b7a">https://pan.quark.cn/s/cd3562338b7a</a>`,
	)

	results, _, err := ParseSearchResults(html, "testchannel")
	if err != nil {
		t.Fatal(err)
	}
	if len(results) != 1 {
		t.Fatalf("got %d results, want 1", len(results))
	}

	result := results[0]
	wantTitle := "【库库光影】【亡灵之村】【2026】 夸克网盘"
	if result.Title != wantTitle {
		t.Fatalf("title = %q, want %q", result.Title, wantTitle)
	}
	if len(result.Links) != 1 || result.Links[0].WorkTitle != wantTitle {
		t.Fatalf("unexpected links: %+v", result.Links)
	}
	if !strings.Contains(result.Content, "9月9日\n\n🎬") {
		t.Fatalf("content should preserve Telegram line breaks: %q", result.Content)
	}
}

func TestParseSearchResultsKeepsFirstLineTitle(t *testing.T) {
	html := telegramMessageHTML(
		"testchannel/2",
		`【国漫】<mark class="highlight">师兄啊师兄</mark> 年番 (2024) 更新157集<br/><br/>简介内容<br/><br/>链接：<br/><a href="https://pan.quark.cn/s/fee3d596eb7a">https://pan.quark.cn/s/fee3d596eb7a</a>`,
	)

	results, _, err := ParseSearchResults(html, "testchannel")
	if err != nil {
		t.Fatal(err)
	}
	if len(results) != 1 {
		t.Fatalf("got %d results, want 1", len(results))
	}
	wantTitle := "【国漫】师兄啊师兄 年番 (2024) 更新157集"
	if results[0].Title != wantTitle {
		t.Fatalf("title = %q, want %q", results[0].Title, wantTitle)
	}
}

func telegramMessageHTML(dataPost, messageHTML string) string {
	return fmt.Sprintf(`<div class="tgme_widget_message_wrap"><div class="tgme_widget_message" data-post="%s"><div class="tgme_widget_message_bubble"><div class="tgme_widget_message_text">%s</div><a class="tgme_widget_message_date"><time datetime="2026-09-09T10:40:49+00:00"></time></a></div></div></div>`, dataPost, messageHTML)
}
