package test

import (
	"context"
	"fmt"
	"github.com/cucumber/godog"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
)

type apiFeature struct {
	resp *httptest.ResponseRecorder
}

func (a *apiFeature) resetResponse() {
	a.resp = httptest.NewRecorder()
}

func (a *apiFeature) iSendrequestTo(method, endpoint string) error {
	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		return err
	}

	// handle panic
	defer func() {
		switch t := recover().(type) {
		case string:
			err = fmt.Errorf(t)
		case error:
			err = t
		}
	}()

	cli := http.Client{}

	resp, err := cli.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	bodyBytes, readErr := ioutil.ReadAll(resp.Body)

	if readErr != nil {
		return readErr
	}
	_, err = a.resp.Write(bodyBytes)
	if err != nil {
		return err
	}

	a.resp.Code = resp.StatusCode

	return nil
}

//
//func (a *apiFeature) iSendRequestToWithEmail(method, endpoint, param string) error {
//	bodyContent := fmt.Sprintf(`{"email": "%s"}`, param)
//	req, err := http.NewRequest(method, endpoint, strings.NewReader(bodyContent))
//	if err != nil {
//		return err
//	}
//	req.Header.Set("Content-Type", "application/json")
//
//	defer func() {
//		switch t := recover().(type) {
//		case string:
//			err = fmt.Errorf(t)
//		case error:
//			err = t
//		}
//	}()
//
//	cli := http.Client{}
//
//	resp, err := cli.Do(req)
//	if err != nil {
//		return err
//	}
//
//	defer resp.Body.Close()
//
//	bodyBytes, readErr := ioutil.ReadAll(resp.Body)
//
//	if readErr != nil {
//		return readErr
//	}
//	_, err = a.resp.Write(bodyBytes)
//	if err != nil {
//		return err
//	}
//
//	a.resp.Code = resp.StatusCode
//
//	return nil
//}

func (a *apiFeature) theResponseCodeShouldBe(code int) error {
	if code != a.resp.Code {
		return fmt.Errorf("expected response code to be: %d, but actual is: %d", code, a.resp.Code)
	}
	return nil
}

func (a *apiFeature) theResponseShouldMatchJson(arg1 *godog.DocString) (err error) {
	var expected, actual []byte

	expected = []byte(arg1.Content)
	sExpected := strings.TrimSpace(string(expected))

	actual = a.resp.Body.Bytes()
	sActual := strings.TrimSpace(string(actual))

	if sActual != sExpected {
		err = fmt.Errorf("expected json: |%s| does not match actual: |%s|", sExpected, sActual)
	}

	return err
}

func (a *apiFeature) theResponseShouldContain(arg1 *godog.DocString) (err error) {
	expected := arg1.Content
	sExpected := strings.TrimSpace(expected)

	actual := a.resp.Body.Bytes()
	sActual := strings.TrimSpace(string(actual))

	if !strings.Contains(sActual, sExpected) {
		err = fmt.Errorf("actual response: |%s| does not contain expected: |%s|", sActual, sExpected)
	}

	return err
}

//
//func theResponseShouldMatchJson(arg1 *godog.DocString) error {
//	return godog.ErrPending
//}

func InitializeScenario(s *godog.ScenarioContext) {
	api := &apiFeature{}

	s.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		api.resetResponse()
		return ctx, nil
	})

	s.Step(`^I send "(GET|POST)" request to "([^"]*)"$`, api.iSendrequestTo)
	//s.Step(`^I send "(GET|POST)" request to "([^"]*)" with email "([^"]*)"$`, api.iSendRequestToWithEmail)
	s.Step(`^the response code should be (\d+)$`, api.theResponseCodeShouldBe)
	s.Step(`^the response should match json:$`, api.theResponseShouldMatchJson)
	s.Step(`^the response should contain:$`, api.theResponseShouldContain)
}
