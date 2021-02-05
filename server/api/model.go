package api

type WordCount struct {
	Word  string `json:"word"`
	Count int    `json:"count"`
}

type OAuthAccessResponse struct {
	AccessToken string `json:"access_token"`
}

type Repos struct {
	ID            int    `json:"id"`
	BranchesURL   string `json:"branches_url"`
	DefaultBranch string `json:"default_branch"`
	FullName      string `json:"full_name"`
	GitRefsURL    string `json:"git_refs_url"`
	GitURL        string `json:"git_url"`
	Name          string `json:"name"`
	NodeID        string `json:"node_id"`
	Owner         Owner  `json:"owner"`
	Private       bool   `json:"private"`
	SSHURL        string `json:"ssh_url"`
	URL           string `json:"url"`
}

type Owner struct {
	ID               int    `json:"id"`
	Login            string `json:"login"`
	NodeID           string `json:"node_id"`
	AvatarURL        string `json:"avatar_url"`
	GravatarID       string `json:"gravatar_id"`
	Name             string `json:"name"`
	OrganizationsURL string `json:"organizations_url"`
	ReposURL         string `json:"repos_url"`
}

type Branch struct {
	Name      string `json:"name"`
	Commit    Commit `json:"commit"`
	Protected bool   `json:"protected"`
	// ProtectionUrl string `json:"protectionUrl"`
}

type Commit struct {
	Sha string `json:"sha"`
	URL string `json:"url"`
}

type NewBranchReq struct {
	Ref   string `json:"ref"`
	Sha   string `json:"sha"`
	Repos Repos  `json:"repos"`
}

type NewBranch struct {
	Ref string `json:"ref"`
	Sha string `json:"sha"`
}

type NewBranchResponse struct {
	Ref    string `json:"ref"`
	URL    string `json:"url"`
	Object Object `json:"object"`
}

type Object struct {
	TypesOf string `json:"type"`
	Sha     string `json:"sha"`
	URL     string `json:"url"`
}

type CreatePullTempRequest struct {
	Head  string `json:"head"`
	Base  string `json:"base"`
	Repos Repos  `json:"repos"`
}

type CreatePullRequest struct {
	Head  string `json:"head"`
	Base  string `json:"base"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

//TODO:remove struct -CreatePullResponse
type CreatePullResponse struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	URL   string `json:"url"`
}

type GetCommitResponse struct {
	Sha  string `json:"sha"`
	URL  string `json:"url"`
	Tree Tree   `json:"tree"`
}

type Tree struct {
	Sha     string `json:"sha,omitempty"`
	URL     string `json:"url,omitempty"`
	Path    string `json:"path,omitempty"`
	Mode    string `json:"mode,omitempty"`
	TypeOf  string `json:"type,omitempty"`
	Size    int    `json:"size,omitempty"`
	Content string `json:"content,omitempty"`
}

type GetTreeResponse struct {
	Sha       string `json:"sha"`
	URL       string `json:"url"`
	Tree      []Tree `json:"tree"`
	Truncated bool   `json:"truncated"`
}

type GetBlobResponse struct {
	Content  string `json:"content"`
	Encoding string `json:"encoding"`
	URL      string `json:"url"`
	Sha      string `json:"sha"`
	Size     int    `json:"size"`
}

//TODO:- createBlobRequest struct to be made
type CreateBlobResponse struct {
	URL string `json:"url"`
	Sha string `json:"sha"`
}

type CreateTreeRequest struct {
	ArrayOfTree []Tree `json:"tree"`
}

type CreateCommitRequest struct {
	Parents []string `json:"parents"`
	Message string   `json:"message"`
	TreeSha string   `json:"tree"`
}

//TODO:- updateReferenceRequest struct to be made
type UpdateReferenceRequest struct {
	Sha string `json:"sha"`
}
