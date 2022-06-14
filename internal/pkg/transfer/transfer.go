package transfer

import (
	"encoding/json"
	"log"
)

type Tess struct {
	Info struct {
		Name              string `json:"name"`
		ReservationNumber string `json:"reservationNumber"`
		CenterCode        string `json:"centerCode"`
		ExamDate          string `json:"examDate"`
		ReportDate        string `json:"reportDate"`
	} `json:"info"`
	Koss struct {
		Autonomy struct {
			Distribution string `json:"distribution"`
			Score        int    `json:"score"`
		} `json:"autonomy"`
		Compensation struct {
			Distribution string `json:"distribution"`
			Score        int    `json:"score"`
		} `json:"compensation"`
		Culture struct {
			Distribution string `json:"distribution"`
			Score        int    `json:"score"`
		} `json:"culture"`
		JobInstability struct {
			Distribution string `json:"distribution"`
			Score        int    `json:"score"`
		} `json:"jobInstability"`
		Relationship struct {
			Distribution string `json:"distribution"`
			Score        int    `json:"score"`
		} `json:"relationship"`
		Requirements struct {
			Distribution string `json:"distribution"`
			Score        int    `json:"score"`
		} `json:"requirements"`
		System struct {
			Distribution string `json:"distribution"`
			Score        int    `json:"score"`
		} `json:"system"`
		Total struct {
			Distribution string `json:"distribution"`
			Score        int    `json:"score"`
			Level        string `json:"level"`
			Result       string `json:"result"`
			ByYear       int    `json:"byYear"`
			Details      string `json:"details"`
		} `json:"total"`
	} `json:"KOSS"`
	Phq9 struct {
		Level       string `json:"level"`
		Result      string `json:"result"`
		Score       int    `json:"score"`
		RiskRate    int    `json:"riskRate"`
		Description string `json:"description"`
		Guidance    string `json:"guidance"`
		Advice      string `json:"advice"`
		ByYear      int    `json:"byYear"`
		Details     string `json:"details"`
	} `json:"PHQ9"`
	Gad7 struct {
		Level       string `json:"level"`
		Result      string `json:"result"`
		Score       int    `json:"score"`
		RiskRate    int    `json:"riskRate"`
		Description string `json:"description"`
		Guidance    string `json:"guidance"`
		Advice      string `json:"advice"`
		ByYear      int    `json:"byYear"`
		Details     string `json:"details"`
	} `json:"GAD7"`
	Adnm4 struct {
		Level       string  `json:"level"`
		Result      string  `json:"result"`
		Score       float64 `json:"score"`
		RiskRate    int     `json:"riskRate"`
		Description string  `json:"description"`
		Guidance    string  `json:"guidance"`
		Advice      string  `json:"advice"`
		ByYear      float64 `json:"byYear"`
		Details     string  `json:"details"`
	} `json:"ADNM4"`
	Pcptsd5 struct {
		Level       string `json:"level"`
		Result      string `json:"result"`
		Score       int    `json:"score"`
		RiskRate    int    `json:"riskRate"`
		Description string `json:"description"`
		Guidance    string `json:"guidance"`
		Advice      string `json:"advice"`
		ByYear      int    `json:"byYear"`
		Details     string `json:"details"`
	} `json:"PCPTSD5"`
	Isi struct {
		Level       string  `json:"level"`
		Result      string  `json:"result"`
		Score       float64 `json:"score"`
		RiskRate    int     `json:"riskRate"`
		Description string  `json:"description"`
		Guidance    string  `json:"guidance"`
		Advice      string  `json:"advice"`
		ByYear      float64 `json:"byYear"`
		Details     string  `json:"details"`
	} `json:"ISI"`
	CSS struct {
		Result      string `json:"result"`
		Description string `json:"description"`
		Guidance    string `json:"guidance"`
		Advice      string `json:"advice"`
		Details     string `json:"details"`
	} `json:"CSS"`
}

func Transfer(d []byte) Tess {

	var err error


	var tess Tess
	err = json.Unmarshal(d, &tess)
	if err != nil {
		log.Fatal("Title - Error during Unmarshal(): ", err)
	}

	return tess

}
