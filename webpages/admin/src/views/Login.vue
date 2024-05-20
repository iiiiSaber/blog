<template>
    <div class="container">
        <div class="loginBox">
            <a-form-model ref="loginFormRef" :rules="rules" :model="formdata" class="loginForm">
                <a-form-model-item prop="username">
                    <a-input v-model="formdata.username" placeholder="请输入用户名" v-on:keyup.enter="loginAdmin">
                        <a-icon slot="prefix" type="user" style="color:rgba(0,0,0,.25)" />
                    </a-input>
                </a-form-model-item>

                <a-form-model-item prop="password">
                    <a-input v-model="formdata.password" placeholder="请输入密码" type="password" v-on:keyup.enter="loginAdmin">
                        <a-icon slot="prefix" type="lock" style="color:rgba(0,0,0,.25)" />
                    </a-input>
                </a-form-model-item>

                <a-form-model-item class="loginButtom">
                    <a-button type="primary" style="margin:20px" @click="loginAdmin">登录</a-button>
                    <a-button type="info" style="margin:20px" @click="resetForm">取消</a-button>
                </a-form-model-item>
            </a-form-model>
        </div>
    </div>
</template>

<script>
export default{
    data(){
        return{
            formdata:{
                username:'',
                password:'',
            },
            rules:{
                username: [
                    { required: true, message: '请输入用户名', trigger: 'blur' },
                    {
                        min: 4,
                        max: 12,
                        message: '用户名必须在4到12个字符之间',
                        trigger: 'blur',
                        },
                    ],
                password: [
                    { required: true, message: '请输入密码', trigger: 'blur' },
                    {
                        min: 6,
                        max: 20,
                        message: '密码必须在6到20个字符之间',
                        trigger: 'blur',
                    },
                ],
            }
        }
    },
    methods:{
        resetForm(){
            this.$refs.loginFormRef.resetFields()
        },
        loginAdmin(){
            this.$refs.loginFormRef.validate(async valid=>{
                //console.log(valid)
                if(!valid) {
                    return this.$message.error("未输入符合要求的数据，请重新输入")
                }
                // console.log(this.formdata.username)
                // console.log(this.formdata.password)
                const {data:res} = await this.$http.post('login',this.formdata)
                // console.log(res)

                if(res.status!=200){
                    return this.$message.error(res.message)
                }
                window.sessionStorage.setItem('token',res.token)
                this.$router.push('admin')
            })
        }
    }
}
</script>

<style scoped>
.container{
    height:100%;
    background-image: url('../assets/back_login_pic.png');
    /* 背景图不平铺 */
    background-repeat: no-repeat;
    /* 让背景图基于容器大小伸缩 */
    background-size: cover;
}
.loginBox{
    width:450px;
    height: 300px;
    background-color: rgba(255,255,255,0.3);
    position: absolute;
    top:50%;
    left:75%;
    transform: translate(-50%,-50%);
    border-radius: 10px;
}
.loginForm{
    width:100%;
    position:absolute;
    bottom:10%;
    padding:0 20px;
    box-sizing: border-box;
}
.loginButtom{
    display: flex;
    justify-content: flex-end;
}
</style>
