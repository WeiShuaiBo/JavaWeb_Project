<template>
    <div>
        <!-- 利用a标签的点击事件  遍历对象内部的数据 这里遍历的是第二个切换栏部分 随后给定一个点击事件,这里的点击事件传两个参数
    一个参数是下面的的methods方法内部会采用排他的思想,对每个导航部分进行设置样式,第二个参数传入的是index下标,主要是未了子传父

    从而使子组件内部的点击事件在父组件中也会起到作用,那样点击某个部分的按钮就会出现某个部分的页面 -->

        <!-- 随后中间采用的是插值语法将数据插入进去 -->
        <a href="javascript:;" v-for="(item, index) in this.navs" :key="index" @click="change($event, index)">{{ item }}
        </a>
    </div>
</template>

<script>

export default {
    name: "TiaoZhuanYe1",
    data() {
        return {
            // 这里是对数据进行归总在一个数组内部 从而通过上方对该数组的遍历,将内容渲染到页面中
            navs: ["零食店", "网上开店", "奶茶铺"],
        };
    },
    methods: {
        // 在a标签中有解析过两个参数的作用
        change(e, index) {
            // 利用遍历对该导航进行排他思想
            for (const i in this.navs) {
                // 筛选出当前父级中的某个自己的类为空,意思是不设置样式
                e.target.parentNode.children[i].className = "";
            }
            // 通过上面的遍历排他思想,下面便是将不符合条件的(反之就是符合条件,我们所点击的那一部分) 设置对当前元素进行添加样式
            //classList.add()这个是添加样式的方法 内部添加的是类名
            e.target.classList.add("tablet");
            //  下面是子传父  传给父组件的是下标index  第一个是别名(设置的名)  第二个是传入的index数据
            this.$emit("btnClick", index);

        }

    },

};

</script>

<style scoped>
div {

    font-size: 19px;

    width: 700px;

    height: 100%;

    margin-left: 100px;

    margin-top: 30px;

    padding: 0px 0px 20px 0px;

    border-bottom: 2px solid #eee;

}

a {

    width: 100px;

    height: 100px;

    color: #000;

    text-decoration: none;

    margin-left: 30px;

    padding: 0px 0px 19px 0px;

}

.tablet {

    color: orangered;

    border-bottom: 3px solid orangered;

}
</style>
