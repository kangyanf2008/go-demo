package main

import (
	"fmt"
	"github.com/ynqa/wego/pkg/embedding"
	"os"
)

func main() {
	/*
		model, err := word2vec.New(
				word2vec.Window(200),
				word2vec.Model(word2vec.SkipGram),
				word2vec.Optimizer(word2vec.NegativeSampling),
				word2vec.NegativeSampleSize(5),
				word2vec.Verbose(),
			)
			f, fe := os.Open(".//src//word2vec//data")
			if fe != nil {
				fmt.Println(fe)
				return
			}
			defer f.Close()

			fmt.Println("学习", model.Train(f))
			modelFile, _ := os.OpenFile("..//src//word2vec//data.bin", os.O_CREATE, 0666)
			fmt.Println(model.Save(modelFile, vector.Single))
			modelFile.Close()
	*/
	modelFile2, modelFile2Err := os.OpenFile(".//src//word2vec//data.bin", os.O_RDONLY, 0666)
	if modelFile2Err != nil {
		fmt.Println(modelFile2Err)
		return
	}

	embed, err := embedding.Load(modelFile2)
	modelFile2.Close()
	fmt.Println(embed, err)

	fmt.Println(embed.Find("不知道"))
	/*	searcher, err := search.New(embed...)
		fmt.Println()
		neighbors, _ := searcher.SearchInternal("好", 3)
		neighbors.Describe()*/

}
