<template>
  <div class="display">
    <div class="display-data-page" v-for="(item, index) in formData" :key="index">
      <h2>已提交申请信息</h2>
      <div class="data-table">
        <div>
          <strong>组长姓名：</strong>{{ item.name }}
        </div>
        <div>
          <strong>学校：</strong>{{ item.university }}
        </div>
        <div>
          <strong>学院：</strong>{{ item.college }}
        </div>
        <div>
          <strong>专业：</strong>{{ item.major }}
        </div>
        <div>
          <strong>邮箱：</strong>{{ item.email }}
        </div>
        <div>
          <strong>手机号码：</strong>{{ item.phone }}
        </div>
        <div>
          <strong>申报类型：</strong>{{ item.projectDirection }}
        </div>
        <div>
          <strong>项目类型：</strong>{{ item.projectSort }}
        </div>
        <div>
          <strong>项目名称：</strong>{{ item.projectName }}
        </div>
        <div>
          <strong>项目成员：</strong>{{ item.member }}
        </div>
        <div>
          <strong>项目介绍：</strong>{{ item.introduction }}
        </div>
        <div>
          <strong>项目创意：</strong>{{ item.creativity }}
        </div>
        <div>
          <strong>申请优势：</strong>{{ item.advantage }}
        </div>
        <div>
          <strong>指导老师：</strong>{{ item.instructor }}
        </div>
        <div>
          <strong style="color: brown;">申请状态：</strong>{{ item.instatus }}
        </div>
        <div>
          <strong style="color: brown;">审批描述：</strong>{{ item.content }}
        </div>
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

  data() {
    return {
      isSubmitting: false,
      formData: []
    };
  },
  mounted() {
    this.fetchData();
  },
  methods: {
    fetchData() {
      axios.get('/listProject')
        .then(response => {

          if (response.data.code === 1000) {
            this.formData = response.data.data
            console.log(response.data.data)
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