package chunker

import (
	"strings"

	"github.com/yuin/goldmark/ast"
)

// Chunk represents a semantically coherent block of markdown text with its hierarchical context.
type Chunk struct {
	Path    []string
	Content string
}

func findLineStart(source []byte, offset int) int {
	for i := offset - 1; i >= 0; i-- {
		if source[i] == '\n' {
			return i + 1
		}
	}
	return 0
}

// ChunkDocument parses a markdown document and splits it into semantic chunks based on headings.
// It also ensures no chunk exceeds maxChars by falling back to paragraph/sentence splitting.
func ChunkDocument(source []byte, maxChars int) []Chunk {
	root := ParseMarkdown(source)
	contexts := ExtractContext(root, source)

	var chunks []Chunk
	chunkStart := 0
	var currentPath []string

	for _, ctx := range contexts {
		if ctx.Node.Kind() == ast.KindHeading {
			lines := ctx.Node.Lines()
			if lines.Len() > 0 {
				startOffset := lines.At(0).Start
				lineStart := findLineStart(source, startOffset)

				content := string(source[chunkStart:lineStart])
				chunks = append(chunks, processChunk(currentPath, content, maxChars)...)

				chunkStart = lineStart
				currentPath = ctx.Path
			}
		}
	}

	if chunkStart < len(source) {
		content := string(source[chunkStart:])
		chunks = append(chunks, processChunk(currentPath, content, maxChars)...)
	}

	return chunks
}

func processChunk(path []string, content string, maxChars int) []Chunk {
	content = strings.TrimSpace(content)
	if content == "" {
		return nil
	}

	prefix := formatChunk(path, "")
	limit := maxChars - len(prefix)
	if limit < 50 {
		limit = 50 // Minimum sensible limit to avoid infinite loops if path is huge
	}

	subContents := splitContent(content, limit)
	var chunks []Chunk
	for _, sc := range subContents {
		chunks = append(chunks, Chunk{
			Path:    path,
			Content: formatChunk(path, sc),
		})
	}
	return chunks
}

func splitContent(content string, limit int) []string {
	if len(content) <= limit {
		return []string{content}
	}

	var result []string
	paragraphs := strings.Split(content, "\n\n")

	var currentChunk string
	for _, p := range paragraphs {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}

		if len(p) > limit {
			// Flush current
			if currentChunk != "" {
				result = append(result, strings.TrimSpace(currentChunk))
				currentChunk = ""
			}
			// Split the large paragraph by sentences (rough approximation)
			sentences := strings.Split(p, ". ")
			var currentSent string
			for i, s := range sentences {
				sText := s
				if i < len(sentences)-1 {
					sText += ". "
				}
				if len(currentSent)+len(sText) > limit {
					if currentSent != "" {
						result = append(result, strings.TrimSpace(currentSent))
						currentSent = sText
					} else {
						// Single sentence is larger than limit, force append
						result = append(result, strings.TrimSpace(sText))
						currentSent = ""
					}
				} else {
					currentSent += sText
				}
			}
			if currentSent != "" {
				result = append(result, strings.TrimSpace(currentSent))
			}
		} else {
			// Try to add paragraph to current chunk
			addedLen := len(p)
			if currentChunk != "" {
				addedLen += 2 // account for \n\n
			}
			if len(currentChunk)+addedLen > limit {
				result = append(result, strings.TrimSpace(currentChunk))
				currentChunk = p
			} else {
				if currentChunk == "" {
					currentChunk = p
				} else {
					currentChunk += "\n\n" + p
				}
			}
		}
	}
	if currentChunk != "" {
		result = append(result, strings.TrimSpace(currentChunk))
	}
	return result
}

func formatChunk(path []string, content string) string {
	content = strings.TrimSpace(content)
	var cleanPath []string
	for _, p := range path {
		if p != "" {
			cleanPath = append(cleanPath, p)
		}
	}

	if len(cleanPath) == 0 {
		return content
	}

	pathStr := strings.Join(cleanPath, " > ")
	if content == "" {
		return "[Path: " + pathStr + "]"
	}
	return "[Path: " + pathStr + "]\n\n" + content
}
