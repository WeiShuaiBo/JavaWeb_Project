<template>
  <div class="display">
    <div class="display-data-page">
      <h2>已提交申请信息</h2>
      <div v-if="formData" class="data-table">
        <div>
          <strong>组长姓名：</strong>{{ formData.name }}
        </div>
        <div>
          <strong>学校：</strong>{{ formData.university }}
        </div>
        <div>
          <strong>学院：</strong>{{ formData.college }}
        </div>
        <div>
          <strong>专业：</strong>{{ formData.major }}
        </div>
        <div>
          <strong>邮箱：</strong>{{ formData.email }}
        </div>
        <div>
          <strong>手机号码：</strong>{{ formData.phone }}
        </div>
        <div>
          <strong>申报类型：</strong>{{ formData.projectDirection }}
        </div>
        <div>
          <strong>项目类型：</strong>{{ formData.projectSort }}
        </div>
        <div>
          <strong>项目名称：</strong>{{ formData.projectName }}
        </div>
        <div>
          <strong>项目成员：</strong>{{ formData.Member }}
        </div>
        <div>
          <strong>项目介绍：</strong>{{ formData.Introduction }}
        </div>
        <div>
          <strong>项目创意：</strong>{{ formData.Creativity }}
        </div>
        <div>
          <strong>申请优势：</strong>{{ formData.Advantage }}
        </div>
        <div>
          <strong>指导老师：</strong>{{ formData.Instructor }}
        </div>
      </div>
      <div v-else class="no-data-message">
        暂无数据
      </div>
    </div>
    <div class="show">
      <button :disabled="isSubmitting" type="submit" @click="goToHomePage">
        {{ isSubmitting ? '返回中...' : '返回主页' }}
      </button>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: "DisplayDataPage",

  computed: {
    formData() {
      return this.$store.state.formData;
    },
  },
  data() {
    return {
      isSubmitting: false,
      formData1: []
    };
  },
  mounted() {
    this.fetchData();
  },
  methods: {
    fetchData() {
      axios.get('/listProject')
          .then(response => {
            if (response.code === 1000) {
              const data = response.data;
              this.formData.projectSort = data.projectSort;
              this.formData.projectName = data.projectName;
              this.formData.Member = data.Member;
              this.formData.Introduction = data.Introduction;
              this.formData.Creativity = data.Creativity;
              this.formData.Advantage = data.Advantage;
              this.formData.Instructor = data.Instructor;
            } else {
              console.log(response.message);
            }
          })
          .catch(error => {
            console.log("An error occurred while fetching data.", error);
          });
    },
    goToHomePage() {
      this.$router.push({ name: 'Home' });
    }
  }
};
</script>

<style scoped>
.display-data-page {
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

.data-table {
  font-size: 14px;
}

.data-table>div {
  padding: 10px;
  border-bottom: 1px solid #ddd;
}

.no-data-message {
  text-align: center;
  padding: 20px;
  font-size: 16px;
  color: #999;
}

.show {
  display: flex;
  justify-content: center;
  align-items: center;
}

button {
  padding: 10px;
  background-color: #4caf50;
  color: #fff;
  border: none;
  border-radius: 3px;
  cursor: pointer;
  justify-content: center;
  align-items: center;
}
</style>