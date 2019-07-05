// project project main.go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"
	"unicode"
	_ "unsafe"
)

func ListDir(folder string) []string {
	filePaths := make([]string, 0)
	files, errDir := ioutil.ReadDir(folder)
	if errDir != nil {
		log.Fatal(errDir)
	}
	for _, file := range files {
		if file.IsDir() {
			ListDir(folder + "/" + file.Name())
			continue

		} else {
			strAbsPath, errPath := filepath.Abs(folder + "/" + file.Name()) // 输出绝对路径
			if errPath != nil {
				fmt.Println(errPath)
			}
			filePaths = append(filePaths, strAbsPath)
		}
	}
	return filePaths
}

func WalkDir(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 30)
	suffix = strings.ToUpper(suffix)                                                     //忽略后缀匹配的大小写
	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		if fi.IsDir() { // 忽略目录
			return nil
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, filename)
		}
		return nil
	})
	return files, err
}

func MoveFiles(filePaths []string, destination string) { //}, flag string) {
	for _, v := range filePaths {
		input, err := ioutil.ReadFile(v)
		_, name := filepath.Split(v)
		if err != nil {
			fmt.Println(err)
		}
		//filepath.Ext(v)
		err = ioutil.WriteFile(destination+"/"+name, input, 0644)
		if err != nil {
			fmt.Println("Error creating", destination)
			fmt.Println(err)
		}
	}
	return
}

func ReadFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file_stream, err := ioutil.ReadAll(file)
	return string(file_stream)
}

func CutLine(file_stream string) []string {
	return strings.FieldsFunc(file_stream, split)
}

func split(singleChar rune) bool {
	if unicode.IsLetter(singleChar) ||
		unicode.IsNumber(singleChar) {
		return false
	}
	return true
}

func Load(filePaths []string) map[string][]string {
	wordsMap := make(map[string][]string)
	for _, v := range filePaths {
		wordsMap[v] = CutLine(ReadFile(v))
	}
	return wordsMap
}

func GetFeature(wordStream map[string][]string) (map[string]map[string]float64, map[string]float64) {
	TFMap := make(map[string]map[string]float64) //各单词在各文本中的TF map[文本名]map[单词]TF
	words_sum := 0
	for filePath, WordSlice := range wordStream {
		DocFrequencyMap := make(map[string]float64) //各文章中word的出现次数 map[单词]出现次数
		for _, word := range WordSlice {
			DocFrequencyMap[word] += 1
			words_sum++
		}
		SubTFMap := make(map[string]float64) //内部map再次声明才可以使用 map[单词]TF
		for word, v := range DocFrequencyMap {
			//TFMap[filePath][word] = v / float64(words_sum)
			SubTFMap[word] = v / float64(words_sum)
		}
		TFMap[filePath] = SubTFMap
		DocFrequencyMap = nil //map由make创建，作为引用一旦为空就会被回收
		words_sum = 0         //别忘了清空
	}

	docs_count := len(wordStream)
	IDFMap := make(map[string]float64)       //各单词在文本库中的IDF map[单词]IDF
	LibFrequency := make(map[string]float64) //在库中引用到各单词的文本数量 map[单词]文本数量
	for _, wordsTFMap := range TFMap {
		for word, v := range wordsTFMap {
			if v != 0 {
				LibFrequency[word]++
			}
		}
	}
	for word, v := range LibFrequency {
		IDFMap[word] = math.Log(float64(docs_count)/v + 1)
	}
	LibFrequency = nil

	return TFMap, IDFMap
}

func CalculateTFIDFMap(TFMap map[string]map[string]float64, IDFMap map[string]float64) map[string]map[string]float64 {
	TFIDFMap := make(map[string]map[string]float64) //map[文件]map[单词]tf-idf指标
	for filePath, wordTFMap := range TFMap {
		DocTFIDFMap := make(map[string]float64) //嵌套类型初始化，稍后计算完成赋值  map[单词]tf-idf指标
		for word, tf := range wordTFMap {
			DocTFIDFMap[word] = tf * IDFMap[word]
		}
		TFIDFMap[filePath] = DocTFIDFMap
		DocTFIDFMap = nil //清空计数map，可有可无
	}
	return TFIDFMap
}

func CalculateVectorCos(fileAPath string, fileBPath string, TFIDFMap map[string]map[string]float64) float64 {
	var sum_AB, sum_AA, sum_BB float64
	for word, tfidf := range TFIDFMap[fileAPath] {
		sum_AB += tfidf * TFIDFMap[fileBPath][word]
		sum_AA += math.Pow(tfidf, 2)
		sum_BB += math.Pow(TFIDFMap[fileBPath][word], 2)
	}
	return (sum_AB / ((math.Sqrt(sum_AA)) * (math.Sqrt(sum_BB))))
}

type CosSimilarInfo struct {
	MarkFilePath string
	LibFilePath  string
	CosSimilar   float64
}

func CalculateCosInDirs(AimDirPath string, ReferDirPath string, AimSuffix string, ReferSuffix string) []CosSimilarInfo {
	CosSimilarInfos := make([]CosSimilarInfo, 0)
	WaitFilePaths, _ := WalkDir(AimDirPath, AimSuffix)      //, AimSuffix)
	ReferFilePaths, _ := WalkDir(ReferDirPath, ReferSuffix) //, ReferSuffix)
	WaitTFIDFMap := CalculateTFIDFMap(GetFeature(Load(WaitFilePaths)))
	ReferTFIDFMap := CalculateTFIDFMap(GetFeature(Load(ReferFilePaths)))

	for waitFile, waitVector := range WaitTFIDFMap {
		for referFile, referVector := range ReferTFIDFMap {
			var sum_AB, sum_AA, sum_BB float64
			for word, tfidf := range waitVector {
				sum_AB += tfidf * referVector[word]
				sum_AA += math.Pow(tfidf, 2)
				sum_BB += math.Pow(referVector[word], 2)
			}
			info := CosSimilarInfo{
				MarkFilePath: waitFile,
				LibFilePath:  referFile,
				CosSimilar:   0,
			}
			Denominator := (math.Sqrt(sum_AA)) * (math.Sqrt(sum_BB))
			if Denominator != 0 {
				info.CosSimilar = (sum_AB / Denominator)
			}
			if info.CosSimilar > float64(0.4) {
				CosSimilarInfos = append(CosSimilarInfos, info)
			}
		}
	}
	CosSimilarInfos = SortByCosSimilar(CosSimilarInfos)
	for _, v := range CosSimilarInfos {
		fmt.Println(v)
	}
	return CosSimilarInfos
}

func SortByCosSimilar(infos []CosSimilarInfo) []CosSimilarInfo {
	for i := 0; i < len(infos); i++ {
		for j := 0; j < len(infos)-1; j++ {
			if infos[j].CosSimilar < infos[j+1].CosSimilar {
				infos[j], infos[j+1] = infos[j+1], infos[j]
			}
		}
	}
	return infos
}

func Print(s []string) {
	for _, v := range s {
		fmt.Println(v)
	}
}

/*
func main() {
	fmt.Println("--------------------------")
	fileStream := CutLine(ReadFile("cut.txt"))
	Print(fileStream)

	fmt.Println("--------------------------")
	fileStream = strings.Fields(ReadFile("cut.txt"))
	Print(fileStream)

	fmt.Println("--------------------------")
	filePaths := []string{"test.txt", "test1.txt", "test2.txt", "test3.txt", "test4.txt"}
	fmt.Println(Load(filePaths))

	fmt.Println("--------------------------")
	fmt.Println(GetFeature(Load(filePaths)))

	//fmt.Println("--------------------------")
	//fmt.Println(CalculateTFIDFMap(GetFeature(Load(filePaths))))

	//fmt.Println("--------------------------")
	//fmt.Println(CalculateVectorCos("test3.txt", "test4.txt", CalculateTFIDFMap(GetFeature(Load(filePaths)))))

	fmt.Println("--------------------------")
	files, _ := WalkDir("./", "")
	fmt.Println(files)

	var x = make([]struct{}, 10)
	var y = make(map[struct{}]struct{}, 0)
	fmt.Println(unsafe.Sizeof(x), " ", unsafe.Sizeof(y)) // prints 12 in the playground

	//fmt.Println("--------------------------")
	//files, err = WalkDir("D://project//Zilliqa", "")
	//fmt.Println(files, err)

	// fmt.Println("--------------------------")
	// files, err = WalkDir("D://project//Zilliqa", ".cpp")
	// MoveFiles(files, "D://project//moss.py//submission//base") //, "a")

	// fmt.Println("--------------------------")
	// files, err = WalkDir("D://project//bitcoin", ".cpp")
	// MoveFiles(files, "D://project//moss.py//submission") //, "b")

	// fmt.Println("==========================")
	// dirPath, _ := WalkDir("D://project//Zilliqa", "") //使用双斜杠注意转义字符
	// fmt.Println(CalculateTFIDFMap(GetFeature(Load(dirPath))))

	//fmt.Println("==========================")
	//waitDirPath := "D://project//Zilliqa//src"
	//referDirPath := "D://project//aleth"
	//fmt.Println(CalculateCosInDirs(waitDirPath, referDirPath, ".cpp", ".cpp"))
}
*/
