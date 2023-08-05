<template>
    <div id="root">
        <!-- 以下所有的标签都是通过子组件引入后 使用的子组件内部里的自定义name标签 -->
        <TiaoZhuan1></TiaoZhuan1>
        <!-- 接收子组件传入的 数据 @btnClick子组件传入的数据(不变) cpnClcik传入之后，在父组件中使用的自定义方法名-->
        <TiaoZhuanYe1 @btnClick="cpnClcik"></TiaoZhuanYe1>
        <!-- 采用v-if判断该导航按钮是否是索引值所指向的第三个导航 下标为2若是那么返回的是true真则展现，否则隐藏不展示-->
        <YingPingT1 v-if="cindex == 2"></YingPingT1>
        <YingPingS1 v-if="cindex == 1"></YingPingS1>
        <YingPingF1 v-if="cindex == 0"></YingPingF1>
    </div>
</template>
<script>
// 引入子组件数据
import TiaoZhuan1 from "./Tiaozhuan1.vue";
import TiaoZhuanYe1 from "./Tiaozhuanye1.vue";
import YingPingT1 from "./YingPingT1.vue";
import YingPingS1 from "./YingPingS1.vue";
import YingPingF1 from "./YingPingF1.vue";
export default {
    name: "Fu1",
    data() {
        return {
            // 随便设置一个变量给定初始值为0
            cindex: 0,
        };
    },
    // 将引入的子组件中name标签也就是引入的文件名 添加该部分包裹后方可使用该标签
    components: {
        TiaoZhuan1,
        TiaoZhuanYe1,
        YingPingT1,
        YingPingS1,
        YingPingF1
    },
    // 添加一个方法
    methods: {
        // 这里的数据是通过子组件中的二级切换栏部分传入的数据 采用方法，将index下标传给父组件
        cpnClcik(index) {
            // 将传入的index数据传给当前父组件return中所设置的变量
            this.cindex = index;
        },
    },
};

</script>