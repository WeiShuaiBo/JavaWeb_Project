<template>
    <div class="lyear-login-box">
        <div class="lyear-login-left">
            <div class="lyear-overlay"></div>
            <div class="lyear-featured">
                <h4>鱼上了岸，也就不再是鱼。 <small> --三体II: 黑暗森林</small></h4>
            </div>
        </div>
        <div class="lyear-login-right">

            <div class="lyear-logo text-center">
                <a href="#!"><img src="../img/logo-sidebar.png" alt="logo" /></a>
            </div>

            <form>
                <div class="form-group">
                    <label for="username">用户名</label>
                    <input type="text" class="form-control" v-model="username" id="username" name="username"
                        placeholder="请输入您的用户名">
                </div>

                <div class="form-group">
                    <label for="password">密码</label>
                    <input type="password" class="form-control" v-model="password" id="password" name="password"
                        placeholder="请输入您的密码">
                </div>

                <div class="form-group">


                    <div class="form-group">
                        <button class="btn btn-block btn-primary" type="submit" @click="submit()">立即登录</button>
                    </div>
                </div>
                <footer class="text-center">
                    <p class="m-b-0">版权所有 © 2023 数据202. 保留所有权利</p>
                </footer>
            </form>

        </div>
    </div>
</template>

<script>
import axios from "axios";
import Swallow from "sweetalert2";
export default {
    name: 'admlogin',
    // 可以在这里添加组件的其他属性和方法
    data() {
        return {
            username: '',
            password: '',
            rules: {
                username: [
                    { required: true, message: '请输入用户名', trigger: 'blur' },
                ],
                password: [
                    { required: true, message: '请输入密码', trigger: 'blur' },
                ],
            },
        };
    },
    methods: {
        submit() {
            axios.post("/login1", {
                username: this.username,
                password: this.password,
            })
                .then((res) => {
                    console.log(res + "管理员数据")
                    if (res.code == 1000) {
                        localStorage.setItem("loginResult", JSON.stringify(res.data));
                        this.$store.commit("login", res.data);
                        // this.$router.push({ path: this.redirect || "/" });
                        console.log('signup success');
                        // debugger
                        this.$router.push({ name: "admindex" });
                        Swallow.fire({
                            icon: 'success',
                            title: '登录成功',
                            text: '进入主页'
                        })
                    } else {
                        console.log(333333, res)
                        alert("登录失败")
                        // console.log(res.data.msg);
                        Swallow.fire({
                            icon: 'error',
                            title: '登录失败',
                            text: res.data.msg
                        }).then(() => {
                            this.username = "";
                            this.password = "";
                        });
                    }
                })
                .catch((error) => {
                    console.log(error);
                });
        },
    }
}
</script>

<style scoped>
body {
    background-color: #f8f8f8;
}

.lyear-login-box {
    display: flex;
    height: 100vh;
}

.lyear-login-left {
    position: relative;
    width: 50%;
    background-image: url('../img/login.jpg');
    background-size: cover;
    background-repeat: no-repeat;
    background-position: center center;
}

.lyear-overlay {
    position: absolute;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    background: rgba(0, 0, 0, 0.5);
}

.lyear-featured {
    position: absolute;
    bottom: 30px;
    left: 30px;
    color: #fff;
}

.lyear-featured h4 {
    font-size: 24px;
    line-height: 30px;
}

.lyear-featured small {
    display: block;
    font-size: 14px;
    text-align: right;
    margin-top: 15px;
}

.lyear-login-right {
    position: relative;
    width: 50%;
    padding: 50px;
    background-color: #fff;
}

.lyear-logo {
    text-align: center;
    margin-bottom: 50px;
}

.form-group {
    margin-bottom: 20px;
}

.form-control {
    width: 100%;
    height: 40px;
    padding: 10px 15px;
    font-size: 14px;
    line-height: 1.42857143;
    border: 1px solid #ccc;
    border-radius: 4px;
    transition: border-color ease-in-out 0.15s, box-shadow ease-in-out 0.15s;
}

.form-control:focus {
    border-color: #66afe9;
    outline: 0;
    box-shadow: 0 0 8px rgba(102, 175, 233, 0.6);
}

.form-group label {
    font-weight: bold;
}

.text-center {
    text-align: center;
}

.btn {
    display: inline-block;
    padding: 6px 12px;
    margin-bottom: 0;
    font-size: 14px;
    font-weight: 400;
    line-height: 1.42857143;
    text-align: center;
    white-space: nowrap;
    vertical-align: middle;
    touch-action: manipulation;
    cursor: pointer;
    user-select: none;
    background-image: none;
    border: 1px solid transparent;
    border-radius: 4px;
    transition: color ease-in-out 0.15s, background-color ease-in-out 0.15s, border-color ease-in-out 0.15s, box-shadow ease-in-out 0.15s;
}

.btn-primary {
    color: #fff;
    background-color: #337ab7;
    border-color: #2e6da4;
}

.btn-primary:hover,
.btn-primary:focus,
.btn-primary:active,
.btn-primary.active,
.open .dropdown-toggle.btn-primary {
    color: #fff;
    background-color: #286090;
    border-color: #204d74;
}

.m-b-0 {
    margin-bottom: 0;
}

/* 添加其他自定义样式 */
</style>