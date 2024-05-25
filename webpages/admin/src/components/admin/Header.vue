<template>
    <div class="head">
        <span class="backTitle">博客后台管理</span>
        <a-button type="danger" @click="loginExit">退出</a-button>
    </div>
</template>

<script>
export default{
    methods:{
        loginExit(){
            window.sessionStorage.clear('token')
            this.$router.push('/login')
        }
    }
}
</script>
<style>
.head{
    display: flex;
    justify-content: space-between;
    align-items: center;
}
.backTitle{
    font-size: 24px;
    font-weight: bold;
}
</style>