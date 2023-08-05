<template>
    <div class="all">
        <div class="first">
            <p class="lingshi">
                网上开店对成本的要求低，占用资金少，经营方式灵活，不受时间、地域等因素的限制，客户范围广泛。网上开店成本低，回收快，适合小商店和个人创业。网上开店所需的启动资金少，网上商店比同等规模的地面商店要低得多，不需要增加租金和延长营业时间。网上开店是简便、实惠的创业模式，只要创新和持之以恒，成功就在眼前。
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
    name: "YingPingS1",
    data() {
        return {
            // 将所有的数据放入对象中  随后将每个对象为一部分，添加到数组中
            navs: [
                {
                    id: 1,
                    name: "高校大学生的创业创新项目3篇",
                    link: 'http://w.51tietu.net/wenku/811977.html',
                    title: "行业分析网上开店对成本的要求低它占用资金少,也基本不需要水、电、管理费等方面的支出;这样就解决了大学生没有创业资金的问题; 网店经营方式灵活, 基本不受时间、地域等因素的限制, 只要有时间、能上网就能正常营业, 而且只要少量存货就能开店, 所以也可以随时更换经营其他商品; 网店的客户范围十分广, 只要是..",
                    time: "发布时间:2020-11-27",
                    price: "100.00",
                },
                {
                    id: 2,
                    name: "大学生创新创业项目汇总",
                    link: 'https://www.cnrencai.com/goldjob/xiangmu/760028.html',
                    title: "网上开店网上创业、网上购物已经不是新鲜的事情开网店的人群比例也在增加，无论网店店主有多少,对于缺少创业资金又有创业想法的人来讲。在网上开店都是一个不错的选择这种简便、实惠的创业模式还将在继续延续下去，要想成功只要创新和持之以恒就可以。",
                    time: "发布时间:2020-11-27",
                    price: "300.00",
                },
                {
                    id: 3,
                    name: "大学生创新创业项目计划书(精选84篇)",
                    link: 'https://www.cnrencai.com/goldjob/xiangmu/760028.html',
                    title: "网上开店网上创业、网上购物已经不是新鲜的事情开网店的人群比例也在增加，无论网店店主有多少,对于缺少创业资金又有创业想法的人来讲。在网上开店都是一个不错的选择这种简便、实惠的创业模式还将在继续延续下去，要想成功只要创新和持之以恒就可以。",
                    time: "发布时间:2020-11-13",
                    price: "723.00",
                },
                {
                    id: 4,
                    name: "大学生创新创业项目(推荐8篇)",
                    link: 'https://www.cnrencai.com/goldjob/xiangmu/760028.html',
                    title: "市场环境网上开店,好处多多,最主要的有以几个突出的优势: (一)投资少,回收快,并且无所谓存货，所以特别适合小商店和个人在网上创业。(1)从启动资金来看, 传统的地面开店，门]面租金加装修费还有首批进货资金,少说也需要几万元的启动资金。而网上开店所需的启动资金却少得多按照我们的计划, 你只要有一台能够...",
                    time: "发布时间:2020-10-01",
                    price: "280.00",
                },
                {
                    id: 5,
                    name: "大学生创新创业项目计划书(精选16篇)",
                    link: 'https://www.cnrencai.com/goldjob/xiangmu/760028.html',
                    title: "部分创新及其特色(线上与线下结合)网上开店,好处多多,最主要的有以下几个突出的优势: (- -)投资少,回收快, 而且无所谓存货, 所以特别适合小商店和个人在网上创业。(1)从启动资金 来看传统的地面开店，门面租金加装修费, 还有首批进货资金，少说也需要几万元的启动资金。而网上开店所需的启动资金却少得多按照我...",
                    time: "发布时间:2023-07-21",
                    price: "310.00",
                },
                {
                    id: 6,
                    name: "最新大学生创新创业项目计划书(精选5篇)",
                    link: 'https://www.cnrencai.com/goldjob/xiangmu/760028.html',
                    title: "网上开店网上创业、网_上购物已经不是新鲜的事情开网店的人群比例也在增加，无论网店店主有多少，对于缺少创业资金又有创业想法的人来讲。在网上开店都是一个不错的选择, 这种简便、实惠的创业模式还将在继续延续下去, 要想成功只要创新和持之以恒就可以。",
                    time: "发布时间:2023-07-25",
                    price: "835.00",
                },
                {
                    id: 7,
                    name: "大学生创新创业项目计划书6篇",
                    link: 'https://www.cnrencai.com/goldjob/xiangmu/760028.html',
                    title: "大学生创新创业项目:网上开店,一分耕耘一分收获从全球网上经营的情况看,以下几类较易取得成功:",
                    time: "发布时间:2020-11-09",
                    price: "859.00",
                },
                {
                    id: 8,
                    name: "大学生创新创业项目集锦",
                    link: 'https://www.cnrencai.com/goldjob/xiangmu/760028.html',
                    title: "市场环境网上开店,好处多多,最主要的有以几个突出的优势: (一)投资少,回收快,并且无所谓存货，所以特别适合小商店和个人在网上创业。(1)从启动资金来看, 传统的地面开店，门]面租金加装修费还有首批进货资金,少说也需要几万元的启动资金。而网上开店所需的启动资金却少得多按照我们的计划, 你只要有一台能够...",
                    time: "发布时间:2020-11-13",
                    price: "312.00",
                },
                {
                    id: 9,
                    name: "最新大学生创新创业项目计划书_百度文库",
                    link: 'https://www.cnrencai.com/goldjob/xiangmu/760028.html',
                    title: "市场环境网上开店,好处多多,最主要的有以几个突出的优势: (一)投资少,回收快,并且无所谓存货，所以特别适合小商店和个人在网上创业。(1)从启动资金来看, 传统的地面开店，门]面租金加装修费还有首批进货资金,少说也需要几万元的启动资金。而网上开店所需的启动资金却少得多按照我们的计划, 你只要有一台能够...",
                    time: "发布时间:2020-11-13",
                    price: "648.00",
                },
                {
                    id: 10,
                    name: "适合大学生的创新创业项目——知乎",
                    link: 'https://www.cnrencai.com/goldjob/xiangmu/760028.html',
                    title: "市场环境网上开店,好处多多,最主要的有以几个突出的优势: (一)投资少,回收快,并且无所谓存货，所以特别适合小商店和个人在网上创业。(1)从启动资金来看, 传统的地面开店，门]面租金加装修费还有首批进货资金,少说也需要几万元的启动资金。而网上开店所需的启动资金却少得多按照我们的计划, 你只要有一台能够...",
                    time: "发布时间:2020-11-30",
                    price: "319.00",
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