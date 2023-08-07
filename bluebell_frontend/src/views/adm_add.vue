<template>
    <div class="lyear-layout-web">
        <div class="lyear-layout-container">
            <!--左侧导航-->
            <aside class="lyear-layout-sidebar">
                <!-- logo -->
                <div id="logo" class="sidebar-header">
                    <a href="index.html"><img src="../img/2.gif" title="LightYear" alt="LightYear" /></a>
                </div>
                <div class="lyear-layout-sidebar-scroll">
                    <nav class="sidebar-main">
                        <ul class="nav nav-drawer">
                            <li class="nav-item active"> <router-link to="/admindex">
                                    后台首页</router-link> </li>
                            <li class="nav-item nav-item-has-subnav">
                                <a href="javascript:void(0)">功能列表</a>
                            </li>
                        </ul>
                    </nav>
                    <div class="sidebar-footer">
                        <p class="copyright">版权所有 &copy; 2023. <a target="_blank" href="http://lyear.itshubao.com">数据202</a>
                            All rights reserved.</p>
                    </div>
                </div>
            </aside>
            <!--End 左侧导航-->

            <!--头部信息-->
            <header class="lyear-layout-header">
                <nav class="navbar navbar-default">
                    <div class="topbar">
                        <div class="topbar-left">
                            <div class="lyear-aside-toggler">
                                <span class="lyear-toggler-bar"></span>
                                <span class="lyear-toggler-bar"></span>
                                <span class="lyear-toggler-bar"></span>
                            </div>
                            <span class="navbar-page-title"> 功能页面 - 审批项目 </span>
                        </div>

                        <ul class="topbar-right">
                            <li class="dropdown dropdown-profile">
                                <a href="javascript:void(0)" data-toggle="dropdown">
                                    <img class="img-avatar img-avatar-48 m-r-10" src="../img/user.jpg" alt="小波" />
                                    <span>小波 <span class="caret"></span></span>
                                    <button class="button" @click="sendTohome">退出登录</button>
                                </a>


                            </li>
                            <!--切换主题配色-->
                            <li class="dropdown dropdown-skin">
                                <span data-toggle="dropdown" class="icon-palette"></span>
                                <ul class="dropdown-menu dropdown-menu-right" data-stopPropagation="true">
                                    <!-- 这里省略主题配色的部分 -->
                                </ul>
                            </li>
                            <!--切换主题配色-->
                        </ul>
                    </div>
                </nav>
            </header>
            <!--End 头部信息-->

            <!--页面主要内容-->
            <main class="lyear-layout-content">
                <div class="container-fluid">
                    <div class="row">
                        <div class="col-lg-12">
                            <div class="card">
                                <div class="card-body">
                                    <form @submit.prevent="submitForm" class="row">
                                        <div class="form-group col-md-12">
                                            <strong>项目名称：</strong>{{ projectName }}
                                        </div>
                                        <div class="form-group col-md-12">
                                            <label for="type">审批结果</label>
                                            <div class="form-controls">
                                                <select v-model="tickKind" name="tickKind" class="form-control" id="type">
                                                    <option value="同意">同意</option>
                                                    <option value="拒绝">拒绝</option>
                                                </select>
                                            </div>
                                        </div>
                                        <div class="form-group col-md-12">
                                            <label for="title">审批原因</label>
                                            <input v-model="tickReason" type="text" class="form-control" id="title"
                                                name="tickReason" placeholder="请输入审批原因" />
                                        </div>
                                        <div class="form-group col-md-12">
                                            <label for="seo_description">描述</label>
                                            <textarea v-model="context" class="form-control" id="seo_description"
                                                name="context" rows="5" placeholder="描述"></textarea>
                                        </div>
                                        <div class="form-group col-md-12">
                                            <label for="start-time">审批时间:</label>
                                            <input v-model="ticketCreateTime" type="datetime-local" id="start-time"
                                                name="ticketCreateTime">
                                            <br>
                                        </div>
                                        <div class="form-group col-md-12 button-group">
                                            <button type="submit" class="btn btn-primary ajax-post"
                                                target-form="add-form">确定</button>
                                            <button type="button" class="btn btn-default" @click="goBack">返回</button>
                                        </div>
                                    </form>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </main>
            <!--End 页面主要内容-->
        </div>
    </div>
</template>

<script>
export default {
    name: 'admadd',
    data() {
        return {
            projectName: '',
            tickKind: '', // 投票类型
            tickReason: '', // 投票标题
            context: '', // 投票描述
            ticketCreateTime: '', // 开始时间

        };
    },
    created() {
        // 使用$route.params来获取传递的参数
        const projectName = this.$route.params.projectName;
        this.projectName = projectName
        // 在这里处理projectName的逻辑
    },
    methods: {
        // 处理表单提交
        // submitForm() {
        //     const formData = {
        //         tickKind: this.tickKind,
        //         tickName: this.tickName,
        //         context: this.context,
        //         ticketCreateTime: this.ticketCreateTime,
        //         ticketEndTime: this.ticketEndTime,
        //     };
        //     // 发起提交表单请求，处理逻辑
        //     // 示例：
        //     // axios.post('/api/add', formData)
        //     //   .then((response) => {
        //     //     alert(response.data);
        //     //     this.goBack();
        //     //   })
        //     //   .catch((error) => {
        //     //     console.error(error);
        //     //     alert('表单提交失败，请重试。');
        //     //   });
        // },
        // 处理退出登录
        logout() {
            // 处理退出登录逻辑
            // 示例：
            // axios.post('/api/logout')
            //   .then(() => {
            //     // 退出成功，跳转到登录页或其他处理
            //   })
            //   .catch((error) => {
            //     console.error(error);
            //   });
        },
        // 返回上一页
        goBack() {
            this.$router.go(-1);
        },
        sendTohome() {
            this.$router.push({ name: "Login" })
        },
    },
    mounted() {
        // 初始化开始时间和结束时间输入框
        this.submitTimeRange();
    },

};
</script>

<style>
/* 这里省略样式部分 */
.button-group {
    display: flex;
    justify-content: space-evenly;
}
</style>