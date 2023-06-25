<template>
    <div class="display-data-page">
        <h2>已提交申请信息</h2>
        <table class="data-table">
            <tr>
                <th>姓名</th>
                <th>学校</th>
                <th>学院</th>
                <th>专业</th>
                <th>邮箱</th>
                <th>手机号码</th>
                <th>项目创意</th>
            </tr>
            <tr v-for="form in formData" v-bind:key="form" >
                <td>{{ form.text }}</td>
                <!-- <td>{{ university }}</td>
                <td>{{ college }}</td>
                <td>{{ major }}</td>
                <td>{{ formData.email }}</td>
                <td>{{ formData.phone }}</td>
                <td>{{ formData.projectIdea }}</td> -->
            </tr>
        </table>
    </div>
</template>

<script>
export default {
    name: 'DisplayDataPage',
    props: ['formData']
};
</script>

<style scoped>
.data-table {
    width: 100%;
    border-collapse: collapse;
}

.data-table th,
.data-table td {
    padding: 10px;
    text-align: left;
    border-bottom: 1px solid #ddd;
}
</style>