<template>
  <div class="user-profile">
    <h2>个人信息</h2>
    <div class="user-details">
      <div class="field">
        <label>姓名:</label>
        <span>{{ user.user_name }}</span>
      </div>
      <div class="field">
        <label>年龄:</label>
        <span>{{ user.user_age }}</span>
      </div>
      <div class="field">
        <label>性别:</label>
        <span>{{ user.user_sex }}</span>
      </div>
      <div class="field">
        <label>地址:</label>
        <span>{{ user.user_address }}</span>
      </div>
      <div class="field">
        <label>身份证号:</label>
        <span>{{ user.user_postId }}</span>
      </div>
      <div class="field">
        <label>邮箱:</label>
        <span>{{ user.user_email }}</span>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name:"UserInformation",
  data() {
    return {
      user: {
        user_name: '',
        user_age: '',
        user_sex: '',
        user_address: '',
        user_postId: '',
        user_email:''
      },
    };
  },
  methdos:{

  }
};
</script>

<style>
.user-profile {
  margin: 20px;
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 5px;
}

.field {
  margin-bottom: 10px;
}

label {
  font-weight: bold;
}

span {
  margin-left: 10px;
}
</style>