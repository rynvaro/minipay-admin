<template>
  <div>

    <el-row>
      <el-col :span="12">
        <el-tabs v-model="activeName" type="card" @tab-click="handleClick">
          <el-tab-pane label="本月总流水" name="first">
            <el-table
              :data="financial.cashflow.month"
              style="width: 35%">
              <el-table-column
                prop="in"
                label="收入"
                width="100">
              </el-table-column>
              <el-table-column
                prop="out"
                label="支出"
                width="100">
              </el-table-column>
            </el-table>
          </el-tab-pane>
          <el-tab-pane label="今日流水" name="second">
            <el-table
              :data="financial.cashflow.day"
              style="width: 35%">
              <el-table-column
                prop="in"
                label="收入"
                width="100">
              </el-table-column>
              <el-table-column
                prop="out"
                label="支出"
                width="100">
              </el-table-column>
            </el-table>
          </el-tab-pane>
          <el-tab-pane label="历史流水" name="third">
            <el-table
              :data="financial.cashflow.histories"
              style="width: 70%">
              <el-table-column
                prop="_id"
                label="日期"
                width="180"
                :formatter="dateFormater">
              </el-table-column>
              <el-table-column
                prop="in"
                label="收入"
                width="100">
              </el-table-column>
              <el-table-column
                prop="out"
                label="支出"
                width="100">
              </el-table-column>
            </el-table>
          </el-tab-pane>
        </el-tabs>
      </el-col>
      <el-col :span="12">
        <el-tabs v-model="activeNameOrder" type="card" @tab-click="handleClick">
          <el-tab-pane label="订单" name="first">
            <el-table
              :data="financial.order.statistics"
              style="width: 60%">
              <el-table-column
                prop="month"
                label="本月订单数"
                width="180">
              </el-table-column>
              <el-table-column
                prop="day"
                label="今日订单数"
                width="180">
              </el-table-column>
            </el-table>
          </el-tab-pane>
          <el-tab-pane label="历史" name="third">
            <el-table
              :data="financial.order.histories"
              style="width: 90%">
              <el-table-column
                prop="timestamp"
                label="日期"
                width="120"
                :formatter="dateFormater">
              </el-table-column>
              <el-table-column
                prop="userName"
                label="用户名"
                width="100">
              </el-table-column>
              <el-table-column
                prop="finalAmount"
                label="消费金额"
                width="100">
              </el-table-column>
              <el-table-column
                prop="storeName"
                label="门店"
                width="100">
              </el-table-column>
              <el-table-column
                prop="coupon"
                label="优惠券使用"
                width="100">
              </el-table-column>
            </el-table>
          </el-tab-pane>
        </el-tabs>
      </el-col>
    </el-row>

    <el-row>
      <el-col :span="12">
        <el-tabs v-model="activeNameCoupon" type="card" @tab-click="handleClick">
          <el-tab-pane label="优惠券" name="first">
            <el-table
              :data="financial.coupon.statistics"
              style="width: 40%">
              <el-table-column
                prop="month"
                label="本月总计"
                width="100">
              </el-table-column>
              <el-table-column
                prop="day"
                label="今日使用"
                width="100">
              </el-table-column>
            </el-table>
          </el-tab-pane>
          <el-tab-pane label="历史" name="third">
            <el-table
              :data="financial.coupon.histories"
              style="width: 80%">
              <el-table-column
                prop="timestamp"
                label="日期"
                width="120"
                :formatter="dateFormater">
              </el-table-column>
              <el-table-column
                prop="userName"
                label="用户名"
                width="100">
              </el-table-column>
              <el-table-column
                prop="coupon"
                label="优惠券面额"
                width="100">
              </el-table-column>
              <el-table-column
                prop="storeName"
                label="使用门店"
                width="100">
              </el-table-column>
            </el-table>
          </el-tab-pane>
        </el-tabs>

      </el-col>
      <el-col :span="12">
        <el-tabs v-model="activeNameSubsidy" type="card" @tab-click="handleClick">
          <el-tab-pane label="商家补贴" name="first">
            <el-table
              :data="financial.subsidy.statistics"
              style="width: 40%">
              <el-table-column
                prop="month"
                label="本月总补贴"
                width="100">
              </el-table-column>
              <el-table-column
                prop="day"
                label="今日补贴"
                width="100">
              </el-table-column>
            </el-table>
          </el-tab-pane>
          <el-tab-pane label="历史" name="third">
            <el-table
              :data="financial.subsidy.histories"
              style="width: 90%">
              <el-table-column
                prop="timestamp"
                label="日期"
                width="120"
                :formatter="dateFormater">
              </el-table-column>
              <el-table-column
                prop="storeName"
                label="商家门店"
                width="100">
              </el-table-column>
              <el-table-column
                prop="userName"
                label="新用户名"
                width="100">
              </el-table-column>
              <el-table-column
                prop="subsidy"
                label="补贴金额"
                width="100">
              </el-table-column>
            </el-table>
          </el-tab-pane>
        </el-tabs>
      </el-col>
    </el-row>

    <el-tabs v-model="activeNameUser" type="card" @tab-click="handleClick">
      <el-tab-pane label="新增用户" name="first">
        <el-table
          :data="financial.user.statistics"
          style="width: 30%">
          <el-table-column
            prop="month"
            label="本月新增"
            width="180">
          </el-table-column>
          <el-table-column
            prop="day"
            label="今日新增"
            width="180">
          </el-table-column>
        </el-table>
      </el-tab-pane>
      <el-tab-pane label="历史" name="third">
        <el-table
          :data="financial.user.histories"
          style="width: 80%">
          <el-table-column
            prop="createdAt"
            label="日期"
            width="180"
            :formatter="dateFormater">
          </el-table-column>
          <el-table-column
            prop="data.name"
            label="用户名"
            width="180">
          </el-table-column>
          <el-table-column
            prop="data.firstPayStoreName"
            label="首单消费门店"
            width="180">
          </el-table-column>
          <el-table-column
            prop="data.payAmount"
            label="消费金额"
            width="180">
          </el-table-column>
        </el-table>
      </el-tab-pane>
    </el-tabs>


  </div>
</template>

<script>
  import axios from 'axios'
  import moment from 'moment'
  export default {
    data() {
      return {
        activeName: 'first',
        activeNameOrder: 'first',
        activeNameCoupon: 'first',
        activeNameSubsidy: 'first',
        activeNameUser: 'first',
        financial: {},
        stores: [],
      }
    },
    methods: {
      handleClick() {

      },
      dateFormater: function(row, column) {
        var date = row[column.property];
        if (date == 0) {
          return ''
        }
        return moment(date).format("YYYY-MM-DD")
      },
    },
    mounted() {
      let data = {
        tp: 'financial'
      }
      axios.post('http://localhost:8090/proxy', data).then(resp => {
        this.financial = resp.data
        console.log(resp.data)
      })
    }
  }
</script>

<style>
</style>
