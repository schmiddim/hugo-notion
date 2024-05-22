package services

import (
	"fmt"
	"github.com/dstotijn/go-notion"
	"strings"
)

type Markdown struct {
}

func (m *Markdown) richi(p []notion.RichText) string {
	str := ""
	for _, text := range p {
		if text.HRef != nil {
			str += fmt.Sprintf("[%s](%s)", text.PlainText, *text.HRef)
		} else {
			str += text.PlainText
		}
		// possible todo: combinations of annotations
		if text.Annotations != nil {
			if text.Annotations.Bold {
				str = fmt.Sprintf("**%s** ", strings.TrimSpace(str))
			}
			if text.Annotations.Italic {
				str = fmt.Sprintf("_%s_ ", strings.TrimSpace(str))
			}
			if text.Annotations.Strikethrough {
				str = fmt.Sprintf("~~%s~~ ", strings.TrimSpace(str))
			}
			if text.Annotations.Underline {
				str = fmt.Sprintf("<u>%s</u> ", strings.TrimSpace(str))
			}
			if text.Annotations.Code {
				str = fmt.Sprintf("`%s` ", strings.TrimSpace(str))
			}

		}
	}
	return str
}

// deprecated
func (m *Markdown) richTextToMarkdown(r notion.RichText) string {
	return fmt.Sprint(r.PlainText)
}

func (m *Markdown) mapLinkPreviewBlock(l *notion.LinkPreviewBlock) string {
	return fmt.Sprintf("[Link](%s)", l.URL)
}

func (m *Markdown) mapTableBlock(t *notion.TableBlock) string {
	str := "+---+Implement the Table---+---+\n"

	return str
}

func (m *Markdown) mapEquationBlock(t *notion.EquationBlock) string {
	return fmt.Sprintf("$$%s$$", t.Expression)
}

func (m *Markdown) mapParagraphBlock(p *notion.ParagraphBlock) string {
	str := m.richi(p.RichText)
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
	return "## " + m.richi(t.RichText)
}
func (m *Markdown) mapBulletedListItemBlock(b *notion.BulletedListItemBlock) string {
	str := "* " + m.richi(b.RichText) + "\n"

	return str
}
func (m *Markdown) mapCodeBlock(b *notion.CodeBlock) string {
	language := ""
	if b.Language != nil {
		language = *b.Language
	}

	return fmt.Sprintf("```%s\n%s\n```", language, strings.TrimSpace(m.richi(b.RichText)))
}

func (m *Markdown) mapEmbedBlock(b *notion.EmbedBlock) string {
	return fmt.Sprintf("[Embed](%s)", b.URL)
}

func (m *Markdown) mapDividerBlock() string {
	return "---"
}
func (m *Markdown) mapQuoteBlock(b *notion.QuoteBlock) string {
	return "> " + m.richi(b.RichText)
}

func (m *Markdown) mapToDoBlock(b *notion.ToDoBlock) string {
	return "TODO: " + m.richi(b.RichText)
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
	return "# " + m.richi(h.RichText)

}
func (m *Markdown) mapHeading2Block(h *notion.Heading2Block) string {
	return "## " + m.richi(h.RichText)
}
func (m *Markdown) mapHeading3Block(h *notion.Heading3Block) string {
	return "### " + m.richi(h.RichText)
}
