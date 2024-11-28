import { createRouter, createWebHistory } from 'vue-router'  
import Home from '../views/Home.vue'  
import Terraform from '../views/Terraform.vue'  

const routes = [  
  { path: '/', name: 'Home', component: Home },  
  { path: '/terraform', name: 'Terraform', component: Terraform },  
]  

const router = createRouter({  
  history: createWebHistory(),  
  routes,  
})  

export default router  