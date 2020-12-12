<template>

  <div>
    <el-row>
      <el-col :span="10">
        <el-input placeholder="搜索商家" v-model="q" @change="search" clearable></el-input>
      </el-col>
      <el-col :span="4">
        <el-button @click="allQRcode">统一生成二维码</el-button>
      </el-col>
      <el-col :span="4">
        <span>本页条数：{{stores.length}}</span>
      </el-col>
    </el-row>
    <el-dialog
      width="30%"
      @close="closeCode"
      :visible.sync="innerVisible"
      append-to-body>
        <div class="qrcode">
          <div id="qrcode" ref="qrcode"></div>
      </div>
    </el-dialog>
    <el-table
          :data="stores"
          style="width: 100%">
          <el-table-column
            prop="merchantName"
            label="商家姓名"
            width="180">
          </el-table-column>
          <el-table-column
            prop="merchantPhone"
            label="联系方式"
            width="180">
          </el-table-column>
          <el-table-column
            prop="storeName"
            label="店铺名称"
            width="180">
          </el-table-column>
          <el-table-column
            prop="balance"
            label="余额">
            <template slot-scope="scope">
              {{scope.row.balance | rounding}}
            </template>
          </el-table-column>
          <el-table-column
            prop="canUseBalance"
            label="是否可用余额支付">
            <template slot-scope="scope">
              <span style="margin-left: 10px" v-if="scope.row.canUseBalance || scope.row.canUseBalance === undefined">可用</span>
              <span style="margin-left: 10px" v-if="scope.row.canUseBalance === false" >不可用</span>
            </template>
          </el-table-column>
          <el-table-column
            prop="fanxian100"
            label="是否参与返现100%活动">
            <template slot-scope="scope">
              <span style="margin-left: 10px" v-if="scope.row.fanxian100">已参与</span>
              <span style="margin-left: 10px" v-if="!scope.row.fanxian100">不参与</span>
            </template>
          </el-table-column>
          <el-table-column
            prop="maxFanxian"
            label="最高返现">
            <template slot-scope="scope">
              <span style="margin-left: 10px" v-if="scope.row.maxFanxian">{{scope.row.maxFanxian}}</span>
              <span style="margin-left: 10px" v-if="!scope.row.maxFanxian">未设置</span>
            </template>
          </el-table-column>
          <el-table-column
            prop="address"
            label="地址">
          </el-table-column>
          <el-table-column
            prop="deleted"
            label="是否显示">
            <template slot-scope="scope">
              <span style="margin-left: 10px" v-if="scope.row.deleted === 1" >已隐藏</span>
              <span style="margin-left: 10px" v-if="scope.row.deleted === 0" >正常显示</span>
            </template>
          </el-table-column>
          <el-table-column
            prop="norake"
            label="平台参与抽成">
            <template slot-scope="scope">
              <span style="margin-left: 10px" v-if="scope.row.norake" >不参与</span>
              <span style="margin-left: 10px" v-if="!scope.row.norake" >参与</span>
            </template>
          </el-table-column>
          <el-table-column
            fixed="right"
            label="操作"
            width="300">
            <template slot-scope="scope">
              <el-button @click="handleClick(scope.row)" type="text" size="small">编辑</el-button>
              <el-button @click="handleDelete(scope.row, scope.$index)" type="text" size="small">隐藏</el-button>
              <el-button @click="handleUnDelete(scope.row, scope.$index)" type="text" size="small">显示</el-button>
              <el-button @click="showQRCode(scope.row)" type="text" size="small">查看二维码</el-button>
              <el-button @click="handleHardDelete(scope.row, scope.$index)" type="text" size="small">硬删除</el-button>
              <el-button @click="rake(scope.row, scope.$index,0)" type="text" size="small">参与抽成</el-button>
              <el-button @click="rake(scope.row, scope.$index,1)" type="text" size="small">不参与抽成</el-button>
              <el-button @click="logout(scope.row, scope.$index)" type="text" size="small">清除登录状态</el-button>
              <el-button @click="canUseBalance(scope.row, scope.$index,0)" type="text" size="small">禁用余额支付</el-button>
              <el-button @click="canUseBalance(scope.row, scope.$index,1)" type="text" size="small">启用余额支付</el-button>

              <el-button @click="fanxian100(scope.row, scope.$index,0)" type="text" size="small">退出返现100%活动</el-button>
              <el-button @click="fanxian100(scope.row, scope.$index,1)" type="text" size="small">参与返现100%活动</el-button>


              <el-button @click="maxFanxian(scope.row, scope.$index)" type="text" size="small">设置最高返现</el-button>


              <el-button @click="resetBalance(scope.row, scope.$index)" type="text" size="small">清空余额</el-button>


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

            <el-dialog title="返现设置" :visible.sync="dialogFormVisible">
              <el-form :model="form">
                <el-form-item label="最高返现">
                  <el-input v-model.number="form.maxFanxian"></el-input>
                </el-form-item>
              </el-form>
              <div slot="footer" class="dialog-footer">
                <el-button @click="dialogFormVisible = false">取 消</el-button>
                <el-button type="primary" @click="doMaxFanxian">确 定</el-button>
              </div>
            </el-dialog>
  </div>
</template>

<script>
  import axios from 'axios'
  import QRCode from 'qrcodejs2'
  import { Loading } from 'element-ui';
  export default {
    data() {
      return {
        dialogFormVisible: false,
        stores:[],
        q: '',
        qrcode: '',
        innerVisible: false,
        currentPage: 1,
        pageSize: 10,
        totalCount: 0,
        form: {},
      }
    },
    filters:{
      rounding (value) {
       return value.toFixed(2)
       }
    },
    methods: {
      logout(row, index) {
        let loadingInstance = Loading.service({ fullscreen: true });
        let data = {
          id: row._id,
          tp: "logout"
        }
        axios.post('http://localhost:8090/proxy',data).then(resp => {
          loadingInstance.close()
          this.$message.success("清除成功")
        })
      },
      maxFanxian(row, index){
          this.dialogFormVisible = true
          this.form.index = index
      },
      doMaxFanxian() {
        console.log(this.form)
        let loadingInstance = Loading.service({ fullscreen: true });
        let data = {
          id: this.stores[this.form.index]._id,
          maxFanxian: this.form.maxFanxian,
          tp: "maxfanxian"
        }
        axios.post('http://localhost:8090/proxy',data).then(resp => {
          this.dialogFormVisible = false
          axios.post('http://localhost:8090/stores',{q: this.q,currentPage: this.currentPage, pageSize: this.pageSize}).then(resp => {
            console.log(resp.data)
            this.stores = resp.data.data
            this.currentPage = resp.data.currentPage
            this.pageSize = resp.data.pageSize
            this.totalCount = resp.data.totalCount
            loadingInstance.close()
          })
        })
        
      },
      resetBalance(row, index) {
        let loadingInstance = Loading.service({ fullscreen: true });
        let data = {
          id: row._id,
          tp: "resetbalance"
        }
        axios.post('http://localhost:8090/proxy',data).then(resp => {
          this.stores[index].balance = 0
          loadingInstance.close()
        })
      },
      fanxian100(row, index,e) {
        let loadingInstance = Loading.service({ fullscreen: true });
        let data = {
          id: row._id,
          fanxian100: e,
          tp: "fanxian100"
        }
        axios.post('http://localhost:8090/proxy',data).then(resp => {
          // if (e == 1) {
          //   row.fanxian100 = true
          //   this.stores[index].fanxian100 = true
          // }else {
          //   this.stores[index].fanxian100 = false
          // }
          axios.post('http://localhost:8090/stores',{q: this.q,currentPage: this.currentPage, pageSize: this.pageSize}).then(resp => {
            console.log(resp.data)
            this.stores = resp.data.data
            this.currentPage = resp.data.currentPage
            this.pageSize = resp.data.pageSize
            this.totalCount = resp.data.totalCount
            loadingInstance.close()
          })
        })
      },
      canUseBalance(row, index,e){
        let loadingInstance = Loading.service({ fullscreen: true });
        let data = {
          id: row._id,
          canUseBalance: e,
          tp: "canusebalance"
        }
        axios.post('http://localhost:8090/proxy',data).then(resp => {
          axios.post('http://localhost:8090/stores',{q: this.q,currentPage: this.currentPage, pageSize: this.pageSize}).then(resp => {
            console.log(resp.data)
            this.stores = resp.data.data
            this.currentPage = resp.data.currentPage
            this.pageSize = resp.data.pageSize
            this.totalCount = resp.data.totalCount
            loadingInstance.close()
          })
        })
      },
      rake(row, index,e){
        let loadingInstance = Loading.service({ fullscreen: true });
        let data = {
          id: row._id,
          norake: e,
          tp: "rake"
        }
        axios.post('http://localhost:8090/proxy',data).then(resp => {
          loadingInstance.close()
          this.afterRake()
        })
      },
      afterRake(){
        let loadingInstance = Loading.service();
        axios.post('http://localhost:8090/stores',{q: this.q,currentPage: this.currentPage, pageSize: this.pageSize}).then(resp => {
          this.stores = resp.data.data
          this.currentPage = resp.data.currentPage
          this.pageSize = resp.data.pageSize
          this.totalCount = resp.data.totalCount
          loadingInstance.close()
        })
      },
      handleSizeChange(e) {
        this.pageSize = e
        this.currentPage = 1
        axios.post('http://localhost:8090/stores',{q: '',currentPage: this.currentPage, pageSize: this.pageSize}).then(resp => {
          console.log(resp.data)
          this.stores = resp.data.data
          this.currentPage = resp.data.currentPage
          this.pageSize = resp.data.pageSize
          this.totalCount = resp.data.totalCount
        })
      },
      handleCurrentChange(e) {
        this.currentPage = e
        axios.post('http://localhost:8090/stores',{q: '',currentPage: this.currentPage, pageSize: this.pageSize}).then(resp => {
          console.log(resp.data)
          this.stores = resp.data.data
          this.currentPage = resp.data.currentPage
          this.pageSize = resp.data.pageSize
          this.totalCount = resp.data.totalCount
        })
      },
      allQRcode() {
        this.$router.push({name: 'qrcodes', params: {stores: this.stores}})
      },
      handleClick(row) {
        this.$router.push({name: 'pubstores', params: {_id: row._id}})
      },
      handleHardDelete(row,index) {
        let loadingInstance = Loading.service({ fullscreen: true });
        axios.post('http://localhost:8090/storedelete',{tp:'storedeletehard',_id: row._id}).then(resp => {
          if (resp.status == 200) {
            this.$message.success("删除成功")
            this.stores.splice(index,1)
          }else {
            this.$message.error("删除失败")
          }
          loadingInstance.close()
        })
      },
      handleDelete(row,index){
         axios.post('http://localhost:8090/storedelete',{tp:'storedelete',_id: row._id,deleted: 1}).then(resp => {
           if (resp.status == 200) {
             this.$message.success("隐藏成功")
             this.stores[index].deleted = 1
           }else {
             this.$message.error("隐藏失败")
           }
         })
      },
      handleUnDelete(row,index){
         axios.post('http://localhost:8090/storedelete',{tp:'storedelete',_id: row._id,deleted: 0}).then(resp => {
           if (resp.status == 200) {
             this.$message.success("显示成功")
             this.stores[index].deleted = 0
           }else {
             this.$message.error("显示失败")
           }
         })
      },
      search() {
        axios.post('http://localhost:8090/stores',{q: this.q,currentPage: this.currentPage, pageSize: this.pageSize}).then(resp => {
          this.stores = resp.data.data
        })
      },
      showQRCode(row) {
        this.innerVisible = true
        this.qrcode = row._id
        this.$nextTick(() => {
          this.createQrcode()
        })
      },
      createQrcode () {
        this.qr = new QRCode('qrcode', {
          width: 150,
          height: 150,
          text: this.qrcode
        })
      },
      closeCode () {
        this.$refs.qrcode.innerHTML = ''
      },
    },
    mounted() {
      axios.post('http://localhost:8090/stores',{q: '',currentPage: this.currentPage, pageSize: this.pageSize}).then(resp => {
        console.log(resp.data)
        this.stores = resp.data.data
        this.currentPage = resp.data.currentPage
        this.pageSize = resp.data.pageSize
        this.totalCount = resp.data.totalCount
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

  .qrcode {
    display: flex;
    justify-content: center;
    align-items: center;
    border-radius: 10px;
  }

  .title {
    margin-top: 10px;
  }

</style>
