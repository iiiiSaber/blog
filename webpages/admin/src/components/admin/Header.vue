<template>
    <a-button type="danger" @click="loginExit">退出</a-button>
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