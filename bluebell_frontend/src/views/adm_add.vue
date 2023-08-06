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
                            <li class="nav-item active"> <router-link to="/index"><i class="mdi mdi-home"></i>
                                    后台首页</router-link> </li>
                            <li class="nav-item nav-item-has-subnav">
                                <a href="javascript:void(0)"><i class="mdi mdi-file-outline"></i> 功能列表</a>
                                <ul class="nav nav-subnav">
                                    <li> <router-link to="/list">投票信息</router-link> </li>
                                    <li> <router-link to="/add">发起投票</router-link> </li>
                                    <li> <router-link to="/step">表单向导</router-link> </li>
                                </ul>
                            </li>
                        </ul>
                    </nav>
                    <div class="sidebar-footer">
                        <p class="copyright">版权所有 &copy; 2023. <a target="_blank" href="http://lyear.itshubao.com">haha</a>
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
                            <span class="navbar-page-title"> 功能页面 - 发起投票活动 </span>
                        </div>

                        <ul class="topbar-right">
                            <li class="dropdown dropdown-profile">
                                <a href="javascript:void(0)" data-toggle="dropdown">
                                    <img class="img-avatar img-avatar-48 m-r-10" src="../img/user.jpg" alt="小波" />
                                    <span>小波 <span class="caret"></span></span>
                                </a>
                                <ul class="dropdown-menu dropdown-menu-right">
                                    <li> <router-link to="/user"><i class="mdi mdi-account"></i> 个人信息</router-link> </li>
                                    <li> <router-link to="/edit"><i class="mdi mdi-lock-outline"></i> 修改密码</router-link>
                                    </li>
                                    <li> <a href="javascript:void(0)"><i class="mdi mdi-delete"></i> 清空缓存</a></li>
                                    <li class="divider"></li>
                                    <li> <a @click="logout" href="#"><i class="mdi mdi-logout-variant"></i> 退出登录</a> </li>
                                </ul>
                            </li>
                            <!--切换主题配色-->
                            <li class="dropdown dropdown-skin">
                                <span data-toggle="dropdown" class="icon-palette"><i class="mdi mdi-palette"></i></span>
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
                                            <label for="type">投票类型</label>
                                            <div class="form-controls">
                                                <select v-model="tickKind" name="tickKind" class="form-control" id="type">
                                                    <option value="音乐">音乐</option>
                                                    <option value="书籍">书籍</option>
                                                    <option value="运动">运动</option>
                                                    <option value="游戏">游戏</option>
                                                </select>
                                            </div>
                                        </div>
                                        <div class="form-group col-md-12">
                                            <label for="title">标题</label>
                                            <input v-model="tickName" type="text" class="form-control" id="title"
                                                name="tickName" placeholder="请输入标题" />
                                        </div>
                                        <div class="form-group col-md-12">
                                            <label for="seo_description">描述</label>
                                            <textarea v-model="context" class="form-control" id="seo_description"
                                                name="context" rows="5" placeholder="描述"></textarea>
                                        </div>
                                        <div class="form-group col-md-12">
                                            <label for="start-time">开始时间:</label>
                                            <input v-model="ticketCreateTime" type="datetime-local" id="start-time"
                                                name="ticketCreateTime">
                                            <br>
                                        </div>
                                        <div class="form-group col-md-12">
                                            <label for="end-time">结束时间:</label>
                                            <input v-model="ticketEndTime" type="datetime-local" id="end-time"
                                                name="ticketEndTime">
                                        </div>
                                        <div class="form-group col-md-12">
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
            status: 0, // 剩余投票数
            tickKind: '音乐', // 投票类型
            tickName: '', // 投票标题
            context: '', // 投票描述
            ticketCreateTime: '', // 开始时间
            ticketEndTime: '', // 结束时间
        };
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
    },
    mounted() {
        // 初始化开始时间和结束时间输入框
        this.submitTimeRange();
    },

};
</script>

<style>
/* 这里省略样式部分 */
</style>