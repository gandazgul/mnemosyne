package search

import (
	"math"
	"strings"
	"unicode"
)

const (
	defaultBM25K1 = 1.5
	defaultBM25B  = 0.75
)

func bm25Scores(query string, documents []string) []float64 {
	nDocs := len(documents)
	if nDocs == 0 {
		return nil
	}

	queryTerms := uniqueLexicalTerms(query)
	if len(queryTerms) == 0 {
		return make([]float64, nDocs)
	}

	tokenized := make([][]string, nDocs)
	docLens := make([]int, nDocs)
	df := make(map[string]int, len(queryTerms))

	for i, doc := range documents {
		tokens := lexicalTokens(doc)
		tokenized[i] = tokens
		docLens[i] = len(tokens)

		seen := make(map[string]bool, len(queryTerms))
		for _, token := range tokens {
			if _, ok := queryTerms[token]; ok {
				seen[token] = true
			}
		}
		for token := range seen {
			df[token]++
		}
	}

	var totalLen int
	for _, length := range docLens {
		totalLen += length
	}
	if totalLen == 0 {
		return make([]float64, nDocs)
	}
	avgDocLen := float64(totalLen) / float64(nDocs)

	idf := make(map[string]float64, len(queryTerms))
	for term := range queryTerms {
		termDF := df[term]
		idf[term] = math.Log((float64(nDocs)-float64(termDF)+0.5)/(float64(termDF)+0.5) + 1)
	}

	scores := make([]float64, nDocs)
	for i, tokens := range tokenized {
		if docLens[i] == 0 {
			continue
		}

		tf := make(map[string]int, len(queryTerms))
		for _, token := range tokens {
			if _, ok := queryTerms[token]; ok {
				tf[token]++
			}
		}

		docLen := float64(docLens[i])
		for term, freq := range tf {
			f := float64(freq)
			numerator := f * (defaultBM25K1 + 1)
			denominator := f + defaultBM25K1*(1-defaultBM25B+defaultBM25B*docLen/avgDocLen)
			scores[i] += idf[term] * numerator / denominator
		}
	}

	return scores
}

func normalizeMax(scores []float64) []float64 {
	normalized := make([]float64, len(scores))
	var maxScore float64
	for _, score := range scores {
		if score > maxScore {
			maxScore = score
		}
	}
	if maxScore <= 0 {
		return normalized
	}
	for i, score := range scores {
		normalized[i] = score / maxScore
	}
	return normalized
}

func uniqueLexicalTerms(text string) map[string]struct{} {
	tokens := lexicalTokens(text)
	terms := make(map[string]struct{}, len(tokens))
	for _, token := range tokens {
		terms[token] = struct{}{}
	}
	return terms
}

func lexicalTokens(text string) []string {
	var tokens []string
	var b strings.Builder
	flush := func() {
		if b.Len() == 0 {
			return
		}
		token := strings.ToLower(b.String())
		b.Reset()
		if len(token) < 2 || isLexicalStopword(token) {
			return
		}
		tokens = append(tokens, token)
	}

	for _, r := range text {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteRune(r)
			continue
		}
		flush()
	}
	flush()

	return tokens
}

func isLexicalStopword(token string) bool {
	switch token {
	case "a", "an", "and", "are", "as", "at", "be", "but", "by", "can", "do", "does", "for", "from",
		"had", "has", "have", "how", "i", "if", "in", "into", "is", "it", "its", "of", "on", "or",
		"should", "that", "the", "their", "then", "there", "these", "this", "to", "was", "were",
		"what", "when", "where", "which", "who", "why", "with", "would", "you", "your":
		return true
	default:
		return false
	}
}
