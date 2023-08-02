<template>
    <div id="root">
        <!-- 以下所有的标签都是通过子组件引入后 使用的子组件内部里的自定义name标签 -->
        <TiaoZhuan></TiaoZhuan>
        <!-- 接收子组件传入的 数据 @btnClick子组件传入的数据(不变) cpnClcik传入之后，在父组件中使用的自定义方法名-->
        <TiaoZhuanYe @btnClick="cpnClcik"></TiaoZhuanYe>
        <!-- 采用v-if判断该导航按钮是否是索引值所指向的第三个导航 下标为2若是那么返回的是true真则展现，否则隐藏不展示-->
        <YingPingT v-if="cindex == 2"></YingPingT>
        <YingPingS v-if="cindex == 1"></YingPingS>
        <YingPingF v-if="cindex == 0"></YingPingF>
    </div>
</template>
<script>
// 引入子组件数据
import TiaoZhuan from "./Tiaozhuan.vue";
import TiaoZhuanYe from "./Tiaozhuanye.vue";
import YingPingT from "./YingPingT.vue";
import YingPingS from "./YingPingS.vue";
import YingPingF from "./YingPingF.vue";
export default {
    name: "Fu",
    data() {
        return {
            // 随便设置一个变量给定初始值为0
            cindex: 0,
        };
    },
    // 将引入的子组件中name标签也就是引入的文件名 添加该部分包裹后方可使用该标签
    components: {
        TiaoZhuan,
        TiaoZhuanYe,
        YingPingT,
        YingPingS,
        YingPingF
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