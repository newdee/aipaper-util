package common

import "github.com/newdee/aipaper-util/config"

// WordsPackage 字数套餐包
type WordsPackage struct {
	PackageId          int        `json:"package_id"`
	PackageName        string     `json:"package_name"`
	PackageDescription string     `json:"package_description"`
	PackageProductId   string     `json:"package_product_id"`
	GiftList           []GiftData `json:"giftList"`
	OriginalPrice      float64    `json:"original_price"` // 产品原价，单位元，需要业务根据产品ID从product_info获取
	Price              float64    `json:"price"`          // 产品售价，单位元，需要业务根据产品ID从product_info获取
}

type GiftData struct {
	GiftName      string `json:"gift_name"`
	GiftProductId string `json:"gift_product_id"`
	Quantity      int    `json:"quantity"`
}

func GetWordsPackageList() ([]WordsPackage, error) {
	cfg, err := config.Get(config.Common)
	if err != nil {
		return nil, err
	}
	var wordsPackage []WordsPackage
	err = cfg.GetWithUnmarshal("words_package", &wordsPackage, &config.JSONUnmarshaler{})
	if err != nil {
		return nil, err
	}
	return wordsPackage, nil
}

// ChatallPackage Chatall套餐包
type ChatallPackage struct {
	PackageId          int         `json:"package_id"`
	PackageName        string      `json:"package_name"`
	PackageDescription string      `json:"package_description"`
	PackageProductId   string      `json:"package_product_id"`
	ModelList          []ModelData `json:"model_list"`
	OriginalPrice      float64     `json:"original_price"` // 产品原价，单位元，需要业务根据产品ID从product_info获取
	Price              float64     `json:"price"`          // 产品售价，单位元，需要业务根据产品ID从product_info获取
}

type ModelData struct {
	ModelName        string `json:"model_name"`
	ModelDescription string `json:"model_description"`
	AvailableTimes   int    `json:"available_times"`
}

func GetChatallPackageList() ([]ChatallPackage, error) {
	cfg, err := config.Get(config.Common)
	if err != nil {
		return nil, err
	}
	var chatallPackage []ChatallPackage
	err = cfg.GetWithUnmarshal("chatall_package", &chatallPackage, &config.JSONUnmarshaler{})
	if err != nil {
		return nil, err
	}
	return chatallPackage, nil
}
