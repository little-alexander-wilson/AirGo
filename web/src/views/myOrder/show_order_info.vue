<template>
  <el-dialog v-model="state.isShowDialog" :title="state.title" width="80%" destroy-on-close align-center>
   <div>
     <el-descriptions
         column="1"
         border
     >
       <el-descriptions-item label="订单编号">{{ orderStoreData.orderManageData.value.currentOrder.out_trade_no }}</el-descriptions-item>
       <el-descriptions-item label="创建时间">{{ DateStrtoTime(orderStoreData.orderManageData.value.currentOrder.created_at) }}</el-descriptions-item>
       <el-descriptions-item label="用户">{{ orderStoreData.orderManageData.value.currentOrder.user_name }}</el-descriptions-item>
       <el-descriptions-item label="商品ID">{{ orderStoreData.orderManageData.value.currentOrder.goods_id }}</el-descriptions-item>
       <el-descriptions-item label="商品类型">
         <el-tag class="ml-2" v-if="orderStoreData.orderManageData.value.currentOrder.goods_type === 'subscribe'">订阅</el-tag>
         <el-tag class="ml-2" v-if="orderStoreData.orderManageData.value.currentOrder.goods_type === 'recharge'">充值</el-tag>
         <el-tag class="ml-2" v-if="orderStoreData.orderManageData.value.currentOrder.goods_type === 'general'">普通商品</el-tag>
        </el-descriptions-item>
       <el-descriptions-item label="商品标题">{{ orderStoreData.orderManageData.value.currentOrder.subject }}</el-descriptions-item>
       <el-descriptions-item label="商品价格">{{ orderStoreData.orderManageData.value.currentOrder.price }}</el-descriptions-item>
       <el-descriptions-item label="订单金额">{{ orderStoreData.orderManageData.value.currentOrder.total_amount }}</el-descriptions-item>
       <el-descriptions-item v-if="orderStoreData.orderManageData.value.currentOrder.deliver_type !== 'none'" label="发货内容">
         <v-md-preview :text="orderStoreData.orderManageData.value.currentOrder.deliver_text"></v-md-preview></el-descriptions-item>
     </el-descriptions>
   </div>
  </el-dialog>
</template>

<script setup lang="ts">

import {reactive} from "vue";
import {useOrderStore} from "/@/stores/orderStore";
import {storeToRefs} from "pinia";
import {DateStrtoTime} from "/@/utils/formatTime"

const orderStore = useOrderStore()
const orderStoreData = storeToRefs(orderStore)

const state = reactive({
  isShowDialog: false,
  title: "订单详情",
})


// 打开弹窗
const openDialog = (row: Order) => {
  state.isShowDialog = true
  orderStoreData.orderManageData.value.currentOrder = row
}
// 关闭弹窗
const closeDialog = () => {
  state.isShowDialog = false
};

//确认提交
const onSubmit=() =>{


}

// 暴露变量
defineExpose({
  openDialog,   // 打开弹窗
});

</script>

<style scoped lang="scss">

</style>