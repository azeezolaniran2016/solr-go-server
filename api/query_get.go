package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sendgrid/go-solr"
)

func (a *api) query(e echo.Context) error {
	sResp, err := a.SolrClient.Select([]string{a.SolrServerURL}, solr.Query("*:*"))
	if err != nil {
		a.Log.WithError(err).Error("failed query to solr")
	} else {
		e.JSON(http.StatusOK, sResp.Response.Docs)
	}
	return nil
}
