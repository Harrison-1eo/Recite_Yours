<template>
    <div>
        <Navbar />
        <div class="container">
            <div class="row justify-content-center mt-5">
                <div class="col-lg-4 col-md-6 col-sm-8">
                    <h2 class="text-center">注册</h2>
                    <form @submit.prevent="handleRegister">
                        <div class="mb-3">
                            <label for="username" class="form-label">用户名</label>
                            <input type="text" class="form-control" id="username" v-model="username" required />
                        </div>
                        <div class="mb-3">
                            <label for="password" class="form-label">密码</label>
                            <input type="password" class="form-control" id="password" v-model="password" required />
                        </div>
                        <button type="submit" class="btn btn-primary w-100">注册</button>
                    </form>
                    <div class="text-center mt-3">
                        <router-link to="/login">跳转登录</router-link>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import Navbar from '../components/Navbar.vue'
import axios_instance from "@/utils/axios_config";
import {showSuccessToast} from "@/utils/showMessage";

export default {
    name: 'Register',
    components: {
        Navbar
    },
    data() {
        return {
            username: '',
            password: '',
        }
    },
    methods: {
        async handleRegister() {
            try {
                await axios_instance.post('/register', {
                    username: this.username,
                    password: this.password
                })
                this.$router.push('/login')
                showSuccessToast('注册成功')
            } catch (error) {
                console.error('Registration failed:', error)
            }
        }
    }
}
</script>

<style scoped>
/* Additional styles if needed */
</style>
