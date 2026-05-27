package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/kurokirasama/cmdg-image-render/pkg/render"
)

type imageMetadata struct {
	Index  int    `json:"index"`
	Source string `json:"source"`
}

type renderOutput struct {
	RenderedText string          `json:"rendered_text"`
	InlineImages []imageMetadata `json:"inline_images"`
}

func main() {
	width := flag.Int("width", 80, "Terminal width")
	startIdx := flag.Int("start-index", 0, "Starting index for image markers")
	flag.Parse()

	// Currently 'width' is not used by the underlying C engine, but we keep it
	// for future-proofing table rendering or word-wrapping.
	_ = width

	body, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Failed to read stdin: %v", err)
	}

	elems, ok := render.HTMLToElements(string(body))
	if !ok {
		// If rendering fails, just output the raw input as text.
		output := renderOutput{
			RenderedText: string(body),
		}
		if err := json.NewEncoder(os.Stdout).Encode(output); err != nil {
			log.Fatalf("Failed to encode JSON: %v", err)
		}
		return
	}

	var out strings.Builder
	var images []imageMetadata
	idx := *startIdx

	for _, elem := range elems {
		switch elem.Type {
		case render.HElemText:
			out.WriteString(elem.Text)
		case render.HElemH1:
			out.WriteString("\n\n=== ")
			out.WriteString(strings.ToUpper(elem.Text))
			out.WriteString(" ===\n\n")
		case render.HElemH2:
			out.WriteString("\n\n-- ")
			out.WriteString(elem.Text)
			out.WriteString(" --\n\n")
		case render.HElemLink:
			out.WriteString(elem.Text)
			if elem.Attr1 != "" {
				out.WriteString(fmt.Sprintf(" [%s]", elem.Attr1))
			}
		case render.HElemImage:
			images = append(images, imageMetadata{
				Index:  idx,
				Source: elem.Attr1,
			})
			fmt.Fprintf(&out, " ##IMG_%d_## ", idx)
			idx++
		case render.HElemBlockquote:
			out.WriteString("\n> ")
			if elem.Attr2 != "" {
				out.WriteString(elem.Attr2)
				out.WriteString("\n> ")
			}
			out.WriteString(strings.ReplaceAll(elem.Text, "\n", "\n> "))
			out.WriteString("\n")
		case render.HElemTable:
			out.WriteString("\n")
			out.WriteString(elem.Text)
			out.WriteString("\n")
		}
	}

	output := renderOutput{
		RenderedText: out.String(),
		InlineImages: images,
	}
	if err := json.NewEncoder(os.Stdout).Encode(output); err != nil {
		log.Fatalf("Failed to encode JSON: %v", err)
	}
}
