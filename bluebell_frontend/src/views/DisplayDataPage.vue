html
<template>
    <div class="display-data-page">
        <h2>已提交申请信息</h2>
        <table class="data-table">
            <thead>
                <tr>
                    <th>姓名</th>
                    <th>学校</th>
                    <th>学院</th>
                    <th>专业</th>
                    <th>邮箱</th>
                    <th>手机号码</th>
                    <th>项目创意</th>
                </tr>
            </thead>
            <tbody>
                <tr v-if="formData">
                    <td>{{ formData.name }}</td>
                    <td>{{ formData.university }}</td>
                    <td>{{ formData.college }}</td>
                    <td>{{ formData.major }}</td>
                    <td>{{ formData.email }}</td>
                    <td>{{ formData.phone }}</td>
                    <td>{{ formData.projectIdea }}</td>
                </tr>
            </tbody>
        </table>
    </div>
</template>

<script>
export default {
    name: 'DisplayDataPage',
    computed: {
        formData() {
            return this.$store.state.formData
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
    width: 100%;
    border-collapse: collapse;
    font-size: 14px;
}

.data-table th,
.data-table td {
    padding: 10px;
    text-align: left;
    border-bottom: 1px solid #ddd;
}

.data-table th {
    background-color: #f2f2f2;
}
</style>