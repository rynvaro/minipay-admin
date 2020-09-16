<template>

  <div>

    <el-table
        :data="users"
        border
        style="width: 100%">
        <!-- <el-table-column
          prop="_id"
          label="ID"
          width="280">
        </el-table-column> -->
        <el-table-column
          prop="data.name"
          label="昵称"
          width="100">
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
          prop="data.signs"
          label="累计签到"
          width="50">
        </el-table-column>

        <el-table-column
          prop="data.sevenSigns"
          label="连续签到"
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
          width="500">
          <template slot-scope="scope">
            <el-button @click="handleClick(scope.row, 1)" type="text" size="small">发放优惠券</el-button>
            <el-button @click="handleClick(scope.row, 2)" type="text" size="small">发放经验</el-button>
            <el-button @click="handleClick(scope.row, 3)" type="text" size="small">增加签到天数</el-button>
            <el-button @click="handleClick(scope.row, 4)" type="text" size="small">发放积分</el-button>
            <el-button @click="handleClick(scope.row, 5)" type="text" size="small">注册日期前移N天</el-button>
            <el-button @click="handleClick(scope.row, 6)" type="text" size="small">签到日期迁移N天</el-button>
          </template>
        </el-table-column>
      </el-table>
  </div>
</template>

<script>
  import axios from 'axios'
  import moment from 'moment'
  export default {
    data() {
      return {
        users:[]
      }
    },
    methods: {
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
      axios.get('http://localhost:8090/users').then(resp => {
        this.users = resp.data.data
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
