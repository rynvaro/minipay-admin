<template>
  <div>
    <el-divider content-position="left">优惠券</el-divider>
    <el-table
        :data="coupons"
        border
        style="width: 100%; margin-top: 10px;">
        <el-table-column
          prop="code"
          label="编码"
          width="180">
        </el-table-column>
        <el-table-column
          prop="desc"
          label="描述"
          width="100">
        </el-table-column>

        <el-table-column
          prop="value"
          label="面额"
          width="200">
          <template slot-scope="scope">
            {{scope.row.value / 100}}元
          </template>
        </el-table-column>

        <el-table-column
          prop="man"
          label="满多少可用"
          width="200">
          <template slot-scope="scope">
            {{scope.row.man / 100}}元
          </template>
        </el-table-column>

        <el-table-column
          prop="vpday"
          label="有效期"
          width="200">
          <template slot-scope="scope">
            {{scope.row.vpday}}天
          </template>
        </el-table-column>

        <el-table-column
          fixed="right"
          label="操作"
          width="300">
          <template slot-scope="scope">
            <el-button @click="handleClick(scope.row,scope.$index)" type="text" size="small">删除</el-button>
          </template>
        </el-table-column>

      </el-table>
      <el-button style="width: 100%;margin-top: 5px;" icon="el-icon-plus" @click="addCoupon">新建优惠券</el-button>

      <el-divider content-position="left">发放规则</el-divider>
      <el-table
          :data="rules"
          border
          style="width: 100%; margin-top: 10px;">
          <el-table-column
            prop="type"
            label="类型"
            width="180">
            <template slot-scope="scope">
              <span v-if="scope.row.type == 'register'">新用户注册</span>
              <span v-if="scope.row.type == 'consumption'">用户消费</span>
            </template>
          </el-table-column>
          <el-table-column
            prop="threshold"
            label="消费阈值"
            width="100">
            <template slot-scope="scope">
              {{scope.row.threshold / 100}}元
            </template>
          </el-table-column>

          <el-table-column
            prop="code"
            label="发放编码"
            width="200">
          </el-table-column>

          <el-table-column
          fixed="right"
          label="操作"
          width="300">
          <template slot-scope="scope">
            <el-button @click="handleClick1(scope.row,scope.$index)" type="text" size="small">删除</el-button>
          </template>
        </el-table-column>
        </el-table>
        <el-button style="width: 100%;margin-top: 5px;" icon="el-icon-plus" @click="addRule">添加规则</el-button>


        <el-dialog title="添加优惠券" :visible.sync="dialogFormVisible">

            <el-row>
              <el-col :span="3">
                <div class="title">编码：</div>
              </el-col>
              <el-col :span="6">
                <el-input placeholder="输入优惠券编码" v-model="coupon.code" clearable></el-input>
              </el-col>
              <el-col :span="3">
                <div class="title">描述：</div>
              </el-col>
              <el-col :span="6">
                <el-input placeholder="输入描述" v-model="coupon.desc" clearable></el-input>
              </el-col>
            </el-row>

            <el-row>
              <el-col :span="3">
                <div class="title">面额：</div>
              </el-col>
              <el-col :span="6">
                <el-input placeholder="输入面额" v-model="coupon.value" clearable></el-input>
              </el-col>
              <el-col :span="3">
                <div class="title">使用门槛：</div>
              </el-col>
              <el-col :span="6">
                <el-input placeholder="输入使用门槛" v-model="coupon.man" clearable></el-input>
              </el-col>
            </el-row>

            <el-row>
              <el-col :span="6">
                <div class="title">有效期(单位：天)：</div>
              </el-col>
              <el-col :span="6">
                <el-input placeholder="输入有效期(天)" v-model="coupon.vpday" clearable></el-input>
              </el-col>
            </el-row>


          <div slot="footer" class="dialog-footer">
            <el-button @click="dialogFormVisible = false">取 消</el-button>
            <el-button type="primary" @click="doAddCoupon()">确 定</el-button>
          </div>
        </el-dialog>

        <el-dialog title="添加规则" :visible.sync="dialogFormVisible1">

            <el-row>
              <el-col :span="3">
                <div class="title">类型：</div>
              </el-col>
              <el-col :span="6">
                <el-select v-model="rule.type" filterable placeholder="规则类型">
                    <el-option
                      v-for="item in typeList"
                      :key="item.id"
                      :label="item.label"
                      :value="item.value">
                    </el-option>
                  </el-select>
              </el-col>
            </el-row>

            <el-row>
              <el-col :span="3">
                <div class="title">消费阈值：</div>
              </el-col>
              <el-col :span="6">
                <el-input placeholder="输入阈值" v-model="rule.threshold" clearable></el-input>
              </el-col>
              <el-col :span="3">
                <div class="title">发放编码：</div>
              </el-col>
              <el-col :span="6">
                <el-input placeholder="输入发放编码" v-model="rule.code" clearable></el-input>
              </el-col>
            </el-row>

          <div slot="footer" class="dialog-footer">
            <el-button @click="dialogFormVisible1 = false">取 消</el-button>
            <el-button type="primary" @click="doAddRule()">确 定</el-button>
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
        coupon:{},
        rule: {},
        q: '',
        dialogFormVisible: false,
        dialogFormVisible1: false,
        coupons: [],
        rules: [],
        typeList: [
          {
            value: "register",
            label: '新用户注册'
          },
          {
            value: "consumption",
            label: '用户消费'
          }
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
      handleClick(row,index){
        let loadingInstance = Loading.service({ fullscreen: true });
        axios.post('http://localhost:8090/proxy',{id: row._id,tp:'coupondicdelete'}).then(resp => {
          loadingInstance.close()
          this.coupons.splice(index,1)
        })
      },
      handleClick1(row,index){
        let loadingInstance = Loading.service({ fullscreen: true });
        axios.post('http://localhost:8090/proxy',{id: row._id, tp:'couponruledelete'}).then(resp => {
          loadingInstance.close()
          this.rules.splice(index,1)
        })
      },
      addRule() {
        this.dialogFormVisible1 = true
      },
      doAddRule() {
        if (!this.rule.type) {
          this.$message.error("请选择类型")
          return
        }
        if (!this.rule.threshold) {
          this.$message.error("请输入消费阈值")
          return
        }
        if (!this.rule.code) {
          this.$message.error("请输入发放编码")
          return
        }
        let loadingInstance = Loading.service({ fullscreen: true });
        this.rule.tp = 'couponruleadd'
        axios.post('http://localhost:8090/proxy',this.rule).then(resp => {
          console.log(resp.data)
          this.dialogFormVisible1 = false
          axios.post('http://localhost:8090/proxy',{tp: 'couponrules'}).then(resp => {
            this.rules = resp.data.data
          })
          loadingInstance.close()
        })
      },
      addCoupon() {
         this.dialogFormVisible = true
      },
      doAddCoupon() {
        if (!this.coupon.code) {
          this.$message.error("请输入编码")
          return
        }
        if (!this.coupon.desc) {
          this.$message.error("请输入描述")
          return
        }
        if (!this.coupon.value) {
          this.$message.error("请输入面额")
          return
        }
        if (!this.coupon.man) {
          this.$message.error("请输入使用门槛")
          return
        }
        if (!this.coupon.vpday) {
          this.$message.error("请输入有效期")
          return
        }
        let loadingInstance = Loading.service({ fullscreen: true });
        this.coupon.tp = 'coupondicadd'
        axios.post('http://localhost:8090/proxy',this.coupon).then(resp => {
          axios.post('http://localhost:8090/proxy',{tp: 'coupondic'}).then(resp => {
            console.log(resp.data)
            this.coupons = resp.data.data
          })
          loadingInstance.close()
          this.dialogFormVisible = false

        })
      }
    },
    mounted() {
      let loadingInstance = Loading.service({ fullscreen: true });
      axios.post('http://localhost:8090/proxy',{tp: 'coupondic'}).then(resp => {
        console.log(resp.data)
        this.coupons = resp.data.data
      })
      axios.post('http://localhost:8090/proxy',{tp: 'couponrules'}).then(resp => {
        console.log(resp.data)
        this.rules = resp.data.data
        loadingInstance.close()
      })

    }
  }
</script>

<style>
</style>
