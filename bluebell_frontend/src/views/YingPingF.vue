<template>
    <div class="all">
      <div class="top-message">
        <p class="message-text">
          请注意，具体的创新创业项目方向会根据你的兴趣、技能和市场需求来确定。在申请之前，建议你深入研究所选方向的市场、竞争情况以及可行性分析，以确保项目的成功。
        </p>
      </div>
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
                    src: require("../img/img2.jpg"),
                    name: "大学生创业项目一: 技术创新项目\n",
                    title: "涉及新技术、新产品或新服务的研发和实施，例如人工智能、生物技术、清洁能源等领域。",
                    time: "发布时间:2020-11-13",
                    price: "312.00",
                },
                {
                    id: 2,
                    src: require("../img/img5.png"),
                    name: "大学生创业项目二：社会问题解决项目",
                    title: "针对社会问题提出创新解决方案，如教育、健康、环保等领域的项目。",
                    time: "发布时间:2023-07-21",
                    price: "310.00",
                },
                {
                    id: 3,
                    src: require("../img/img6.jpg"),
                    name: "大学生创业项目三：文化创意项目",
                    title: "围绕艺术、文化、设计等方面的创新创业项目，如数字媒体、文化传承等",
                    time: "发布时间:2020-10-01",
                    price: "280.00",
                },
                {
                    id: 4,
                    src: require("../assets/images/avatar.png"),
                    name: "大学生创业项目四：环保可持续项目",
                    title: "致力于环境保护、可持续发展的项目，如循环经济、绿色科技等。",
                    time: "发布时间:2020-11-27",
                    price: "100.00",
                },
                {
                    id: 5,
                    src: require("../assets/images/avatar.png"),
                    name: "大学生创业项目五：创业孵化器项目",
                    title: "建立和运营创业孵化器，支持创业者发展和成长。",
                    time: "发布时间:2023-07-25",
                    price: "835.00",
                },
                {
                    id: 6,
                    src: require("../assets/images/avatar.png"),
                    name: "大学生创业项目六：跨界合作项目",
                    title: "不同领域之间的合作，例如科技与艺术、医疗与技术等，创造出新的商业模式和价值。",
                    time: "发布时间:2020-11-27",
                    price: "300.00",
                },
                {
                    id: 7,
                    src: require("../assets/images/avatar.png"),
                    name: "大学生创业项目七：教育创新项目",
                    title: "提出新的教育模式、在线学习平台、职业培训等方面的创新项目。",
                    time: "发布时间:2020-11-09",
                    price: "859.00",
                },
                //请注意，具体的创新创业项目方向会根据你的兴趣、技能和市场需求来确定。在申请之前，建议你深入研究所选方向的市场、竞争情况以及可行性分析，以确保项目的成功。

            ],
        };
    },
};
</script>
<style scoped>

/* Style for the top message */
.top-message {
  background-color: #f5f5f5;
  border-bottom: 1px solid #ddd;
  padding: 20px 0;
  text-align: center;
  font-size: 30px;
  color: #333;
}

/* Style for the message text within top message */
.message-text {
  max-width: 800px;
  margin: 0 auto;
}

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
