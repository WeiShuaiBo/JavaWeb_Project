<template>
    <div>
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
                                <li class="nav-item active"> <a href="/index">后台首页</a> </li>
                                <li class="nav-item nav-item-has-subnav">
                                    <a href="javascript:void(0)"> 功能列表</a>
                                    <ul class="nav nav-subnav">
                                        <li> <a href="/list">投票信息</a> </li>
                                        <li> <a href="/add">新增投票</a> </li>
                                        <li> <a href="/step">表单向导</a> </li>
                                    </ul>
                                </li>
                            </ul>
                        </nav>
                        <div class="sidebar-footer">
                            <p class="copyright">版权所有 &copy; 2019. <a target="_blank"
                                    href="http://lyear.itshubao.com">IT书包</a> 保留所有权利。</p>
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
                                <span class="navbar-page-title"> 信息列表 </span>
                            </div>
                            <span class="navbar-page-title">剩余投票数: {{ status }}票</span>
                            <div class="topbar-right">
                                <ul class="dropdown dropdown-profile">
                                    <li>
                                        <img class="img-avatar img-avatar-48 m-r-10" src="../img/user.jpg" alt="笔下光年" />
                                        <span>小波 <span class="caret"></span></span>
                                        <button class="button" @click="sendTohome">退出登录</button>
                                        <ul class="dropdown-menu dropdown-menu-right">
                                            <li><a href="user.html">个人信息</a></li>
                                            <li><a href="edit.html"> 修改密码</a></li>
                                            <li><a href="javascript:void(0)">清空缓存</a></li>
                                            <li><a @click="logout">退出登录</a></li>
                                        </ul>
                                    </li>
                                </ul>
                            </div>
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
                                    <div class="card-toolbar clearfix">
                                        <div class="pull-right search-bar" method="get" action="#!" role="form">
                                            <div class="input-group">
                                                <div class="input-group-btn">
                                                    <input type="hidden" name="search_field" id="search-field"
                                                        value="title">
                                                    <button class="btn btn-default dropdown-toggle" id="search-btn"
                                                        data-toggle="dropdown" type="button" aria-haspopup="true"
                                                        aria-expanded="false">
                                                        标题 <span class="caret"></span>
                                                    </button>
                                                    <ul class="dropdown-menu">
                                                        <li> <a tabindex="-1" href="javascript:void(0)"
                                                                @click="setSearchField('title')">类型</a> </li>
                                                        <li> <a tabindex="-1" href="javascript:void(0)"
                                                                @click="setSearchField('cat_name')">名字</a> </li>
                                                    </ul>
                                                </div>
                                                <input type="text" class="form-control" v-model="keyword"
                                                    placeholder="请输入名称">
                                            </div>
                                        </div>

                                    </div>
                                    <div class="card-body">
                                        <div class="table-responsive">
                                            <table class="table table-bordered">
                                                <thead>
                                                    <tr>
                                                        <th>
                                                            <label class="lyear-checkbox checkbox-primary">
                                                                <input type="checkbox" v-model="selectAll"><span></span>
                                                            </label>
                                                        </th>
                                                        <th>编号</th>
                                                        <th>类型</th>
                                                        <th>名字</th>
                                                        <th>内容</th>
                                                        <th>发起时间</th>
                                                        <th>截止时间</th>
                                                        <th>剩余时间</th>
                                                        <th>投票票数</th>
                                                        <th>操作</th>
                                                    </tr>
                                                </thead>
                                                <tbody>
                                                    <tr v-for="item in res" :key="item.Id">
                                                        <td>
                                                            <label class="lyear-checkbox checkbox-primary">
                                                                <input type="checkbox" v-model="selectedItems"
                                                                    :value="item.Id"><span></span>
                                                            </label>
                                                        </td>
                                                        <td>{{ item.Id }}</td>
                                                        <td>{{ item.TicketKind }}</td>
                                                        <td>{{ item.TickName }}</td>
                                                        <td>{{ item.Context }}</td>
                                                        <td>{{ item.TicketCreateTime }}</td>
                                                        <td>{{ item.TicketEndTime }}</td>
                                                        <td>{{ item.CountdownFormatted }}</td>
                                                        <td>{{ item.TicketNum }}</td>
                                                        <td>
                                                            <button type="button" @click="submitForm(item.Id)">投票</button>
                                                        </td>
                                                    </tr>
                                                </tbody>
                                            </table>
                                        </div>
                                        <ul class="pagination">
                                            <li class="disabled"><span>«</span></li>
                                            <li class="active"><span>1</span></li>
                                            <li><a href="#1">2</a></li>
                                            <li><a href="#1">3</a></li>
                                            <li><a href="#!">»</a></li>
                                        </ul>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </main>
                <!--End 页面主要内容-->
            </div>
        </div>
    </div>
</template>

<script>
import axios from 'axios';
// import "../../public/static/css/bootstrap.min.css";
// import "../../public/static/css/materialdesignicons.min.css";
// import "../../public/static/css/style.min.css";
// import "../../public/static/css/eye.css";
export default {
    name: 'admindex',
    data() {
        return {
            status: 0, // 剩余投票数
            res: [], // 投票信息列表
            keyword: '', // 搜索关键词
            selectAll: false, // 是否全选
            selectedItems: [], // 已选中的投票项
            showHeadBar: false, // 控制是否显示 HeadBar 组件
        };
    },
    methods: {
        // 设置搜索字段
        setSearchField(field) {
            document.getElementById('search-field').value = field;
            document.getElementById('search-btn').innerText = field === 'title' ? '标题' : '名字';
        },
        // 处理表单提交
        submitForm(itemId) {
            const form = document.getElementById('form-' + itemId);
            const formData = new FormData(form);
            axios.post('/submit', formData)
                .then((response) => {
                    alert(response.data);
                    this.refreshData();
                })
                .catch((error) => {
                    console.error(error);
                    alert('表单提交失败，请重试。');
                });
        },
        sendTohome() {
            this.$router.push({ name: "SignUp" })
        },
        // 处理删除按钮点击
        handleDelete() {
            if (this.selectedItems.length === 0) {
                alert('请勾选要删除的记录');
                return;
            }

            axios.post('/delete', { ids: this.selectedItems })
                .then(() => {
                    alert('删除成功');
                    this.refreshData();
                    this.selectedItems = [];
                })
                .catch((error) => {
                    console.error(error);
                    alert('删除失败，请稍后重试。');
                });
        },
        // 刷新数据
        refreshData() {
            // 从服务器获取数据并更新res和status
            // 示例：
            // axios.get('/api/votes')
            //   .then((response) => {
            //     this.res = response.data;
            //     this.status = response.data.status;
            //   })
            //   .catch((error) => {
            //     console.error(error);
            //   });
        },
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
        // 更新倒计时
        updateCountdown() {
            this.res.forEach((item) => {
                if (item.Countdown > 0) {
                    item.Countdown -= 1000;
                    const days = Math.floor(item.Countdown / (24 * 60 * 60 * 1000));
                    const hours = Math.floor((item.Countdown % (24 * 60 * 60 * 1000)) / (60 * 60 * 1000));
                    const minutes = Math.floor((item.Countdown % (60 * 60 * 1000)) / (60 * 1000));
                    const seconds = Math.floor((item.Countdown % (60 * 1000)) / 1000);
                    item.CountdownFormatted = `${days}天${hours}时${minutes}分${seconds}秒`;
                } else {
                    item.CountdownFormatted = '已结束';
                }
            });
        },
    },
    mounted() {
        this.refreshData();
        setInterval(this.updateCountdown, 1000);
    },
};
</script>
<style>
.button {
    width: 60px;
    height: 40px;
    background-color: aquamarine;
}
</style>