package message

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

// サーガの実施状況の監視と役割の割り振り
// ユーザと会社のTopicにメッセージの送信
const STEP1_URL = "http://172.22.0.2:1000/user_id"
const STEP2_URL = "http://172.22.0.4:1010/company_id"

type SagaStep interface {
	Execute() error
	Undo() error
}
type BaseStep struct{}

func (b *BaseStep) Undo() error {
	fmt.Println("Default Undo called")
	return nil
}

type Step1 struct {
	BaseStep
	UserID string
}

func (s *Step1) Execute() error {
	userId := s.UserID

	queryParams := url.Values{}
	queryParams.Add("userId", userId)
	fullURL := fmt.Sprintf("%s?%s", STEP1_URL, queryParams.Encode())

	resp, err := http.Get(fullURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		return errors.New("invalid userID")
	}
	return nil
}

type Step2 struct {
	BaseStep
	CompnayId string
}

func (s *Step2) Execute() error {
	compnayId := s.CompnayId

	queryParams := url.Values{}
	queryParams.Add("companyId", compnayId)
	fullURL := fmt.Sprintf("%s?%s", STEP2_URL, queryParams.Encode())

	resp, err := http.Get(fullURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("invalid companyID")
	}
	return nil
}

func RunSaga(steps []SagaStep) error {
	for i, step := range steps {
		if err := step.Execute(); err != nil {

			fmt.Printf("Error undoing step: %v\n", err)
			for j := i - 1; j >= 0; j-- {
				if undoErr := steps[j].Undo(); undoErr != nil {
					fmt.Printf("Error undoing step %d: %v\n", j+1, undoErr)
					return undoErr
				}
			}
			return err
		}
	}

	fmt.Println("Saga completed successfully")
	return nil
}
