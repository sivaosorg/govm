package blueprint

type IconText string

type card struct {
	Icon        string   `json:"icon"`
	IconText    IconText `json:"icon_text,omitempty"`
	Title       string   `json:"title,omitempty"`
	Description string   `json:"description,omitempty"`
	ImageUrl    string   `json:"image_url,omitempty"`
	ButtonText  string   `json:"button_text,omitempty"`
	ButtonUrl   string   `json:"button_url,omitempty"`
}
