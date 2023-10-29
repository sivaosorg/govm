package md

import (
	"fmt"
	"strings"
)

// NewMarkdownBuilder creates a new MarkdownBuilder instance.
func NewMarkdownBuilder() *MarkdownBuilder {
	return &MarkdownBuilder{}
}

// AddHeader adds a Markdown header to the content.
func (builder *MarkdownBuilder) AddHeader(text string, level int) *MarkdownBuilder {
	if level < 1 || level > 6 {
		level = 1
	}
	fmt.Fprintf(&builder.content, "%s %s\n", strings.Repeat("#", level), text)
	return builder
}

// AddText adds plain text to the content.
func (builder *MarkdownBuilder) AddText(text string) *MarkdownBuilder {
	builder.content.WriteString(text)
	return builder
}

// AddLink adds a Markdown link to the content.
func (builder *MarkdownBuilder) AddLink(text, url string) *MarkdownBuilder {
	fmt.Fprintf(&builder.content, "[%s](%s)", text, url)
	return builder
}

// AddList adds a Markdown list to the content.
func (builder *MarkdownBuilder) AddList(items []string, ordered bool) *MarkdownBuilder {
	listType := "*"
	if ordered {
		listType = "1."
	}
	for _, item := range items {
		fmt.Fprintf(&builder.content, "%s %s\n", listType, item)
	}
	return builder
}

// AddCodeBlock adds a code block to the content.
func (builder *MarkdownBuilder) AddCodeBlock(code, language string) *MarkdownBuilder {
	builder.content.WriteString("```")
	if language != "" {
		builder.content.WriteString(language)
	}
	builder.content.WriteString("\n")
	builder.content.WriteString(code)
	builder.content.WriteString("\n```\n")
	return builder
}

// AddImage adds an image to the content.
func (builder *MarkdownBuilder) AddImage(altText, imageUrl string) *MarkdownBuilder {
	fmt.Fprintf(&builder.content, "![%s](%s)\n", altText, imageUrl)
	return builder
}

// AddTable adds a table to the content.
func (builder *MarkdownBuilder) AddTable(headers []string, rows [][]string) *MarkdownBuilder {
	numColumns := len(headers)

	// Add table headers
	for i, header := range headers {
		builder.content.WriteString(header)
		if i < numColumns-1 {
			builder.content.WriteString(" | ")
		}
	}
	builder.content.WriteString("\n")
	// Add table header separator
	for i := 0; i < numColumns; i++ {
		builder.content.WriteString("---")
		if i < numColumns-1 {
			builder.content.WriteString(" | ")
		}
	}
	builder.content.WriteString("\n")
	// Add table rows
	for _, row := range rows {
		for i, cell := range row {
			builder.content.WriteString(cell)
			if i < numColumns-1 {
				builder.content.WriteString(" | ")
			}
		}
		builder.content.WriteString("\n")
	}
	return builder
}

// AddBlockQuote adds a block quote to the content.
func (builder *MarkdownBuilder) AddBlockQuote(quote string) *MarkdownBuilder {
	builder.content.WriteString("> ")
	builder.content.WriteString(quote)
	builder.content.WriteString("\n")
	return builder
}

// AddHorizontalRule adds a horizontal rule to the content.
func (builder *MarkdownBuilder) AddHorizontalRule() *MarkdownBuilder {
	builder.content.WriteString("\n---\n")
	return builder
}

// AddNestedList adds a nested list to the content.
func (builder *MarkdownBuilder) AddNestedList(items [][]string, ordered bool) *MarkdownBuilder {
	listType := "*"
	if ordered {
		listType = "1."
	}
	for _, item := range items {
		builder.content.WriteString("  ") // Indent for nesting
		for i, subItem := range item {
			builder.content.WriteString(listType + " " + subItem)
			if i < len(item)-1 {
				builder.content.WriteString("\n")
			}
		}
		builder.content.WriteString("\n")
	}
	return builder
}

// AddFootnote adds a footnote to the content.
func (builder *MarkdownBuilder) AddFootnote(footnoteID, footnoteText string) *MarkdownBuilder {
	builder.content.WriteString("[^")
	builder.content.WriteString(footnoteID)
	builder.content.WriteString("]: ")
	builder.content.WriteString(footnoteText)
	builder.content.WriteString("\n")
	return builder
}

// AddTaskList adds a task list (checklist) to the content.
func (builder *MarkdownBuilder) AddTaskList(items []string) *MarkdownBuilder {
	for _, item := range items {
		builder.content.WriteString("- [ ] ")
		builder.content.WriteString(item)
		builder.content.WriteString("\n")
	}
	return builder
}

// AddDefinitionList adds a definition list to the content.
func (builder *MarkdownBuilder) AddDefinitionList(definitions map[string]string) *MarkdownBuilder {
	for term, definition := range definitions {
		builder.content.WriteString(term)
		builder.content.WriteString(":\n")
		builder.content.WriteString("  ")
		builder.content.WriteString(definition)
		builder.content.WriteString("\n")
	}
	return builder
}

// AddInlineCode adds an inline code span to the content.
func (builder *MarkdownBuilder) AddInlineCode(code string) *MarkdownBuilder {
	builder.content.WriteString("`")
	builder.content.WriteString(code)
	builder.content.WriteString("`")
	return builder
}

// AddSuperscript adds a superscript to the content.
func (builder *MarkdownBuilder) AddSuperscript(text string) *MarkdownBuilder {
	builder.content.WriteString("<sup>")
	builder.content.WriteString(text)
	builder.content.WriteString("</sup>")
	return builder
}

// AddStrike-through adds strike-through text to the content.
func (builder *MarkdownBuilder) AddStrikeThrough(text string) *MarkdownBuilder {
	builder.content.WriteString("~~")
	builder.content.WriteString(text)
	builder.content.WriteString("~~")
	return builder
}

// AddInlineLink adds an inline link to the content.
func (builder *MarkdownBuilder) AddInlineLink(text, url string) *MarkdownBuilder {
	builder.content.WriteString("[")
	builder.content.WriteString(text)
	builder.content.WriteString("](")
	builder.content.WriteString(url)
	builder.content.WriteString(")")
	return builder
}

// Build returns the constructed Markdown content as a string.
func (builder *MarkdownBuilder) Build() string {
	return builder.content.String()
}
