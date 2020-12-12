<template>

  <div>
    <el-row>
      <el-col :span="4">
        <el-select v-model="storeIndex"
        filterable
        remote
        :remote-method="storeSearch"
        :loading="loading"
        placeholder="选择商家">
            <el-option
              v-for="(item, index) in stores"
              :key="item._id"
              :label="item.storeName"
              :value="index">
            </el-option>
          </el-select>
      </el-col>

      <el-col span="8">
        <el-date-picker
              v-model="timerange"
              type="datetimerange"
              range-separator="至"
              start-placeholder="开始日期"
              end-placeholder="结束日期">
            </el-date-picker>
      </el-col>

      <el-col :span="3">
        <el-input placeholder="用户名" v-model="q" clearable></el-input>
      </el-col>

      <el-col :span="3">
        <el-button @click="search">查询</el-button>
      </el-col>

      <el-col :span="3">
        <el-button @click="clear">重置</el-button>
      </el-col>

    </el-row>
    <el-table
        :data="orders"
        border
        stripe
        show-summary
        style="width: 100%">
       <!-- <el-table-column
          prop="_id"
          label="订单号"
          width="280">
        </el-table-column> -->
        <el-table-column
          prop="storeName"
          label="商家"
          width="200">
        </el-table-column>

        <el-table-column
          prop="realCoupon"
          label="优惠券"
          width="100">
        </el-table-column>

        <el-table-column
          prop="finalAmount"
          label="用户实付"
          width="100">
          <template slot-scope="scope">
            <span style="margin-left: 10px">{{scope.row.payAmount + scope.row.mustPayAmount - scope.row.realCoupon}}</span>
          </template>
        </el-table-column>

        <el-table-column
          prop="totalAmount"
          label="订单金额"
          width="100">
          <template slot-scope="scope">
            <span style="margin-left: 10px">{{scope.row.payAmount + scope.row.mustPayAmount - scope.row.realCoupon}}</span>
          </template>
        </el-table-column>

        <el-table-column
          prop="subsidy"
          label="赏金令"
          width="100">
        </el-table-column>

        <el-table-column
          prop="income7"
          label="平台收入"
          width="100">
        </el-table-column>

        <el-table-column
          prop="merchantIncome"
          label="商家收入"
          width="100">
          <!-- <template slot-scope="scope">
            {{scope.row.totalAmount + scope.row.subsidy - scope.row.income7}}
          </template> -->
        </el-table-column>

        <el-table-column
          prop="userName"
          label="用户名"
          width="200">
        </el-table-column>

        <el-table-column
          prop="timestamp"
          label="下单时间"
          width="180"
          :formatter="dateFormater">
        </el-table-column>

        <!-- <el-table-column
          fixed="right"
          label="操作"
          width="300">
          <template slot-scope="scope">
            <el-button @click="handleClick(scope.row, 1)" type="text" size="small">查看</el-button>
          </template>
        </el-table-column> -->
      </el-table>
      <el-pagination
            @size-change="handleSizeChange"
            @current-change="handleCurrentChange"
            :current-page="currentPage"
            :page-sizes="[10, 20, 30, 50, 100]"
            :page-size="10"
            layout="total, sizes, prev, pager, next, jumper"
            :total="totalCount">
          </el-pagination>
  </div>
</template>

<script>
  import axios from 'axios'
  import moment from 'moment'
  import { Loading } from 'element-ui';
  export default {
    data() {
      return {
        timerange: '',
        q: '',
        orders:[],
        currentPage: 1,
        pageSize: 10,
        totalCount: 0,

        loading: false,
        stores: [],
        storeIndex: '搜索商家',
      }
    },
    methods: {
      handleSizeChange(e) {
        let loadingInstance = Loading.service({ fullscreen: true });
        this.pageSize = e
        this.currentPage = 1
        let data = {
          tp: 'orders',
          q: this.q,
          currentPage: this.currentPage,
          pageSize: this.pageSize,
        }
        if (this.storeIndex != '搜索商家') {
          data.storeId = this.stores[this.storeIndex]._id
        }
        if (this.timerange.length > 1) {
          data.startTime = new Date(this.timerange[0]).getTime()
          data.endTime = new Date(this.timerange[1]).getTime()
        }
        axios.post('http://localhost:8090/proxy',data).then(resp => {
          console.log(resp.data)
          this.setOrders(resp.data.data)
          this.currentPage = resp.data.currentPage
          this.pageSize = resp.data.pageSize
          this.totalCount = resp.data.totalCount
          loadingInstance.close()
        })
      },
      handleCurrentChange(e) {
        let loadingInstance = Loading.service({ fullscreen: true });
        this.currentPage = e
        let data = {
          tp: 'orders',
          q: this.q,
          currentPage: this.currentPage,
          pageSize: this.pageSize,
        }
        if (this.storeIndex != '搜索商家') {
          data.storeId = this.stores[this.storeIndex]._id
        }
        if (this.timerange.length > 1) {
          data.startTime = new Date(this.timerange[0]).getTime()
          data.endTime = new Date(this.timerange[1]).getTime()
        }
        axios.post('http://localhost:8090/proxy',data).then(resp => {
          console.log(resp.data)
          this.setOrders(resp.data.data)
          this.currentPage = resp.data.currentPage
          this.pageSize = resp.data.pageSize
          this.totalCount = resp.data.totalCount
          loadingInstance.close()
        })
      },
      storeSearch(q) {
        this.loading = true
        axios.post('http://localhost:8090/stores',{q: q,currentPage: 1, pageSize: 20}).then(resp => {
          this.stores = resp.data.data
          this.loading = false
        })
      },
      clear() {
        this.storeIndex = '搜索商家'
        this.timerange = []
        this.q = ''
      },
      search() {
        let loadingInstance = Loading.service({ fullscreen: true });
        let data = {
          tp: 'orders',
          q: this.q,
          currentPage: this.currentPage,
          pageSize: this.pageSize,
        }
        if (this.storeIndex != '搜索商家') {
          data.storeId = this.stores[this.storeIndex]._id
        }
        if (this.timerange.length > 1) {
          data.startTime = new Date(this.timerange[0]).getTime()
          data.endTime = new Date(this.timerange[1]).getTime()
        }
        axios.post('http://localhost:8090/proxy', data).then(resp => {
          console.log(resp.data)
          this.setOrders(resp.data.data)
          this.currentPage = resp.data.currentPage
          this.pageSize = resp.data.pageSize
          this.totalCount = resp.data.totalCount
          loadingInstance.close()
        })
      },
      dateFormater: function(row, column) {
        var date = row[column.property];
        return moment(date).format("YYYY-MM-DD HH:mm:ss")
      },
      setOrders(orders) {
        for (var i = 0; i< orders.length; i ++) {
          orders[i].merchantIncome = (orders[i].totalAmount + orders[i].subsidy - orders[i].income7).toFixed(2)
        }
        this.orders = orders
      }
    },
    mounted() {
      let loadingInstance = Loading.service({ fullscreen: true });
      axios.post('http://localhost:8090/proxy',{tp: 'orders' ,q: '',currentPage: this.currentPage, pageSize: this.pageSize}).then(resp => {
        console.log(resp.data)
        this.setOrders(resp.data.data)
        this.currentPage = resp.data.currentPage
        this.pageSize = resp.data.pageSize
        this.totalCount = resp.data.totalCount
        loadingInstance.close()
      })
      // axios.post('http://localhost:8090/stores',{q: this.q,currentPage: 1, pageSize: 10}).then(resp => {
      //   this.stores = resp.data.data
      //   loadingInstance.close()
      // })
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
    margin-top: 10px;
  }

</style>
