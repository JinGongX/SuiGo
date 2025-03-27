package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	//apiKey   = "sk-ybnvXQ0WlwazLEmPvIpY5pAEPcKNclBq92LJUogVKEpPswUA" // 替换为您的 Moonshot API Key
	olbaseURL  = "http://127.0.0.1:11434/api"
	olmodelID  = "qwen2" // 选择合适的模型
	olendpoint = "/chat"
)

// ChatCompletionRequest 定义了请求体的结构
type olChatCompletionRequest struct {
	Model    string `json:"model"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
	//Temperature float32 `json:"temperature"`
	Stream bool `json:"stream"`
}

// ChatCompletionResponse 定义了响应体的结构
type olChatCompletionResponse struct {
	Message struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"message"`
}

// sendRequest 发送请求到 Moonshot AI API 并处理响应
func olsendRequest(client *http.Client, requestBody []byte) (*olChatCompletionResponse, error) {
	req, err := http.NewRequest("POST", olbaseURL+olendpoint, ioutil.NopCloser(bytes.NewBuffer(requestBody)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	//req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusTooManyRequests {
		retryAfter := resp.Header.Get("Retry-After")
		if retryAfter != "" {
			duration, _ := time.ParseDuration(retryAfter)
			time.Sleep(duration)
		} else {
			time.Sleep(5 * time.Second) // 默认等待5秒
		}
		return olsendRequest(client, requestBody) // 递归重试
	}
	var response olChatCompletionResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	return &response, err
}

var olhistory = []struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}{
	{"system", "你是一位百科全书,科学家。你叫小随同学"},
}

func ChatlocalAI(c *gin.Context) {
	var request olChatCompletionRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := &http.Client{
		Timeout: 50 * time.Second, // 设置请求超时时间
	}

	// 将用户的消息添加到历史记录中
	olhistory = append(olhistory, struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	}{
		Role:    "user",
		Content: request.Messages[0].Content, // 假设用户的消息是第一个
	})

	// 构建完整的请求体，包含历史消息
	request.Messages = append([]struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	}{
		{
			Role:    "system",
			Content: "你是一位百科全书,科学家。",
		},
	}, olhistory...)

	// 序列化 request 结构体到 JSON
	requestBodyJSON, err := json.Marshal(request)
	if err != nil {
		// 处理序列化错误
		c.JSON(http.StatusInternalServerError, gin.H{"error": "请求体序列化失败"})
		return
	}
	response, err := olsendRequest(client, requestBodyJSON)
	if err != nil {
		var apiError struct {
			Message string `json:"message"`
		}
		if jsonErr := json.Unmarshal([]byte(err.Error()), &apiError); jsonErr == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": apiError.Message})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "请求处理失败"})
		}
		return
	}
	// 将用户的消息添加到历史记录中
	olhistory = append(olhistory, struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	}{
		Role:    response.Message.Role,
		Content: response.Message.Content, // 假设用户的消息是第一个
	})
	fmt.Println(olhistory)

	// 为了保持历史记录的简洁，我们只保留最近的一定数量的历史消息
	// 例如，这里我们只保留最近的两个系统消息和最后一个用户消息
	////if len(history) > 3 {
	////	history = history[len(history)-3:] // 根据需要调整
	////}
	c.JSON(http.StatusOK, response)
}
