<template>
    <div>
        <Navbar />
        <div class="container">
            <div class="row justify-content-center mt-5">
                <div class="col-lg-4 col-md-6 col-sm-8">
                    <h2 class="text-center">登录</h2>
                    <form @submit.prevent="handleLogin">
                        <div class="mb-3">
                            <label for="username" class="form-label">用户名</label>
                            <input type="text" class="form-control" id="username" v-model="username" required />
                        </div>
                        <div class="mb-3">
                            <label for="password" class="form-label">密码</label>
                            <input type="password" class="form-control" id="password" v-model="password" required />
                        </div>
                        <div class="mb-3 form-check">
                            <input type="checkbox" class="form-check-input" id="rememberMe" v-model="rememberMe"/>
                            <label class="form-check-label" for="rememberMe">记住我</label>
                        </div>
                        <button type="submit" class="btn btn-primary w-100">登录</button>
                    </form>
                    <div class="text-center mt-3">
                        <router-link to="/register">点击注册</router-link>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import Navbar from '../components/Navbar.vue'
import axios_instance from '../utils/axios_config'
import {API} from "@/utils/axios_config";
import {showSuccessToast} from "@/utils/showMessage";

export default {
    name: 'Login',
    components: {
        Navbar
    },
    data() {
        return {
            username: '',
            password: '',
            rememberMe: false
        }
    },
    created() {
        if (localStorage.getItem('rememberMe')) {
            this.username = localStorage.getItem('username')
            this.password = localStorage.getItem('password')
            this.rememberMe = true
        }
    },
    methods: {
        async handleLogin() {
            try {
                const res = await axios_instance.post(API.LOGIN,
                        {
                            username: this.username,
                            password: this.password
                        })
                if (!res) return;
                if (this.rememberMe) {
                    localStorage.setItem('rememberMe', this.rememberMe)
                    localStorage.setItem('username', this.username)
                    localStorage.setItem('password', this.password)
                }
                localStorage.setItem('username', this.username)
                localStorage.setItem('token', res.token)
                this.$router.push('/user')
                showSuccessToast('登录成功')
            } catch (error) {
                console.error('Login failed:', error)
            }
        }
    }
}
</script>

<style scoped>
/* Additional styles if needed */
</style>
