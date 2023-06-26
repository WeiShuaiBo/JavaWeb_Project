<template>
    <div class="background">
        <div class="registration-page">
        <h2>大学生创新创业平台申请报名</h2>
        <form @submit.prevent="submitForm">
            <div class="form-group">
                <label for="name">姓名：</label>
                <input type="text" id="name" v-model.trim="formData.name" required />
            </div>

            <div class="form-group">
                <label for="university">学校：</label>
                <input type="text" id="university" v-model.trim="formData.university" required />
            </div>

            <div class="form-group">
                <label for="college">学院：</label>
                <select id="college" v-model.trim="formData.college" required>
                    <option value="">请选择学院</option>
                    <option value="计算机科学与技术学院">计算机学院</option>
                    <option value="信息工程学院">信息工程学院</option>
                    <option value="软件学院">软件学院</option>
                    <!-- 其他学院选项 -->
                </select>
            </div>

            <div class="form-group">
                <label for="major">专业：</label>
                <select id="major" v-model.trim="formData.major" required>
                    <option value="">请选择专业</option>
                    <option value="计算机科学与技术">计算机科学与技术</option>
                    <option value="数据科学与大数据技术">数据科学与大数据技术</option>
                    <option value="软件工程">软件工程</option>
                    <option value="人工智能">人工智能</option>
                    <!-- 其他专业选项 -->
                </select>
            </div>

            <div class="form-group">
                <label for="email">邮箱：</label>
                <input type="email" id="email" v-model.trim="formData.email" required />
            </div>

            <div class="form-group">
                <label for="phone">手机号码：</label>
                <input type="tel" id="phone" v-model.trim="formData.phone" required />
            </div>

            <div class="form-group">
                <label for="projectIdea">项目创意：</label>
                <textarea id="projectIdea" v-model.trim="formData.projectIdea" required></textarea>
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
export default {
    name:"Registration",
    data() {
        return {
            formData: {
                name: '',
                university: '',
                college:'',
                major: '',
                email: '',
                phone: '',
                projectIdea: ''
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
                this.$router.push({ path: '/display', props: { formData: this.formData } });
            }, 1000);

            this.$store.commit('setFormData', this.formData)
        },
        validateForm() {
            // 进行表单验证，确保所有字段都填写正确
            // ...

            return true; // 返回true表示验证通过，可以提交表单
        }
    }
};
</script>
<style scoped>
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
