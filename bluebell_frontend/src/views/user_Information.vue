<template>
  <div class="container">
    <div class="personal-info">
      <h2>个人信息</h2>
      <div class="avatar-container">
          <img class="avatar" src="avatar.jpg" alt="Avatar">
          <button class="btn">设置头像</button>
      </div>
      <form @submit.prevent="submitForm">
        <div>
          <label for="name">姓名:</label>
          <input type="text" id="name" v-model="name" required />
        </div>
        <div>
          <label for="gender">性别:</label>
          <select id="gender" v-model="gender" required>
            <option value="">请选择</option>
            <option value="male">男</option>
            <option value="female">女</option>
          </select>
        </div>
        <div>
          <label for="birthDate">出生日期:</label>
          <input type="date" id="birthDate" v-model="birthDate" required />
        </div>
        <div>
          <label for="idCard">身份证号码：</label>
          <input type="text" id="idCard" v-model="idCard" required />
        </div>

        <!-- <label for="avatar">头像：</label>
          <input type="file" accept="image/*" @change="handleFileUpload">
          <img :src="avatarUrl" alt="头像" v-if="avatarUrl">

          <br><br> -->
        <div>
          <label for="address">地址:</label>
          <textarea id="address" v-model="address" required></textarea>
        </div>
        
        <div>
          <label for="email">邮箱:</label>
          <input type="email" id="email" v-model="email" required />
        </div>
        <div class="btn-container">
            <button class="btn">保存</button>
            <button class="btn">修改</button>
        </div>
      </form>
    </div>
  </div>

</template>

<script>
export default {
  name:"UserInformation",
   data() {
    return {
      name: '',
      gender: '',
      birthDate: '',
      idCard: '',
      avatarUrl: null,
      address: '',
      email: ''
    };
  },
  methods: {
    handleFileUpload(event) {
      const file = event.target.files[0];
      this.avatarUrl = URL.createObjectURL(file);
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
.avatar-container {
  position: relative;
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-bottom: 20px;
}

.avatar {
  width: 120px;
  height: 120px;
  border-radius: 50%;
  object-fit: cover;
  margin-bottom: 10px;
}
.personal-info {
  width: 600px;
  margin: 0 auto;
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 5px;
  font-family: Arial, sans-serif;
}

.personal-info h2 {
  font-size: 24px;
  text-align: center;
  margin-bottom: 20px;
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

/* .personal-info button[type="submit"]:hover {
  background-color: #3071a9;
} */
/* .btn && .submit {
  background-color: #428bca;
  color: #fff;
  padding: 10px 20px;
  font-size: 16px;
  border: none;
  border-radius: 3px;
  cursor: pointer;
  display: block;
  margin: 0 auto;
} */
.container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}

.personal-info {
  max-width: 400px;
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 5px;
  font-family: Arial, sans-serif;
}

.personal-info h2 {
  font-size: 24px;
  text-align: center;
  margin-bottom: 20px;
}

</style>