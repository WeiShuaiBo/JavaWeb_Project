<template>
    <div class="all">
        <div class="first">
            <p class="lingshi">
                奶茶铺是适合大学生创业的项目。奶茶铺是小规模店铺，但市场规模惊人。投资少、见效快、收入稳定、风险小、无技术难度。奶茶店的成本低、回报快、获利高，适合大学毕业生加入创业好项目。奶茶行业投资小，无需太大的门面，装修也没有过多要求。开一家奶茶店在学校里面，吸引很多学生。
            </p>
        </div>
        <ul>
            <li v-for="(item, index) in navs" :key="index">
                <span :class="index + 1 > 3 ? 'geshu1' : 'geshu'">{{ index + 1 }}</span>
                <img :src="item.src" />
                <div class="hang">
                    <p class="pingming">
                        <a :href="item.link" target="_blank">{{ item.name }}</a>
                    </p>
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
    name: "YingPing1",
    data() {
        return {
            // 将所有的数据放入对象中  随后将每个对象为一部分，添加到数组中
            navs: [
                {
                    id: 1,
                    name: "适合大学生的创新创业项目推荐(精选8篇)",
                    link: 'https://www.wcqjyw.com/wendang/qitafanwen/188294.html',
                    title: "奶茶店现在的饮品加盟店做的越来越火,不断有不同的品牌走向市场。因为人们生活水平的提高,使得人们对于相关消费的追去越来越高。而加盟奶茶店的的成本低、回报快、获利高, 正是适合当今大学毕业生加入的创业好项目。",
                    time: "发布时间:2023-07-21",
                    price: "310.00",
                },
                {
                    id: 2,
                    name: "大学生创新创业项目有哪些项目-百度知了好学",
                    link: 'https://localsite.baidu.com/okam/pages/article/index?categoryLv1=%E6%95%99%E8%82%B2%E5%9F%B9%E8%AE%AD&ch=54&srcid=10004&strategyId=127502045499766&source=natural',
                    title: "奶茶店国内奶茶饮食的传统和历史源远流长,奶茶已经成了人们生活中不可或缺的一部分。相比很多创业项目来说奶茶行业的投资是非常小的。无需太大的门面, 装修方面也没有过多要求，奶茶成本也不高。要说大学生做什么生意好?开什么店赚钱投资小?非奶茶店莫属。",
                    time: "发布时间:2020-11-13",
                    price: "312.00",
                },
                {
                    id: 3,
                    name: "适合大学生创业的几个项目",
                    link: 'http://www.ddnx.com/zhichang/889872.html',
                    title: "奶茶店如果你将你的目标消费群体定在学生群体,你可以开一家奶茶店在学校里面现在年轻人都非常喜欢喝奶茶, 开一家奶茶店会吸弓| 很多学生的。",
                    time: "发布时间:2020-11-27",
                    price: "300.00",
                },
                {
                    id: 4,
                    name: "8个大学生创新创业项目",
                    link: 'https://tongxiehui.net/bm/62616bd27eaa4.html',
                    title: "大学生创业项目:奶茶铺时下年轻人最钟爱的休闲饮品非珍珠奶茶莫属。目前，奶茶铺大多都是以外卖形式经营的小型店铺, 虽然店铺规模不大, 但市场规模惊人。投资分析: 奶茶铺具有投入少、见效快、收入稳定、风险小、无技术难度等特点。奶茶铺的门面一般10至15平方米即可, 月租金约2000 ~5000元。开店所需的..",
                    time: "发布时间:2020-10-01",
                    price: "280.00",
                },
                {
                    id: 5,
                    name: "8个大学生创新创业项目",
                    link: 'https://tongxiehui.net/bm/62616bd27eaa4.html',
                    title: "大学生创业项目:奶茶铺时下年轻人最钟爱的休闲饮品非珍珠奶茶莫属。目前，奶茶铺大多都是以外卖形式经营的小型店铺, 虽然店铺规模不大, 但市场规模惊人。投资分析: 奶茶铺具有投入少、见效快、收入稳定、风险小、无技术难度等特点。奶茶铺的门面一般10至15平方米即可, 月租金约2000 ~5000元。开店所需的..",
                    time: "发布时间:2020-10-01",
                    price: "100.00",
                },
                {
                    id: 6,
                    name: "大学生创新创业项目汇总",
                    link: 'http://www.pincai.com/article/2537331.htm',
                    title: "大学生创新创业项目:奶茶店国内奶茶饮食的传统和历史源远流长,奶茶已经成了人们生活中不可或缺的一部分。相比很多创业项目来说奶茶行业的投资是非常小的。无需太大的i门面, 装修方面也没有过多要求，奶茶成本也不高。要说大学生做什么生意好?开什么店最赚钱投资小?非奶茶店莫属。",
                    time: "发布时间:2020-11-09",
                    price: "859.00",
                },
                {
                    id: 7,
                    name: "大学生创业项目方案(精选11篇)",
                    link: 'http://www.yjbys.com/qiuzhizhinan/show-582010.html',
                    title: "市场调查奶茶店就是一个资金投入低、 消费人群广、回收获本快，而且i ]面非常好找的创业项目,正常除了保存3个月左右的店租、人工和日常开销外，奶茶店经营治理不必太多周转金, 非常适合小本自主创业。奶茶, 发展趋势很好, 而且有非常宏大的潜力。",
                    time: "发布时间:2020-11-13",
                    price: "723.00",
                },
                {
                    id: 8,
                    name: "关于大学生创业项目介绍(通用12篇)",
                    link: 'https://www.sohu.com/a/128431161_572178',
                    title: "大学生创新创业项目奶茶店国内奶茶饮食的传统和历史源远流长，奶茶已经成了人们生活中不可或缺的一部分。相比很多创业项目来说奶茶行业的投资是非常小的。无需太大的]面, 装修方面也没有过多要求，奶茶成本也不高。要说大学生做什么生意好?开什么店最赚钱投资小?非奶茶店莫属。",
                    time: "发布时间:2020-11-13",
                    price: "648.00",
                },
                {
                    id: 9,
                    name: "大学生创业创新项目有哪些",
                    link: 'https://www.sohu.com/a/128431161_572178',
                    title: "适合大学生创业的项目3大学生创业项目:奶茶铺时下年轻人最钟爱的休闲饮品非珍珠奶茶莫属。目前，奶茶铺大多都是以外卖形式经营的小型店铺, 虽然店铺规模不大, 但市场规模惊人。投资分析: 奶茶铺具有投入少、见效快、收入稳定、风险小、无技术难度等特点。奶茶铺的门面一般10至1 5平方米即可, 月租金约20...",
                    time: "发布时间:2020-11-30",
                    price: "319.00",
                },
                {
                    id: 10,
                    name: "精选8个大学生创新创业项目，总有一个满意的",
                    link: 'https://www.sohu.com/a/128431161_572178',
                    title: "大学生创新创业项目:奶茶店国内奶茶饮食的传统和历史源远流长，奶茶已经成了人们生活中不可或缺的一部分。相比很多创业项目来说，奶茶行业的投资是非常小的。无需太大的i ]面, 装修方面也没有过多要求，奶茶成本也不高。要说大学生做什么生意好?开什么店最赚钱投资小?非奶茶店莫属。",
                    time: "发布时间:2023-07-25",
                    price: "835.00",
                },
            ],
        };
    },
};
</script>
<style scoped>
.all {
    width: 100%;
    position: relative;
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

.first {
    position: absolute;
    top: -110px;
    /* 调整距离ul的上方距离 */
    left: 120px;
    width: 600px;
    height: 100%;
    bottom: 300px;
}


.lingshi {
    font-size: 17px;
    margin: 0;
    padding: 10px 5px;
    border-radius: 24px;
    background-color: #fff;
}

ul {
    margin-top: 100px;
    /* 调整ul与上方内容之间的距离 */
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


/* 对右侧包裹的div进行设置样式 */
.hang {
    width: 600px;
    height: 100%;
    position: absolute;
    left: 120px;
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
