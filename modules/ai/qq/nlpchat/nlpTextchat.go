package qq

import (
	"chs/modules/ai/qq/pojo"
	"chs/modules/ai/qq/util"
	"log"
)

const nlpTextchatUrl = "https://api.ai.qq.com/fcgi-bin/nlp/nlp_textchat"

type AiNlpTextchat struct {
	config   *Configuration
	answer   pojo.Answer
	session  string
	question string
}

//type DataJson struct {
//	Session    string `json:"session"`
//	AnswerData string `json:"answer"`
//}

var nlpTextchat *AiNlpTextchat

func NewNlpTextchat(conf ...*Configuration) *AiNlpTextchat {
	if nlpTextchat == nil {
		if conf == nil {
			conf[0] = DefaultConfiguration()
		}
		nlpTextchat = &AiNlpTextchat{
			config: conf[0],
			answer: pojo.Answer{},
		}
	}
	return nlpTextchat
}

func (this *AiNlpTextchat) ToMap() map[string]string {
	if this.session == "" || this.question == "" {
		log.Println("AiNlpTextchat question or session err")
		return nil
	}
	config := this.config.ToMap()
	config["session"] = this.session
	config["question"] = this.question
	return config
}

func (this *AiNlpTextchat) Question(question, session string) *AiNlpTextchat {
	this.question = question
	this.session = session
	return this
}

func (this *AiNlpTextchat) Answer() string {
	requestBody := GetRequestBody(this)
	err := util.HttpPostJSON(nlpTextchatUrl, requestBody, &this.answer)
	if err != nil || this.answer.ErrCode != 0 {
		log.Printf("QQ Ai request err :%v answer {%v}", err, this.answer)
		return ""
	}
	answerData := this.answer.DataJson.(map[string]interface{})
	return answerData["answer"].(string)
}
