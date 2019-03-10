package youporn

type YoupornEmbedCode struct {
	Embed YoupornCode `json:"embed,omitempty"`
}

type YoupornCode struct {
	Code string `json:"code,omitempty"`
}
