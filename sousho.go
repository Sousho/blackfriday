package blackfriday

import (
	_ "embed"
	"log"
	"regexp"
	"strings"

	"github.com/jdkato/prose"
	"rogchap.com/v8go"
)

var ctx *v8go.Context

//go:embed ascii_math.js
var javascript string

/*go:embed ascii_symbols.js
var symbols string*/

var initialized bool = false

var math_regex []*regexp.Regexp = []*regexp.Regexp{
	regexp.MustCompile("\\\\((.*)\\\\)"),
	regexp.MustCompile("\\\\[(.*)\\\\]"),
}

func init_math() {
	if !initialized {
		//Info.Println("Ascii Math Initialized")
		ctx, _ = v8go.NewContext()
		//_, err := ctx.RunScript(symbols, "ascii_math.js")
		//e(err)
		_, err := ctx.RunScript(javascript, "ascii_math.js")
		e(err)
		initialized = true
	}
}

func render_shorthand(eq []byte) []byte {
	log.Println(string(eq))
	latex, err := ctx.RunScript("AMTparseAMtoTeX('"+string(eq)+"')", "ascii_math.js")
	e(err)
	return []byte("$" + latex.String() + "$")
}

var sousho_directives map[string]string = map[string]string{
	"start lemma":      "\\begin{lemma}",
	"end lemma":        "\\end{lemma}",
	"start theorem":    "\\begin{theorem}",
	"end theorem":      "\\end{theorem}",
	"start remark":     "\\begin{remark}",
	"end remark":       "\\end{remark}",
	"start exercise":   "\\begin{exercise}",
	"end exercise":     "\\end{exercise}",
	"start definition": "\\begin{definition}",
	"end definition":   "\\end{definition}",
}

/*func process_sousho(text []byte) []byte {
	result := []byte{}
	for _, regex := range math_regex {
		result = regex.ReplaceAllFunc(result, func(match []byte) []byte {
			matches := regex.FindSubmatch(match)
			//Info.Println("Listing submatches for " + string(match))
			//for _, submatch := range matches {
			//Info.Println(submatch)
			//}
			if len(matches) > 1 {
				return []byte(render_shorthand(string(matches[1])))
			}
			return match
		})
	}
	return []byte(result)
}*/

var part_of_speech map[string]uint16 = map[string]uint16{
	"(":    0,
	")":    1,
	",":    2,
	":":    3,
	".":    4,
	"''":   5,
	"#":    6,
	"$":    7,
	"CC":   8,
	"CD":   9,
	"DT":   10,
	"EX":   11,
	"FW":   12,
	"IN":   13,
	"JJ":   14,
	"JJR":  15,
	"JJS":  16,
	"LS":   17,
	"MD":   18,
	"NN":   19,
	"NNP":  20,
	"NNPS": 21,
	"NNS":  22,
	"PDT":  23,
	"POS":  24,
	"PRP":  25,
	"PRP$": 26,
	"RB":   27,
	"RBR":  28,
	"RBS":  29,
	"RP":   30,
	"SYM":  31,
	"TO":   32,
	"UH":   33,
	"VB":   34,
	"VBD":  35,
	"VBG":  36,
	"VBN":  37,
	"VBP":  38,
	"VBZ":  39,
	"WDT":  40,
	"WP":   41,
	"WP$":  42,
	"WRB":  42,
}

func label_freq(text []byte) []byte {
	text_str := string(text)
	result := ""
	doc, err := prose.NewDocument(text_str)
	e(err)
	for _, tok := range doc.Tokens() {
		log.Println(tok.Text, word_freq[0][[2]uint16{token_list[tok.Text], 0}])
		if word_freq[0][[2]uint16{token_list[tok.Text], 0}] <= 1e-10 {
			result += "\\hl{" + tok.Text + "} "
		} else {
			result += tok.Text + " "
		}
		//part_of_speech[tok.Tag]
	}
	return []byte(strings.Trim(result, " "))
}
