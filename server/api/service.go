package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func getRepoListService(token string) (repoList []Repos, err error) {

	httpClient := http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/user/repos", nil)
	if err != nil {
		log.Print("could not create HTTP request: %v", err)
		return repoList, err
	}
	req.Header.Set("Authorization", "token "+token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	res, err := httpClient.Do(req)
	if err != nil {
		log.Print("could not send HTTP request: %v", err)
		return repoList, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Println("expected status 200 got - ", res.StatusCode)
		return repoList, errors.New("something went wrong")
	}
	if err := json.NewDecoder(res.Body).Decode(&repoList); err != nil {
		log.Print("could not parse JSON response: %v", err)
		return repoList, err
	}
	return repoList, err
}

func getBranchListByRepoNameService(token string, repoObj Repos) (branchList []Branch, err error) {
	httpClient := http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/repos/"+repoObj.Owner.Login+"/"+repoObj.Name+"/branches", nil)
	if err != nil {
		log.Print("could not create HTTP request: %v", err)
		return branchList, err
	}
	req.Header.Set("Authorization", "token "+token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	res, err := httpClient.Do(req)
	if err != nil {
		log.Print("could not send HTTP request: %v", err)
		return branchList, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Println("expected status 200 got - ", res.StatusCode)
		return branchList, errors.New("something went wrong")
	}
	if err := json.NewDecoder(res.Body).Decode(&branchList); err != nil {
		log.Print("could not parse JSON response: %v", err)
		return branchList, err
	}
	return branchList, err
}

func createNewBranchService(token string, NewBranchReqObj NewBranchReq) (branchObj NewBranchResponse, err error) {
	var newObj NewBranch
	newObj.Ref = NewBranchReqObj.Ref
	newObj.Sha = NewBranchReqObj.Sha
	byteData, _ := json.Marshal(newObj)
	httpClient := http.Client{}
	req, err := http.NewRequest("POST", "https://api.github.com/repos/"+NewBranchReqObj.Repos.Owner.Login+"/"+NewBranchReqObj.Repos.Name+"/git/refs", bytes.NewReader(byteData))
	if err != nil {
		log.Print("could not create HTTP request: %v", err)
		return branchObj, err
	}
	req.Header.Set("Authorization", "token "+token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	res, err := httpClient.Do(req)
	if err != nil {
		log.Print("could not send HTTP request: %v", err)
		return branchObj, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusCreated {
		log.Println("expected status 201 got - ", res.StatusCode)
		return branchObj, errors.New("something went wrong")
	}
	if err := json.NewDecoder(res.Body).Decode(&branchObj); err != nil {
		log.Print("could not parse JSON response: %v", err)
		return branchObj, err
	}
	return branchObj, err
}

func getCommitService(token, owner, repo, commitSha string) (obj GetCommitResponse, err error) {
	httpClient := http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/repos/"+owner+"/"+repo+"/git/commits/"+commitSha, nil)
	if err != nil {
		log.Print("could not create HTTP request: %v", err)
		return obj, err
	}
	req.Header.Set("Authorization", "token "+token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	res, err := httpClient.Do(req)
	if err != nil {
		log.Print("could not send HTTP request: %v", err)
		return obj, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Println("expected status 200 got - ", res.StatusCode)
		return obj, errors.New("something went wrong")
	}
	if err := json.NewDecoder(res.Body).Decode(&obj); err != nil {
		log.Print("could not parse JSON response: %v", err)
		return obj, err
	}
	return obj, err
}
func getTreeService(token, owner, repo, treeSha string) (obj GetTreeResponse, err error) {
	httpClient := http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/repos/"+owner+"/"+repo+"/git/trees/"+treeSha, nil)
	if err != nil {
		log.Print("could not create HTTP request: %v", err)
		return obj, err
	}
	req.Header.Set("Authorization", "token "+token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	res, err := httpClient.Do(req)
	if err != nil {
		log.Print("could not send HTTP request: %v", err)
		return obj, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Println("expected status 200 got - ", res.StatusCode)
		return obj, errors.New("something went wrong")
	}
	if err := json.NewDecoder(res.Body).Decode(&obj); err != nil {
		log.Print("could not parse JSON response: %v", err)
		return obj, err
	}
	return obj, err
}
func getBlobService(token, owner, repo, fileSha string) (obj GetBlobResponse, err error) {
	httpClient := http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/repos/"+owner+"/"+repo+"/git/blobs/"+fileSha, nil)
	if err != nil {
		log.Print("could not create HTTP request: %v", err)
		return obj, err
	}
	req.Header.Set("Authorization", "token "+token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	res, err := httpClient.Do(req)
	if err != nil {
		log.Print("could not send HTTP request: %v", err)
		return obj, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Println("expected status 200 got - ", res.StatusCode)
		return obj, errors.New("something went wrong")
	}
	if err := json.NewDecoder(res.Body).Decode(&obj); err != nil {
		log.Print("could not parse JSON response: %v", err)
		return obj, err
	}
	return obj, err
}
func createBlobService(token, owner, repo string, reqObj GetBlobResponse) (obj CreateBlobResponse, err error) {
	byteData, _ := json.Marshal(reqObj)
	httpClient := http.Client{}
	req, err := http.NewRequest("POST", "https://api.github.com/repos/"+owner+"/"+repo+"/git/blobs", bytes.NewReader(byteData))
	if err != nil {
		log.Print("could not create HTTP request: %v", err)
		return obj, err
	}
	req.Header.Set("Authorization", "token "+token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	res, err := httpClient.Do(req)
	if err != nil {
		log.Print("could not send HTTP request: %v", err)
		return obj, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusCreated {
		log.Println("expected status 201 got - ", res.StatusCode)
		return obj, errors.New("something went wrong")
	}
	if err := json.NewDecoder(res.Body).Decode(&obj); err != nil {
		log.Print("could not parse JSON response: %v", err)
		return obj, err
	}
	return obj, err
}
func createTreeService(token, owner, repo string, reqData CreateTreeRequest) (obj GetTreeResponse, err error) {
	byteData, _ := json.Marshal(reqData)
	httpClient := http.Client{}
	req, err := http.NewRequest("POST", "https://api.github.com/repos/"+owner+"/"+repo+"/git/trees", bytes.NewReader(byteData))
	if err != nil {
		log.Print("could not create HTTP request: %v", err)
		return obj, err
	}
	req.Header.Set("Authorization", "token "+token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	res, err := httpClient.Do(req)
	if err != nil {
		log.Print("could not send HTTP request: %v", err)
		return obj, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusCreated {
		log.Print("expected status 201 got", res.StatusCode)
		return obj, errors.New("something went wrong")
	}
	if err := json.NewDecoder(res.Body).Decode(&obj); err != nil {
		log.Print("could not parse JSON response: %v", err)
		return obj, err
	}
	return obj, err
}
func createCommitService(token, owner, repo string, reqObj CreateCommitRequest) (obj GetCommitResponse, err error) {
	byteData, _ := json.Marshal(reqObj)
	httpClient := http.Client{}
	req, err := http.NewRequest("POST", "https://api.github.com/repos/"+owner+"/"+repo+"/git/commits", bytes.NewReader(byteData))
	if err != nil {
		log.Print("could not create HTTP request: %v", err)
		return obj, err
	}
	req.Header.Set("Authorization", "token "+token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	res, err := httpClient.Do(req)
	if err != nil {
		log.Print("could not send HTTP request: %v", err)
		return obj, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusCreated {
		log.Println("expected status 201 got - ", res.StatusCode)
		return obj, errors.New("something went wrong")
	}
	if err := json.NewDecoder(res.Body).Decode(&obj); err != nil {
		log.Print("could not parse JSON response: %v", err)
		return obj, err
	}
	return obj, err
}
func updateReferenceService(token, owner, repo, ref string, reqObj UpdateReferenceRequest) (branchObj NewBranchResponse, err error) {
	byteData, _ := json.Marshal(reqObj)
	httpClient := http.Client{}
	// req, err := http.NewRequest("PATCH", "https://api.github.com/repos/"+owner+"/"+repo+"/git/refs/"+ref, bytes.NewReader(byteData))
	req, err := http.NewRequest("PATCH", "https://api.github.com/repos/"+owner+"/"+repo+"/git/refs/heads/"+ref, bytes.NewReader(byteData))
	if err != nil {
		log.Print("could not create HTTP request: %v", err)
		return branchObj, err
	}
	req.Header.Set("Authorization", "token "+token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Content-Type", "application/json")

	res, err := httpClient.Do(req)
	if err != nil {
		log.Print("could not send HTTP request: %v", err)
		return branchObj, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Println("expected status 200 got - ", res.StatusCode)
		return branchObj, errors.New("something went wrong")
	}
	if err := json.NewDecoder(res.Body).Decode(&branchObj); err != nil {
		log.Print("could not parse JSON response: %v", err)
		return branchObj, err
	}
	return branchObj, err
}
func createPullRequestService(token, owner, repo string, reqObj CreatePullRequest) (resObj CreatePullResponse, err error) {
	byteData, _ := json.Marshal(reqObj)
	httpClient := http.Client{}
	req, err := http.NewRequest("POST", "https://api.github.com/repos/"+owner+"/"+repo+"/pulls", bytes.NewReader(byteData))
	if err != nil {
		log.Print("could not create HTTP request: %v", err)
		return resObj, err
	}
	req.Header.Set("Authorization", "token "+token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	res, err := httpClient.Do(req)
	if err != nil {
		log.Print("could not send HTTP request: %v", err)
		return resObj, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusCreated {
		log.Println("expected status 201 got - ", res.StatusCode)
		return resObj, errors.New("something went wrong")
	}
	if err := json.NewDecoder(res.Body).Decode(&resObj); err != nil {
		log.Print("could not parse JSON response: %v", err)
		return resObj, err
	}
	return resObj, err
}
