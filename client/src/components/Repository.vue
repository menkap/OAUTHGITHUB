<template>
<div>
  <div>
    <ul>
      <li v-if="!token"><a href="https://github.com/login/oauth/authorize?client_id=8b18b2678e8888094a3b&scope=repo write:repo_hook admin:repo_hook public_repo gist&redirect_uri=http://localhost:4887/oauth/redirect">Login with github</a></li>
    </ul>
  </div>
  <section class="forms" v-if="token">
    <div class="container">
    <div class="card p-3">
      <div>
         <div class="row">
              <label class="">
                Select Repository
                <span style="color:red;">*</span>
              </label>
              <multiselect
                v-model="selectedRepo"
                :options="repoList"
                :searchable="true"
                :close-on-select="true"
                open-direction="bottom"
                label="name"
                track-by="name"
                @select="getBranchesList"
                name="repo"
              />
            </div><br/>
            <!-- <div class="row"> -->
            <div class="row" v-if="!newBranchName">
              <label>
                Select Branch
                <span style="color:red;">*</span>
              </label>
              <multiselect
                v-model="selectedBranch"
                :options="branchList"
                :searchable="true"
                :close-on-select="true"
                open-direction="bottom"
                label="name"
                track-by="name"
                @select="ViewBranch"
                name="branch"
              />
            </div>
            <div class="row" v-else>
              <label>
                Selected Branch
                <span style="color:red;">*</span>
              </label>
             <input
                type="text"
                class="form-control mb-2"
                v-model="newBranchName"
                name="newBranchName"
                disabled
              >
            </div><br/>

            <button
              v-if="selectedRepo"
              id="newBranchNameBtn"
              class="btn btn-success"
              v-b-modal.createBranchModal>
              Create Branch
            </button>
      </div>
    </div>
    </div>
    </section>
    <b-modal
      id="createBranchModal"
      modal-class="custom-modal"
      no-close-on-esc
      no-close-on-backdrop
      size="lg"
      @ok="createBranch()"
      cancel-variant="secondary btn-rounded"
      ok-variant="success btn-rounded"
      footer-class="justify-content-center pb-4"
      title="Create Branch"
      >
      <div class="course-option-wrapper">
        <!-- <h6 class="mt-4 mb-3">
          Create Branch
        </h6> -->
        <div class="row course-row">
          <div class="col-md-12">
            <label for>Base Branch</label>
            <multiselect
              v-model="selectedBaseBranch"
              :options="branchList"
              :searchable="true"
              :close-on-select="true"
              :show-labels="true"
              label="name"
              track-by="name"
              class="multiselect-white"
            />
          </div>
        </div>
        <div class="row course-row">
          <div class="col-md-12">
            <label for>New Branch Name</label>
            <input
                type="text"
                class="form-control mb-2"
                v-model="newBranchName"
                name="newBranchName"
              >
          </div>
        </div>
       
      </div>
    </b-modal>
</div>
</template>
<script>
import axios from 'axios'
// import Multiselect from 'vue-multiselect'
export default {
  data() {
    return {
    //  query :window.location.search.substring(1),
    token :window.location.search.substring(1).split('access_token=')[1],
    repoList: [],
    branchList:[],
    selectedRepo:"",
    newBranchName:"",
    selectedBaseBranch:"",
    selectedBranch:"",
    isBranchCreated:false,
    }
  },
  // components: {Multiselect},
  computed: {

  },
  methods: {
    getRepoList() {
      axios
        .get('getRepoList',{
          headers: {
            Authorization: this.token 
          }
        })
        .then((response) => {
          if (response.status === 200) {
            this.repoList = response.data
          }
        })
        .catch((err) => {
          console.error(err)
        })
    },
    ViewBranch(selectedBranch) {
      this.$store.commit("setCreatedBranch",selectedBranch)
      this.$router.push("/ViewBranch")
    },
    createBranch(){
      let data = {
        sha: this.selectedBaseBranch.commit.sha,
        ref: "refs/heads/" + this.newBranchName,
        repos: this.selectedRepo
      }
      axios
        .post('/createBranch',data,{
          headers: {
            Authorization: this.token 
          }
        })
        .then((response) => {
          if (response.status === 200) {
            this.$store.commit("setCreatedBranch",response.data)
            this.selectedBranch = response.data
            this.isBranchCreated = true
            this.$store.commit("setIsBranchCreated",true)
            this.$router.push("/ViewBranch")

          }
        })
        .catch((err) => {
          console.error(err)
        })
    },
    
   
    getBranchesList(repoObj) {
      this.selectedRepo = repoObj
      this.$store.commit("setSelectedRepo",repoObj)
      let data = {
        name: repoObj.name,
        owner: repoObj.owner
      }
      axios
        .post('getBranchListByRepoName',data,{
          headers: {
            Authorization: this.token 
          }
        })
        .then((response) => {
          if (response.status === 200) {
            this.branchList = response.data
            this.$store.commit("setBranches",response.data)
          }
        })
        .catch((err) => {
          console.error(err)
        })
    },
  },
  created() {
    this.getRepoList()
  },
  mounted() {
    this.$store.commit("setToken",this.token)
  },
}
</script>


<style>

/* table, th, td {
  border: 1px solid black;
} */
* {
  font-family: 'Lato', 'Avenir', sans-serif;
}
</style>
