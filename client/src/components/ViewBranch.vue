<template>
  <div>
    <section class="forms">
      <div class="container">
        <div class="card p-3">
          <div class="row criterion-question mb-3 pb-3 text-center text-lg-left">
            <div class="col-md-12">
              <div class="row grid-wrapper-xl rounded bg-lt-grey m-0 mb-3">
                <div class="col-lg-6">
                  <div class="info-wrapper" v-for="(file,index) in treeData.tree" :key="index">
                    <div class="info-data" @click="getBlob(file)">{{ file.path }}</div>
                  </div>
                </div>

                <div class="col-md-6">
                  <p>Text Editor</p>
                  <p>Selected File - <strong>{{selectedBlob.path}}</strong></p>
                  <vue-editor v-model="content"></vue-editor>
                   <div class>
                    <button @click="createBlob()">Update</button>
                  </div>
                </div>
              </div>
            </div>
           
          </div>
          <div class="row">
           
            <div class>
              <button v-b-modal.confirmModal>Commit</button>
            </div>
             <div v-if="isAllowedToRaiseMergeReq">
              <button v-b-modal.raiseMergeReqModal>Raise Merge Request</button>
            </div>
          </div>
        </div>

       
      </div>
    </section>
    <b-modal
      id="confirmModal"
      title="Commit Changes"
      modal-class="custom-modal"
      no-close-on-esc
      no-close-on-backdrop
      size="lg"
      ok-title=" Confirm"
      @ok="commitChanges()"
      cancel-variant="secondary btn-rounded"
      ok-variant="success btn-rounded"
      footer-class="border-0 justify-content-center"
    >
      <div class="course-option-wrapper">
        
        <div class="row course-row">
          <div class="col-md-12">
            <label for>Message</label>
            <input type="text" class="form-control mb-2" v-model="message" name="message" />
          </div>
        </div>
      </div>
    </b-modal>
    <b-modal
      id="raiseMergeReqModal"
      title="Raise Merge Request"
      modal-class="custom-modal"
      no-close-on-esc
      no-close-on-backdrop
      size="lg"
      ok-title="Ok"
      @ok="raiseMergeReq()"
      cancel-variant="secondary btn-rounded"
      ok-variant="success btn-rounded"
      footer-class="justify-content-center pb-4"
    >
      <div class="course-option-wrapper">
        <!-- <h6 class="mt-4 mb-3">Raise Merge Request</h6> -->
        <div class="row course-row">
          <div class="col-md-12">
            <label for>Source Branch</label>
            <multiselect
              v-model="sourceBranch"
              :options="branchList"
              :searchable="false"
              :close-on-select="true"
              :show-labels="false"
              class="multiselect-white"
            />
          </div>
        </div>
        <div class="row course-row">
          <div class="col-md-12">
            <label for>Target Branch</label>
            
            <multiselect
              v-model="targetBranch"
              :options="branchList"
              :searchable="true"
              :close-on-select="true"
              :show-labels="true"
              open-direction="bottom"
              label="name"
              track-by="name"
              class="multiselect-white"
            />
          </div>
        </div>
        <div class="row course-row">
          <div class="col-md-12">
            <label for>Title</label>
            <input type="text" class="form-control mb-2" v-model="title" name="title" />
          </div>
        </div>
        
      </div>
    </b-modal>
  </div>
</template>
<script>
var base64 = require('base-64');
import axios from 'axios'
// import Multiselect from 'vue-multiselect'
import { VueEditor } from "vue2-editor";
export default {
  data() {
    return {
      message: "",
      content: "<h1>Some initial content</h1>",
      isAllowedToRaiseMergeReq: false,
      token:this.$store.getters.getToken,
      treeData:"",
      selectedBlob:"",
      newTreeSha:"",
      newCommitSha:"",
      sourceBranch:"",
      targetBranch:"",
      title:"",
      branchList:this.$store.getters.getBranches,
      updatedBlobMap:{}
    };
  },
  components: {
    VueEditor
  },
  computed: {
    getSelectedRepo() {
      return this.$store.getters.getSelectedRepo
    },
     getCreatedBranch() {
      return this.$store.getters.getCreatedBranch
    }

  },
  methods: {
    getCommit() {
      let sha =""
      if (this.getCreatedBranch && this.getCreatedBranch.object) {
        sha = this.getCreatedBranch.object.sha
      } else {
        sha = this.getCreatedBranch.commit.sha
      }
      axios
        .get('/getCommit/'+this.getSelectedRepo.owner.login+'/'+this.getSelectedRepo.name+'/'+sha,{
          headers: {
            Authorization: this.token 
          }
        })
        .then((response) => {
          if (response.status === 200) {
            // this.repoList = response.data
            this.getTree(response.data.tree.sha)
          }
        })
        .catch((err) => {
          console.error(err)
        })
    },
    getTree(treeSha) {
      axios
        .get('/getTree/'+this.getSelectedRepo.owner.login+'/'+this.getSelectedRepo.name+'/'+treeSha,{
          headers: {
            Authorization: this.token 
          }
        })
        .then((response) => {
          if (response.status === 200) {
            this.treeData = response.data
          }
        })
        .catch((err) => {
          console.error(err)
        })
    },
    getBlob(file) {
      this.selectedBlob = file
      axios
        .get('/getBlob/'+this.getSelectedRepo.owner.login+'/'+this.getSelectedRepo.name+'/'+file.sha,{
          headers: {
            Authorization: this.token 
          }
        })
        .then((response) => {
          if (response.status === 200) {
            this.content = base64.decode(response.data.content);
          }
        })
        .catch((err) => {
          console.error(err)
        })
    },
    commitBranch() {
      this.isAllowedToRaiseMergeReq = true;
    },
    raiseMergeReq() {
      let data = {
        head: this.sourceBranch,  //source or feature branch
        base: this.targetBranch.name, //destination or master
        title: this.title
      }
      axios
        .post('/createPullRequest/'+this.getSelectedRepo.owner.login+'/'+this.getSelectedRepo.name,data,{
          headers: {
            Authorization: this.token 
          }
        })
        .then((response) => {
          if (response.status === 200) {
            alert("success")
          }
        })
        .catch((err) => {
          console.error(err)
        })
    },
    
    createBlob(){
      let data = {
        // content: base64.encode(this.content)
        content: this.content
      }
      let self = this
      axios
        .post('/createBlob/'+this.getSelectedRepo.owner.login+'/'+this.getSelectedRepo.name,data,{
          headers: {
            Authorization: this.token 
          }
        })
        .then((response) => {
          if (response.status === 200) {
            let selectedFileName = self.selectedBlob.path
            self.updatedBlobMap[selectedFileName] = response.data.sha
            // this.createTree(response.data.sha)
            console.log(self.updatedBlobMap)
          }
        })
        .catch((err) => {
          console.error(err)
        })
    },
    createTree(){
      return new Promise((resolve,reject)=>{
        let tree = []
        for (const key in this.updatedBlobMap) {
          let blobObj ={
            path: key,
            sha:this.updatedBlobMap[key],
            mode: "100644",
            type: 'blob'
          }
          tree.push(blobObj)
        }

        let data = {
          tree:tree
          
        }
        axios
          .post('/createTree/'+this.getSelectedRepo.owner.login+'/'+this.getSelectedRepo.name,data,{
            headers: {
              Authorization: this.token 
            }
          })
          .then((response) => {
            if (response.status === 200) {
              // this.newTreeSha = response.data.sha
              this.updateBlobMap = {}
              resolve(response.data.sha)
            }
          })
          .catch((err) => {
            console.error(err)
            reject()
          })

      }) 
    },
    createCommit(newTreeSha){
      return new Promise((resolve,reject)=>{
      
      let parentRef = []

      if (this.getCreatedBranch && this.getCreatedBranch.object) {
        parentRef.push(this.getCreatedBranch.object.sha)
      } else {
        parentRef.push(this.getCreatedBranch.commit.sha)
      }
        let data = {
          message: this.message,
          tree:newTreeSha,
          parents: parentRef
        }
        axios
          .post('/createCommit/'+this.getSelectedRepo.owner.login+'/'+this.getSelectedRepo.name,data,{
            headers: {
              Authorization: this.token 
            }
          })
          .then((response) => {
            if (response.status === 200) {
              // this.newCommitSha = response.data.sha
              resolve(response.data.sha)
              // this.updateReference(response.data.sha)
            }
          })
          .catch((err) => {
            console.error(err)
            reject()
          })
      }) 
    },
    updateReference(newCommitSha){
      return new Promise((resolve,reject)=>{
        let ref = ""
        if (this.getCreatedBranch && this.getCreatedBranch.object) {
          ref = this.getCreatedBranch.ref.split("/")[2]
        // ref = this.getCreatedBranch.ref
      } else {
        // ref = "refs/heads/"+this.getCreatedBranch.name
      
        ref = this.getCreatedBranch.name
      }
        // let ref = this.getCreatedBranch.ref
        let data = {
          sha: newCommitSha
        }
        axios
          .post('/updateReference/'+this.getSelectedRepo.owner.login+'/'+this.getSelectedRepo.name+'/'+ref,data,{
            headers: {
              Authorization: this.token 
            }
          })
          .then((response) => {
            if (response.status === 200) {
              this.isAllowedToRaiseMergeReq = true;
              let tempRef = ""
              if (this.getCreatedBranch && this.getCreatedBranch.object) {
                tempRef = this.getCreatedBranch.ref.split("/")
                this.sourceBranch = tempRef[2]
              } else {
                tempRef = this.getCreatedBranch.name
                this.sourceBranch = tempRef
              }
              // let tempRef = this.getCreatedBranch.ref.split("/")
              // this.sourceBranch = tempRef[2]
              resolve()
            }
          })
          .catch((err) => {
            console.error(err)
            reject()
          })
      })
    },
    commitChanges(){
      this.createTree().then(newTreeSha=>{
        this.createCommit(newTreeSha).then(newCommitSha=>{
          this.updateReference(newCommitSha)
        })
      }).catch(()=>{
        alert("something went wrong")
      })
    },
  },
  created() {
  },
  mounted() {
     this.getCommit()
  }
};
</script>


<style>


/* table,
th,
td {
  border: 1px solid black;
} */
* {
  font-family: "Lato", "Avenir", sans-serif;
}
</style>
