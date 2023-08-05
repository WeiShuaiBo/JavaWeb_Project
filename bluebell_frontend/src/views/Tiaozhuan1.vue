<template>
    <div>
        <div id="deoo">
            <!-- 使用ul包裹li进行浮动 -->
            <ul class="tell">
                <li><img src="../assets/images/avatar.png" /></li>
                <li>热门项目榜单</li>
                <li>榜单</li>
            </ul>
        </div>
    </div>
</template>

<script >

export default {
    name: 'TiaoZhuan1',
    data() {
        return {
        }
    },
};

</script>

<style scoped>
/* scoped 指的是当前的css样式不会受到其它组件内部的css样式污染 */
/* 给定外层的div设置一个盒子大小 */
#deoo {
    width: 800px;
    height: 100%;
}

img {
    width: 50px;
    height: 50px;
}

/* 对ul内部的内容进行添加样式 */

.tell {
    width: 700px;
    margin: auto;
    padding: 50px;
    border-bottom: 2px solid #000;
}

/* ul内部li的样式 */
.tell li {
    float: left;
    list-style: none;
}

/* 设置ul内部的指定为第二个li样式 */

.tell li:nth-of-type(2) {
    font-size: 24px;
    margin-top: 10px;
    margin-left: 30px;
}

/* 设置ul内部的指定为第三个li样式 */
.tell li:nth-of-type(3) {
    position: relative;
    left: 400px;
    top: 9px;
    height: 100%;
    font-size: 19px;
    padding: 0px 0px 15px 0px;
    color: orangered;
    border-bottom: 4px solid orangered;
}
</style>
