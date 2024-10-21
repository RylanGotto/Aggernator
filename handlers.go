package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func GameScoreHandler(c echo.Context) error {
	scores, err := GetGameScores()
	if err != nil {
		errorMesg := map[string]string{
			"Error": err.Error(),
		}
		return c.JSON(http.StatusBadRequest, errorMesg)
	}
	return c.JSON(http.StatusOK, scores)
}

func NewsHandler(c echo.Context) error {
	n := NewsParams{
		Q: "Election 2024",
	}
	news, err := n.GetTopHeadlines()
	if err != nil {
		errorMesg := map[string]string{
			"Error": err.Error(),
		}
		return c.JSON(http.StatusBadRequest, errorMesg)
	}
	return c.JSON(http.StatusOK, news)
}

func RedditHandler(c echo.Context) error {
	reddits, err := GetRedditFeeds()
	if err != nil {
		errorMesg := map[string]string{
			"Error": err.Error(),
		}
		return c.JSON(http.StatusBadRequest, errorMesg)
	}
	return c.JSON(http.StatusOK, reddits)
}
