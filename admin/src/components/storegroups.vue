<template>

  <div>
    <el-table
          :data="groups"
          style="width: 100%">
          <el-table-column
            prop="title"
            label="名称"
            width="180">
          </el-table-column>
          <el-table-column
            prop="discount"
            label="折扣"
            width="100">
          </el-table-column>
          <el-table-column
            prop="enabled"
            label="已启用"
            width="100">
          </el-table-column>
          <el-table-column
            label="操作"
            width="200">
            <template slot-scope="scope">
              <el-button @click="handleEdit(scope.$index)" type="text" size="small">编辑</el-button>
              <el-button @click="handleDelete(scope.$index)" type="text" size="small">删除</el-button>
              <el-button @click="handleDisable(scope.$index)" type="text" size="small">禁用</el-button>
              <el-button @click="handleEnable(scope.$index)" type="text" size="small">启用</el-button>
            </template>
          </el-table-column>
        </el-table>

        <el-dialog title="添加分组" :visible.sync="dialogFormVisible" @close="clearDialog">

            <el-row>
              <el-col :span="3">
                <div class="title">分组名称：</div>
              </el-col>
              <el-col :span="6">
                <el-input placeholder="请输入分组名称" v-model="group.title" clearable></el-input>
              </el-col>
              <el-col :span="3">
                <div class="title">分组折扣：</div>
              </el-col>
              <el-col :span="6">
                <el-input placeholder="请输入分组折扣" v-model="group.discount" clearable></el-input>
              </el-col>
            </el-row>

            <el-divider content-position="center">已添加的商家</el-divider>

          <!-- <el-row > -->
            <!-- <el-col :span="24">
              <span>已添加的商家</span>
            </el-col>
          </el-row> -->

          <el-row>
            <el-col :span="24">
              <el-table
                    :data="group.stores"
                    style="width: 100%">
                    <el-table-column
                      prop="storeName"
                      label="店铺名称"
                      width="180">
                    </el-table-column>
                    <el-table-column
                      prop="address"
                      label="地址">
                    </el-table-column>
                    <el-table-column
                      prop="discount"
                      label="原折扣">
                    </el-table-column>
                    <el-table-column
                      fixed="right"
                      label="操作"
                      width="50">
                      <template slot-scope="scope">
                        <el-button @click="handleRemove(scope.row, scope.$index)" type="text" size="small">删除</el-button>
                      </template>
                    </el-table-column>
                  </el-table>
            </el-col>
          </el-row>
<el-divider content-position="center">商家列表</el-divider>
          <el-row>
            <el-col :span="24">
              <el-input placeholder="搜索商家" v-model="q" @change="search" clearable></el-input>
            </el-col>
          </el-row>

          <el-row>
            <el-col :span="24">
              <el-table
                    :data="stores"
                    style="width: 100%">
                    <el-table-column
                      prop="storeName"
                      label="店铺名称"
                      width="180">
                    </el-table-column>
                    <el-table-column
                      prop="address"
                      label="地址">
                    </el-table-column>
                    <el-table-column
                      prop="discount"
                      label="折扣">
                    </el-table-column>
                    <el-table-column
                      fixed="right"
                      label="操作"
                      width="50">
                      <template slot-scope="scope">
                        <el-button @click="handleClick(scope.row)" type="text" size="small">添加</el-button>
                      </template>
                    </el-table-column>
                  </el-table>
            </el-col>
          </el-row>

          <div slot="footer" class="dialog-footer">
            <el-button @click="dialogFormVisible = false">取 消</el-button>
            <el-button type="primary" @click="doAddEvent(type)">确 定</el-button>
          </div>
        </el-dialog>

        <el-button style="width: 100%;margin-top: 5px;" icon="el-icon-plus" @click="addEvent">添加分组</el-button>
  </div>
</template>

<script>
  import axios from 'axios'
  import { Loading } from 'element-ui';
  import uuidv1 from 'uuid/v1'
  export default {
    data() {
      return {
        q: '',
        stores:[],
        groups:[],
        dialogFormVisible: false,
        group: {
          stores: []
        },
        removes: [],
        type: 1, // 1 add 2 update
      }
    },
    methods: {
      clearDialog(done) {
        this.group = {
          stores: []
        }
        this.removes = []
      },
      handleRemove(row,index) {
        this.group.stores.splice(index,1)
        this.removes.push(row)
      },
      handleClick(row) {
        if (!this.group.stores) {
          this.group.stores = []
        }
        console.log(row)
        this.group.stores.push(row)
      },
      search() {
        let loadingInstance = Loading.service({ fullscreen: true });
        axios.post('http://localhost:8090/stores',{q: this.q,currentPage: 1, pageSize: 5}).then(resp => {
          this.stores = resp.data.data
          loadingInstance.close()
        })
      },
      addEvent() {
         this.dialogFormVisible = true
         this.type = 1
      },
      doAddEvent(type){

        if (!this.group.title) {
          this.$message.error("请输入分组名称")
          return
        }
        if (!this.group.discount) {
          this.$message.error("请输入分组折扣")
          return
        }
        let loadingInstance = Loading.service({ fullscreen: true });
        if (type == 1) {
          this.group.tp = 'groupadd'
          this.group._id = uuidv1()
          this.group.enabled = 1
          axios.post('http://localhost:8090/proxy',this.group).then(resp => {
            if (resp.status == 200) {
              this.$message.success("添加成功")
              this.groups.push(this.group)
              this.group = {}
            }else {
              this.$message.error("添加失败")
            }
            this.dialogFormVisible = false
          })
        }else if (type == 2) {
          this.group.tp = 'groupupdate'
          this.group.removes = this.removes
          axios.post('http://localhost:8090/proxy',this.group).then(resp => {
            if (resp.status == 200) {
              this.$message.success("更新成功")
            }else {
              this.$message.error("更新失败")
            }
            this.dialogFormVisible = false
          })
        }
        loadingInstance.close()
      },
      handleDelete(index) {
        let group = this.groups[index]
        group.tp = 'groupdelete'
        axios.post('http://localhost:8090/proxy',group).then(resp => {
          if (resp.status == 200) {
            this.groups.splice(index,1)
            this.$message.success("删除成功")
          }else {
            this.$message.error("删除失败")
          }
        })
      },
      handleEdit(index) {
        this.group = this.groups[index]
        this.type = 2
        this.dialogFormVisible = true
      },
      handleDisable(index) {
        let group = this.groups[index]
        let data = {
          tp: 'groupdisable',
          id: group._id,
        }
        let loadingInstance = Loading.service({ fullscreen: true });
        axios.post('http://localhost:8090/proxy',data).then(resp => {
          if (resp.status == 200) {
            this.groups[index].enabled = 0
            this.$message.success("禁用成功")
          }else {
            this.$message.error("禁用失败")
          }
          loadingInstance.close()
        })
      },
      handleEnable(index){
        let group = this.groups[index]
        let data = {
          tp: 'groupenable',
          id: group._id,
        }
        let loadingInstance = Loading.service({ fullscreen: true });
        axios.post('http://localhost:8090/proxy',data).then(resp => {
          if (resp.status == 200) {
            this.groups[index].enabled = 1
            this.$message.success("启用成功")
          }else {
            this.$message.error("启用失败")
          }
          loadingInstance.close()
        })
       }
    },


    mounted() {
      axios.post('http://localhost:8090/proxy',{tp: 'groups'}).then(resp => {
        console.log(resp)
        this.groups = resp.data.data
      })
      axios.post('http://localhost:8090/stores',{q: '',currentPage: 1, pageSize: 5}).then(resp => {
        console.log(resp.data)
        this.stores = resp.data.data
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
