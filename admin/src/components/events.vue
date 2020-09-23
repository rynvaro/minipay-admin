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
            prop="image"
            label="图片">
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

        <el-dialog title="添加活动" :visible.sync="dialogFormVisible">

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
              <div class="title">活动图片：</div>
            </el-col>
            <el-col :span="8">
               <el-upload class="upload-demo" ref="item" action="http://localhost:8090/upload" :on-success="onItemUploadSuccess" :auto-upload="false">
                 <el-button slot="trigger" size="small" type="primary">选取文件</el-button>
                 <el-button style="margin-left: 10px;" size="small" type="success" @click="submitUpload">上传到服务器</el-button>
                 <div slot="tip" class="el-upload__tip"></div>
               </el-upload>
            </el-col>
          </el-row>

          <el-row>
            <el-col :span="3">
              <div class="title">跳转连接：</div>
            </el-col>
            <el-col :span="18">
              <el-input placeholder="活动跳转连接" v-model="event.link" clearable></el-input>
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
  export default {
    data() {
      return {
        events:[],
        dialogFormVisible: false,
        event: {},
        type: 1,
      }
    },
    methods: {
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
        if (!this.event.link) {
          this.$message.error("请填写跳转连接")
          return
        }

        if (type == 1) {
          this.event.tp = 'eventadd'
          axios.post('http://localhost:8090/eventadd',this.event).then(resp => {
            if (resp.status == 200) {
              this.$message.success("添加成功")
              this.events.push(this.event)
            }else {
              this.$message.error("添加失败")
            }
            this.dialogFormVisible = false
          })
        }else if (type == 2) {
          this.event.tp = 'eventupdate'
          axios.post('http://localhost:8090/eventupdate',this.event).then(resp => {
            if (resp.status == 200) {
              this.$message.success("更新成功")
            }else {
              this.$message.error("更新失败")
            }
            this.dialogFormVisible = false
          })
        }

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
      onItemUploadSuccess(e){
        this.event.image = e
      },
      submitUpload(){
        this.$refs.item.submit()
      }
    },
    mounted() {
      axios.get('http://localhost:8090/events').then(resp => {
        console.log(resp)
        this.events = resp.data.data
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
