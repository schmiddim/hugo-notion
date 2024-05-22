package services

import (
	"blog-golang/data"
	"context"
	"fmt"
	"github.com/dstotijn/go-notion"
	log "github.com/sirupsen/logrus"
	"sync"
)

type NotionWrapper struct {
	token  string
	dbId   string
	client *notion.Client
}

func NewNotionWrapper(token string, dbId string) *NotionWrapper {

	return &NotionWrapper{token: token, dbId: dbId, client: notion.NewClient(token)}
}

func (n *NotionWrapper) GetClient() *notion.Client {

	return n.client

}
func (n *NotionWrapper) GetPostsFromNotionDB() []data.Post {
	var posts []data.Post
	postCh := make(chan data.Post)
	var wg sync.WaitGroup

	d := notion.DatabaseQuery{
		Filter: nil,
	}
	dbResults, err := n.client.QueryDatabase(context.Background(), n.dbId, &d)
	if err != nil {
		fmt.Println(err)
	}

	for _, page := range dbResults.Results {
		wg.Add(1)
		go func(page notion.Page) {
			defer wg.Done()
			p := data.Post{}

			properties, ok := page.Properties.(notion.DatabasePageProperties)
			if !ok {
				panic("Type Assertion failed")
			}
			p.DateCreated = page.CreatedTime
			titleProperty := properties["Name"].Title
			for _, text := range titleProperty {
				p.Title = text.PlainText
			}
			p.DateCreated = page.CreatedTime
			children, err := n.client.FindBlockChildrenByID(context.Background(), page.ID, nil)
			if err != nil {
				fmt.Println(err)
			}

			//Content Is a Child of the Page
			p.Content = n.readChildrenToMarkDown(children.Results)
			postCh <- p
		}(page)
	}

	go func() {
		wg.Wait()
		close(postCh)
	}()

	for post := range postCh {
		if post.Title != "" {
			posts = append(posts, post)
		}
	}

	return posts
}

func (n *NotionWrapper) readChildrenToMarkDown(blocks []notion.Block) string {
	s := ""
	md := Markdown{}
	for _, child := range blocks {
		switch block := child.(type) {
		case *notion.ParagraphBlock:
			s += md.mapParagraphBlock(block)
		case *notion.Heading1Block:
			s += md.mapHeading1Block(block)
		case *notion.Heading2Block:
			s += md.mapHeading2Block(block)
		case *notion.Heading3Block:
			s += md.mapHeading3Block(block)
		case *notion.ImageBlock:
			s += md.mapImageBlock(block)
		case *notion.BookmarkBlock:
			s += md.mapBookmarkBlock(block)
		case *notion.BulletedListItemBlock:
			s += md.mapBulletedListItemBlock(block)
		case *notion.DividerBlock:
			s += md.mapDividerBlock()
		case *notion.EmbedBlock:
			s += md.mapEmbedBlock(block)
		case *notion.NumberedListItemBlock:
			s += md.mapNumberedListItemBlock(block)
		case *notion.QuoteBlock:
			s += md.mapQuoteBlock(block)
		case *notion.VideoBlock:
			s += md.mapVideoBlock(block)
		case *notion.ToDoBlock:
			s += md.mapToDoBlock(block)
		case *notion.CodeBlock:
			s += md.mapCodeBlock(block)
		case *notion.TableBlock:
			s += md.mapTableBlock(block)
		case *notion.LinkPreviewBlock:
			s += md.mapLinkPreviewBlock(block)
		case *notion.EquationBlock:
			s += md.mapEquationBlock(block)
		default:
			log.Error("Unsupported block type")
			log.Error(block)

		}
	}
	return s + "\n"
}
