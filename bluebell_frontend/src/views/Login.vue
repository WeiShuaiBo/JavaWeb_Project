<template>
  <div id="login">
    <transition appear name="opacitytrans">
      <div class="container" id="container">
        <div class="form-container sign-in-container">
          <form action="#">
            <h1>大学生创新创业平台</h1>

            <span>Version 1.0.0</span>
            <input type="text" placeholder="账号" v-model="username" />
            <input type="password" placeholder="密码" v-model="password" />
            <input type="text" class="form-control" v-model="captchaId" name="captcha" id="captcha" placeholder="请输入验证码" />
            <div class="row">
              <div>
                <img id="captchaId" :src="captchaUrl" alt="验证码" >
                <button @click="refresh">刷新</button>
              </div>
            </div>
            <div class="button" @click="submit()">登录</div>
          </form>
        </div>
        <div class="overlay-container">
          <div class="overlay">
            <div class="overlay-panel overlay-right">
              <img class="logo" src="../img/studentlogo.jpg" alt="" />
              <p>
                「 Innovation leads the future 」
              </p>
              <!-- <button class="ghost" id="signUp">About Us</button>n -->
            </div>
          </div>
        </div>
      </div>
    </transition>
  </div>
</template>


<script>
import axios from "axios";
import Swallow from "sweetalert2";
export default {
  name: "SignUp",
  data() {
    return {
      username: "",
      password: "",
      captchaId: "",
      captchaUrl: "/api/v1/captcha",
    };
  },

  methods: {
    refresh() {
      this.captchaUrl = "/api/v1/captcha?timestamp=" + Date.now(); // 添加时间戳参数实现刷新验证码
    },
    submit() {
      axios
          .post("/login", {
            username: this.username,
            password: this.password,
            captchaId: this.captchaId // 将 captchaId 作为参数传递给后端
          })
          .then((res) => {
            console.log(res.data);
            if (res.code == 1000) {
              localStorage.setItem("loginResult", JSON.stringify(res.data));
              this.$store.commit("login", res.data);
              this.$router.push({ path: this.redirect || "/" });
              console.log('signup success');
              Swallow.fire({
                icon: 'success',
                title: '登录成功',
                text: '进入主页'
              }).then(() => {
                this.$router.push({ name: "Home" });
              });
            } else {
              alert("登录失败")
              console.log(res.data.msg);
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
  },
};
</script>
<style  scoped>
#login {
  font-family: "Montserrat", sans-serif;
  background: #f6f5f7;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 100vh;
  /* margin: -20px 0 50px; */
  /* background-image: url("https://www.17sucai.com/preview/1749733/2019-06-22/%E7%99%BB%E5%BD%95/img/img1.png"); */
  background-image: url('https://togathers3.s3.ap-northeast-1.amazonaws.com/static/adminV3/login.jpg');
  background-size: cover;
}

.logo {
  width: 160px;
  height: auto;
}

h1 {
  font-weight: bold;
  margin: 0;
  color: beige;
}

p {
  font-size: 14px;
  font-weight: bold;
  line-height: 20px;
  letter-spacing: 0.5px;
  margin: 20px 0 30px;
  color: black;
}

span {
  font-size: 12px;
  color: beige;
}

a {
  color: #fff;
  font-size: 14px;
  text-decoration: none;
  margin: 15px 0;
}

.container {
  border-radius: 10px;
  box-shadow: 0 14px 28px rgba(0, 0, 0, 0.25), 0 10px 10px rgba(0, 0, 0, 0.22);
  position: absolute;
  overflow: hidden;
  width: 768px;
  max-width: 100%;
  min-height: 480px;
  opacity: 0.8;
}

.form-container form {
  background: rgb(18, 21, 21);
  display: flex;
  flex-direction: column;
  padding: 0 50px;
  height: 100%;
  justify-content: center;
  align-items: center;
  text-align: center;
}

.social-container {
  margin: 20px 0;
}

.social-container a {
  border: 1px solid #ddd;
  border-radius: 50%;
  display: inline-flex;
  justify-content: center;
  align-items: center;
  margin: 0 5px;
  height: 40px;
  width: 40px;
}

.form-container input {
  background: #eee;
  border: none;
  padding: 12px 15px;
  margin: 8px 0;
  width: 100%;
}

.button {
  cursor: pointer;
  border-radius: 20px;
  /* border: 1px solid #ff4b2b;
  background: #ff4b2b; */
  /* border: 1px solid #fa8817;
  background: #fa8817; */
  border: 1px solid #1BBFB4;
  background: #1BBFB4;
  color: #fff;
  font-size: 12px;
  font-weight: bold;
  padding: 12px 45px;
  margin-top: 20px;
  letter-spacing: 1px;
  text-transform: uppercase;
  transition: transform 80ms ease-in;
}

input[type="text"] {
  width: 240px;
  text-align: center;
  background: transparent;
  border: none;
  border-bottom: 1px solid #fff;
  font-family: "PLay", sans-serif;
  font-size: 16px;
  font-weight: 200px;
  padding: 10px 0;
  transition: border 0.5s;
  outline: none;
  color: #fff;
  font-weight: bold;
}

input[type="password"] {
  width: 240px;
  text-align: center;
  background: transparent;
  border: none;
  border-bottom: 1px solid #fff;
  font-family: "PLay", sans-serif;
  font-size: 16px;
  font-weight: bold;
  padding: 10px 0;
  transition: border 0.5s;
  outline: none;
  color: #fff;
}

input[type="email"] {
  width: 240px;
  text-align: center;
  background: transparent;
  border: none;
  border-bottom: 1px solid #fff;
  font-family: "PLay", sans-serif;
  font-size: 16px;
  font-weight: 200px;
  padding: 10px 0;
  transition: border 0.5s;
  outline: none;
  color: #fff;
  font-weight: bold;
}

.button:active {
  transform: scale(0.95);
}

.button:focus {
  outline: none;
}

.button.ghost {
  background: transparent;

  /* border-color: #fa8817;
  background-color: #fa8817; */
  border-color: #1BBFB4;
  background-color: #1BBFB4;
  margin: 0;
}

.form-container {
  position: absolute;
  top: 0;
  height: 100%;
  transition: all 0.6s ease-in-out;
}

.sign-in-container {
  left: 0;
  width: 50%;
  z-index: 2;
}

.sign-up-container {
  left: 0;
  width: 50%;
  z-index: 1;
  opacity: 0;
}

.overlay-container {
  position: absolute;
  top: 0;
  left: 50%;
  width: 50%;
  height: 100%;
  overflow: hidden;
  transition: transform 0.6s ease-in-out;
  z-index: 100;
}

.overlay {
  background: transparent;
  background: linear-gradient(to right, #ff4b2b, #ff416c) no repeat 0 0 / cover;
  color: #fff;
  position: absolute;
  left: -100%;
  height: 100%;
  width: 200%;
  transform: translateX(0);
  transition: transform 0.6s ease-in-out;
}

.overlay-panel {
  position: absolute;
  top: 0;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;

  height: 100%;
  width: 50%;
  text-align: center;
  transform: translateX(0);
  transition: transform 0.6s ease-in-out;
}

.overlay-right {
  right: 0;
  transform: translateX(0);
}
</style>