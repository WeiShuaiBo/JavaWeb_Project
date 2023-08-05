<template>
  <div class="container">
    <h2 class="personal-info-title">个人信息</h2>
    <div class="content-container">
      <div class="personal-info">
        <form @submit.prevent="submitForm">
          <div>
            <label for="name">姓名:</label>
            <input type="text" id="name" v-model="name" :readonly="!editing" required />
          </div>
          <div>
            <label for="gender">性别:</label>
            <select id="gender" v-model="gender" :disabled="!editing" required>
              <option value="">请选择</option>
              <option value="male">男</option>
              <option value="female">女</option>
            </select>
          </div>
          <div>
            <label for="birthDate">出生日期:</label>
            <input type="date" id="birthDate" v-model="birthDate" :disabled="!editing" required />
          </div>
          <div>
            <label for="idCard">身份证号码：</label>
            <input type="text" id="idCard" v-model="idCard" :disabled="!editing" required />
          </div>
          <div>
            <label for="address">地址:</label>
            <textarea id="address" v-model="address" :disabled="!editing" required></textarea>
          </div>
          <div>
            <label for="email">邮箱:</label>
            <input type="email" id="email" v-model="email" :disabled="!editing" required />
          </div>
          <div class="btn-container">
            <button class="btn" v-if="!editing" @click="editForm">修改</button>
            <button class="btn" v-else @click="saveForm">保存</button>
          </div>
        </form>
      </div>
      <div class="avatar-container">
        <img class="avatar" :src="{ avatarUrl }" alt="Avatar">
        <input type="file" ref="fileInput" style="display: none" @change="handleFileUpload">
        <button class="btn" @click="selectImage">设置头像</button>

      </div>
    </div>

  </div>
</template>

<script>
import Axios from 'axios';

export default {
  name: "UserInformation",
  data() {
    return {
      name: '',
      gender: '',
      birthDate: '',
      idCard: '',
      avatarUrl: "",
      address: '',
      email: '',
      editing: false // 添加editing变量
    };
  },
  created() {
    this.getData()
  },
  methods: {
    // 触发文件选择的函数
    getData() {
      Axios.get("/getInf").then((res) => {
        console.log(res)
        this.name = res.data.name;
        this.gender = res.data.gender;
        this.birthDate = res.data.birthDate;
        this.idCard = res.data.idCard,

          this.address = res.data.address;
        this.email = res.data.email;
      })
    },
    getData1() {
      Axios.post("/")
    },
    // 点击修改按钮切换编辑状态
    editForm() {
      this.editing = true;
    },
    // 点击保存按钮保存表单内容并切换回非编辑状态
    saveForm() {
      this.editing = false;
      // 发送保存请求
      Axios.post("/updateInf", {
        name: this.name,
        gender: this.gender,
        birthDate: this.birthDate,
        idCard: this.idCard,
        address: this.address,
        email: this.email
      }).then((res) => {
        console.log(res);
      });
    },
    handleFileUpload(event) {
      const file = event.target.files[0];
      const formData = new FormData();
      formData.append('file', file);
      // Use Axios to send the file to the backend
      Axios.post("/upload", formData, {
        headers: {
          'Content-Type': 'multipart/form-data'
        }
      }).then((res) => {
        console.log(res);
        console.log(111)
        console.log(res.data)
        console.log(res.data.data)
        // Set avatarUrl from the response data
        if (res.data.code === 1000) {
          this.avatarUrl = res.data.data;
        }
      }).catch((error) => {
        console.error(error);
      });
    },

    selectImage() {
      // Trigger the file input dialog
      this.$refs.fileInput.click();
    },
    submitForm() {
      // 在这里可以处理表单提交逻辑
      console.log('提交表单');
      console.log(this.name, this.gender, this.birthDate, this.age, this.address, this.email);
    }
  }
};
</script>

<style>
.container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100vh;
  background-image: url("../img/6.jpg");
  background-size: 150%;
  background-position: center;
}

.personal-info-title {
  font-size: 24px;
  text-align: center;
  margin-bottom: 20px;
}

.content-container {
  display: flex;
  justify-content: center;
  align-items: flex-start;
  width: 80%;
  margin-top: 20px;
}

.personal-info {
  width: 50%;
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 5px;
  font-family: Arial, sans-serif;
}

.avatar-container {
  width: 45%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  margin-left: 20px;
}

.avatar {
  width: 300px;
  height: 300px;
  border-radius: 50%;
  object-fit: cover;
  margin-bottom: 10px;
}

.personal-info form div {
  margin-bottom: 15px;
}

.personal-info label {
  display: block;
  font-weight: bold;
  margin-bottom: 5px;
}

.personal-info input[type="text"],
.personal-info select,
.personal-info input[type="date"],
.personal-info input[type="number"],
.personal-info input[type="email"],
.personal-info textarea {
  width: 100%;
  padding: 8px;
  font-size: 16px;
  border-radius: 3px;
  border: 1px solid #ccc;
}

.btn-container {
  display: flex;
  justify-content: space-between;
}

.btn {
  background-color: #428bca;
  color: #fff;
  padding: 10px 20px;
  font-size: 16px;
  border: none;
  border-radius: 3px;
  cursor: pointer;
  display: block;
  margin: 0 auto;
}
</style>