<!-- src/views/UserHome.vue -->
<template>
    <div class="container mt-5">
        <h1 style="font-weight: bold">Recite Yours!</h1>
        <div class="button-box">
            <button class="btn btn-primary button-box-button" @click="goToAddWord">添加生词</button>
            <button class="btn btn-secondary button-box-button" @click="goToShowWords">导出生词本</button>
            <button class="btn btn-danger button-box-button" @click="logout">退出登录</button>
        </div>

    </div>
</template>

<script>
export default {
    name: 'UserHome',
    methods: {
        goToAddWord() {
            this.$router.push({ name: 'AddWord' }); // 跳转到添加生词页面
        },
        goToShowWords() {
            this.$router.push({ name: 'ShowWords' }); // 跳转到导出生词本页面
        },
        logout() {
            localStorage.removeItem('token'); // 清除 token
            this.$router.push({ name: 'Login' }); // 跳转到登录页面
        }
    }
}
</script>

<style scoped>
.button-box{
    display: flex;
    justify-content: center;
    flex-direction: column;
    align-items: flex-start;
    margin-top: 20px;
}

.button-box-button{
    margin: 10px;
}
</style>
