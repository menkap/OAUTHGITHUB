export const state = {
  createdBranch:"",
  selectedRepo:"",
  token:"",
  branchList:[],
  isBranchCreated:false
}

export const getters = {
    getCreatedBranch: state => state.createdBranch,
    getSelectedRepo: state => state.selectedRepo,
    getToken: state => state.token,
    getBranches: state => state.branchList,
    getIsBranchCreated: state => state.isBranchCreated
}

export const mutations = {
  setCreatedBranch(state, createdBranch) {
    state.createdBranch = createdBranch
  },
  setSelectedRepo(state, selectedRepo) {
    state.selectedRepo = selectedRepo
  },
  setToken(state, token) {
    state.token = token
  },
  setBranches(state, branchList) {
    state.branchList = branchList
  },
  setIsBranchCreated(state, isBranchCreated) {
    state.isBranchCreated = isBranchCreated
  },
}

export const actions = {
}
