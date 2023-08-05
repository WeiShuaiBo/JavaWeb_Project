<template>
    <div class="all">
        <div class="first">
            <p class="lingshi">
                近年来，随着零食市场需求量的不断上涨，零食店成为了万人青睐的创业项目。市场需求量大，而且零食店的加盟也是很火的。休闲食品是大学生的专属食品，符合他们的消费和购买习惯。低热量零食店的出现弥补了市场上比较大的缺陷，满足大学生的需求。因此，大学生创新创业项目中的零食店是一个具有广阔市场前景的项目。
            </p>
        </div>
        <ul>
            <li v-for="(item, index) in navs" :key="index">
                <span :class="index + 1 > 3 ? 'geshu1' : 'geshu'">{{ index + 1 }}</span>
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
    name: "YingPian1",
    data() {
        return {
            // 将所有的数据放入对象中  随后将每个对象为一部分，添加到数组中
            navs: [
                {
                    id: 1,
                    name: "适合大学生创业的项目有哪些(精选9篇)",
                    link: 'http://www.gdyjs.com/chuangye/145749.html',
                    title: "甜食小吃店开甜食小吃店相对比较简单但由于开的人多，竞争就相对的激烈。甜食小吃店比较适合女性开。由于这类店的特点决定了顾客以女性为主，所以在选址时要开在女性来往较多的地方, 譬如临近商店、女子学校或工厂, 以及街口附近。开店前, 首先要分析顾客的身 份，如家庭主妇、女学生或女职员等。",
                    time: "发布时间:2022-11-13",
                    price: "312.00",
                },
                {
                    id: 2,
                    name: "适合大学生的创业项目推荐(精选8篇)",
                    link: 'https://www.wcqjyw.com/wendang/qitafanwen/188294.html',
                    title: "零食店近年来随着休闲食品市场需求量的不断上涨加之本小利大的优势,零食店成为了万人青睐的创业项目。众所周知休闲食品主要的目标消费群体为年轻人, 因此对于大学毕业生来说开一家零食店是一 个非常不错的创业方向。很多人倾向于这类商品的消费和购买, 因此这种零食店的加盟也是很火的。",
                    time: "发布时间:2023-07-21",
                    price: "310.00",
                },
                {
                    id: 3,
                    name: "大学生创业创新项目有",
                    link: 'https://jiameng.baidu.com/content/detail/106962826846?from=search&rid=1.24.27.89',
                    title: "零食店近年来随着休闲食品市场需求量的不断上涨,加之本小利大的优垫，零食店成为了万人青睐的创业项目, 众所周知休闲食品主要的目标消费群体为年轻人, 因此对于大学毕业生来说开一家零食店是一个非常不错的创业方向。很多人倾向于这类商品的消费和购买因此这种零食店的加盟也是很火的。",
                    time: "发布时间:2020-10-01",
                    price: "280.00",
                },
                {
                    id: 4,
                    name: "大学生创新创业项目有哪些项目-百度知了好学",
                    link: 'https://jiameng.baidu.com/content/detail/106962826846?from=search&rid=1.24.27.89',
                    title: "低热量零食店零食对于小孩还有女孩子都是他们专属的东西,相信很多人都会有吃这些东西的习惯，可是零食里面的热量含量高, 只要多吃的话, 就会造成发胖的现象, 那么现在就是需要含量低的东西。弥补了市场当中比较大的缺陷。目前的市场需求还是在不断的扩大当中, 他们主打健康, 而且在很多方面比较的便利, 也是大学...",
                    time: "发布时间:2020-11-27",
                    price: "100.00",
                },
                {
                    id: 5,
                    name: "大学生创业创新项目3篇",
                    link: 'https://localsite.baidu.com/okam/pages/article/index?categoryLv1=%E6%95%99%E8%82%B2%E5%9F%B9%E8%AE%AD&ch=54&srcid=10004&strategyId=127502045499766&source=natural',
                    title: "低热量零食店零食对于小孩还有女孩子都是他们专属的东西,相信很多人都会有吃这些东西的习惯，可是零食里面的热量含量高, 只要多吃的话, 就会造成发胖的现象, 那么现在就是需要含量低的东西。弥补了市场当中比较大的缺陷。目前的市场需求还是在不断的扩大当中, 他们主打健康, 而且在很多方面比较的便利, 也是大学...",
                    time: "发布时间:2023-07-25",
                    price: "835.00",
                },
                {
                    id: 6,
                    name: "适合大学生创业的几个项目",
                    link: 'http://wen.51tietu.net/gushi/1303253.html',
                    title: "小吃店或者可以开一家小吃店,可以开在学校内,也可以选择走出学校开在学校外面,例如商圈附近或者是居民楼附近都是可以的。",
                    time: "发布时间:2020-11-27",
                    price: "300.00",
                },
                {
                    id: 7,
                    name: "8个大学生创新创业项目",
                    link: 'http://www.ddnx.com/zhichang/889872.html',
                    title: "大学生创新创业项目:小吃店大学生如何快速赚钱?选对行业选对创业项目,其实赚钱很容易!今年1 -9月，我国小吃业的营业收入达到了14737.1亿元, 同比增长达到16.5%, 占社会消费品零售总额比重11.3%。有没有快速致富好项目?现在小吃行业的发展前景十分广阔投资者在创业时, 选择经营一家小吃加盟店可以获得好的..",
                    time: "发布时间:2020-11-09",
                    price: "859.00",
                },
                {
                    id: 8,
                    name: "大学生创新创业项目哪些好(集锦8篇)",
                    link: 'https://www.gerenjianli.cn/fwdq/shiyongwen/10332667.html',
                    title: "通过观察控制台输出的SQL发现页面传递过来的员工id的值和数据库中的id值不一致，这是怎么回事呢?分页查询时服务端响应给页面的数据中id的值为19位数字，类型为long页面中js处理long型数字只能精确到前16位，所以最终通过ajax请求提交给服务端的时候id就改变了前面我们已经发现了问题的原因，即js对long型数据进行处理时丢失精度，导致提交的id和数据库中的id不一致。",
                    time: "发布时间:2020-11-13",
                    price: "723.00",
                },
                {
                    id: 9,
                    name: "大学生的创业项目有哪些_百度知道",
                    link: 'https://zhidao.baidu.com/question/1804164917568011987.html?device=mobile&ssid=0&from=0&uid=0&pu=usm%400%2Csz%40224_220%2Cta%40sf___24_114.0&bd_page_type=1&baiduid=5F9014AE4050DF91F50204E22D26BF90&tj=tc',
                    title: "炸鸡小吃店金黄香脆的外皮鲜嫩多汁的鸡肉,满口香辣咸麻,掀起了-股全民吃炸鸡的热浪。对于大学生来说, 不管是在学校附近还是在休闲娱乐地段, 开一家炸鸡小吃店都是相当棒的选择。开一家炸鸡小吃店的投资也不多, 5-7万元即可开业。大学生资金不足, 可以和三五亲朋好友合伙开店, 各自分: 工店面运营的相关事宜。",
                    time: "发布时间:2020-11-30",
                    price: "319.00",
                },
                {
                    id: 10,
                    name: "大学生创新创业项目汇总",
                    link: 'http://www.cnrencai.com/goldjob/xiangmu/760028.html',
                    title: "炸鸡小吃店金黄香脆的外皮鲜嫩多汁的鸡肉,满口香辣咸麻,掀起了-股全民吃炸鸡的热浪。对于大学生来说, 不管是在学校附近还是在休闲娱乐地段, 开一家炸鸡小吃店都是相当棒的选择。开一家炸鸡小吃店的投资也不多, 5-7万元即可开业。大学生资金不足, 可以和三五亲朋好友合伙开店, 各自分: 工店面运营的相关事宜。",
                    time: "发布时间:2020-11-13",
                    price: "648.00",
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
