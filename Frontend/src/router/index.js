import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";
import UserPage from "../views/userpage.vue";
import SignupView from "../views/SignupView.vue";
import ResumePage from "../views/Resumepage.vue";

const routes =[
  {
    path:'/',
   name:'UserPage',
   component: UserPage
 },

 {
  path: '/signup',
  name: 'SignupView',
  component: SignupView,
},

  {
    path:'/login',
    name:'HomeView',
    component:HomeView
  },

  {
    path:'/resume',
    name:'ResumePage',
    component:ResumePage
  },

];

const router = createRouter({
  history: createWebHistory(),
  routes,
});


export default router;

// import { createRouter, createWebHistory } from 'vue-router'
// import HomeView from '../views/HomeView.vue'

// const router = createRouter({
//   history: createWebHistory(import.meta.env.BASE_URL),
//   routes: [
//     {
//       path: '/',
//       name: 'home',
//       component: HomeView,
//     },
//   ],
// })


// export default router

