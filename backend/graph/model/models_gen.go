// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type AssetType struct {
	ID         string                `json:"id"`
	Name       string                `json:"name"`
	Extends    *AssetType            `json:"extends"`
	Attributes []*AssetTypeAttribute `json:"attributes"`
}

type AssetTypeAttribute struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type *Type  `json:"type"`
}

type AssetTypeAttributeInput struct {
	Name        string `json:"name"`
	Type        *Type  `json:"type"`
	AssetTypeID string `json:"AssetTypeId"`
}

type AssetTypeInput struct {
	Name    string  `json:"name"`
	Extends *string `json:"extends"`
}

type Type string

const (
	TypeString Type = "STRING"
	TypeNumber Type = "NUMBER"
)

var AllType = []Type{
	TypeString,
	TypeNumber,
}

func (e Type) IsValid() bool {
	switch e {
	case TypeString, TypeNumber:
		return true
	}
	return false
}

func (e Type) String() string {
	return string(e)
}

func (e *Type) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Type(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Type", str)
	}
	return nil
}

func (e Type) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
