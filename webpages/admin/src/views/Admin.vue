<template>
    <a-layout class="container">

        <Nav></Nav>

        <a-layout>

            <a-layout-header class="headerBtnExit">
                <Header></Header>
            </a-layout-header>

            <a-layout-content style="background-color: rgba(255,255,255,0.3);">
                <router-view>

                </router-view>
            </a-layout-content>

            <a-layout-footer style="background-color: rgba(255,255,255,0.3);">
                <Footer></Footer>
            </a-layout-footer>
            
        </a-layout>
    </a-layout>
</template>

<script>
import Nav from '../components/admin/Nav'
import Footer from '../components/admin/Footer'
import Header from '../components/admin/Header'
export default{
    components:{Nav,Footer,Header},
}
</script>

<style scoped>
.container{
    height:100%;
}
.headerBtnExit{
    background-color: rgba(255,255,255,0.3);
    display:flex;
    justify-content:flex-end;
    align-items:center;
}

</style>
