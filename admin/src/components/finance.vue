<template>
  <div>

    <el-tabs v-model="activeName" type="border-card" @tab-click="handleClick">
      <el-tab-pane label="汇总" name="1">
        <el-divider content-position="left">收支</el-divider>
        <el-row type="flex" justify="start">
          <el-col :span="4">
            <el-tag type="success">本月收入：{{financialsummary.m.in}} 元</el-tag>
          </el-col>
          <el-col :span="4">
            <el-tag type="danger">本月支出：{{financialsummary.m.out}} 元</el-tag>
          </el-col>

          <el-col :span="4">
            <el-tag type="success">今日收入：{{financialsummary.d.in}} 元</el-tag>
          </el-col>
          <el-col :span="4">
            <el-tag type="danger">今日支出：{{financialsummary.d.out}} 元</el-tag>
          </el-col>
        </el-row>

        <el-divider content-position="left">订单</el-divider>
        <el-row type="flex" justify="start">
          <el-col :span="4">
            <el-tag type="success">本月订单数：{{financialsummary.m.orders}}</el-tag>
          </el-col>
          <el-col :span="4">
            <el-tag type="success">今日订单数：{{financialsummary.d.orders}}</el-tag>
          </el-col>
        </el-row>

        <el-divider content-position="left">用户</el-divider>
        <el-row type="flex" justify="start">
          <el-col :span="4">
            <el-tag type="success">本月新增用户数：{{financialsummary.m.users}} 个</el-tag>
          </el-col>
          <el-col :span="4">
            <el-tag type="success">今日新增用户数：{{financialsummary.d.users}} 个</el-tag>
          </el-col>
        </el-row>

        <el-divider content-position="left">优惠券</el-divider>
        <el-row type="flex" justify="start">
          <el-col :span="5">
            <el-tag type="danger">本月优惠券使用：{{financialsummary.m.coupon}} 元</el-tag>
          </el-col>
          <el-col :span="5">
            <el-tag type="danger">今日优惠券使用：{{financialsummary.d.coupon}} 元</el-tag>
          </el-col>
        </el-row>

        <el-divider content-position="left">赏金令</el-divider>
        <el-row type="flex" justify="start">
          <el-col :span="5">
            <el-tag type="danger">本月商家补贴：{{financialsummary.m.subsidy}} 元</el-tag>
          </el-col>
          <el-col :span="5">
            <el-tag type="danger">今日商家补贴：{{financialsummary.d.subsidy}} 元</el-tag>
          </el-col>
        </el-row>


      </el-tab-pane>
      <el-tab-pane label="收支历史" name="2">
        <el-table
          :data="financialinouthis">
          <el-table-column
            prop="_id"
            label="日期"
            width="180"
            :formatter="dateFormater">
          </el-table-column>
          <el-table-column
            prop="in"
            label="收入"
            width="180">
            <template slot-scope="scope">
              {{scope.row.in | rounding}}
            </template>
          </el-table-column>
          <el-table-column
            prop="out"
            label="支出"
            width="180">
            <template slot-scope="scope">
              {{scope.row.out | rounding}}
            </template>
          </el-table-column>
        </el-table>
        <el-pagination
              @current-change="handleCurrentChangeInout"
              :current-page="currentPage"
              :page-sizes="[10]"
              :page-size="10"
              layout="total, sizes, prev, pager, next, jumper"
              :total="totalCount">
            </el-pagination>
      </el-tab-pane>

      <el-tab-pane label="优惠券记录" name="3">
        <el-table
          :data="couponhis"
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
            width="180">
          </el-table-column>
          <el-table-column
            prop="coupon"
            label="优惠券面额"
            width="100">
          </el-table-column>
          <el-table-column
            prop="storeName"
            label="使用门店"
            width="280">
          </el-table-column>
        </el-table>
        <el-pagination
              @current-change="handleCurrentChangeCoupon"
              :current-page="currentPage"
              :page-sizes="[10]"
              :page-size="10"
              layout="total, sizes, prev, pager, next, jumper"
              :total="totalCount">
            </el-pagination>
      </el-tab-pane>

      <el-tab-pane label="补贴记录" name="4">
        <el-table
          :data="subsidyhis"
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
            width="280">
          </el-table-column>
          <el-table-column
            prop="userName"
            label="新用户名"
            width="180">
          </el-table-column>
          <el-table-column
            prop="subsidy"
            label="补贴金额"
            width="00">
          </el-table-column>
        </el-table>
        <el-pagination
              @current-change="handleCurrentChangeSubsidy"
              :current-page="currentPage"
              :page-sizes="[10]"
              :page-size="10"
              layout="total, sizes, prev, pager, next, jumper"
              :total="totalCount">
            </el-pagination>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script>
  import axios from 'axios'
  import moment from 'moment'
  import { Loading } from 'element-ui';
  export default {
    data() {
      return {
        activeName: '1',
        financialsummary: {},

        currentPage: 1,
        pageSize: 10,
        totalCount: 0,
        financialinouthis: [],
        couponhis: [],
        subsidyhis: []
      }
    },
    filters:{
      rounding (value) {
       return value.toFixed(2)
       }
    },
    methods: {
      handleCurrentChangeInout(e) {
        let loadingInstance = Loading.service({ fullscreen: true });
        this.currentPage = e
        axios.post('http://localhost:8090/proxy', {tp: 'financialinouthis',currentPage: this.currentPage, pageSize: this.pageSize}).then(resp => {
          console.log(resp.data)
          this.financialinouthis = resp.data.data
          this.currentPage = resp.data.currentPage
          this.pageSize = resp.data.pageSize
          this.totalCount = resp.data.totalCount
          loadingInstance.close()
        })
      },
      handleCurrentChangeCoupon(e) {
        let loadingInstance = Loading.service({ fullscreen: true });
        this.currentPage = e
        axios.post('http://localhost:8090/proxy', {tp: 'couponhis',currentPage: this.currentPage, pageSize: this.pageSize}).then(resp => {
          console.log(resp.data)
          this.couponhis = resp.data.data
          this.currentPage = resp.data.currentPage
          this.pageSize = resp.data.pageSize
          this.totalCount = resp.data.totalCount
          loadingInstance.close()
        })
      },
      handleCurrentChangeSubsidy(e) {
        let loadingInstance = Loading.service({ fullscreen: true });
        this.currentPage = e
        axios.post('http://localhost:8090/proxy', {tp: 'subsidyhis',currentPage: this.currentPage, pageSize: this.pageSize}).then(resp => {
          console.log(resp.data)
          this.subsidyhis = resp.data.data
          this.currentPage = resp.data.currentPage
          this.pageSize = resp.data.pageSize
          this.totalCount = resp.data.totalCount
          loadingInstance.close()
        })
      },
      handleClick() {
        this.currentPage = 1
        let loadingInstance = Loading.service({ fullscreen: true });
        switch (this.activeName) {
          case "1":
            axios.post('http://localhost:8090/proxy', {tp: 'financialsummary'}).then(resp => {
              this.financialsummary = resp.data
              console.log(resp.data)
              loadingInstance.close()
            })
          break;
          case "2":
            axios.post('http://localhost:8090/proxy', {tp: 'financialinouthis',currentPage: this.currentPage, pageSize: this.pageSize}).then(resp => {
              console.log(resp.data)
              this.financialinouthis = resp.data.data
              this.currentPage = resp.data.currentPage
              this.pageSize = resp.data.pageSize
              this.totalCount = resp.data.totalCount
              loadingInstance.close()
            })
          break;
          case "3":
            axios.post('http://localhost:8090/proxy', {tp: 'couponhis',currentPage: this.currentPage, pageSize: this.pageSize}).then(resp => {
              console.log(resp.data)
              this.couponhis = resp.data.data
              this.currentPage = resp.data.currentPage
              this.pageSize = resp.data.pageSize
              this.totalCount = resp.data.totalCount
              loadingInstance.close()
            })
          break;
          case "4":
            axios.post('http://localhost:8090/proxy', {tp: 'subsidyhis',currentPage: this.currentPage, pageSize: this.pageSize}).then(resp => {
              console.log(resp.data)
              this.subsidyhis = resp.data.data
              this.currentPage = resp.data.currentPage
              this.pageSize = resp.data.pageSize
              this.totalCount = resp.data.totalCount
              loadingInstance.close()
            })
          break;
        }
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
      let loadingInstance = Loading.service({ fullscreen: true });
      axios.post('http://localhost:8090/proxy', {tp: 'financialsummary'}).then(resp => {
        this.financialsummary = resp.data
        console.log(resp.data)
        loadingInstance.close()
      })
    }
  }
</script>

<style>
  .el-row {
    margin-bottom: 20px;

    &:last-child {
      margin-bottom: 0;
    }
  }

  .el-col {
    border-radius: 4px;
  }

  .title {
    margin-top: 5px;
  }
</style>
