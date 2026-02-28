package common

import "time"

var StartTime = time.Now().Unix() // unit: second
var Version = "v0.0.11"           // this hard coding will be replaced automatically when building, no need to manually change

type ModelInfo struct {
	Model string
	ID    string
	Type  string
}

// 创建映射表（假设用 model 名称作为 key）
var ModelRegistry = map[string]ModelInfo{
	// Claude 4.6 系列 (最新)
	"claude-opus-4-6":          {"claude-opus-4-6", "019c2fac-13de-7550-a751-f5f593c77c72", "chat"},
	"claude-opus-4-6-thinking": {"claude-opus-4-6-thinking", "019c2f86-74db-7cc3-baa5-6891bebb5999", "chat"},
	"claude-sonnet-4-6":        {"claude-sonnet-4-6", "019c6d29-a30c-7e20-9bd0-6650af926623", "chat"},

	// Claude 4.5 系列
	"claude-opus-4-5-20251101":                 {"claude-opus-4-5-20251101", "019adbec-8396-71cc-87d5-b47f8431a6a6", "chat"},
	"claude-opus-4-5-20251101-thinking-32k":    {"claude-opus-4-5-20251101-thinking-32k", "019ab8b2-9bcf-79b5-9fb5-149a7c67b7c0", "chat"},
	"claude-sonnet-4-5-20250929":               {"claude-sonnet-4-5-20250929", "019a2d13-28a5-7205-908c-0a58de904617", "chat"},
	"claude-sonnet-4-5-20250929-thinking-32k":  {"claude-sonnet-4-5-20250929-thinking-32k", "b0ea1407-2f92-4515-b9cc-b22a6d6c14f2", "chat"},
	"claude-haiku-4-5-20251001":                {"claude-haiku-4-5-20251001", "0199e8e9-01ed-73e0-96ba-cf43b286bf10", "chat"},

	// Claude 4.1 系列
	"claude-opus-4-1-20250805":                {"claude-opus-4-1-20250805", "96ae95fd-b70d-49c3-91cc-b58c7da1090b", "chat"},
	"claude-opus-4-1-20250805-thinking-16k":   {"claude-opus-4-1-20250805-thinking-16k", "f1a2eb6f-fc30-4806-9e00-1efd0d73cbc4", "chat"},

	// Claude 4 系列
	"claude-opus-4-20250514":                {"claude-opus-4-20250514", "ee116d12-64d6-48a8-88e5-b2d06325cdd2", "chat"},
	"claude-opus-4-20250514-thinking-16k":   {"claude-opus-4-20250514-thinking-16k", "3b5e9593-3dc0-4492-a3da-19784c4bde75", "chat"},
	"claude-sonnet-4-20250514":              {"claude-sonnet-4-20250514", "ac44dd10-0666-451c-b824-386ccfea7bcc", "chat"},
	"claude-sonnet-4-20250514-thinking-32k": {"claude-sonnet-4-20250514-thinking-32k", "4653dded-a46b-442a-a8fe-9bb9730e2453", "chat"},

	// Claude 3.x 系列
	"claude-3-7-sonnet-20250219":              {"claude-3-7-sonnet-20250219", "c5a11495-081a-4dc6-8d9a-64a4fd6f7bbc", "chat"},
	"claude-3-7-sonnet-20250219-thinking-32k": {"claude-3-7-sonnet-20250219-thinking-32k", "be98fcfd-345c-4ae1-9a82-a19123ebf1d2", "chat"},
	"claude-3-5-sonnet-20241022":              {"claude-3-5-sonnet-20241022", "f44e280a-7914-43ca-a25d-ecfcc5d48d09", "chat"},
	"claude-3-5-haiku-20241022":               {"claude-3-5-haiku-20241022", "f6fbf06c-532c-4c8a-89c7-f3ddcfb34bd1", "chat"},

	// Search 模型
	"claude-opus-4-6-search":   {"claude-opus-4-6-search", "019c6f55-308b-71ac-95af-f023a48253cf", "chat"},
	"claude-sonnet-4-6-search": {"claude-sonnet-4-6-search", "019c6f55-70d6-7a9c-b89b-9a0db36a3582", "chat"},
	"claude-opus-4-5-search":   {"claude-opus-4-5-search", "019bda1f-2da5-7e9b-b357-2ed018d39393", "chat"},
	"claude-sonnet-4-5-search": {"claude-sonnet-4-5-search", "019bda1f-3664-75d3-91b9-ede8f7561e44", "chat"},
	"claude-opus-4-1-search":   {"claude-opus-4-1-search", "d942b564-191c-41c5-ae22-400a930a2cfe", "chat"},
	"claude-opus-4-search":     {"claude-opus-4-search", "25bcb878-749e-49f4-ac05-de84d964bcee", "chat"},

	// GPT 系列
	"chatgpt-4o-latest-20250326": {"chatgpt-4o-latest-20250326", "9513524d-882e-4350-b31e-e4584440c2c8", "chat"},
	"gpt-4.1-2025-04-14":         {"gpt-4.1-2025-04-14", "14e9311c-94d2-40c2-8c54-273947e208b0", "chat"},
	"gpt-4.1-mini-2025-04-14":    {"gpt-4.1-mini-2025-04-14", "6a5437a7-c786-467b-b701-17b0bc8c8231", "chat"},

	// O 系列
	"o3-2025-04-16":      {"o3-2025-04-16", "cb0f1e24-e8e9-4745-aabc-b926ffde7475", "chat"},
	"o3-mini":            {"o3-mini", "c680645e-efac-4a81-b0af-da16902b2541", "chat"},
	"o4-mini-2025-04-16": {"o4-mini-2025-04-16", "f1102bbf-34ca-468f-a9fc-14bcf63f315b", "chat"},

	// Gemini 系列
	"gemini-2.0-flash-001":                    {"gemini-2.0-flash-001", "7a55108b-b997-4cff-a72f-5aa83beee918", "chat"},
	"gemini-2.5-flash-preview-04-17":          {"gemini-2.5-flash-preview-04-17", "7fff29a7-93cc-44ab-b685-482c55ce4fa6", "chat"},
	"gemini-2.5-pro-preview-05-06":            {"gemini-2.5-pro-preview-05-06", "0337ee08-8305-40c0-b820-123ad42b60cf", "chat"},

	// Llama 系列
	"llama-3.3-70b-instruct":             {"llama-3.3-70b-instruct", "dcbd7897-5a37-4a34-93f1-76a24c7bb028", "chat"},
	"llama-4-maverick-03-26-experimental": {"llama-4-maverick-03-26-experimental", "49bd7403-c7fd-4d91-9829-90a91906ad6c", "chat"},
	"llama-4-maverick-17b-128e-instruct":  {"llama-4-maverick-17b-128e-instruct", "b5ad3ab7-fc56-4ecd-8921-bd56b55c1159", "chat"},

	// Qwen 系列
	"qwen-max-2025-01-25": {"qwen-max-2025-01-25", "fe8003fc-2e5d-4a3f-8f07-c1cff7ba0159", "chat"},
	"qwen3-235b-a22b":     {"qwen3-235b-a22b", "2595a594-fa54-4299-97cd-2d7380d21c80", "chat"},
	"qwen3-30b-a3b":       {"qwen3-30b-a3b", "9a066f6a-7205-4325-8d0b-d81cc4b049c0", "chat"},
	"qwq-32b":             {"qwq-32b", "885976d3-d178-48f5-a3f4-6e13e0718872", "chat"},

	// DeepSeek
	"deepseek-v3-0324": {"deepseek-v3-0324", "2f5253e4-75be-473c-bcfc-baeb3df0f8ad", "chat"},

	// Grok
	"grok-3-preview-02-24": {"grok-3-preview-02-24", "551ba709-049c-4883-b4aa-8df24174e676", "chat"},
	"grok-3-mini-beta":     {"grok-3-mini-beta", "7699c8d4-0742-42f9-a117-d10e84688dab", "chat"},

	// 其他
	"gemma-3-27b-it":       {"gemma-3-27b-it", "789e245f-eafe-4c72-b563-d135e93988fc", "chat"},
	"amazon.nova-pro-v1:0": {"amazon.nova-pro-v1:0", "a14546b5-d78d-4cf6-bb61-ab5b8510a9d6", "chat"},
	"command-a-03-2025":    {"command-a-03-2025", "0f785ba1-efcb-472d-961e-69f7b251c7e3", "chat"},
	"mistral-medium-2505":  {"mistral-medium-2505", "27b9f8c6-3ee1-464a-9479-a8b3c2a48fd4", "chat"},

	// 图像模型
	"gemini-2.0-flash-preview-image-generation": {"gemini-2.0-flash-preview-image-generation", "69bbf7d4-9f44-447e-a868-abc4f7a31810", "image"},
	"imagen-3.0-generate-002":                   {"imagen-3.0-generate-002", "51ad1d79-61e2-414c-99e3-faeb64bb6b1b", "image"},
	"ideogram-v2":                               {"ideogram-v2", "34ee5a83-8d85-4d8b-b2c1-3b3413e9ed98", "image"},
	"gpt-image-1":                               {"gpt-image-1", "6e855f13-55d7-4127-8656-9168a9f4dcc0", "image"},
	"photon":                                    {"photon", "17e31227-36d7-4a7a-943a-7ebffa3a00eb", "image"},
	"dall-e-3":                                  {"dall-e-3", "bb97bc68-131c-4ea4-a59e-03a6252de0d2", "image"},
	"recraft-v3":                                {"recraft-v3", "b70ab012-18e7-4d6f-a887-574e05de6c20", "image"},
	"flux-1.1-pro":                              {"flux-1.1-pro", "9e8525b7-fe50-4e50-bf7f-ad1d3d205d3c", "image"},
}

// 通过 model 名称查询的方法
func GetModelInfo(modelName string) (ModelInfo, bool) {
	info, exists := ModelRegistry[modelName]
	return info, exists
}

func GetModelList() []string {
	var modelList []string
	for k := range ModelRegistry {
		modelList = append(modelList, k)
	}
	return modelList
}
