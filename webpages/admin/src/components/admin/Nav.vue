<template>
    <a-layout-sider style="background: rgba(255,255,255,0.3)" breakpoint="lg" v-model="collapsed">
        <div class="logo">
            <span>{{collapsed? '博客':'博客管理'}}</span>
        </div>
        
        <a-menu style="background-color: rgba(255,255,255,0.3);" mode="inline" @click="goToPage">
            <a-menu-item key="dashboard">
                <a-icon type="dashboard"></a-icon><span>仪表盘</span>
            </a-menu-item>
            <a-sub-menu>
                <span slot="title">
                    <a-icon type="file"/>
                    <span>文章列表</span>
                </span>
                <a-menu-item key="addarticle">
                    <a-menu-item ><a-icon type="form"></a-icon><span>写文章</span></a-menu-item>
                </a-menu-item>
                <a-menu-item key="articlelist">
                    <a-menu-item ><a-icon type="snippets"></a-icon><span>文章列表</span></a-menu-item>
                </a-menu-item>
            </a-sub-menu>
            <a-menu-item key="categorylist">
                <a-menu-item><a-icon type="book"></a-icon><span>分类列表</span></a-menu-item>
            </a-menu-item>
            <a-menu-item key="userlist">
                <a-menu-item><a-icon type="user"></a-icon><span>用户列表</span></a-menu-item>
            </a-menu-item>
        </a-menu>
    </a-layout-sider>
</template>

<script>
export default {
    data(){
        return {
            collapsed:false
        }
    },
    methods:{
        goToPage(item){
            console.log(item.key)
            this.$router.push(item.key).catch(err=>err)
        }
    }
}
</script>

<style>
.logo{
    height: 32px;
    margin:16px;
    border-radius: 10px;
    background-color: #dde3e9;
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 16px;
}

</style>