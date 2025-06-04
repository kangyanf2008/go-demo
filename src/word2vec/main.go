package main

import (
	"fmt"
	"github.com/ynqa/wego/pkg/embedding"
	"github.com/ynqa/wego/pkg/model/modelutil/vector"
	"github.com/ynqa/wego/pkg/model/word2vec"
	"github.com/ynqa/wego/pkg/search"
	"os"
)

func main() {

	createModelFile()
	return
	modelFile2, modelFile2Err := os.OpenFile(".//src//word2vec//data.bin", os.O_RDONLY, 0666)
	if modelFile2Err != nil {
		fmt.Println(modelFile2Err)
		return
	}
	embed, err := embedding.Load(modelFile2)
	if err != nil {
		fmt.Println("embedding.load", err)
	}
	modelFile2.Close()
	//fmt.Println(embed, err)

	/*	r, e := embed.Find("不知道")
		fmt.Println(r, e)*/

	fmt.Println(embed.Find("多少"))

	searcher, err := search.New(embed...)
	fmt.Println()
	/*
		neighbors, _ := searcher.SearchInternal("多少", 1)
		fmt.Println(neighbors)
	*/
	e, _ := embed.Find("这个多少钱。应该是10元")
	fmt.Println(searcher.SearchVector(e.Vector, 1))
}

func createModelFile() error {
	model, err := word2vec.New(
		word2vec.Window(200),
		word2vec.Model(word2vec.SkipGram),
		word2vec.Optimizer(word2vec.NegativeSampling),
		word2vec.NegativeSampleSize(5),
		word2vec.Verbose(),
	)
	if err != nil {
		println(err)
		return err
	}

	f, fe := os.Open(".//src//word2vec//data")
	if fe != nil {
		fmt.Println(fe)
		return fe
	}
	defer f.Close()

	fmt.Println("学习", model.Train(f))
	modelFile, _ := os.OpenFile("..//src//word2vec//data.bin", os.O_CREATE, 0666)
	fmt.Println(model.Save(modelFile, vector.Single))
	modelFile.Close()

	return nil
}
