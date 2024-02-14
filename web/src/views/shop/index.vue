<template>
  <div style="margin-left: 15px; margin-right: 15px;">
    <div class="home-card-one mb15" style="margin-top: 15px;">
      <el-radio-group
        v-model="state.goods_type"
        size="large"
        border
        @change="getAllEnabledGoods"
      >
        <el-radio-button label="subscribe">订阅</el-radio-button>
        <el-radio-button label="general">普通商品</el-radio-button>
        <el-radio-button label="recharge">充值卡</el-radio-button>
      </el-radio-group>
    </div>

    <dev>
      <el-row :gutter="15" class="home-card-one">
        <el-col
          :xs="24"
          :sm="12"
          :md="12"
          :lg="8"
          :xl="8"
          v-for="(v, k) in goodsList"
          :key="k"
        >
          <div class="home-card-item">
            <el-card>
              <!-- <template #header>
                <div>
                  <el-text class="card-header-left">{{ v.subject }}</el-text>
                </div>
              </template> -->
              <!--            订阅商品开始-->
              <div v-if="v.goods_type === 'subscribe'">
              </div>
              <!--            订阅商品结束-->
              <!--            充值商品开始-->
              <div v-if="v.goods_type === 'recharge'">
                <div class="card-text">
                  <el-tag class="card-text-left">充值金额</el-tag>
                  <span class="card-text-right">{{ v.recharge_amount }}</span>
                </div>
              </div>
              <!--            充值商品结束-->
              <!-- <div class="card-text">
                <el-tag class="card-text-left" type="warning">价格</el-tag>
                <span class="card-text-right">¥{{ v.total_amount }}</span>
              </div> -->
              <div v-html="v.des"></div>
              <div>
                
                <div style="display: flex; justify-content: center; align-items: center;">
                  <el-button
                    size="large"
                    @click="openSubmitOrderDialog(v)"
                    color="#FF2E63"
                    >立即购买</el-button
                  >
                </div>
              </div>
            </el-card>
          </div>
        </el-col>
      </el-row>
      <!--引入提交订单弹窗-->
      <SubmitOrderDialog
        ref="SubmitOrderDialogRef"
        @openPurchaseDialog="openPurchaseDialog"
      ></SubmitOrderDialog>
      <!-- 引入确认支付弹窗组件 -->
      <PurchaseDialog
        ref="PurchaseDialogRef"
        @openQRDialog="openQRDialog"
      ></PurchaseDialog>
      <!-- 引入二维码弹窗 -->
      <QRDialog ref="QRDialogRef"></QRDialog>
    </dev>
  </div>
</template>

<script setup lang="ts">
import { defineAsyncComponent, onMounted, reactive, ref } from "vue";
import { storeToRefs } from "pinia";
import { useShopStore } from "/@/stores/shopStore";

const shopStore = useShopStore();
const { goodsList, shopData } = storeToRefs(shopStore);
const SubmitOrderDialog = defineAsyncComponent(
  () => import("/@/views/shop/dialog_submit_order.vue")
);
const PurchaseDialog = defineAsyncComponent(
  () => import("/@/views/shop/dialog_purchase.vue")
);
const QRDialog = defineAsyncComponent(
  () => import("/@/views/shop/dialog_QR.vue")
);
const PurchaseDialogRef = ref();
const QRDialogRef = ref();
const SubmitOrderDialogRef = ref();
const state = reactive({
  goods_type: "subscribe",
});
const getAllEnabledGoods = () => {
  shopStore.getAllEnabledGoods({ goods_type: state.goods_type });
};
//加载时获取全部已启用商品
onMounted(() => {
  getAllEnabledGoods();
});
//打开提交订单弹窗
const openSubmitOrderDialog = (goood: Goods) => {
  shopData.value.currentGoods = goood;
  SubmitOrderDialogRef.value.openDialog();
};
//打开确认购买弹窗
const openPurchaseDialog = (goods: Goods) => {
  PurchaseDialogRef.value.openDialog();
};
//打开二维码弹窗
const openQRDialog = () => {
  QRDialogRef.value.openDialog();
};
</script>

<style scoped>
/* 
.card-show-area {
  margin-left: 20px;
} */

.home-card-item {
  width: 100%;
  height: 100%;
  border-radius: 10px;
  transition: all ease 0.3s;
  padding: 0px;
  overflow: hidden;
  background: var(--el-color-white);
  color: var(--el-text-color-primary);
  border: 2px solid var(--next-border-color-light);
}

.el-card {
  background-repeat: no-repeat;
  background-position: 100%, 100%;
  padding: 40px;
}

.card-text {
  display: flex;
  justify-content: space-between;
  height: 35px;
}

.card-text-left {
  margin-top: auto;
  margin-bottom: auto;
}

.card-text-right {
  margin-top: auto;
  margin-bottom: auto;
  font-size: 20px;
}

.card-header-left {
  /* font-size: 20px;
  font-weight: bold;
  font-family: "Microsoft YaHei", sans-serif;
  color: #D72323; */
  color: black; font-size: 24px; font-family: Inter; font-weight: 900; word-wrap: break-word;
}
</style>
