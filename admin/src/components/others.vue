<template>
  <div>
    <el-row>
      <el-col :span="3">
        <div class="title">Home背景图片：</div>
      </el-col>
      <el-col :span="4">
         <el-upload class="upload-demo" ref="item" action="http://localhost:8090/upload" :on-success="onItemUploadSuccess" :auto-upload="false">
           <el-button slot="trigger" size="small" type="primary">选取文件</el-button>
           <el-button style="margin-left: 10px;" size="small" type="success" @click="submitUpload(1)">上传到服务器</el-button>
           <div slot="tip" class="el-upload__tip"></div>
         </el-upload>
      </el-col>
      <el-col :span="4">
        <el-button style="width: 100%;" @click="update">更新</el-button>
      </el-col>
    </el-row>
  </div>
</template>

<script>
  import axios from 'axios'
  export default {
    data() {
      return {
        image: ''
      }
    },
    methods: {
      update(){
        if (!this.image) {
          this.$message.error("请上传图片")
          return
        }
        let data = {
          tp: 'homeheader',
          image: this.image
        }
        axios.post('http://localhost:8090/proxy',data).then(resp => {
          if (resp.status == 200) {
            this.$message.success("更新成功")
          }else {
            this.$message.error("更新成功")
          }
        })
      },
      onItemUploadSuccess(e){
        this.image = e
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
    }
  }
</script>

<style>
  .title{
    margin-top: 10px;
  }
</style>
