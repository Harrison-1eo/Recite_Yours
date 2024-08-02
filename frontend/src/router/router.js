import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import Login from '../views/Login.vue'
import Register from '../views/Register.vue'
import UserHome from '../views/UserHome.vue';
import AddWord from '../views/AddWord.vue';
import ShowWords from "../views/ShowWords.vue";

const routes = [
    { path: '/', name: 'Home', component: Home },
    { path: '/login', name: 'Login', component: Login },
    { path: '/register', name: 'Register', component: Register },
    { path: '/user', name: 'UserHome', component: UserHome },
    { path: '/add-word', name: 'AddWord', component: AddWord },
    { path: '/show-words', name: 'ShowWords', component: ShowWords },
    { path: '/:pathMatch(.*)*', redirect: '/' },

]

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes,
})

export default router
