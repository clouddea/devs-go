package modeling

type Content struct {
	Source     string
	SourcePort string
	Target     string
	TargetPort string
	Payload    interface{}
}

type Message struct {
	contents []Content
}

func NewMessage() *Message {
	return &Message{contents: make([]Content, 0)}
}

/** 消息包是否为空 */
func (receiver Message) IsEmpty() bool {
	return len(receiver.contents) == 0
}

/** 清空消息包 */
func (receiver *Message) Clear() {
	receiver.contents = make([]Content, 0)
}

/** 添加多个消息 */
func (receiver *Message) Add(message Message) {
	receiver.contents = append(receiver.contents, message.contents...)
}

/** 添加一个消息 */
func (receiver *Message) AddContent(content Content) {
	receiver.contents = append(receiver.contents, content)
}
