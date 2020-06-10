package littleaspen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

const (
	API_ROOT = "https://www.littleaspen.com/api"
)

type Client struct {
	apiKey     string
	apiRoot    string
	httpClient *http.Client
}

func New(apiKey string) *Client {
	return &Client{
		apiKey:     apiKey,
		apiRoot:    API_ROOT,
		httpClient: http.DefaultClient,
	}
}

func (c *Client) GetDocuments() ([]Document, error) {
	resp, err := c.get("/documents")
	if err != nil {
		return nil, err
	}

	docs := []Document{}
	if err := c.responseData(resp, "documents", &docs); err != nil {
		return nil, err
	}

	return docs, nil
}

func (c *Client) GetContentVersions(documentSlug string) ([]ContentVersion, error) {
	resp, err := c.get("/documents/" + documentSlug + "/content_versions")
	if err != nil {
		return nil, err
	}

	versions := []ContentVersion{}
	if err := c.responseData(resp, "content_versions", &versions); err != nil {
		return nil, err
	}

	return versions, nil
}

func (c *Client) GetContentVersion(docSlug, versionSlug string) (ContentVersion, error) {
	cv := ContentVersion{}

	resp, err := c.get("/documents/" + docSlug + "/content_versions/" + versionSlug)
	if err != nil {
		return cv, err
	}

	err = c.responseData(resp, "content_version", &cv)
	return cv, err
}

func (c *Client) GetAcceptances(docSlug, versionSlug string) ([]Acceptance, error) {
	accs := []Acceptance{}

	resp, err := c.get("/documents/" + docSlug + "/content_versions/" + versionSlug + "/acceptances")
	if err != nil {
		return accs, err
	}

	err = c.responseData(resp, "acceptances", &accs)
	return accs, err
}

func (c *Client) GetAcceptance(docSlug, versionSlug, externalID string) (Acceptance, error) {
	acc := Acceptance{}

	resp, err := c.get("/documents/" + docSlug + "/content_versions/" + versionSlug + "/acceptances/" + externalID)
	if err != nil {
		return acc, err
	}

	err = c.responseData(resp, "acceptance", &acc)
	return acc, err
}

func (c *Client) CreateAcceptance(docSlug, versionSlug, externalID string, data map[string]interface{}) (Acceptance, error) {
	if data == nil {
		data = map[string]interface{}{
			"user_type": "business",
		}
	}
	acc := Acceptance{
		ExternalID: externalID,
		Metadata:   data,
	}
	type outter struct {
		Acceptance Acceptance `json:"acceptance"`
	}

	bts, _ := json.Marshal(outter{acc})
	fmt.Println(string(bts))
	req, err := http.NewRequest(
		"POST",
		c.fullUrl("/documents/"+docSlug+"/content_versions/"+versionSlug+"/acceptances"),
		bytes.NewBuffer(bts),
	)
	if err != nil {
		return acc, err
	}

	c.prepare(req)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return acc, err
	}

	v, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(v))
	err = json.NewDecoder(resp.Body).Decode(&acc)
	return acc, err
}

func (c *Client) get(path string) (*http.Response, error) {
	req, err := http.NewRequest("GET", c.fullUrl(path), nil)
	if err != nil {
		return nil, err
	}

	c.prepare(req)
	return c.httpClient.Do(req)
}

func (c *Client) responseData(resp *http.Response, key string, dest interface{}) error {
	type outter struct {
		Data map[string]json.RawMessage `json:"data"`
	}

	d := outter{}
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return err
	}

	inner, ok := d.Data[key]
	if !ok {
		return errors.Errorf("missing key in data: %s", key)
	}

	return json.Unmarshal(inner, dest)
}

func (c *Client) prepare(r *http.Request) {
	r.Header.Set("Authorization", "Bearer "+c.apiKey)
}

func (c *Client) fullUrl(path string) string {
	return c.apiRoot + path
}
