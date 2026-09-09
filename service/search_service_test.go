package service

import (
	"testing"

	"pansou/model"
)

func TestMergeResultsByTypeFallsBackToTelegramContent(t *testing.T) {
	results := []model.SearchResult{
		{
			UniqueID: "testchannel_3018717",
			Channel:  "testchannel",
			Title:    "9月9日",
			Content:  "9月9日\n\n【库库光影】【劣探德克尔】【2026】 夸克网盘",
			Links: []model.Link{
				{
					Type:      "quark",
					URL:       "https://pan.quark.cn/s/53400d555905",
					WorkTitle: "9月9日",
				},
			},
		},
	}

	merged := mergeResultsByType(results, "劣探德克尔", nil)
	if len(merged["quark"]) != 1 {
		t.Fatalf("Telegram content match was filtered out: %+v", merged)
	}
}
