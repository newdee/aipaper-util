package common

import (
	"fmt"
	"github.com/newdee/aipaper-util/config"
)

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	IsSupported bool    `json:"is_supported"`
}

type Language struct {
	Language string `json:"language"`
	Value    string `json:"value"`
}

// Category 描述信息结构体
type Category struct {
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	TotalPrice    float64 `json:"total_price"`
	OriginalPrice float64 `json:"original_price"`
	MinWordNum    int     `json:"min_word_num"`
	MaxWordNum    int     `json:"max_word_num"`
}

// Subcategory 子分类结构体
type Subcategory struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// Subject 学科信息结构体
type Subject struct {
	Code          string        `json:"code"`
	Name          string        `json:"name"`
	Subcategories []Subcategory `json:"subcategories"`
}

// Feature 功能特性结构体
type Feature struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	URL  string `json:"url"`
}

// HomeData 响应数据结构体
type HomeData struct {
	Categories []Category `json:"descriptions"`
	Subjects   []Subject  `json:"subjects"`
	Features   []Feature  `json:"features"`
}

// RechargeInfo 充值积分配置
type RechargeInfo struct {
	RechargeList []struct {
		Index       int     `json:"index"`
		Description string  `json:"description"`
		Price       float64 `json:"price,omitempty"`
		GiftPoints  float64 `json:"gift_points,omitempty"`
	} `json:"recharge_list"`
}

func GetFeatureList() ([]Feature, error) {
	cfg, err := config.Get(config.Common)
	if err != nil {
		return nil, err
	}
	var featureList []Feature
	err = cfg.GetWithUnmarshal("feature", &featureList, &config.JSONUnmarshaler{})
	if err != nil {
		return nil, err
	}
	return featureList, nil
}

func GetSubjectList() ([]Subject, error) {
	cfg, err := config.Get(config.Common)
	if err != nil {
		return nil, err
	}
	var subjectList []Subject
	err = cfg.GetWithUnmarshal("subject", &subjectList, &config.JSONUnmarshaler{})
	if err != nil {
		return nil, err
	}
	return subjectList, nil
}

func GetCategoryProductList() ([]Category, error) {
	cfg, err := config.Get(config.Common)
	if err != nil {
		return nil, err
	}
	var categoryList []Category
	err = cfg.GetWithUnmarshal("category_product", &categoryList, &config.JSONUnmarshaler{})
	if err != nil {
		return nil, err
	}
	return categoryList, nil
}

func GetRechargeInfo() (*RechargeInfo, error) {
	cfg, err := config.Get(config.Common)
	if err != nil {
		return nil, err
	}
	var rechargeInfo *RechargeInfo
	err = cfg.GetWithUnmarshal("recharge_info", &rechargeInfo, &config.JSONUnmarshaler{})
	if err != nil {
		return nil, err
	}
	return rechargeInfo, nil
}

func GetCategoryList() ([]Category, error) {
	cfg, err := config.Get(config.Common)
	if err != nil {
		return nil, err
	}
	var categoryList []Category
	err = cfg.GetWithUnmarshal("category", &categoryList, &config.JSONUnmarshaler{})
	if err != nil {
		return nil, err
	}
	return categoryList, nil
}

func GetCategoryByName(name string) (*Category, error) {
	list, err := GetCategoryList()
	if err != nil {
		return nil, fmt.Errorf("get category failed")
	}

	for _, item := range list {
		if item.Name == name {
			return &item, nil
		}
	}

	return nil, fmt.Errorf("category not found")
}

func GetLanguageList() ([]Language, error) {
	cfg, err := config.Get(config.AIPaper)
	if err != nil {
		return nil, err
	}
	var languageList []Language
	err = cfg.GetWithUnmarshal("language", &languageList, &config.JSONUnmarshaler{})
	if err != nil {
		return nil, err
	}
	return languageList, nil
}

func GetProductList() ([]Product, error) {
	cfg, err := config.Get(config.AIPaper)
	if err != nil {
		return nil, err
	}
	var productList []Product
	err = cfg.GetWithUnmarshal("product", &productList, &config.JSONUnmarshaler{})
	if err != nil {
		return nil, err
	}
	return productList, nil
}
