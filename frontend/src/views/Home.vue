<template>
    <div>
        <Navbar />
        <div class="container">
            <div class="row justify-content-center mt-5">
                <div class="col-lg-6 col-md-8 col-sm-10">
                    <h1 class="text-center">Welcome to Recite Yours!</h1>
                    <div class="text-center mt-4 button-box">
                        <router-link class="btn btn-primary button-box-button" to="/login">登录</router-link>
                        <router-link class="btn btn-secondary ml-3 button-box-button" to="/register">注册</router-link>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import Navbar from '../components/Navbar.vue'

export default {
    name: 'Home',
    components: {
        Navbar
    }
}
</script>

<style scoped>
.button-box{
    display: flex;
    justify-content: center;
    flex-direction: column;
    align-items: center;
    margin-top: 20px;
}

.button-box-button{
    margin: 10px;
}
</style>
