<template>
  <div>
    <el-row>
        <el-col :span="6">
          <div>本月提现：{{month}}</div>
        </el-col>
        <el-col :span="6">
          <div>今日提现：{{day}}</div>
        </el-col>
    </el-row>
    <el-row style="margin-top: 20px;">
      <el-col :span="24">
         <div>提现申请</div>
      </el-col>
    </el-row>
    <el-table
        :data="withdraws"
        border
        style="width: 100%; margin-top: 10px;">
        <el-table-column
          prop="publishedAt"
          label="申请日期"
          width="120"
          :formatter="dateFormater">
        </el-table-column>
        <el-table-column
          prop="storeName"
          label="提现商家"
          width="180">
        </el-table-column>
        <el-table-column
          prop="merchantName"
          label="商家姓名"
          width="100">
        </el-table-column>

        <el-table-column
          prop="bankCard"
          label="收款账户"
          width="200">
        </el-table-column>

        <el-table-column
          prop="bank"
          label="开户行"
          width="100">
        </el-table-column>

        <el-table-column
          prop="withdrawAmount"
          label="提现金额"
          width="100">
        </el-table-column>
        <el-table-column
          prop="status"
          label="当前状态"
          width="150">
          <template slot-scope="scope">
            <span style="margin-left: 10px" v-if="scope.row.status === 0" >待审核</span>
            <span style="margin-left: 10px" v-if="scope.row.status === 1" >审核中</span>
            <span style="margin-left: 10px" v-if="scope.row.status === 2" >付款中</span>
            <span style="margin-left: 10px" v-if="scope.row.status === 3" >已付款</span>
            <span style="margin-left: 10px" v-if="scope.row.status === 4" >已撤销</span>
          </template>
        </el-table-column>
        <el-table-column
          fixed="right"
          label="操作"
          width="300">
          <template slot-scope="scope">
            <el-button @click="handleClick(scope.row)" type="text" size="small">更新状态</el-button>
          </template>
        </el-table-column>
      </el-table>

      <el-row style="margin-top: 20px;">
        <el-col :span="24">
           <div>提现历史</div>
        </el-col>
      </el-row>
      <el-row style="margin-top: 20px;">
        <el-col :span="10">
          <el-input placeholder="输入商家姓名" v-model="q" @change="search" clearable></el-input>
        </el-col>
        <el-col :span="3">
          <el-button @click="search">查询</el-button>
        </el-col>
      </el-row>

      <el-table
          :data="histories"
          border
          show-summary
          style="width: 100%; margin-top: 10px;">
          <el-table-column
            prop="publishedAt"
            label="申请日期"
            width="120"
            :formatter="dateFormater">
          </el-table-column>
          <el-table-column
            prop="storeName"
            label="提现商家"
            width="180">
          </el-table-column>
          <el-table-column
            prop="merchantName"
            label="商家姓名"
            width="100">
          </el-table-column>

          <el-table-column
            prop="bankCard"
            label="收款账户"
            width="200">
          </el-table-column>

          <el-table-column
            prop="bank"
            label="开户行"
            width="100">
          </el-table-column>

          <el-table-column
            prop="withdrawAmount"
            label="提现金额"
            width="100">
          </el-table-column>
          <el-table-column
            prop="status"
            label="当前状态"
            width="150">
            <template slot-scope="scope">
              <span style="margin-left: 10px" v-if="scope.row.status === 0" >待审核</span>
              <span style="margin-left: 10px" v-if="scope.row.status === 1" >审核中</span>
              <span style="margin-left: 10px" v-if="scope.row.status === 2" >付款中</span>
              <span style="margin-left: 10px" v-if="scope.row.status === 3" >已付款</span>
              <span style="margin-left: 10px" v-if="scope.row.status === 4" >已撤销</span>
            </template>
          </el-table-column>
          <el-table-column
            fixed="right"
            label="操作"
            width="300">
            <template slot-scope="scope">
              <el-button @click="handleClick(scope.row, 1)" type="text" size="small">更新状态</el-button>
            </template>
          </el-table-column>
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

        <el-dialog title="新增板块" :visible.sync="dialogFormVisible">

          <el-row>
            <el-col :span="6">
              <div class="title" style="text-align: center;">更新状态到：</div>
            </el-col>
            <el-col :span="6">
              <el-select v-model="status" filterable placeholder="跳转连接">
                  <el-option
                    v-for="item in statuslist"
                    :key="item.id"
                    :label="item.label"
                    :value="item.value">
                  </el-option>
                </el-select>
            </el-col>
          </el-row>
          <el-row style="margin-top: 20px;">
            <el-col :span="6">
              <div class="title" style="text-align: center;">实际付款金额：</div>
            </el-col>
            <el-col :span="6">
              <el-input placeholder="最后状态需要填写" v-model="realWithdrawAmount" clearable></el-input>
            </el-col>
          </el-row>
          <div slot="footer" class="dialog-footer">
            <el-button @click="dialogFormVisible = false">取 消</el-button>
            <el-button type="primary" @click="doUpdate">确 定</el-button>
          </div>
        </el-dialog>
  </div>
</template>

<script>
  import axios from 'axios'
  import moment from 'moment'
  import { Loading } from 'element-ui';
  export default {
    data() {
      return {
        q: '',
        dialogFormVisible: false,
        withdraws: [],
        histories: [],
        status: 0,
        realWithdrawAmount: 0,
        currentId: '',
        month: 0,
        day: 0,
        currentPage: 1,
        pageSize: 10,
        totalCount: 0,
        statuslist: [
          {
            value: 0,
            label: '待审核'
          },
          {
            value: 1,
            label: '审核中'
          },
          {
            value: 2,
            label: '付款中'
          },
          {
            value: 3,
            label: '已付款'
          },
        ],
      }
    },
    methods: {
      dateFormater: function(row, column) {
        var date = row[column.property];
        if (date == 0) {
          return ''
        }
        return moment(date).format("YYYY-MM-DD")
      },
      handleSizeChange(e) {
        let loadingInstance = Loading.service({ fullscreen: true });
        this.pageSize = e
        this.currentPage = 1
        axios.post('http://localhost:8090/proxy',{q: this.q, tp: 'withdrawhis',currentPage: this.currentPage, pageSize: this.pageSize}).then(resp => {
          console.log(resp.data)
          this.histories = resp.data.data
          this.currentPage = resp.data.currentPage
          this.pageSize = resp.data.pageSize
          this.totalCount = resp.data.totalCount
          loadingInstance.close()
        })
      },
      handleCurrentChange(e) {
        let loadingInstance = Loading.service({ fullscreen: true });
        this.currentPage = e
        axios.post('http://localhost:8090/proxy',{q: this.q, tp: 'withdrawhis',currentPage: this.currentPage, pageSize: this.pageSize}).then(resp => {
          console.log(resp.data)
          this.histories = resp.data.data
          this.currentPage = resp.data.currentPage
          this.pageSize = resp.data.pageSize
          this.totalCount = resp.data.totalCount
          loadingInstance.close()
        })
      },
      search() {
        let loadingInstance = Loading.service({ fullscreen: true });
        axios.post('http://localhost:8090/proxy',{q: this.q, tp: 'withdrawhis',currentPage: this.currentPage, pageSize: this.pageSize}).then(resp => {
          console.log(resp.data)
          this.histories = resp.data.data
          this.currentPage = resp.data.currentPage
          this.pageSize = resp.data.pageSize
          this.totalCount = resp.data.totalCount
          loadingInstance.close()
        })
      },
      handleClick(row) {
        this.currentId = row._id
        this.dialogFormVisible = true
      },
      doUpdate() {
        let data = {
          tp: 'withdrawupdate',
          _id: this.currentId,
          status: this.status,
          realAmount: this.realWithdrawAmount,
        }
        if (this.status == 3) {
          if (!this.realWithdrawAmount) {
            this.$message.error("请输入实际打款金额")
            return
          }
        }
        axios.post('http://localhost:8090/proxy',data).then(resp => {
          if (resp.status == 200) {
            this.$message.success("更新成功")
            this.currentId = ''
            this.dialogFormVisible = false
            this.$router.go(0)
          }else {
            this.$message.error("更新失败")
          }
        })
      }
    },
    mounted() {
      let loadingInstance = Loading.service({ fullscreen: true });
      axios.post('http://localhost:8090/proxy',{tp: 'withdraws'}).then(resp => {
        this.withdraws = resp.data.reqs
        this.month = resp.data.month
        this.day = resp.data.day
      })
      axios.post('http://localhost:8090/proxy',{tp: 'withdrawhis',currentPage: this.currentPage, pageSize: this.pageSize}).then(resp => {
        console.log(resp.data)
        this.histories = resp.data.data
        this.currentPage = resp.data.currentPage
        this.pageSize = resp.data.pageSize
        this.totalCount = resp.data.totalCount
        loadingInstance.close()
      })

    }
  }
</script>

<style>
</style>
