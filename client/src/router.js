import Vue from "vue";
import Router from "vue-router";
import login from "./views/login.vue";
import axios from "axios"
import store from "./store.js"

Vue.use(Router);

const router = new Router({
  mode: "history",
  base: process.env.BASE_URL,
  routes: [
    {
      path: "/",
      name: "login",
      component: login
    },
    {
      path: "/home",
      name: "home",
      component: () =>
        import("./components/home.vue"),
      meta: {
        auth: true
      }
    }
  ]
});

router.beforeEach((to,from,next)=>{
  if(to.matched.some(record => record.meta.auth)){
      axios({
        method: 'post',
        url: "http://localhost:1000/auth/check",
        headers: {"X-Access-Token": store.state.token}
      })
      .then(()=>next())
      .catch(err=>{
        if(err.response){
          alert(err.response.data)
          next("/")
          return
        }
        console.log(err)
      })
      
      return
  }
  next()
})
export default router