<template>
  <div class="container">
    <h2 class="personal-info-title">个人信息</h2>
    <div class="content-container">
      <div class="personal-info">
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
      <div class="avatar-container">
        <img class="avatar" src="../assets/images/avatar.png" alt="Avatar">
        <button class="btn" @click="selectImage">设置头像</button>
      </div>
    </div>
    
  </div>
</template>

<script>
export default {
  name: "UserInformation",
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
    selectImage() {
      // 创建一个<input>元素
      var input = document.createElement('input');
      input.type = 'file';

      // 当用户选择文件时，触发change事件
      input.addEventListener('change', function (event) {
        var file = event.target.files[0];

        var reader = new FileReader();

        // 当文件加载完成时，将图片显示到<img>元素中
        reader.onload = function (e) {
          var img = document.querySelector('.avatar');
          img.src = e.target.result;
        };

        // 读取图片文件的内容
        reader.readAsDataURL(file);
      });

      // 触发<input>元素的点击事件，打开文件选择对话框
      input.click();
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