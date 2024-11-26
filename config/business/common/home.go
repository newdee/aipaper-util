package common

import (
	"github.com/newdee/aipaper-util/config"
)

// Category 描述信息结构体
type Category struct {
	Name        string `json:"name"`
	Description string `json:"description"`
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
