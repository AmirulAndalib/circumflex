package comment

import (
	"regexp"
	"strings"

	text "github.com/MichaelMure/go-term-text"
)

type Comment struct {
	Sections []*section
}

type section struct {
	IsCodeBlock bool
	IsQuote     bool
	Text        string
}

const (
	singleSpace = " "
	doubleSpace = "  "
	tripleSpace = "   "
)

func ParseComment(c string, commentWidth int, availableScreenWidth int) (string, []string) {
	if c == "[deleted]" {
		return Dimmed + "[deleted]" + Normal, []string{}
	}

	c = strings.Replace(c, "<p>", "", 1)
	c = strings.ReplaceAll(c, "\n</code></pre>\n", "<p>")
	paragraphs := strings.Split(c, "<p>")

	comment := new(Comment)
	comment.Sections = make([]*section, len(paragraphs))

	for i, paragraph := range paragraphs {
		s := new(section)
		s.Text = replaceCharacters(paragraph)

		if strings.Contains(s.Text, "<pre><code>") {
			s.IsCodeBlock = true
		}

		if isQuote(s.Text) {
			s.IsQuote = true
		}

		comment.Sections[i] = s
	}

	output := ""

	var URLs []string

	for i, s := range comment.Sections {
		paragraph := s.Text

		switch {
		case s.IsQuote:
			paragraph = strings.ReplaceAll(paragraph, "<i>", "")
			paragraph = strings.ReplaceAll(paragraph, "</i>", "")
			paragraph = strings.ReplaceAll(paragraph, "</a>", "")
			paragraph = strings.ReplaceAll(paragraph, tripleSpace, singleSpace)
			paragraph = strings.ReplaceAll(paragraph, doubleSpace, singleSpace)
			paragraph = strings.Replace(paragraph, ">>", "", 1)
			paragraph = strings.Replace(paragraph, ">", "", 1)
			paragraph = strings.ReplaceAll(paragraph, "@dang", green("@dang"))
			paragraph = strings.TrimLeft(paragraph, " ")
			paragraph = trimURLs(paragraph)

			paragraph = Italic + Dimmed + paragraph + Normal

			indentBlock := " ▎"
			padding := text.WrapPad(Dimmed + indentBlock)
			wrappedAndPaddedComment, _ := text.Wrap(paragraph, commentWidth, padding)
			paragraph = wrappedAndPaddedComment

		case s.IsCodeBlock:
			paragraph = replaceHTML(paragraph)

			paddingWithBlock := text.WrapPad("")
			wrappedAndPaddedComment, _ := text.Wrap(paragraph, availableScreenWidth, paddingWithBlock)

			codeLines := strings.Split(wrappedAndPaddedComment, NewLine)
			formattedCodeLines := ""

			for j, codeLine := range codeLines {
				isOnLastLine := j == len(codeLines)-1

				if isOnLastLine {
					formattedCodeLines += Dimmed + codeLine + Normal

					break
				}

				formattedCodeLines += Dimmed + codeLine + Normal + NewLine
			}

			paragraph = formattedCodeLines

		default:
			paragraph = strings.ReplaceAll(paragraph, tripleSpace, singleSpace)
			paragraph = strings.ReplaceAll(paragraph, doubleSpace, singleSpace)
			paragraph = highlightReferences(paragraph)
			paragraph = replaceHTML(paragraph)
			paragraph = strings.ReplaceAll(paragraph, "@dang", green("@dang"))
			paragraph = strings.TrimLeft(paragraph, " ")
			paragraph = highlightBackticks(paragraph)

			URLs = append(URLs, extractURLs(paragraph)...)
			paragraph = trimURLs(paragraph)

			padding := text.WrapPad("")
			wrappedAndPaddedComment, _ := text.Wrap(paragraph, commentWidth, padding)
			paragraph = wrappedAndPaddedComment
		}

		separator := getParagraphSeparator(i, len(comment.Sections))
		output += paragraph + separator
	}

	return output, URLs
}

func isQuote(text string) bool {
	quoteMark := ">"

	return strings.HasPrefix(text, quoteMark) ||
		strings.HasPrefix(text, " "+quoteMark) ||
		strings.HasPrefix(text, "<i>"+quoteMark) ||
		strings.HasPrefix(text, "<i> "+quoteMark)
}

func getParagraphSeparator(index int, sliceLength int) string {
	isAtLastParagraph := index == sliceLength-1

	if isAtLastParagraph {
		return ""
	}

	return NewParagraph
}

func replaceCharacters(input string) string {
	input = strings.ReplaceAll(input, "&#x27;", "'")
	input = strings.ReplaceAll(input, "&gt;", ">")
	input = strings.ReplaceAll(input, "&lt;", "<")
	input = strings.ReplaceAll(input, "&#x2F;", "/")
	input = strings.ReplaceAll(input, "&quot;", `"`)
	input = strings.ReplaceAll(input, "&amp;", "&")

	return input
}

func replaceHTML(input string) string {
	input = strings.Replace(input, "<p>", "", 1)

	input = strings.ReplaceAll(input, "<p>", NewParagraph)
	input = strings.ReplaceAll(input, "<i>", Italic)
	input = strings.ReplaceAll(input, "</i>", Normal)
	input = strings.ReplaceAll(input, "</a>", "")
	input = strings.ReplaceAll(input, "<pre><code>", "")
	input = strings.ReplaceAll(input, "</code></pre>", "")

	return input
}

func highlightReferences(input string) string {
	input = strings.ReplaceAll(input, "[0]", "["+white("0")+"]")
	input = strings.ReplaceAll(input, "[1]", "["+red("1")+"]")
	input = strings.ReplaceAll(input, "[2]", "["+yellow("2")+"]")
	input = strings.ReplaceAll(input, "[3]", "["+green("3")+"]")
	input = strings.ReplaceAll(input, "[4]", "["+blue("4")+"]")
	input = strings.ReplaceAll(input, "[5]", "["+cyan("5")+"]")
	input = strings.ReplaceAll(input, "[6]", "["+magenta("6")+"]")
	input = strings.ReplaceAll(input, "[7]", "["+altWhite("7")+"]")
	input = strings.ReplaceAll(input, "[8]", "["+altRed("8")+"]")
	input = strings.ReplaceAll(input, "[9]", "["+altYellow("9")+"]")
	input = strings.ReplaceAll(input, "[10]", "["+altGreen("10")+"]")

	return input
}

func extractURLs(input string) []string {
	expForFirstTag := regexp.MustCompile(`<a href=".*?" rel="nofollow">`)
	URLs := expForFirstTag.FindAllString(input, 10)

	for i := range URLs {
		URLs[i] = strings.ReplaceAll(URLs[i], `<a href="`, "")
		URLs[i] = strings.ReplaceAll(URLs[i], `" rel="nofollow">`, "")
	}

	return URLs
}

func trimURLs(comment string) string {
	expression := regexp.MustCompile(`<a href=".*?" rel="nofollow">`)

	return expression.ReplaceAllString(comment, "")
}

func highlightBackticks(input string) string {
	backtick := "`"
	numberOfBackticks := strings.Count(input, backtick)
	numberOfBackticksIsOdd := numberOfBackticks%2 != 0

	if numberOfBackticks == 0 || numberOfBackticksIsOdd {
		return input
	}

	isOnFirstBacktick := true

	for i := 0; i < numberOfBackticks+1; i++ {
		if isOnFirstBacktick {
			input = strings.Replace(input, backtick, Blue, 1)
		} else {
			input = strings.Replace(input, backtick, Normal, 1)
		}

		isOnFirstBacktick = !isOnFirstBacktick
	}

	return input
}
