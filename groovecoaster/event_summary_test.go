package groovecoaster

import (
	"io/ioutil"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestEventSummary(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	data, err := ioutil.ReadFile("../tests/assets/event_data.json")
	if err != nil {
		t.Error(err)
	}

	httpmock.RegisterResponder(
		"GET",
		scheme+"mypage.groovecoaster.jp/sp/json/event_data.php",
		httpmock.NewBytesResponder(200, data),
	)

	_, err = testClient.EventSummary()
	if err != nil {
		t.Error(err)
	}
}

func TestEventSummary_BadStatus(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		scheme+"mypage.groovecoaster.jp/sp/json/event_data.php",
		httpmock.NewStringResponder(500, ""),
	)

	_, err := testClient.EventSummary()
	if err == nil {
		t.Error(err)
	}
}

func TestEventSummary_EmptyBody(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		scheme+"mypage.groovecoaster.jp/sp/json/event_data.php",
		httpmock.NewStringResponder(200, ""),
	)

	_, err := testClient.EventSummary()
	if err == nil {
		t.Error(err)
	}
}

func TestEventSummary_InvalidJSON(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		scheme+"mypage.groovecoaster.jp/sp/json/event_data.php",
		httpmock.NewStringResponder(200, `{"test": "test"}`),
	)

	_, err := testClient.EventSummary()
	if err == nil {
		t.Error(err)
	}
}