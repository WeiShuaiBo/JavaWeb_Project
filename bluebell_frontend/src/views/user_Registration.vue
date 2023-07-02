<template>
  <div class="background">
    <div class="registration-page">
      <h2>大学生创新创业平台申请报名</h2>
      <form @submit.prevent="submitForm">
        <div class="form-group">
          <label for="username">组长姓名：</label>
          <input type="text" id="username" v-model.trim="name" required />
        </div>

        <div class="form-group">
          <label for="university">学校：</label>
          <select id="university" v-model.trim="university" required>
            <option value="">请选择学校</option>
            <option value="河南科技学院">河南科技学院</option>
            <option value="新乡医学院">新乡医学院</option>
            <option value="新乡学院">新乡学院</option>
            <option value="河南师范大学">河南师范大学</option>
            <option value="985名校">985名校</option>
            <option value="211名校">211名校</option>
          </select>
        </div>

        <div class="form-group">
          <label for="college">学院：</label>
          <select id="college" v-model.trim="college" required>
            <option value="">请选择学院</option>
            <option value="计算机科学与技术学院">计算机学院</option>
            <option value="信息工程学院">信息工程学院</option>
            <option value="软件学院">软件学院</option>
            <option value="其他">其他</option>
            <!-- 其他学院选项 -->
          </select>
        </div>

        <div class="form-group">
          <label for="major">专业：</label>
          <select id="major" v-model.trim="major" required>
            <option value="">请选择专业</option>
            <option value="计算机科学与技术">计算机科学与技术</option>
            <option value="数据科学与大数据技术">数据科学与大数据技术</option>
            <option value="软件工程">软件工程</option>
            <option value="人工智能">人工智能</option>
            <option value="其他">其他</option>
            <!-- 其他专业选项 -->
          </select>
        </div>

        <div class="form-group">
          <label for="email">邮箱：</label>
          <input type="email" id="email" v-model.trim="email" required />
        </div>

        <div class="form-group">
          <label for="phone">手机号码：</label>
          <input type="tel" id="phone" v-model.trim="phone" required />
        </div>

        <div class="form-group">
          <label for="projectDirection">申报类型：</label>
          <select id="projectDirection" v-model.trim="projectDirection" required>
            <option value="">请选择方向</option>
            <option value="科技创新">科技创新</option>
            <option value="社会企业">社会企业</option>
            <option value="数字经济">数字经济</option>
            <option value="乡村振兴">乡村振兴</option>
            <option value="其他">其他</option>
          </select>
        </div>

        <button :disabled="isSubmitting" type="submit">
          {{ isSubmitting ? '提交中...' : '提交申请' }}
        </button>
      </form>
      <div v-if="isSubmitted" class="success-message">申请提交成功！</div>
    </div>
  </div>
</template>
<script>
import axios from "axios";
export default {
  name: "Registration",
  data() {
    return {
      formData: {
        username: "",
        university: "",
        college: "",
        major: "",
        email: "",
        phone: "",
        projectDirection: ""
      },
      isSubmitting: false,
      errorMessage: "",
      showMessage: false
    };
  },
  methods: {
    submitForm() {
      // 表单验证逻辑，确保字段填写正确
      if (!this.validateForm()) {
        return;
      }

          this.isSubmitting = true;
            axios.post('/createProject')
                .then(response => {
                    // 请求成功后的处理逻辑
                    console.log(response.data);  // 输出后端返回的数据

                    this.isSubmitting = false;
                    this.isSubmitted = true;

                    // 表单提交成功后，跳转到 DisplayDataPage 并传递表单数据
                    this.$router.push({ path: '/declaration'});
                })
                .catch(error => {
                    // 处理错误情况
                    console.error(error);
                    this.isSubmitting = false;
                });
      setTimeout(() => {
        // 在这里可以发送请求到服务器保存用户输入的信息

        this.isSubmitting = false;
        this.isSubmitted = true;

        // 表单提交成功后，跳转到 DisplayDataPage 并传递表单数据
        this.$router.push({ path: '/declaration'});
      }, 1000);
    },
    validateForm() {
      // 进行表单验证，确保所有字段都填写正确
      // ...

      return true; // 返回true表示验证通过，可以提交表单
    },
    showSuccessMessage() {
      this.showMessage = true;
      setTimeout(() => {
        this.showMessage = false;
      }, 3000); // Show the success message for 3 seconds
    },
    showErrorMessage(message) {
      this.errorMessage = message;
      this.showMessage = true;
      setTimeout(() => {
        this.showMessage = false;
        this.errorMessage = "";
      }, 3000); // Show the error message for 3 seconds
    }
  }
};
</script><style scoped>
label {
  font-family: Arial, sans-serif; /* 设置字体为 Arial 或者 sans-serif */
  font-size: 16px; /* 设置字体大小为 16 像素 */
  /* 其他样式属性，例如字体颜色、字体加粗等，也可以在这里添加 */
}
.background{
    width: 100%;
    height: 100%;
    background-image: url("../img/background.png");
}
.registration-page {
    max-width: 600px;
    margin: 100px auto;
    padding: 20px;
    background-color: #f8f8f8;
    border-radius: 5px;

}
h2 {
    text-align: center;
    margin-bottom: 20px;
}

form {
    display: grid;
    gap: 10px;
}

.form-group {

  display: grid;
  grid-template-columns: 120px 1fr;
  align-items: center;
  margin-bottom: 15px;
}

.form-group label {
  text-align: right;
  padding-right: 10px;
  font-weight: bold;
}

.form-group select,
.form-group input,
.form-group textarea {
  padding: 5px;
  border: 1px solid #ccc;
  border-radius: 4px;
  width: 100%;
}


label {
    font-weight: bold;
}

input,
textarea {
    width: 100%;
    padding: 10px;
    border: 1px solid #ccc;
    border-radius: 3px;
}

button {
    padding: 10px;
    background-color: #4caf50;
    color: #fff;
    border: none;
    border-radius: 3px;
    cursor: pointer;
}

.success-message {
    margin-top: 20px;
    text-align: center;
    color: green;
}
</style>
