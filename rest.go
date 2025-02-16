package hhooking

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"path"

	jsoniter "github.com/json-iterator/go"
)

const (
	BaseAPIUrl = "https://discord.com/api/v8/"
)

func GetApplicationCommands(appId SnowFlake, authToken string, guildId *SnowFlake) []ApplicationCommand {
	var reqPath string
	switch guildId {
	case nil:
		reqPath = path.Join("applications", fmt.Sprintf("%d", appId), "commands")
	default:
		reqPath = path.Join("applications", fmt.Sprintf("%d", appId), "guilds", fmt.Sprintf("%d", *guildId), "commands")
	}

	var v []ApplicationCommand
	sendRequest(reqPath, "GET", nil, authToken, &v)

	return v
}

type ApplicationCommandPostData struct {
	Name             string                     `json:"name"`
	Description      string                     `json:"description"`
	Options          []ApplicationCommandOption `json:"options,omitempty"`
	DefaultPermisson bool                       `json:"default_permisson,omitempty"`
}

func CreateApplicationCommands(appId SnowFlake, authToken string, guildId *SnowFlake, data ApplicationCommandPostData) ApplicationCommand {
	var reqPath string
	switch guildId {
	case nil:
		reqPath = path.Join("applications", fmt.Sprintf("%d", appId), "commands")
	default:
		reqPath = path.Join("applications", fmt.Sprintf("%d", appId), "guilds", fmt.Sprintf("%d", *guildId), "commands")
	}

	var v ApplicationCommand
	content, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(data)
	if err != nil {
		// TODO: err handling
	}
	sendRequest(reqPath, "POST", content, authToken, &v)

	return v
}

func GetApplicationCommand(appId SnowFlake, authToken string, cmdId SnowFlake, guildId *SnowFlake) ApplicationCommand {
	var reqPath string
	switch guildId {
	case nil:
		reqPath = path.Join("applications", fmt.Sprintf("%d", appId), "commands", fmt.Sprintf("%d", cmdId))
	default:
		reqPath = path.Join("applications", fmt.Sprintf("%d", appId), "guilds", fmt.Sprintf("%d", *guildId), "commands", fmt.Sprintf("%d", cmdId))
	}

	var v ApplicationCommand
	sendRequest(reqPath, "GET", nil, authToken, &v)

	return v
}

// FIXME: 内容は同じだけどtagだけ違うので微妙なところではある "...PostData"との統合を検討.
type ApplicationCommandPatchData struct {
	Name             string                     `json:"name,omitempty"`
	Description      string                     `json:"description,omitempty"`
	Options          []ApplicationCommandOption `json:"options,omitempty"`
	DefaultPermisson bool                       `json:"default_permisson,omitempty"`
}

func EditApplicationCommand(appId SnowFlake, authToken string, cmdId SnowFlake, guildId *SnowFlake, data ApplicationCommandPatchData) ApplicationCommand {
	var reqPath string
	switch guildId {
	case nil:
		reqPath = path.Join("applications", fmt.Sprintf("%d", appId), "commands", fmt.Sprintf("%d", cmdId))
	default:
		reqPath = path.Join("applications", fmt.Sprintf("%d", appId), "guilds", fmt.Sprintf("%d", *guildId), "commands", fmt.Sprintf("%d", cmdId))
	}

	var v ApplicationCommand
	content, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(data)
	if err != nil {
		// TODO: err handling
	}
	sendRequest(reqPath, "PATCH", content, authToken, &v)

	return v
}

func DeleteApplicationCommand(appId SnowFlake, authToken string, cmdId SnowFlake, guildId *SnowFlake) {
	var reqPath string
	switch guildId {
	case nil:
		reqPath = path.Join("applications", fmt.Sprintf("%d", appId), "commands", fmt.Sprintf("%d", cmdId))
	default:
		reqPath = path.Join("applications", fmt.Sprintf("%d", appId), "guilds", fmt.Sprintf("%d", *guildId), "commands", fmt.Sprintf("%d", cmdId))
	}

	sendRequest(reqPath, "DELETE", nil, authToken, nil)
}

func BulkOverwriteApplicationCommands(appId SnowFlake, authToken string, guildId *SnowFlake, data []ApplicationCommandPostData) []ApplicationCommand {
	var reqPath string
	switch guildId {
	case nil:
		reqPath = path.Join("applications", fmt.Sprintf("%d", appId), "commands")
	default:
		reqPath = path.Join("applications", fmt.Sprintf("%d", appId), "guilds", fmt.Sprintf("%d", *guildId), "commands")
	}

	var v []ApplicationCommand
	content, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(data)
	if err != nil {
		// TODO: err handling
	}
	sendRequest(reqPath, "PUT", content, authToken, v)

	return v
}

func CreateInteractionResponse(appId SnowFlake, authToken string, iaId SnowFlake, iaToken string, res InteractionReponse) {
	reqPath := path.Join("interactions", fmt.Sprintf("%d", iaId), fmt.Sprintf("%d", iaToken), "callback")
	sendRequest(reqPath, "POST", nil, authToken, nil)
}

type Message struct{}

func GetOriginalInteractionResponse(appId SnowFlake, authToken string, iaToken string) Message {
	reqPath := path.Join("webhooks", fmt.Sprintf("%d", appId), fmt.Sprint("%d", iaToken), "messages", "@original")
	var v Message
	sendRequest(reqPath, "GET", nil, authToken, &v)

	return v
}

type WebhookMessagePatchData struct{}

// FIXME: msgIdを使用すると"Edit Followup Message"なので命名に問題あり
func EditOriginalInteractionResponse(appId SnowFlake, authToken string, iaToken string, msgId *SnowFlake, data WebhookMessagePatchData) Message {
	var reqPath string
	switch msgId {
	case nil:
		reqPath = path.Join("webhooks", fmt.Sprintf("%d", appId), fmt.Sprintf("%d", iaToken), "messages", "@original")
	default:
		reqPath = path.Join("webhooks", fmt.Sprintf("%d", appId), fmt.Sprintf("%d", iaToken), "messages", fmt.Sprintf("%d", *msgId))
	}

	var v Message
	content, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(data)
	if err != nil {
		// TODO: err handling
	}
	sendRequest(reqPath, "PATCH", content, authToken, &v)

	return v
}

// FIXME: msgIdを使用すると"Delete Followup Message"なので命名に問題あり
func DeleteOriginalInteractionResponse(appId SnowFlake, authToken string, iaToken string, msgId *SnowFlake) {
	var reqPath string
	switch msgId {
	case nil:
		reqPath = path.Join("webhooks", fmt.Sprintf("%d", appId), fmt.Sprintf("%d", iaToken), "messages", "@original")
	default:
		reqPath = path.Join("webhooks", fmt.Sprintf("%d", appId), fmt.Sprintf("%d", iaToken), "messages", fmt.Sprintf("%d", *msgId))
	}

	sendRequest(reqPath, "DELETE", nil, authToken, nil)
}

type WebhookMessagePostData struct{}

func CreateFollowupMessage(appId SnowFlake, authToken string, iaToken string, data WebhookMessagePostData) {
	reqPath := path.Join("webhooks", fmt.Sprintf("%d", appId), fmt.Sprintf("%d", iaToken))
	content, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(data)
	if err != nil {
		// TODO: err handling
	}

	sendRequest(reqPath, "POST", content, authToken, nil)
}

// FIXME: error返しません?
func sendRequest(targetPath string, method string, content []byte, authToken string, rep interface{}) {
	url, err := url.Parse(BaseAPIUrl)
	if err != nil {
		// TODO: err handling
	}

	url.Path = path.Join(url.Path, targetPath)

	var buf *bytes.Buffer
	switch content {
	case nil:
		buf = nil
	default:
		buf = bytes.NewBuffer(content)
	}

	req, err := http.NewRequest(method, url.String(), buf)
	if err != nil {
		// TODO: err handling
	}

	req.Header.Add("Authorization", authToken)

	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		// TODO: err handling
	}

	resBuf := make([]byte, 0)
	for {
		size := 64
		v := make([]byte, size)
		resSize, err := res.Body.Read(v)
		if err != nil {
			// TODO: err handling
		}

		resBuf = append(resBuf, v...)

		if resSize < size {
			break
		}
	}

	// FIXME: pointerがnilだった場合, これは大丈夫?
	err = jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(resBuf, rep)
	if err != nil {
		// TODO: err handling
	}
}
