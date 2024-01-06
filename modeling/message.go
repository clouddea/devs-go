package modeling

type Content struct {
	Source     string
	SourcePort string
	Target     string
	TargetPort string
	Payload    interface{}
}

type Message struct {
	Contents []Content `json:"Contents"`
}

func NewMessage() *Message {
	return &Message{Contents: make([]Content, 0)}
}

/** 消息包是否为空 */
func (receiver Message) IsEmpty() bool {
	return len(receiver.Contents) == 0
}

/** 清空消息包 */
func (receiver *Message) Clear() {
	receiver.Contents = make([]Content, 0)
}

/** 添加多个消息 */
func (receiver *Message) Add(message Message) {
	receiver.Contents = append(receiver.Contents, message.Contents...)
}

/** 添加一个消息 */
func (receiver *Message) AddContent(content Content) {
	receiver.Contents = append(receiver.Contents, content)
}

/** 添加一个消息 */
func (receiver *Message) GetContents() []Content {
	return receiver.Contents
}
