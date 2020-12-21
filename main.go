package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/bluele/mecab-golang"
)

var count map[string]int

func getWords() []string {
	count = make(map[string]int)
	var words []string

	tmp, _ := ioutil.ReadAll(os.Stdin)
	words = strings.Split(string(tmp), ",")
	return words
}

func parseToNode(m *mecab.MeCab) {
	tg, err := m.NewTagger()
	if err != nil {
		panic(err)
	}
	defer tg.Destroy()

	yukarionWords := getWords()

	for _, word := range yukarionWords {
		lt, err := m.NewLattice(word)
		if err != nil {
			panic(err)
		}
		defer lt.Destroy()

		node := tg.ParseToNode(lt)
		for {
			features := strings.Split(node.Feature(), ",")
			if features[0] == "名詞" {
				count[node.Surface()] += 1
			}
			if node.Next() != nil {
				break
			}
		}
	}
}

func main() {
	m, err := mecab.New("-d /usr/lib/mecab/dic/mecab-ipadic-neologd")
	if err != nil {
		panic(err)
	}
	defer m.Destroy()
	parseToNode(m)

	for str, cnt := range count {
		if cnt >= 3 {
			fmt.Printf("%s,%d\n", str, cnt)
		}
	}
}
