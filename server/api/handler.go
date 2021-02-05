package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

//Init - init the Event Router
func Init(e *echo.Echo) {
	e.GET("/oauth/redirect", generateAccessTokenHandler)
	e.GET("/getRepoList", getRepoListHandler)
	e.POST("/getBranchListByRepoName", getBranchListByRepoNameHandler)
	e.POST("/createBranch", createBranchHandler)
	e.GET("/getCommit/:owner/:repo/:commitSha", getCommitHandler)
	e.GET("/getTree/:owner/:repo/:treeSha", getTreeHandler)
	e.GET("/getBlob/:owner/:repo/:fileSha", getBlobHandler)
	e.POST("/createBlob/:owner/:repo", createBlobHandler)
	e.POST("/createTree/:owner/:repo", createTreeHandler)
	e.POST("/createCommit/:owner/:repo", createCommitHandler)
	e.POST("/updateReference/:owner/:repo/:ref", updateReferenceHandler)
	e.POST("/createPullRequest/:owner/:repo", createPullRequestHandler)

}

func getCommitHandler(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	owner := c.Param("owner")
	repo := c.Param("repo")
	commitSha := c.Param("commitSha")
	res, err := getCommitService(token, owner, repo, commitSha)
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusExpectationFailed, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}
func getTreeHandler(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	owner := c.Param("owner")
	repo := c.Param("repo")
	treeSha := c.Param("treeSha")

	res, err := getTreeService(token, owner, repo, treeSha)
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusExpectationFailed, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}
func getBlobHandler(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	owner := c.Param("owner")
	repo := c.Param("repo")
	fileSha := c.Param("fileSha")

	res, err := getBlobService(token, owner, repo, fileSha)
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusExpectationFailed, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}
func createBlobHandler(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	owner := c.Param("owner")
	repo := c.Param("repo")
	reqObj := GetBlobResponse{}
	err := c.Bind(&reqObj)
	if err != nil {
		log.Print("createBlob Bind : ", err)
		return c.JSON(http.StatusExpectationFailed, "USER_PARAMETER_BIND_ERROR")
	}
	res, err := createBlobService(token, owner, repo, reqObj)
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusExpectationFailed, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}
func createTreeHandler(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	owner := c.Param("owner")
	repo := c.Param("repo")
	req := CreateTreeRequest{}
	err := c.Bind(&req)
	if err != nil {
		log.Print("createTree Bind : ", err)
		return c.JSON(http.StatusExpectationFailed, "USER_PARAMETER_BIND_ERROR")
	}
	log.Println(req)
	res, err := createTreeService(token, owner, repo, req)
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusExpectationFailed, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}
func createCommitHandler(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	owner := c.Param("owner")
	repo := c.Param("repo")
	reqObj := CreateCommitRequest{}
	err := c.Bind(&reqObj)
	if err != nil {
		log.Print("createCommit Bind : ", err)
		return c.JSON(http.StatusExpectationFailed, "USER_PARAMETER_BIND_ERROR")
	}
	res, err := createCommitService(token, owner, repo, reqObj)
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusExpectationFailed, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}
func updateReferenceHandler(c echo.Context) error {
	fmt.Println("updateReferenceHandler called")
	token := c.Request().Header.Get("Authorization")
	owner := c.Param("owner")
	repo := c.Param("repo")
	ref := c.Param("ref")
	reqObj := UpdateReferenceRequest{}
	err := c.Bind(&reqObj)
	if err != nil {
		log.Print("updateReference Bind : ", err)
		return c.JSON(http.StatusExpectationFailed, "USER_PARAMETER_BIND_ERROR")
	}
	res, err := updateReferenceService(token, owner, repo, ref, reqObj)
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusExpectationFailed, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}
func createPullRequestHandler(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	owner := c.Param("owner")
	repo := c.Param("repo")
	reqObj := CreatePullRequest{}
	err := c.Bind(&reqObj)
	if err != nil {
		log.Print("createPullRequest Bind : ", err)
		return c.JSON(http.StatusExpectationFailed, "USER_PARAMETER_BIND_ERROR")
	}
	res, err := createPullRequestService(token, owner, repo, reqObj)
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusExpectationFailed, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

func createBranchHandler(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	branchObj := NewBranchReq{}
	err := c.Bind(&branchObj)
	if err != nil {
		log.Print("createBranch Bind : ", err)
		return c.JSON(http.StatusExpectationFailed, "USER_PARAMETER_BIND_ERROR")
	}
	res, err := createNewBranchService(token, branchObj)
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusExpectationFailed, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

func getBranchListByRepoNameHandler(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	repoObj := Repos{}
	err := c.Bind(&repoObj)
	if err != nil {
		log.Print("getBranchListByRepoName Bind : ", err)
		return c.JSON(http.StatusExpectationFailed, "USER_PARAMETER_BIND_ERROR")
	}
	res, err := getBranchListByRepoNameService(token, repoObj)
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusExpectationFailed, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

func getRepoListHandler(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	res, err := getRepoListService(token)
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusExpectationFailed, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

func generateAccessTokenHandler(c echo.Context) error {
	err := c.Request().ParseForm()
	if err != nil {
		log.Print("could not parse query: %v", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	code := c.Request().FormValue("code")
	reqURL := fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s", clientID, clientSecret, code)
	res, err := CallExternalAPI("POST", reqURL, "")
	if err != nil {
		log.Print("Error while calling CallExternalAPI", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	// Parse the request body into the `OAuthAccessResponse` struct
	var t OAuthAccessResponse
	if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
		log.Print("could not parse JSON response: %v", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	c.Response().Header().Set("Location", "http://localhost:8081/Repository?access_token="+t.AccessToken)
	return c.JSON(http.StatusFound, t.AccessToken)
}

func CallExternalAPI(method, reqURL, jsonString string) (res *http.Response, err error) {
	httpClient := http.Client{}
	req, err := http.NewRequest(method, reqURL, strings.NewReader(jsonString))
	if err != nil {
		log.Print("could not create HTTP request: %v", err)
		return res, err
	}
	req.Header.Set("accept", "application/json")

	// Send out the HTTP request
	res, err = httpClient.Do(req)
	if err != nil {
		log.Print("could not send HTTP request: %v", err)
		return res, err
	}
	// defer res.Body.Close()
	return res, err
}
