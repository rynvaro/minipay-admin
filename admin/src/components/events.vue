<template>

  <div>
    <el-table
          :data="events"
          style="width: 100%">
          <el-table-column
            prop="title"
            label="活动名称"
            width="180">
          </el-table-column>
          <el-table-column
            prop="desc"
            label="文案"
            width="100">
          </el-table-column>
          <el-table-column
            prop="btnText"
            label="按钮文字"
            width="100">
          </el-table-column>
          <el-table-column
            label="活动类型"
            width="100">
            <template slot-scope="scope">
              <span style="margin-left: 10px" v-if="scope.row.type === '1'" >快闪店</span>
              <span style="margin-left: 10px" v-if="scope.row.type === '2'" >抽奖</span>
              <span style="margin-left: 10px" v-if="scope.row.type === '3'" >邀请</span>
            </template>
          </el-table-column>
          <el-table-column
            prop="index"
            label="位置">
          </el-table-column>
          <el-table-column
            fixed="right"
            label="操作"
            width="100">
            <template slot-scope="scope">
              <el-button @click="handleEdit(scope.$index)" type="text" size="small">编辑</el-button>
              <el-button @click="handleDelete(scope.$index)" type="text" size="small">删除</el-button>
            </template>
          </el-table-column>
        </el-table>

        <el-dialog title="添加活动" :visible.sync="dialogFormVisible" @close="clearDialog">

            <el-row>
              <el-col :span="3">
                <div class="title">活动名称：</div>
              </el-col>
              <el-col :span="6">
                <el-input placeholder="请输入活动名称" v-model="event.title" clearable></el-input>
              </el-col>
            </el-row>

            <el-row>
              <el-col :span="3">
                <div class="title">活动文案：</div>
              </el-col>
              <el-col :span="18">
                <el-input placeholder="可选" type="textarea" v-model="event.desc" clearable></el-input>
              </el-col>
            </el-row>

            <el-row>
              <el-col :span="3">
                <div class="title">位置：</div>
              </el-col>
              <el-col :span="18">
                <el-input placeholder="请输入位置" v-model="event.index" clearable></el-input>
              </el-col>
            </el-row>

          <el-row>
            <el-col :span="3">
              <div class="title">轮播图片：</div>
            </el-col>
            <el-col :span="8">
               <el-upload class="upload-demo" ref="item" action="http://localhost:8090/upload" :on-success="onItemUploadSuccess" :auto-upload="false">
                 <el-button slot="trigger" size="small" type="primary">选取文件</el-button>
                 <el-button style="margin-left: 10px;" size="small" type="success" @click="submitUpload(1)">上传到服务器</el-button>
                 <div slot="tip" class="el-upload__tip"></div>
               </el-upload>
            </el-col>
          </el-row>

          <el-row>
            <el-col :span="3">
              <div class="title">详情图片：</div>
            </el-col>
            <el-col :span="8">
               <el-upload class="upload-demo" ref="itemDetail" action="http://localhost:8090/upload" multiple="true" :on-success="onDetailUploadSuccess" :auto-upload="false">
                 <el-button slot="trigger" size="small" type="primary">选取文件</el-button>
                 <el-button style="margin-left: 10px;" size="small" type="success" @click="submitUpload(2)">上传到服务器</el-button>
                 <div slot="tip" class="el-upload__tip"></div>
               </el-upload>
            </el-col>
            <el-col :span="3">
              <div class="title">按钮文字：</div>
            </el-col>
            <el-col :span="6">
              <el-input placeholder="详情页按钮文字" v-model="event.btnText" clearable></el-input>
            </el-col>
          </el-row>

          <el-row>
            <el-col :span="3">
              <div class="title">活动类型：</div>
            </el-col>
            <el-col :span="4">
              <el-select v-model="event.type" filterable placeholder="活动类型">
                  <el-option
                    v-for="item in links"
                    :key="item.id"
                    :label="item.label"
                    :value="item.value">
                  </el-option>
                </el-select>
            </el-col>
          </el-row>

          <el-row :hidden.sync="event.type != '1'">
            <el-col :span="24">
              <span>已添加的商家</span>
            </el-col>
          </el-row>

          <el-row  :hidden.sync="event.type != '1'">
            <el-col :span="24">
              <el-table
                    :data="event.stores"
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

          <el-row :hidden.sync="event.type != '1'">
            <el-col :span="24">
              <el-input placeholder="搜索商家" v-model="q" @change="search" clearable></el-input>
            </el-col>
          </el-row>

          <el-row  :hidden.sync="event.type != '1'">
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

          <el-row :hidden.sync="event.type != '2'">
            <el-col :span="4">
             <div class="title">开奖时间：</div>
            </el-col>

            <el-col :span="4">
              <el-date-picker
                    v-model="event.lottoryTime"
                    type="datetime"
                    placeholder="选择日期时间">
                  </el-date-picker>
            </el-col>
          </el-row>

          <el-row :hidden.sync="event.type != '2'">
            <el-col :span="4">
             <div class="title">一等奖人数：</div>
            </el-col>

            <el-col :span="4">
              <el-input placeholder="人数" v-model="event.lottory1Count" clearable></el-input>
            </el-col>

            <el-col :span="4">
             <div class="title">二等奖人数：</div>
            </el-col>

            <el-col :span="4">
              <el-input placeholder="人数" v-model="event.lottory2Count" clearable></el-input>
            </el-col>

            <el-col :span="4">
             <div class="title">三等奖人数：</div>
            </el-col>

            <el-col :span="4">
              <el-input placeholder="人数" v-model="event.lottory3Count" clearable></el-input>
            </el-col>
          </el-row>

          <el-row :hidden.sync="event.type != '2'">
            <el-col :span="4">
             <div class="title">一等奖：</div>
            </el-col>

            <el-col :span="6">
              <el-input placeholder="如: 现金100元" v-model="event.lottory1" clearable></el-input>
            </el-col>
          </el-row>
          <el-row :hidden.sync="event.type != '2'">
            <el-col :span="4">
             <div class="title">二等奖：</div>
            </el-col>

            <el-col :span="6">
              <el-input placeholder="如: 现金50元" v-model="event.lottory2" clearable></el-input>
            </el-col>
          </el-row>
          <el-row :hidden.sync="event.type != '2'">
            <el-col :span="4">
             <div class="title">三等奖：</div>
            </el-col>

            <el-col :span="6">
              <el-input placeholder="如: 现金20元" v-model="event.lottory3" clearable></el-input>
            </el-col>
          </el-row>


          <div slot="footer" class="dialog-footer">
            <el-button @click="dialogFormVisible = false">取 消</el-button>
            <el-button type="primary" @click="doAddEvent(type)">确 定</el-button>
          </div>
        </el-dialog>

        <el-button style="width: 100%;margin-top: 5px;" icon="el-icon-plus" @click="addEvent">添加活动</el-button>
  </div>
</template>

<script>
  import axios from 'axios'
  import { Loading } from 'element-ui';
  export default {
    data() {
      return {
        q: '',
        lottoryTime: '',
        stores:[],
        events:[],
        dialogFormVisible: false,
        event: {
          stores: []
        },
        detailImage: [],
        type: 1,
        links: [
          {
            value: '1',
            label: '快闪店'
          }, {
            value: '2',
            label: '抽奖'
          }, {
            value: '3',
            label: '邀请'
          }],
        deletedIds: []
      }
    },
    methods: {
      clearDialog(done) {
        console.log(done)
        this.event = {
          stores: []
        }
        this.deletedIds = []
      },
      handleRemove(row,index) {
        this.event.stores.splice(index,1)
        this.deletedIds.push(row._id)
      },
      handleClick(row) {
        if (!this.event.stores) {
          this.event.stores = []
        }
        console.log(row)
        this.event.stores.push(row)
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
        if (!this.event.title) {
          this.$message.error("请输入活动名称")
          return
        }
        if (!this.event.index) {
          this.$message.error("请填写轮播位置")
          return
        }
        if (!this.event.image) {
          this.$message.error("请上传活动图片")
          return
        }
        if (this.detailImage.length == 0 && this.event.detailImage.length == 0) {
          this.$message.error("请上传详情图片")
          return
        }

        if (!this.event.btnText) {
          this.$message.error("请输入按钮文字")
          return
        }

        if(!this.event.type) {
          this.$message.error("请选择活动类型")
          return
        }

        if (this.event.type == '1') {
          if (!this.event.stores) {
            this.$message.error("请选择活动商家")
            return
          }
        }

        if (this.event.type == '2') {
          if (!this.event.lottoryTime) {
            this.$message.error("请设置开奖时间")
            return
          }
          if (!this.event.lottory1Count || !this.event.lottory2Count || !this.event.lottory3Count) {
            this.$message.error("请输入各奖项人数")
            return
          }
          if (!this.event.lottory1 || !this.event.lottory2 || !this.event.lottory3) {
            this.$message.error("请设置各奖项奖品")
            return
          }

        }


        let loadingInstance = Loading.service({ fullscreen: true });
        if (type == 1) {
          this.event.tp = 'eventadd'
          this.event.detailImage = this.detailImage
          axios.post('http://localhost:8090/eventadd',this.event).then(resp => {
            if (resp.status == 200) {
              this.$message.success("添加成功")
              this.event._id = resp.data._id
              this.events.push(this.event)
              this.event = {}
            }else {
              this.$message.error("添加失败")
            }
            this.dialogFormVisible = false
          })
        }else if (type == 2) {
          this.event.tp = 'eventupdate'
          this.event.dids = this.deletedIds
          if (this.detailImage.length > 0) {
            this.event.detailImage = this.detailImage
          }
          axios.post('http://localhost:8090/eventupdate',this.event).then(resp => {
            if (resp.status == 200) {
              this.$message.success("更新成功")
            }else {
              this.$message.error("更新失败")
            }
            this.dialogFormVisible = false
          })
        }
        this.detailImage = []
        loadingInstance.close()
      },
      handleDelete(index) {
        let event = this.events[index]
        event.tp = 'eventdelete'
        axios.post('http://localhost:8090/eventdelete',event).then(resp => {
          if (resp.status == 200) {
            this.events.splice(index,1)
            this.$message.success("删除成功")
          }else {
            this.$message.error("删除失败")
          }
        })
      },
      handleEdit(index) {
        this.event = this.events[index]
        this.type = 2
        this.dialogFormVisible = true
      },
      onDetailUploadSuccess(e) {
        this.detailImage.push(e)
      },
      onItemUploadSuccess(e){
        this.event.image = e
      },
      submitUpload(i){
        if (i == 1) {
          this.$refs.item.submit()
        }else if (i == 2) {
          this.$refs.itemDetail.submit()
        }

      }
    },
    mounted() {
      axios.get('http://localhost:8090/events').then(resp => {
        console.log(resp)
        this.events = resp.data.data
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
