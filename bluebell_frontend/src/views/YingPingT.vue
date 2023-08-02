<template>
    <div class="all">
        <ul>
            <li v-for="(item, index) in navs" :key="index">
                <span :class="index + 1 > 3 ? 'geshu1' : 'geshu'">{{ index + 1 }}</span>
                <img :src="item.src" />
                <div class="hang">
                    <p class="pingming">{{ item.name }}</p>
                    <p class="zhuyan">{{ item.title }}</p>
                    <p class="shangying">{{ item.time }}</p>
                    <p class="piao">总票数：<span class="buy">{{ item.price }}</span>万</p>
                    <hr style="margin-top:70px;margin-left:-200px;width:1000px;">
                </div>
            </li>
        </ul>
    </div>
</template>

<script>
export default {
    name: "YingPian",
    data() {
        return {
            // 将所有的数据放入对象中  随后将每个对象为一部分，添加到数组中
            navs: [
                {
                    id: 1,
                    src: require("../assets/images/avatar.png"),
                    name: "2023扩散模型最新技术进展汇总",
                    title: "随着Stable Diffusion和Midjourney等图像生成模型的爆火，今年在模型领域，扩散模型稳占主导地位。因此，与之相关的新技术也层出不穷，短短1个月，扩散模型相关的论文就有上百篇了，可见其发展的火热趋势。这次我整理了一部分最新扩散模型相关论文来和大家分享，有50+ 篇，希望能帮助同学们获取新的思路~",
                    time: "发布时间:2023-07-21",
                    price: "310.00",
                },
                {
                    id: 2,
                    src: require("../assets/images/avatar.png"),
                    name: "2023十大最牛编程语言排行榜以及个语言的优缺点",
                    title: "Tips：TIOBE排行榜 是根据互联网上有经验的程序员、课程和第三方厂商的数量，并使用搜索引擎（如Google、Bing、Yahoo!）以及Wikipedia、Amazon、YouTube和Baidu（百度）统计出排名数据，只是反映某个编程语言的热门程度，并不能说明一门编程语言好不好，或者一门语言所编写的代码数量多少。",
                    time: "发布时间:2023-07-25",
                    price: "835.00",
                },
                {
                    id: 3,
                    src: require("../assets/images/avatar.png"),
                    name: "SystemServer进程",
                    title: "在systemServer启动文章中我们讲了在SystemServer.java的main方法里面调用new SystemServer().run（）方法启动System_server进程。那么我们接着看一下SystemServer.java具体做了哪些事情？",
                    time: "发布时间:2020-11-27",
                    price: "300.00",
                },
                {
                    id: 4,
                    src: require("../assets/images/avatar.png"),
                    name: "【Terraform学习】Terraform-docker部署快速入门（快速入门）",
                    title: "导航到IAM在左侧菜单中，单击角色 。单击创建角色该按钮以创建新的 IAM 角色。在创建角色部分，为角色选择可信实体类型：AWS 服务使用案例:EC2",
                    time: "发布时间:2020-10-01",
                    price: "280.00",
                },
                {
                    id: 5,
                    src: require("../assets/images/avatar.png"),
                    name: "网络安全（黑客）自学笔记",
                    title: "1.计算机基础知识了解了计算机的硬件、软件、操作系统和网络结构等基础知识，可以帮助您更好地理解网络安全的概念和技术。2.网络基础知识了解了网络的结构、协议、服务和安全问题，可以帮助您更好地解决网络安全的原理和技术。3.安全基础知识了解了安全的概念、原则、威吓和攻击方式，可以帮助您更好地理解网络安全的意义和必要性。",
                    time: "发布时间:2020-11-27",
                    price: "100.00",
                },
                {
                    id: 6,
                    src: require("../assets/images/avatar.png"),
                    name: "k8s Label 2",
                    title: "在 k8s 中，我们会轻轻松松的部署几十上百个微服务，这些微服务的版本，副本数的不同进而会带出更多的 pod这么多的 pod ，如何才能高效的将他们组织起来的，如果组织不好便会让管理微服务变得混乱不堪，杂乱无章",
                    time: "发布时间:2020-11-09",
                    price: "859.00",
                },
                {
                    id: 7,
                    src: require("../assets/images/avatar.png"),
                    name: "基于jackson对bean的序列号和反序列化",
                    title: "通过观察控制台输出的SQL发现页面传递过来的员工id的值和数据库中的id值不一致，这是怎么回事呢?分页查询时服务端响应给页面的数据中id的值为19位数字，类型为long页面中js处理long型数字只能精确到前16位，所以最终通过ajax请求提交给服务端的时候id就改变了前面我们已经发现了问题的原因，即js对long型数据进行处理时丢失精度，导致提交的id和数据库中的id不一致。",
                    time: "发布时间:2020-11-13",
                    price: "723.00",
                },
                {
                    id: 8,
                    src: require("../assets/images/avatar.png"),
                    name: "j2ee相关知识点",
                    title: "浏览器栏中，输入的是servlet的mapping映射，请求到servlet中去，jsp路径，会跳转到对应的页面Servlet接口位于最顶端，GenericServlet实现了Servlet，HttpServlet继承了GenericServlet浏览器中访问Servlet映射时，会去调用service()方法，并且会在适当的时候调用doGet",
                    time: "发布时间:2020-11-13",
                    price: "648.00",
                },
                {
                    id: 9,
                    src: require("../assets/images/avatar.png"),
                    name: "Docker consul的容器服务更新与发现",
                    title: "consul服务器 192.168.44.100 运行consul服务、nginx服务、consul-template守护进程registrator服务器 192.168.44.101 运行registrator容器、运行nginx容器",
                    time: "发布时间:2020-11-30",
                    price: "319.00",
                },
                {
                    id: 10,
                    src: require("../assets/images/avatar.png"),
                    name: "微服务 & 云原生：搭建 K8S 集群",
                    title: "基础环境搭建完成后，令当前虚拟机做 Master 主机，利用虚拟机平台，比如 VMWare、VirtualBox 等的克隆功能，克隆一台 Worker 主机（当然也可以克隆多台）。注意，克隆后的多台主机要保证 ip 不同，不同软件的操作不同，具体操作此处此处不做赘述。",
                    time: "发布时间:2020-11-13",
                    price: "312.00",
                },
            ],
        };
    },
};
</script>
<style scoped>
.all {
    width: 100%;
}
/* ul样式 */
ul {
    width: 810px;
    height: 100%;
    /* margin:auto auto; */
}

/* ul内部的所有li样式 */
li {
    width: 1200px;
    height: 280px;
    position: relative;
    list-style: none;
    box-shadow: 0 0 #b8f6ff;
    /* 采用了动画移动效果 */
    transition: all 1s;
}

/* 设置每一个li鼠标移入的效果 */
li:hover {
    box-shadow: 15px 20px 20px rgba(200, 200, 200, .4);
}

/* 所有的li左侧的序列号前三个采用 */
.geshu {
    display: block;
    width: 40px;
    height: 40px;
    text-align: center;
    line-height: 40px;
    background-color: orange;
    position: relative;
    left: 30px;
    top: 110px;
    font-style: italic;
    color: #fff;
}

/* 所有的li左侧的序列号除前三个外 采用 */
.geshu1 {
    display: block;
    width: 40px;
    height: 40px;
    text-align: center;
    line-height: 40px;
    background-color: gray;
    position: relative;
    left: 30px;
    top: 110px;
    font-style: italic;
    color: #fff;
}

/* 设置图片的位置 采用绝对定位 */
img {
    position: absolute;
    left: 100px;
    margin-top: 10PX;
    height: 100px;
    width: 100px;
}

/* 对右侧包裹的div进行设置样式 */
.hang {
    width: 600px;
    height: 100%;
    position: absolute;
    left: 300px;
    top: 70px;
}

/* 对票房部分设置样式 */
.piao {
    width: 600px;
    height: 100%;
    display: block;
    position: absolute;
    left: 600px;
    top: 40px;
    color: red;
}

/* 对票房内部的数据（播放数值）设置样式 */
.buy {
    font-size: 25px;
}

/* 影片名称 */
.pingming {
    font-size: 26px;
}

/* 上映时间 */
.shangying {
    color: #999;
}
</style>
