<template>

  <div>
    <el-row style="margin-top: 20px;">
      <el-col :span="10">
        <el-input placeholder="输入昵称" v-model="q" @change="search" clearable></el-input>
      </el-col>
      <el-col :span="3">
        <el-button @click="search">查询</el-button>
      </el-col>
    </el-row>
    <el-table
        :data="users"
        border
        style="width: 100%">
        <el-table-column
          prop="_id"
          label="ID"
          width="280">
        </el-table-column>
        <el-table-column
          prop="data.name"
          label="昵称"
          width="100">
        </el-table-column>
        <el-table-column
          prop="data.balance"
          label="余额"
          width="100">
          <template slot-scope="scope">
            {{scope.row.data.balance/100}}
          </template>
        </el-table-column>
        <el-table-column
          prop="data.exp"
          label="经验"
          width="50">
        </el-table-column>
        <el-table-column
          prop="data.point"
          label="积分"
          width="50">
        </el-table-column>

        <el-table-column
          prop="data.exp"
          label="等级"
          width="50">
          <!-- <template slot-scope="scope">
            {{level(scope.row.data.exp)}}
          </template> -->
        </el-table-column>

        <el-table-column
          prop="data.phone"
          label="手机号"
          width="120">
        </el-table-column>

        <el-table-column
          prop="data.sevenSigns"
          label="签到"
          width="50">
        </el-table-column>
        <el-table-column
          prop="createdAt"
          label="注册日期"
          width="150"
          :formatter="createAtDateFormater">
        </el-table-column>
        <el-table-column
          prop="signDate"
          label="签到日期"
          width="150"
          :formatter="dateFormater">
        </el-table-column>
        <el-table-column
          fixed="right"
          label="操作"
          width="300">
          <template slot-scope="scope">
            <el-button @click="addBalance(scope.row,scope.$index)" type="text" size="small">修改余额</el-button>
            <el-button @click="handleClick(scope.row, 1)" type="text" size="small">发放优惠券</el-button>
            <el-button @click="handleClick(scope.row, 2)" type="text" size="small">发放经验</el-button>
            <!-- <el-button @click="handleClick(scope.row, 3)" type="text" size="small">增加签到天数</el-button> -->
            <el-button @click="handleClick(scope.row, 4)" type="text" size="small">发放积分</el-button>
            <!-- <el-button @click="handleClick(scope.row, 5)" type="text" size="small">注册日期前移N天</el-button> -->
            <!-- <el-button @click="handleClick(scope.row, 6)" type="text" size="small">签到日期迁移N天</el-button> -->
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
          <el-dialog title="修改余额" :visible.sync="dialogFormVisible">

              <el-row>
                <el-col :span="3">
                  <div class="title">余额：</div>
                </el-col>
                <el-col :span="6">
                  <el-input placeholder="输入余额" v-model="balance" clearable></el-input>
                </el-col>
              </el-row>

            <div slot="footer" class="dialog-footer">
              <el-button @click="dialogFormVisible = false">取 消</el-button>
              <el-button type="primary" @click="doAddBalance()">确 定</el-button>
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
        dialogFormVisible: false,
        balance: 0,
        users:[],
        q: '',
        currentPage: 1,
        pageSize: 10,
        totalCount: 0,
        currentUserId: '',
        currentIndex: -1,
      }
    },
    methods: {
      addBalance(row,index) {
        this.dialogFormVisible = true
        this.currentUserId = row._id
        this.currentIndex = index
      },
      doAddBalance() {
        let balance = this.balance * 100
        let loadingInstance = Loading.service({ fullscreen: true });
        axios.post('http://localhost:8090/proxy',{tp: 'updateuserbalance' ,id:this.currentUserId, balance: balance}).then(resp => {
          console.log(resp)
          this.users[this.currentIndex].data.balance = balance
          this.balance = 0
          this.currentIndex = -1
          this.currentUserId = '',
          this.dialogFormVisible = false
          loadingInstance.close()
        })
      },
      handleSizeChange(e) {
        let loadingInstance = Loading.service({ fullscreen: true });
        this.pageSize = e
        this.currentPage = 1
        axios.post('http://localhost:8090/proxy',{tp: 'userlist' ,q: this.q,currentPage: this.currentPage, pageSize: this.pageSize}).then(resp => {
          console.log(resp.data)
          this.users = resp.data.data
          this.currentPage = resp.data.currentPage
          this.pageSize = resp.data.pageSize
          this.totalCount = resp.data.totalCount
          loadingInstance.close()
        })
      },
      handleCurrentChange(e) {
        let loadingInstance = Loading.service({ fullscreen: true });
        this.currentPage = e
        axios.post('http://localhost:8090/proxy',{tp: 'userlist' ,q: this.q,currentPage: this.currentPage, pageSize: this.pageSize}).then(resp => {
          console.log(resp.data)
          this.users = resp.data.data
          this.currentPage = resp.data.currentPage
          this.pageSize = resp.data.pageSize
          this.totalCount = resp.data.totalCount
          loadingInstance.close()
        })
      },
      search() {
        let loadingInstance = Loading.service({ fullscreen: true });
        axios.post('http://localhost:8090/proxy',{tp: 'userlist' ,q: this.q,currentPage: this.currentPage, pageSize: this.pageSize}).then(resp => {
          console.log(resp.data)
          this.users = resp.data.data
          this.currentPage = resp.data.currentPage
          this.pageSize = resp.data.pageSize
          this.totalCount = resp.data.totalCount
          loadingInstance.close()
        })
      },
      // level(exp){
      //   if (exp >= 0 && exp <1000) {
      //     return 1
      //   }else if (exp >=1000 && exp <10000) {
      //     return 2
      //   } else if (exp > 10000){
      //     return 3
      //   }
      // },
      createAtDateFormater: function(row, column) {
        var date = row[column.property];
        if (date == 0) {
          return ''
        }
        return moment(date).format("YYYY-MM-DD")
      },
      dateFormater: function(row, column) {
        var date = row.data[column.property];
        if (date == 0) {
          return '未签到'
        }
        return moment(date).format("YYYY-MM-DD")
      },
      handleClick(row,index) {
        this.$prompt('输入数量', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
        }).then(({ value }) => {
          switch(index) {
            case 1:
              if (value > 1) {
                this.$message({
                  type: 'error',
                  message: '一次只能发放一个'
                });
                return
              }
              row.tp = 2
              axios.post('http://localhost:8090/addcoupon', row).then(resp => {
                console.log(resp)
                if (resp.status == 200) {
                  this.$message.success("成功")
                }else {
                  this.$message.error("失败")
                }
               })
              return
            case 2:
              row.data.exp +=parseInt(value)
              break
            case 3:
              row.data.signs +=parseInt(value)
              row.data.sevenSigns +=parseInt(value)
              row.data.signDate = Date.parse(new Date())
              break
            case 4:
              row.data.point +=parseInt(value)
              break
            case 5:
              row.createdAt = row.createdAt - (3600 * 24 * 1000) * value
              break
            case 6:
              row.data.signDate = row.data.signDate - (3600 * 24 * 1000) * value
              break
          }
            axios.post('http://localhost:8090/updateuser', row).then(resp => {
              console.log(resp)
              if (resp.status == 200) {
                this.$message.success("成功")
              }else {
                this.$message.error("失败")
              }
            })
          }).catch(() => {
            this.$message({
              type: 'info',
              message: '取消输入'
            });
          });
      }
    },
    mounted() {
      let loadingInstance = Loading.service({ fullscreen: true });
      axios.post('http://localhost:8090/proxy',{tp: 'userlist' ,q: '',currentPage: this.currentPage, pageSize: this.pageSize}).then(resp => {
        console.log(resp.data)
        this.users = resp.data.data
        this.currentPage = resp.data.currentPage
        this.pageSize = resp.data.pageSize
        this.totalCount = resp.data.totalCount
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
    margin-top: 10px;
  }

</style>
