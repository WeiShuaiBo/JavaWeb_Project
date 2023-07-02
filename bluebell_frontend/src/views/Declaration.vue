<template>
  <div class="background">
    <div class="registration-page">
      <h2>大学生创新创业平台项目信息</h2>
      <form @submit.prevent="submitForm">
        <div class="form-group">
          <label for="projectSort">项目类型：</label>
          <select id="projectSort" v-model.trim="formData.projectSort" required>
            <option value="">请选择类型</option>
            <option value="新建项目">新建项目</option>
            <option value="扩展项目">扩展项目</option>
            <option value="改建项目">改建项目</option>
            <option value="迁建项目">迁建项目</option>
            <option value="恢复项目">恢复项目</option>
          </select>
        </div>
      </form>
    </div>
    <div>
      <button>回到主页</button>
    </div>
  </div>
</template>

<script>
export default {
  name: "Registration",
  data() {
    return {
      formData: {
        projectSort: '',
        projectName: '',
        Member: [''], // Initialize with an empty member input
        Introduction: '',
        Creativity: '',
        Advantage: '',
        Instructor: [''] // Initialize with an empty instructor input
      },
      isSubmitting: false,
      isSubmitted: false
    };
  },
  methods: {
    submitForm() {
      // 表单验证逻辑，确保字段填写正确
      if (!this.validateForm()) {
        return;
      }

      this.isSubmitting = true;
      setTimeout(() => {
        // 在这里可以发送请求到服务器保存用户输入的信息

        this.isSubmitting = false;
        this.isSubmitted = true;

        // 表单提交成功后，跳转到 DisplayDataPage 并传递表单数据
        this.$router.push({ path: "/display", props: { formData: this.formData } });
      }, 1000);

      this.$stores.commit("setFormData", this.formData);
    },
    validateForm() {
      // 进行表单验证，确保所有字段都填写正确
      // ...

      return true; // 返回true表示验证通过，可以提交表单
    },
    addMember() {
      this.formData.Member.push('');
    },
    removeMember(index) {
      this.formData.Member.splice(index, 1);
    },
    addInstructor() {
      this.formData.Instructor.push('');
    },
    removeInstructor(index) {
      this.formData.Instructor.splice(index, 1);
    },
  },
};
</script>
<style scoped>
.background {
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
    grid-template-columns: 100px 1fr;
    align-items: center;
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
