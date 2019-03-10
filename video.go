package youporn

type YoupornVideo struct {
	ID           string            `json:"video_id,omitempty"`
	Duration     string            `json:"duration,omitempty"`
	Views        string            `json:"views,omitempty"`
	Rating       float64           `json:"rating,omitempty"`
	Ratings      float64           `json:"ratings,omitempty"`
	Title        string            `json:"title,omitempty"`
	URL          string            `json:"url,omitempty"`
	DefaultThumb string            `json:"default_thumb,omitempty"`
	Thumb        string            `json:"thumb,omitempty"`
	PublishDate  string            `json:"publish_date,omitempty"`
	Thumbs       []YoupornThumb    `json:"thumbs,omitempty"`
	Tags         []YoupornTag      `json:"tags,omitempty"`
	Pornstars    []YoupornPornstar `json:"pornstar,omitempty"`
}

type YoupornThumb struct {
	Size   string  `json:"size,omitempty"`
	Width  float64 `json:"width,omitempty"`
	Height float64 `json:"height,omitempty"`
	Src    string  `json:"src,omitempty"`
}

type YoupornTag struct {
	Name string `json:"tag_name,omitempty"`
}

type YoupornPornstar struct {
	Name string `json:"pornstar_name,omitempty"`
}
