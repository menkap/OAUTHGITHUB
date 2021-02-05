import Vue from 'vue'
import Router from 'vue-router'
import Repository from '@/components/Repository.vue'
import ViewBranch from '@/components/ViewBranch.vue'


Vue.use(Router)
export default new Router({
  routes: [
      {
        path: '/Repository/:access_token',
        name: 'Repository',
        component: Repository,
        meta: { moduleName: 'Repository' }
      },
      {
        path: '/ViewBranch',
        name: 'ViewBranch',
        component: ViewBranch,
        meta: { moduleName: 'ViewBranch' }
      },
  ]
})
