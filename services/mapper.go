package services

import (
	"fmt"
	"github.com/dstotijn/go-notion"
)

type Markdown struct {
}

func (m *Markdown) richTextToMarkdown(r notion.RichText) string {
	return fmt.Sprintln(r.PlainText + "\n\n")
}

func (m *Markdown) mapLinkPreviewBlock(l *notion.LinkPreviewBlock) string {
	return fmt.Sprintf("[Link](%s)", l.URL)
}

func (m *Markdown) mapTableBlock(t *notion.TableBlock) string {
	str := "+---+Implement the Table---+---+\n"

	return str
}

func (m *Markdown) mapParagraphBlock(p *notion.ParagraphBlock) string {
	str := ""
	for _, text := range p.RichText {
		str += m.richTextToMarkdown(text)
	}
	return str
}

func (m *Markdown) mapVideoBlock(v *notion.VideoBlock) string {
	if v.File == nil {
		return fmt.Sprintf("[Video](%s)", v.External.URL)
	} else {
		return fmt.Sprintf("[Video](%s)", "@todododododo")
	}

}
func (m *Markdown) mapToggleBlock(t *notion.ToggleBlock) string {
	str := ""
	for _, text := range t.RichText {
		str += "## " + m.richTextToMarkdown(text)
	}
	return str
}
func (m *Markdown) mapBulletedListItemBlock(b *notion.BulletedListItemBlock) string {
	str := ""
	for _, text := range b.RichText {
		str += "* " + m.richTextToMarkdown(text)
	}
	return str
}
func (m *Markdown) mapCodeBlock(b *notion.CodeBlock) string {
	return fmt.Sprintf("```%s\n%s\n```", b.Language, m.richTextToMarkdown(b.RichText[0]))
}

func (m *Markdown) mapEmbedBlock(b *notion.EmbedBlock) string {
	return fmt.Sprintf("[Embed](%s)", b.URL)
}

func (m *Markdown) mapDividerBlock() string {
	return "---"
}
func (m *Markdown) mapQuoteBlock(b *notion.QuoteBlock) string {
	str := ""
	for _, text := range b.RichText {
		str += "> " + m.richTextToMarkdown(text)
	}
	return str
}

func (m *Markdown) mapToDoBlock(b *notion.ToDoBlock) string {
	str := ""
	for _, text := range b.RichText {
		str += "TODO: " + m.richTextToMarkdown(text)
	}
	return str
}

func (m *Markdown) mapBookmarkBlock(b *notion.BookmarkBlock) string {
	return fmt.Sprintf("[Bookmark](%s)", b.URL)
}

func (m *Markdown) mapImageBlock(b *notion.ImageBlock) string {
	return fmt.Sprintf("![%s](%s)", b.Caption, b.File.URL)
}

func (m *Markdown) mapNumberedListItemBlock(b *notion.NumberedListItemBlock) string {
	str := ""
	for i, text := range b.RichText {
		str += fmt.Sprintf("%d. %s", i+1, m.richTextToMarkdown(text))
	}
	return str
}

func (m *Markdown) mapHeading1Block(h *notion.Heading1Block) string {
	str := ""
	for _, text := range h.RichText {
		str += "# " + m.richTextToMarkdown(text)
	}
	return str
}
func (m *Markdown) mapHeading2Block(h *notion.Heading2Block) string {
	str := ""
	for _, text := range h.RichText {
		str += "## " + m.richTextToMarkdown(text)
	}
	return str
}
func (m *Markdown) mapHeading3Block(h *notion.Heading3Block) string {
	str := ""
	for _, text := range h.RichText {
		str += "### " + m.richTextToMarkdown(text)
	}
	return str
}
