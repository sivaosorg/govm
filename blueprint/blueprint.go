package blueprint

import (
	"bytes"
	"html/template"
	"log"
	"net/url"

	"github.com/sivaosorg/govm/builder"
	"github.com/sivaosorg/govm/utils"
)

func NewCard() *card {
	c := &card{}
	return c
}

func (c *card) SetIconText(value IconText) *card {
	c.IconText = value
	c.Icon = TypeIcons[value]
	return c
}

func (c *card) SetTitle(value string) *card {
	c.Title = value
	return c
}

func (c *card) SetDescription(value string) *card {
	c.Description = value
	return c
}

func (c *card) SetDescriptionWith(value builder.MapBuilder) *card {
	c.SetDescription(value.Json())
	return c
}

func (c *card) SetImageUrl(value string) *card {
	url, err := url.Parse(value)
	if err != nil {
		log.Fatalf("Parse url %v got an error: %v", value, err.Error())
	}
	c.ImageUrl = url.String()
	return c
}

func (c *card) SetButtonText(value string) *card {
	c.ButtonText = value
	return c
}

func (c *card) SetButtonUrl(value string) *card {
	url, err := url.Parse(value)
	if err != nil {
		log.Fatalf("Parse url %v got an error: %v", value, err.Error())
	}
	c.ButtonUrl = url.String()
	return c
}

func (c *card) Json() string {
	return utils.ToJson(c)
}

func (c *card) GenCardDefault() string {
	if utils.IsEmpty(c.Icon) {
		c.Icon = TypeIcons[TypeNotification]
	}
	return c.GenCard(CardDefault)
}

func (c *card) GenCard(layout string) string {
	var docs bytes.Buffer
	t := template.Must(template.New("card").Parse(layout))
	err := t.Execute(&docs, c)
	if err != nil {
		log.Fatalf("Parse template got an error: %v", err.Error())
	}
	return docs.String()
}

func CardValidator(c *card) {
	c.SetImageUrl(c.ImageUrl).
		SetButtonUrl(c.ButtonUrl)
}

func GetCardSample() *card {
	c := NewCard().
		SetTitle("Alert").
		SetIconText(TypeNotification).
		SetDescription("You are the good person").
		SetButtonText("Thumps").
		SetButtonUrl("https://www.google.com")
	return c
}
