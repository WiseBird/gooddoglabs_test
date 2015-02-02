package models

type RestResult struct {
    Data interface{} `json:"data"`
    Error string `json:"error"`
}