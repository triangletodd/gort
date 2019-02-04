package models

import (
	"strconv"

	"encoding/base64"

	"github.com/mayowa/bjf"
	"github.com/triangletodd/gort/internal/k8s"

	v1 "github.com/triangletodd/gort/pkg/apis/gorturl/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type URL struct {
	Name  string
	Long  string
	Short string
}

func (u URL) Create(long string) (*URL, error) {
	c, err := URL{}.GetCount()
	if err != nil {
		return nil, err
	}

	newID := strconv.Itoa(*c + 100 + 1)
	short := bjf.Encode(newID)
	gortURL := &v1.GortURL{
		ObjectMeta: metav1.ObjectMeta{
			Name: newID,
		},
		Spec: v1.GortURLSpec{
			Long:  base64.StdEncoding.EncodeToString([]byte(long)),
			Short: short,
		},
	}

	k8s := k8s.GetClient()
	_, kerr := k8s.MtnV1().GortURLs("default").Create(gortURL)
	if kerr != nil {
		return nil, kerr
	}

	url, merr := URL{}.GetByName(newID)
	if merr != nil {
		return nil, merr
	}

	return url, nil
}

func (u URL) GetCount() (*int, error) {
	list, err := URL{}.List()
	if err != nil {
		return nil, err
	}

	count := len(list)

	return &count, nil
}

func (u URL) GetByName(name string) (*URL, error) {
	k8s := k8s.GetClient()
	gorturl, err := k8s.MtnV1().GortURLs("default").Get(name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	decodedLong, berr := base64.StdEncoding.DecodeString(gorturl.Spec.Long)
	if berr != nil {
		return nil, berr
	}

	url := URL{
		Name:  gorturl.Name,
		Long:  string(decodedLong),
		Short: gorturl.Spec.Short,
	}

	return &url, nil
}

func (u URL) List() ([]*URL, error) {
	k8s := k8s.GetClient()
	gortUrls, err := k8s.MtnV1().GortURLs("default").List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var urls []*URL

	for _, u := range gortUrls.Items {
		decodedLong, berr := base64.StdEncoding.DecodeString(u.Spec.Long)
		if berr != nil {
			return nil, berr
		}

		localUrl := URL{
			Name:  u.Name,
			Long:  string(decodedLong),
			Short: u.Spec.Short,
		}

		urls = append(urls, &localUrl)
	}

	return urls, nil
}
